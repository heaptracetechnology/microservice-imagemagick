# ImageMagick as a microservice
An OMG service for ImageMagick, it is for displaying, converting, and editing raster image and vector image files.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
<!-- [![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com) -->


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Resize image
```sh
$ omg run resize -a input=<IMAGE_BASE64_DATA> -a height=<HEIGHT> -a width=<WIDTH>
```
##### Reflect image
```sh
$ omg run reflect -a input=<IMAGE_BASE64_DATA>
```
##### Extend image
```sh
$ omg run extend -a input=<IMAGE_BASE64_DATA> -a height=<HEIGHT> -a width=<WIDTH> -a background_colour=<COLOUR_NAME>
```
##### Transparent image
```sh
$ omg run transparent -a input=<IMAGE_BASE64_DATA> -a transparent_colour=<COLOUR_NAME>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-imagemagick .
```
### RUN
```
docker run -p 3000:3000 microservice-imagemagick
```
