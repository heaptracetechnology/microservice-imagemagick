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

//ImageMagick struct
type ImageMagick struct {
	InputImage        string `json:"input,omitempty"`
	Height            int    `json:"height,omitempty"`
	Width             int    `json:"width,omitempty"`
	Colour            string `json:"background_colour,omitempty"`
	TransparentColour string `json:"transparent_colour,omitempty"`
	OutputExtension   string `json:"output_extension,omitempty"`
	InputExtension    string `json:"input_extension,omitempty"`
}

//Message struct
type Message struct {
	Success    string `json:"success"`
	Message    string `json:"output"`
	StatusCode int    `json:"statuscode"`
}

//Resize image
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

	f, _ := os.Create("../uploads/input_image.jpg")
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if err := f.Sync(); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.ReadImage("../uploads/input_image.jpg")
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	hWidth := uint(param.Height)
	hHeight := uint(param.Width)

	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	if err := mw.WriteImage("output_image.png"); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	imgFile, _ := os.Open("output_image.png")

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	deleteFile("../uploads/input_image.jpg", "output_image.png")

	message := Message{"true", imgBase64Str, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Reflect image
func Reflect(responseWriter http.ResponseWriter, request *http.Request) {

	imagick.Initialize()
	defer imagick.Terminate()

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

	f, _ := os.Create("../uploads/input_image.jpg")
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if err := f.Sync(); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.ReadImage("../uploads/input_image.jpg")
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	w := mw.GetImageWidth()
	h := mw.GetImageHeight()

	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	mwr := mw.Clone()

	mwr.ResizeImage(w, h/2, imagick.FILTER_LANCZOS, 1)
	mwr.FlipImage()

	mwg := imagick.NewMagickWand()
	mwg.SetSize(w, h/2)
	mwg.ReadImage("gradient:white-black")

	mwr.CompositeImage(mwg, imagick.COMPOSITE_OP_COPY_OPACITY, 0, 0)

	mw.AddImage(mwr)
	mw.SetFirstIterator()

	mwout := mw.AppendImages(true)

	if err := mwout.WriteImage("output_image.png"); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	imgFile, _ := os.Open("output_image.png")

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	deleteFile("../uploads/input_image.jpg", "output_image.png")

	message := Message{"true", imgBase64Str, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Extend image
func Extend(responseWriter http.ResponseWriter, request *http.Request) {

	imagick.Initialize()
	defer imagick.Terminate()

	var err error

	mw := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()

	decoder := json.NewDecoder(request.Body)
	var param ImageMagick
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	pw.SetColor(param.Colour)
	dec, err := base64.StdEncoding.DecodeString(param.InputImage)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	f, _ := os.Create("../uploads/input_image.jpg")
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if err := f.Sync(); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.ReadImage("../uploads/input_image.jpg")
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	w := int(mw.GetImageWidth())
	h := int(mw.GetImageHeight())
	mw.SetImageBackgroundColor(pw)

	err = mw.ExtentImage(uint(param.Width), uint(param.Height), -(param.Width-w)/2, -(param.Height-h)/2)
	if err != nil {
		panic(err)
	}

	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		panic(err)
	}

	if err := mw.WriteImage("output_image.png"); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	imgFile, _ := os.Open("output_image.png")

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	deleteFile("../uploads/input_image.jpg", "output_image.png")

	message := Message{"true", imgBase64Str, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Transparent Image
func Transparent(responseWriter http.ResponseWriter, request *http.Request) {
	imagick.Initialize()
	defer imagick.Terminate()

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

	f, _ := os.Create("../uploads/input_image.jpg")
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if err := f.Sync(); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.ReadImage("../uploads/input_image.jpg")
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	target := imagick.NewPixelWand()
	target.SetColor(param.TransparentColour)
	mw.TransparentPaintImage(target, 0, 10, false)

	if err := mw.WriteImage("output_image.png"); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	imgFile, _ := os.Open("output_image.png")

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	deleteFile("../uploads/input_image.jpg", "output_image.png")

	message := Message{"true", imgBase64Str, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ImageFormat conversion
func ImageFormat(responseWriter http.ResponseWriter, request *http.Request) {
	mw := imagick.NewMagickWand()
	mw.SetSize(640, 480)

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

	inputImageName := "../uploads/input_image." + param.InputExtension
	f, _ := os.Create(inputImageName)
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if err := f.Sync(); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	err = mw.ReadImage(inputImageName)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	formatErr := mw.SetFormat(param.OutputExtension)
	if formatErr != nil {
		result.WriteErrorResponse(responseWriter, formatErr)
		return

	}

	outputImageName := "output_image." + param.OutputExtension
	fmt.Println("outputImageName :::", outputImageName)

	if err := mw.WriteImage(outputImageName); err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	imgFile, _ := os.Open(outputImageName)

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	deleteFile(inputImageName, outputImageName)

	message := Message{"true", imgBase64Str, http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

func deleteFile(inputImage, outputImage string) {
	os.Remove(inputImage)
	os.Remove(outputImage)
}
