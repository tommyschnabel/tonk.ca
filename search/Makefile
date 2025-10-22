
.PHONY: index run unzipresults

unzipresults:
	cd .. && unzip ocr/results.zip

index:
	files=$$(find ../static/ -name "*_OCR.txt" | tr '\n' ' '); \
	go run ./cmd/indexer/main.go $$files

build:
	GOOS=linux GOARCH=amd64 go build -o search_server ./cmd/server/main.go

run:
	go run ./cmd/server/main.go
