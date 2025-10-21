
.PHONY: index run

index:
	files=$$(cd .. && find static/ -name "*_OCR.txt" | tr '\n' ' '); \
	go run ./cmd/indexer/main.go $$files

run:
	go run ./cmd/server/main.go
