package downloadFile

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, pw := io.Pipe()
			r.Header.Add("Content-Type", multipart.NewWriter(pw).FormDataContentType())
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, "<svg></svg>")
		}),
	)

	defer server.Close()
	fileDirectory = os.TempDir()
	fileName := "go_test.svg"
	filePath := filepath.Join(fileDirectory, fileName)
	_ = os.Remove(filePath)

	Download(server.URL, fileName)

	t.Log("should create a file named 'go_test.svg'")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Expected file 'go_test.svg' to exist")
	}
}
