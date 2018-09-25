package apistuff

import "net/http"

// handleHealth simply returns a status struct that says the service is healthy
func HandleHealthOK(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := StatusMessage("OK", "")
	if err != nil {
		HttpError(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
