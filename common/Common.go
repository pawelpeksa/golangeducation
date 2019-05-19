package common

import (
	"encoding/json"
	"net/http"

	uuid "github.com/nu7hatch/gouuid"
)

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func UUID() (string, error) {
	uuid, err := uuid.NewV4()
	return uuid.String(), err
}
