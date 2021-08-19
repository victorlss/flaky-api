# Flaky API

![git hub action](https://github.com/victorlss/flaky-api/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/victorlss/flaky-api)](https://goreportcard.com/report/github.com/victorlss/flaky-api)

This application consumes a flaky API and download photos. 

The objective is handle these errors correctly so that all photos are downloaded.<br />
Downloading photos is slow so this application is about to optimize photos downloads, making use of concurrency.

## How to run
#### Install Dependencies
`go mod tidy`

#### Run
`go run cmd/download-photos/main.go`


