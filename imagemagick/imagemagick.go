package imagemagick

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
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

	f, err := os.Create("../uploads/input_image.jpg")
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

	imgFile, err := os.Open("output_image.png")

	if err != nil {
		os.Exit(1)
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

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

	f, err := os.Create("../uploads/input_image.jpg")
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

	imgFile, err := os.Open("output_image.png")

	if err != nil {
		os.Exit(1)
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

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

	f, err := os.Create("../uploads/input_image.jpg")
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

	imgFile, err := os.Open("output_image.png")

	if err != nil {
		os.Exit(1)
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

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

	f, err := os.Create("../uploads/input_image.jpg")
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

	imgFile, err := os.Open("output_image.png")

	if err != nil {
		os.Exit(1)
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

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
	var deleteOutputImage = os.Remove("output_image.png")
	if deleteOutputImage != nil {
		return deleteOutputImage
	}

	var deleteInputImage = os.Remove("../uploads/input_image.jpg")
	if deleteInputImage != nil {
		return deleteInputImage
	}
	return nil
}
