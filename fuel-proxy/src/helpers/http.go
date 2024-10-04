package helpers

import "net/http"

func HttpWriteError(w http.ResponseWriter, httpStatus int, errText string) (int, error) {
	w.WriteHeader(httpStatus)
	return w.Write([]byte(errText))
}
