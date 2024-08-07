package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Besufikad17/minab_events/models"
	"github.com/skip2/go-qrcode"
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

func GenerateQR(data string) (*string, error) {
	qrCode, _ := qrcode.New(data, qrcode.Medium)
	fileHeader := strconv.FormatInt(time.Now().UnixMilli(), 10) + data
	fileName := fmt.Sprintf("%v.png", fileHeader)
	err := qrCode.WriteFile(256, fileName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(fmt.Sprintf("QR code generated and saved as %v.png", fileHeader))

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sourcePath := filepath.Join(currentDir, fileName)
	destPath := filepath.Join("public/assets/images/", fileName)

	err = os.Rename(sourcePath, destPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &fileName, nil
}

func UploadFile(filePath string) (*string, error) {
	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	uploadPreset := os.Getenv("CLOUDINARY_UPLOAD_PRESET")
	_ = writer.WriteField("upload_preset", uploadPreset)

	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	url := os.Getenv("CLOUDINARY_URL")
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var response models.ImageAsset
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}

	return &response.URL, nil
}
