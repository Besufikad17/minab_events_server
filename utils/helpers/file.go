package helpers

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func isImage(mimeType string) bool {
	return mimeType == "image/jpeg" || mimeType == "image/jpg" || mimeType == "image/png" || mimeType == "image/gif" || mimeType == "image/webp"
}

func SaveImageToFile(input string) (*string, error) {
	b64data := input[strings.IndexByte(input, ',')+1:]
	dec, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		panic(err)
	}

	dir, err := filepath.Abs("public/uploads")
	if err != nil {
		panic(err)
	}

	mimeType := strings.Split(strings.Split(input, ";")[0], ":")[1]

	if !isImage(mimeType) {
		return nil, errors.New("INVALID FILE TYPE")
	}

	fileName := strconv.FormatInt(time.Now().UnixMilli(), 10) + "." + strings.Split(mimeType, "/")[1]
	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.Write(dec); err != nil {
		panic(err)
	}

	if err := file.Sync(); err != nil {
		panic(err)
	}

	_, err = filepath.Abs(filepath.Join(dir, fileName))
	if err != nil {
		panic(err)
	}

	return &fileName, err
}
