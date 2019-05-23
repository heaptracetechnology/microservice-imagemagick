package imagemagick

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
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

	base64data := ""

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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick resize invalid args", func() {

	imageMagick := []byte(`{"status":false}`)
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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

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

var _ = Describe("Imagemagick reflect invalid args", func() {

	imageMagick := []byte(`{"status":false}`)
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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick reflect", func() {

	base64data := ""

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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
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

var _ = Describe("Imagemagick Extend", func() {

	base64data := ""

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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Extend with invalid agrs", func() {

	imageMagick := []byte(`{"status":false}`)
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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Transparent with invalid args", func() {

	imageMagick := []byte(`{"status":false}`)
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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Transparent", func() {

	base64data := ""

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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
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

var _ = Describe("Imagemagick Image Format ", func() {

	base64data := ""

	imageMagick := ImageMagick{InputImage: base64data, InputExtension: "png", OutputExtension: "pdf"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/format", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ImageFormat)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Image Format with invalid args", func() {

	imageMagick := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/format", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ImageFormat)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Image Format", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	imageMagick := ImageMagick{InputImage: base64data, InputExtension: "png", OutputExtension: "pdf"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/format", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ImageFormat)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Oil Paint ", func() {

	base64data := ""

	imageMagick := ImageMagick{InputImage: base64data, Radius: 2.2}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/oilpaint", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(OilPaint)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Oil Paint with invalid args", func() {

	imageMagick := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/oilpaint", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(OilPaint)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Oil Paint", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	imageMagick := ImageMagick{InputImage: base64data, Radius: 2.2}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/oilpaint", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(OilPaint)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Custom ", func() {

	base64data := ""

	imageMagick := CustomArgs{InputImage: base64data}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/oilpaint", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(OilPaint)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Custom with invalid args", func() {

	imageMagick := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/custom", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Custom)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Imagemagick Custom", func() {

	base64data, _ := Encodebase64("../uploads/mask.png")

	var customInput = []Function{
		Function{
			Name:   "resize",
			Height: 400,
			Width:  500,
		},
		Function{
			Name:             "extend",
			Height:           700,
			Width:            600,
			BackgroundColour: "red",
		},
		Function{
			Name: "reflect",
		},
	}

	imageMagick := CustomArgs{InputImage: base64data, CustomInput: customInput}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(imageMagick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/custom", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Custom)
	handler.ServeHTTP(recorder, request)

	Describe("Transparent the image", func() {
		Context("transparent image", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
