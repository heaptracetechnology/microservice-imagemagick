# ImageMagick as a microservice
An OMG service for ImageMagick, it is for displaying, converting, and editing raster image and vector image files.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-imagemagick.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-imagemagick)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-imagemagick/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-imagemagick)
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
##### Format image
```sh
$ omg run format -a input=<IMAGE_BASE64_DATA> -a input_extension=<INPUT_EXTENSION> -a output_extension=<OUTPUT_EXTENSION>
```
##### OilPaint image
```sh
$ omg run oilpaint -a input=<IMAGE_BASE64_DATA> -a radius=<RADIUS>
```
##### Custom image
```sh
$ omg run custom -a input=<IMAGE_BASE64_DATA> -a customize_input=<CUSTOMIZE_INPUT>
```
##### Custom image Example
> omg run custom -a input=<IMAGE_BASE64_DATA> -a customize_input='[{"name":"resize","height": 400,"width": 500},{"name":"extend","background_colour": "red","height": 700,"width": 600},{"name":"reflect"},{"name":"oilpaint","radius":"2.6"}]'

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
