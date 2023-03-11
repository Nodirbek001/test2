package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Error bool        `json:"error"`
	Data  interface{} `json: "data"`
}
type errorInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}
func HandleErrorResponse(w http.ResponseWriter, errCode int, errMessage string) {
	log.Println(errMessage+" code:", errCode)
	w.WriteHeader(errCode)
	writeJSON(w, response{
		Error: true,
		Data: errorInfo{
			Status:  errCode,
			Message: errMessage,
		},
	})
}
func writeJSON(w http.ResponseWriter, date interface{}) {
	bytes, _ := json.MarshalIndent(date, "", "  ") //nima bu
	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)

}
