
.PHONY: indexer

index:
	files=$$(find ../static/ -name "*_OCR.txt" | tr '\n' ' '); \
	go run ./cmd/indexer/main.go $$files
