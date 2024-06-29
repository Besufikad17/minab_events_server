package helpers

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"time"
)

func saveImageToFile(input string) (string, error) {
	dec, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		panic(err)
	}

	dir, err := filepath.Abs("../../public/uploads")
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filepath.Join(dir, time.Now()))
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

	image, err := filepath.Abs(filepath.Join(dir, input))
	if err != nil {
		panic(err)
	}
	return image, err
}
