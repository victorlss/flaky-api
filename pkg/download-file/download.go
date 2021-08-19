package downloadfile

import (
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	fileDirectory = filepath.Join(os.TempDir())
)

const (
	retries = 5
)

// Download a file from URL and store on file path
func Download(url string, filename string) {
	filePath := filepath.Join(fileDirectory, filename)

	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = retries

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	createDirectoryIfNotExists(filePath)
	out, _ := os.Create(filePath)
	_, _ = io.Copy(out, resp.Body)

	_ = resp.Body.Close()
	_ = out.Close()
}

func createDirectoryIfNotExists(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_ = os.Mkdir(filepath.Dir(filePath), 0755)
	}
}
