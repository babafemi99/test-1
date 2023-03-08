package videoHandler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"test1/internal/entities/ResEntity"
	"test1/internal/service/VideoService"
)

type videoHandler struct {
	VideoService.VideoSrv
}

func NewVideoHandler(videoSrv VideoService.VideoSrv) videoHandler {
	return videoHandler{VideoSrv: videoSrv}
}

func (v videoHandler) UploadVideo(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(ResEntity.CustomErrorResponse("something went wrong", http.StatusInternalServerError))
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(ResEntity.CustomErrorResponse("specify a file name please", http.StatusInternalServerError))
		return
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(ResEntity.CustomErrorResponse("something went wrong", http.StatusInternalServerError))
		return
	}
	defer f.Close()

	path := filepath.Join(".", "files")

	_ = os.MkdirAll(path, os.ModePerm)

	fullPath := path + "/" + name

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(ResEntity.CustomErrorResponse("something went wrong", http.StatusInternalServerError))
		return
	}
	defer file.Close()

	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(ResEntity.CustomErrorResponse("something went wrong", http.StatusInternalServerError))
		return
	}

	log.Println(file.Name())
	video, errMsg := v.VideoSrv.UploadVideo(file.Name())
	if errMsg != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errMsg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(ResEntity.CustomSuccessResponse(http.StatusOK, video))

}
