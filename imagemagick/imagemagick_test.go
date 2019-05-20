package imagemagick

import (
	"encoding/base64"
	//"fmt"
	"bytes"
	"encoding/json"
	//"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

//Base64 encoder
func Encodebase64(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buff), nil
}

var _ = Describe("Imagemagick resize", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	imageMagick := ImageMagick{InputImage: base64data, Height: 700, Width: 500}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/resize", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Resize)
	handler.ServeHTTP(recorder, request)

	Describe("Resize the image", func() {
		Context("resize image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick reflect", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	imageMagick := ImageMagick{InputImage: base64data}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/reflect", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Reflect)
	handler.ServeHTTP(recorder, request)

	Describe("Reflect the image", func() {
		Context("reflect image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Extend", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	imageMagick := ImageMagick{InputImage: base64data, Height: 700, Width: 500, Colour: "red"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/extend", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Extend)
	handler.ServeHTTP(recorder, request)

	Describe("Extend the image", func() {
		Context("extend image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Transparent", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	imageMagick := ImageMagick{InputImage: base64data, Height: 700, Width: 500, TransparentColour: "white"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/extend", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Transparent)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
