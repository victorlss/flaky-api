package main

import (
	homeVisionApi "github.com/victorlss/flaky-api/pkg/api/home-vision"
	downloadFile "github.com/victorlss/flaky-api/pkg/download-file"
	"github.com/victorlss/flaky-api/pkg/file"
	"github.com/victorlss/flaky-api/pkg/shared/types"
)

func main() {
	houseChan := make(chan types.House)
	go homeVisionApi.Fetch(houseChan)

	for {
		house, ok := <-houseChan

		if !ok {
			break
		}

		filename := file.GenerateFilename(house)
		go downloadFile.Download(house.PhotoURL, filename)
	}
}
