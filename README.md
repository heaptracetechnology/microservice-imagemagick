# _ImageMagick_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-imagemagick.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-imagemagick)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-imagemagick/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-imagemagick)

An OMG service for ImageMagick, it is for displaying, converting, and editing raster image and vector image files.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Resize image
```coffee
>>> imagemagick resize input:'base64image' height:'resizeHight' width:'resizeWidth'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```
##### Reflect image
```coffee
>>> imagemagick reflect input:'base64image'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```
##### Extend image
```coffee
>>> imagemagick extend input:'base64image' height:'extendHight' width:'extendHight' backgroundColour:'colourName'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```
##### Transparent image
```coffee
>>> imagemagick transparent input:'base64image' transparentColour:'transparentColour'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```
##### Format image
```coffee
>>> imagemagick format input:'base64image' inputExtension:'inputImageExtension' outputExtension:'outputImageExtension'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```
##### OilPaint image
```coffee
>>> imagemagick oilpaint input:'base64image' radius:'radius'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```
##### Custom image
```coffee
>>> imagemagick custom input:'base64image' customizeInput:'[{"name":"resize","height": 400,"width": 500},{"name":"extend","backgroundColour": "red","height": 700,"width": 600},{"name":"reflect"},{"name":"oilpaint","radius":2.6}]'
{"success":"true/false","output":"Base64 output image","statusCode":"HTTPstatusCode"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Resize image
```shell
$ omg run resize -a input=<IMAGE_BASE64_DATA> -a height=<HEIGHT> -a width=<WIDTH>
```
##### Reflect image
```shell
$ omg run reflect -a input=<IMAGE_BASE64_DATA>
```
##### Extend image
```shell
$ omg run extend -a input=<IMAGE_BASE64_DATA> -a height=<HEIGHT> -a width=<WIDTH> -a backgroundColour=<COLOUR_NAME>
```
##### Transparent image
```shell
$ omg run transparent -a input=<IMAGE_BASE64_DATA> -a transparentColour=<COLOUR_NAME>
```
##### Format image
```shell
$ omg run format -a input=<IMAGE_BASE64_DATA> -a inputExtension=<INPUT_EXTENSION> -a outputExtension=<OUTPUT_EXTENSION>
```
##### OilPaint image
```shell
$ omg run oilpaint -a input=<IMAGE_BASE64_DATA> -a radius=<RADIUS>
```
##### Custom image
```shell
$ omg run custom -a input=<IMAGE_BASE64_DATA> -a customizeInput=<CUSTOMIZE_INPUT>
```
##### Custom image Example
```
$ omg run custom -a input=<IMAGE_BASE64_DATA> -a customizeInput='[{"name":"resize","height": 400,"width": 500},{"name":"extend","backgroundColour": "red","height": 700,"width": 600},{"name":"reflect"},{"name":"oilpaint","radius":2.6}]'
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/imagemagick/blob/master/LICENSE).
