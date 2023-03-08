package VideoService

import (
	"context"
	"log"
	"net/http"
	"os"
	"test1/internal/entities"
	"test1/internal/entities/ResEntity"
	"test1/internal/videoRepo"
	"time"
)

type VideoSrv interface {
	UploadVideo(path string) (*entities.VideoRes, *ResEntity.ResponseMsg)
}

type videoSrv struct {
	videoRepo.VideoRepository
}

func NewVideoSrv(videoRepository videoRepo.VideoRepository) videoSrv {
	return videoSrv{VideoRepository: videoRepository}
}

func (v videoSrv) UploadVideo(path string) (*entities.VideoRes, *ResEntity.ResponseMsg) {
	// check if file path is valid.
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Println(err)
		return nil,
			ResEntity.CustomErrorResponse("file not found", http.StatusNotFound)
	}

	// create context
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*90)
	defer cancelFunc()

	//send file path to repo to upload
	video, err := v.PersistVideo(ctx, path)
	if err != nil {
		log.Println(err)

		// TODO handle error if request times out

		return nil,
			ResEntity.CustomErrorResponse("error uploading video", http.StatusInternalServerError)
	}
	// organize data
	data := &entities.VideoRes{
		VideoURL:          video.Url,
		VideoPreviewURL:   video.PreviewUrl,
		VideoThumbnailURL: video.ThumbnailUrl,
	}

	// send finally
	return data, nil

}
