package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/schers/test-mafin/db"
	"github.com/schers/test-mafin/upload"
	"log"
	"net/http"
	"os"
)

var storage string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}

	dataSourceName, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Panic("No db init")
	}

	storage, ok = os.LookupEnv("STORAGE")
	if !ok {
		log.Panic("No storage init")
	}

	err := db.InitDB(dataSourceName)
	if err != nil {

	}
}

func main() {
	host := "localhost:8080"

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(static.Serve("/", static.LocalFile(storage, false)))

	r.POST("/api/v1/image/upload", uploadImage)

	log.Printf("Storage place in: %s", storage)
	log.Printf("Start server on %s", host)
	r.Run(host)
}

func uploadImage(context *gin.Context) {
	originalFile, err := upload.Upload(context.Request, storage)
	if err != nil {
		initErrorResponse(context, err, http.StatusBadRequest)
		return
	}

	file, err := upload.CreateFile(storage, originalFile)
	if err != nil {
		initErrorResponse(context, err, http.StatusInternalServerError)
		return
	}

	err = file.SaveInfo()
	if err != nil {
		initErrorResponse(context, err, http.StatusInternalServerError)
		return
	}

	data := file.ToJson()

	context.JSON(http.StatusCreated, gin.H{"status": "ok", "files": data})
}

func initErrorResponse(context *gin.Context, err error, status int) {
	context.JSON(status, gin.H{
		"status": "error",
		"error":  fmt.Sprintf("Upload error: %q", err.Error()),
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, Content-Range, Content-Disposition, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
	}
}
