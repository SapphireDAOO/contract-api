package server

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func inTempDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	original, err := os.Getwd()
	if err != nil {
		t.Fatalf("reading the working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("changing to the temp directory: %v", err)
	}
	t.Cleanup(func() { os.Chdir(original) })

	return dir
}

func TestLoadEnvSkipsInProduction(t *testing.T) {
	inTempDir(t)
	t.Setenv("PRODUCTION", "1")

	if err := loadEnv(); err != nil {
		t.Errorf("loadEnv returned %v in production, want nil", err)
	}
}

func TestLoadEnvSkipsOnAnEmptyProductionValue(t *testing.T) {
	inTempDir(t)
	t.Setenv("PRODUCTION", "")

	if err := loadEnv(); err != nil {
		t.Errorf("loadEnv returned %v, want nil", err)
	}
}

func TestLoadEnvReadsTheFile(t *testing.T) {
	dir := inTempDir(t)
	os.Unsetenv("PRODUCTION")

	contents := "TEST_LOAD_ENV_KEY=the-value\nTEST_LOAD_ENV_OTHER=second\n"
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte(contents), 0o600); err != nil {
		t.Fatalf("writing the .env: %v", err)
	}
	t.Cleanup(func() {
		os.Unsetenv("TEST_LOAD_ENV_KEY")
		os.Unsetenv("TEST_LOAD_ENV_OTHER")
	})

	if err := loadEnv(); err != nil {
		t.Fatalf("loadEnv returned %v", err)
	}

	if got := os.Getenv("TEST_LOAD_ENV_KEY"); got != "the-value" {
		t.Errorf("TEST_LOAD_ENV_KEY = %q, want the-value", got)
	}
	if got := os.Getenv("TEST_LOAD_ENV_OTHER"); got != "second" {
		t.Errorf("TEST_LOAD_ENV_OTHER = %q, want second", got)
	}
}

func TestLoadEnvFailsWithoutAFile(t *testing.T) {
	inTempDir(t)
	os.Unsetenv("PRODUCTION")

	err := loadEnv()

	if err == nil {
		t.Fatal("loadEnv returned no error for a missing .env")
	}
	if !strings.Contains(err.Error(), "error loading .env file") {
		t.Errorf("error = %q, want it to name the .env", err)
	}
}

func TestLoadEnvDoesNotOverrideTheEnvironment(t *testing.T) {
	dir := inTempDir(t)
	os.Unsetenv("PRODUCTION")
	t.Setenv("TEST_LOAD_ENV_KEY", "from-the-environment")

	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte("TEST_LOAD_ENV_KEY=from-the-file\n"), 0o600); err != nil {
		t.Fatalf("writing the .env: %v", err)
	}

	if err := loadEnv(); err != nil {
		t.Fatalf("loadEnv returned %v", err)
	}

	if got := os.Getenv("TEST_LOAD_ENV_KEY"); got != "from-the-environment" {
		t.Errorf("TEST_LOAD_ENV_KEY = %q, want the environment to win", got)
	}
}
