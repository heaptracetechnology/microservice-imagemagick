package imagemagick

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-imagemagick/result"
	"gopkg.in/gographics/imagick.v2/imagick"
	"net/http"
	"os"
)

type ImageMagick struct {
	InputImage string `json:"input_image,omitempty"`
	Height     int    `json:"height,omitempty"`
	Width      int    `json:"width,omitempty"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

//Resize
func Resize(responseWriter http.ResponseWriter, request *http.Request) {

	imagick.Initialize()
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	decoder := json.NewDecoder(request.Body)
	var param ImageMagick
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	dec, err := base64.StdEncoding.DecodeString(param.InputImage)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	f, err := os.Create("./uploads/inputfile.jpg")
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if err := f.Sync(); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.ReadImage("./uploads/inputfile.jpg")
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	// Get original logo size
	// width := mw.GetImageWidth()
	// height := mw.GetImageHeight()

	// Calculate half the size
	// hWidth := uint(width / 2)
	// hHeight := uint(height / 2)

	hWidth := uint(param.Height)
	hHeight := uint(param.Width)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	//image := mw.GetImageFromMagickWand()

	if err := mw.WriteImage("resized_image.png"); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	imgFile, err := os.Open("resized_image.png") // a QR code image

	if err != nil {
		os.Exit(1)
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	deleteError := deleteFile()
	if deleteError != nil {
		result.WriteErrorResponse(responseWriter, deleteError)
		return
	}

	message := Message{"true", imgBase64Str, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

func deleteFile() (err error) {
	// delete file
	var err1 = os.Remove("resized_image.png")
	if err1 != nil {
		return err1
	}

	fmt.Println("==> done deleting file")

	var err2 = os.Remove("./uploads/inputfile.jpg")
	if err2 != nil {
		return err2
	}

	fmt.Println("==> done deleting file")
	return nil
}
