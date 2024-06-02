package grpcutil

import (
	"net/http"

	"google.golang.org/grpc/status"
)

func GRPCToHTTPError(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}

	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, "Unknown error"
	}

	httpStatus, exists := grpcToHTTPStatus[st.Code()]
	if !exists {
		httpStatus = http.StatusInternalServerError
	}

	return httpStatus, st.Message()
}
