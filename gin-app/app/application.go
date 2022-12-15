package app

import (
    "log"
    "github.com/gin-gonic/gin"
)

var (
    router = gin.Default()
)

func StartApp() {
    mapUrls()

    log.Println("Start App...")
    router.Run()
}