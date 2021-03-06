package route

import (
    "github.com/gorilla/mux"
    imagemagick "github.com/heaptracetechnology/microservice-imagemagick/imagemagick"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Resize",
        "POST",
        "/resize",
        imagemagick.Resize,
    },
    Route{
        "Reflect",
        "POST",
        "/reflect",
        imagemagick.Reflect,
    },
    Route{
        "Extend",
        "POST",
        "/extend",
        imagemagick.Extend,
    },
    Route{
        "Transparent",
        "POST",
        "/transparent",
        imagemagick.Transparent,
    },
    Route{
        "ImageFormat",
        "POST",
        "/format",
        imagemagick.ImageFormat,
    },
    Route{
        "Custom",
        "POST",
        "/custom",
        imagemagick.Custom,
    },
    Route{
        "OilPaint",
        "POST",
        "/oilpaint",
        imagemagick.OilPaint,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
