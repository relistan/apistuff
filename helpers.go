package apistuff

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

// writeTempFile will stream a reader out to a temporary file, with
// the filename being the fname passed, written to a temporary dir.
// The whole filepath is returned.
func WriteTempFile(input io.Reader, fname string, prefix string) (filename string, err error) {
	tmpDir, err := ioutil.TempDir("", prefix)
	if err != nil {
		return "", err
	}

	tmpFname := path.Join(tmpDir, path.Base(fname))
	tmpFile, err := os.OpenFile(tmpFname, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}

	io.Copy(tmpFile, input)
	tmpFile.Close()

	return tmpFile.Name(), nil
}

func StatusMessage(status, message string) ([]byte, error) {
	healthData := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: message,
	}

	return json.MarshalIndent(healthData, "", "  ")
}

func HttpError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	msg, _ := StatusMessage("error", message)
	http.Error(w, string(msg), code)
}
