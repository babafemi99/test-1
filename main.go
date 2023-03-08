package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"test1/cmd/handler/videoHandler"
	"test1/internal/service/VideoService"
	"test1/internal/utils"
	"test1/internal/videoRepo"
	"time"
)

func main() {
	// set up env variables

	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("UNABLE TO LOAD ENVIRONMENT CONFIGURATIONS")
	}

	apiKey := config.ApiKey
	if apiKey == "" {
		log.Fatal("Provide an API key please")
	}

	cloudName := config.CloudName
	if cloudName == "" {
		log.Fatal("Provide a valid cloud name please")
	}

	uploadPreset := config.UploadPreset
	if uploadPreset == "" {
		log.Fatal("Provide a valid upload preset please")
	}

	baseURL := config.BaseUrl
	if baseURL == "" {
		log.Fatal("Provide a valid base URL please")
	}

	port := config.Port
	if port == "" {
		log.Fatal("Provide a valid port please")
	}

	// start application
	cloudinaryObj := videoRepo.CreateCloudinaryObject(
		apiKey,
		uploadPreset,
		baseURL,
		cloudName,
	)

	videoSrv := VideoService.NewVideoSrv(cloudinaryObj)
	handler := videoHandler.NewVideoHandler(videoSrv)

	r := mux.NewRouter()
	r.HandleFunc("/upload-video", handler.UploadVideo).Methods("POST")

	srvDetails := http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     r,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		log.Println("SERVER STARTING ON PORT:", 9095)
		err := srvDetails.ListenAndServe()
		if err != nil {
			log.Printf("ERROR STARTING SERVER: %v", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Printf("Closing now, We've gotten signal: %v", sig)

	ctx := context.Background()
	srvDetails.Shutdown(ctx)
}
