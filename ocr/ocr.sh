#!/bin/bash

files="$(find ../static -name *.pdf)"
for file in ${files[@]}; do
    target="${file%.pdf}_ocr.zip"
    extracted="${file%.pdf}_OCR.txt"
    resultPDF="${file%.pdf}_OCR.pdf"
    echo "Starting $file -> $target"

    if [[ -e $target ]]; then
        echo "$target already exists, skipping ocr..."
    else
        curl -X 'POST' \
            'http://localhost:8080/api/v1/misc/ocr-pdf' \
            -H 'accept: */*' \
            -H 'Content-Type: multipart/form-data' \
            -F 'removeImagesAfter=true' \
            -F 'clean=true' \
            -F 'deskew=true' \
            -F 'cleanFinal=true' \
            -F 'ocrRenderType=hocr' \
            -F "fileInput=@$file;type=application/pdf" \
            -F 'ocrType=skip-text' \
            -F 'languages=eng' \
            -F 'sidecar=true' \
            --output $target
        echo "OCR Finished for $file"
    fi

    echo "Unzipping $target..."
    unzip -o $target -d "$(dirname $target)"
    rm $resultPDF # not needed
    zip ./results.zip $extracted

    echo -e "Finished with $file\n\n"
done
