package paymentprocessorstorage

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func captureLog(t *testing.T, run func()) []map[string]any {
	t.Helper()

	var buf bytes.Buffer
	previous := slog.Default()
	slog.SetDefault(slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(previous) })

	run()

	var records []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
		if line == "" {
			continue
		}
		var record map[string]any
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			t.Fatalf("log line is not JSON: %v (%q)", err, line)
		}
		records = append(records, record)
	}
	return records
}

func watcher() *PaymentProcessorStorage {
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")
	return &PaymentProcessorStorage{address: &address, explorerURL: "https://basescan.org"}
}

func TestReportPauseEnded(t *testing.T) {
	past := big.NewInt(time.Now().Add(-time.Hour).Unix())
	future := big.NewInt(time.Now().Add(time.Hour).Unix())

	tests := []struct {
		name    string
		expiry  *big.Int
		wantMsg string
	}{
		{
			name:    "expiry has passed, so the pause elapsed",
			expiry:  past,
			wantMsg: "emergency pause elapsed",
		},
		{
			name:    "expiry is still ahead, so someone lifted it",
			expiry:  future,
			wantMsg: "pause lifted before it expired, reported by the Unpaused event",
		},
		{
			name:    "no expiry known",
			expiry:  nil,
			wantMsg: "pause ended with no known expiry; leaving it to the Unpaused event",
		},
		{
			name:    "zero expiry is not an expiry",
			expiry:  big.NewInt(0),
			wantMsg: "pause ended with no known expiry; leaving it to the Unpaused event",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := watcher()

			records := captureLog(t, func() { c.reportPauseEnded(tt.expiry) })

			if len(records) == 0 {
				t.Fatal("nothing was logged")
			}
			// The decision is the first line; a lapse then attempts the
			// Discord send, which logs again.
			if records[0]["msg"] != tt.wantMsg {
				t.Errorf("msg = %v, want %q", records[0]["msg"], tt.wantMsg)
			}
		})
	}
}

// An explicit unpause is already announced by the Unpaused event listener, so
// announcing it here too would double-report it.
func TestAnExplicitUnpauseIsNotReportedAsElapsed(t *testing.T) {
	c := watcher()
	future := big.NewInt(time.Now().Add(30 * time.Minute).Unix())

	records := captureLog(t, func() { c.reportPauseEnded(future) })

	for _, record := range records {
		if record["msg"] == "emergency pause elapsed" {
			t.Errorf("an unpause before expiry was reported as a lapse: %v", record)
		}
	}
}

func TestWatcherDisabledWithoutAContract(t *testing.T) {
	var c *PaymentProcessorStorage

	records := captureLog(t, func() { c.WatchEmergencyPause(t.Context()) })

	if len(records) == 0 || records[0]["msg"] != "emergency pause watcher disabled" {
		t.Errorf("records = %v, want the disabled warning", records)
	}
}

func TestPauseEpisode(t *testing.T) {
	future := big.NewInt(time.Now().Add(time.Hour).Unix())
	past := big.NewInt(time.Now().Add(-time.Hour).Unix())

	type reading struct {
		paused bool
		expiry *big.Int
	}

	tests := []struct {
		name       string
		readings   []reading
		wantEnded  bool
		wantExpiry *big.Int
	}{
		{
			name:     "still paused",
			readings: []reading{{true, future}, {true, future}},
		},
		{
			name:       "lapsed once the expiry passed",
			readings:   []reading{{true, past}, {false, past}},
			wantEnded:  true,
			wantExpiry: past,
		},
		{
			name:       "unpaused before the expiry",
			readings:   []reading{{true, future}, {false, big.NewInt(0)}},
			wantEnded:  true,
			wantExpiry: future,
		},
		{
			name:     "never paused",
			readings: []reading{{false, big.NewInt(0)}, {false, big.NewInt(0)}},
		},
		{
			name:     "paused after already being unpaused is not an end",
			readings: []reading{{false, big.NewInt(0)}, {true, future}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var episode pauseEpisode
			var ended bool
			var at *big.Int

			for _, r := range tt.readings {
				ended, at = episode.observe(r.paused, r.expiry)
			}

			if ended != tt.wantEnded {
				t.Fatalf("ended = %v, want %v", ended, tt.wantEnded)
			}
			if !tt.wantEnded {
				return
			}
			if at == nil || at.Cmp(tt.wantExpiry) != 0 {
				t.Errorf("expiry = %v, want %v", at, tt.wantExpiry)
			}
		})
	}
}

// An emergency pause that lapses must not leave its expiry behind: the next
// pause to end would otherwise be reported as having lapsed too, even when
// someone lifted it deliberately.
func TestAnOwnerPauseAfterALapseIsNotReportedAsALapse(t *testing.T) {
	past := big.NewInt(time.Now().Add(-time.Hour).Unix())
	none := big.NewInt(0)

	var episode pauseEpisode

	episode.observe(true, past)               // emergency pause, active
	ended, at := episode.observe(false, past) // it lapses
	if !ended || at.Cmp(past) != 0 {
		t.Fatalf("the lapse was not reported: ended=%v at=%v", ended, at)
	}

	episode.observe(true, none)              // an owner pause, no emergency expiry
	ended, at = episode.observe(false, none) // someone lifts it

	if !ended {
		t.Fatal("the owner pause ending was not observed")
	}
	if at != nil {
		t.Errorf("expiry = %v, want nil so it is not reported as a lapse", at)
	}
}

// The watcher primes itself from the chain before its first tick. Without
// that, a pause lapsing between startup and the first tick shows no
// transition and is never reported.
func TestPrimingCatchesALapseRightAfterStartup(t *testing.T) {
	past := big.NewInt(time.Now().Add(-time.Hour).Unix())

	t.Run("unprimed misses it", func(t *testing.T) {
		var episode pauseEpisode

		ended, _ := episode.observe(false, past)

		if ended {
			t.Error("an unprimed episode reported a transition it could not have seen")
		}
	})

	t.Run("primed catches it", func(t *testing.T) {
		var episode pauseEpisode
		episode.observe(true, past)

		ended, at := episode.observe(false, past)

		if !ended {
			t.Fatal("the lapse went unreported despite priming")
		}
		if at == nil || at.Cmp(past) != 0 {
			t.Errorf("expiry = %v, want %v", at, past)
		}
	})
}

// Priming discards observe's result, which is only safe because a first
// reading can never be a transition.
func TestPrimingNeverReportsATransition(t *testing.T) {
	future := big.NewInt(time.Now().Add(time.Hour).Unix())

	for _, paused := range []bool{true, false} {
		var episode pauseEpisode

		if ended, _ := episode.observe(paused, future); ended {
			t.Errorf("priming with paused=%v reported a transition", paused)
		}
	}
}
