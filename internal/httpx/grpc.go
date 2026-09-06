package httpx

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WriteGRPCError(w http.ResponseWriter, err error, msg string) {
	st, ok := status.FromError(err)
	if !ok {
		WriteHTTPErrorWithStatus(w, http.StatusBadGateway, err, msg)
		return
	}

	var statusCode int
	switch st.Code() {
	case codes.InvalidArgument:
		statusCode = http.StatusBadRequest
	case codes.Unavailable:

		statusCode = http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		statusCode = http.StatusGatewayTimeout
	default:
		statusCode = http.StatusBadGateway
	}

	WriteHTTPErrorWithStatus(w, statusCode, errors.New(st.Message()), msg)
}
