package videoRepo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"test1/internal/entities/cloudinary"
)

type cloudinaryObj struct {
	ApiKey       string
	CloudName    string
	UploadPreset string
	BaseURL      string
}

func CreateCloudinaryObject(apiKey string, uploadPreset string, baseURL, cloudName string) cloudinaryObj {
	return cloudinaryObj{ApiKey: apiKey, UploadPreset: uploadPreset, BaseURL: baseURL, CloudName: cloudName}
}

func (c cloudinaryObj) PersistVideo(ctx context.Context, path string) (cloudinary.Response, error) {
	method := "POST"
	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	file, errFile1 := os.Open(path)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return cloudinary.Response{}, errFile1
	}

	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("file",
		filepath.Base(path))
	if errFile1 != nil {
		fmt.Println(errFile1)
		return cloudinary.Response{}, errFile1
	}

	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return cloudinary.Response{}, errFile1
	}

	_ = writer.WriteField("upload_preset", c.UploadPreset)
	_ = writer.WriteField("api_key", c.ApiKey)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return cloudinary.Response{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL, payload)

	if err != nil {
		fmt.Println(err)
		return cloudinary.Response{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return cloudinary.Response{}, err
	}
	defer res.Body.Close()

	var cloudinaryRes cloudinary.Response

	err = json.NewDecoder(res.Body).Decode(&cloudinaryRes)
	if err != nil {
		return cloudinary.Response{}, err
	}

	// set preview imaage and preview vidoe fields
	cloudinaryRes.PreviewUrl = c.generatePreview(cloudinaryRes)
	cloudinaryRes.ThumbnailUrl = c.generateThumbNail(cloudinaryRes)

	return cloudinaryRes, nil

}

func (c cloudinaryObj) generatePreview(cl cloudinary.Response) string {
	str := fmt.Sprintf(`https://res.cloudinary.com/%s/video/upload/e_preview/v%d/%s`,
		c.CloudName, cl.Version, cl.PublicId)
	return str
}
func (c cloudinaryObj) generateThumbNail(cl cloudinary.Response) string {
	str := fmt.Sprintf(`https://res.cloudinary.com/%s/video/upload/e_preview/v%d/%s.jpg`,
		c.CloudName, cl.Version, cl.PublicId)
	return str
}

//func SendVideo() {
//
//	url := "https://api.cloudinary.com/v1_1/dfifbma3s/video/upload"
//	method := "POST"
//
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//	file, errFile1 := os.Open("./files/gigaz.png")
//	defer file.Close()
//	part1,
//		errFile1 := writer.CreateFormFile("file", filepath.Base("/Users/oreoluwa/Downloads/4ca1ecfe-34f6-45ac-a385-4d6bd0e80544.MP4"))
//	_, errFile1 = io.Copy(part1, file)
//	if errFile1 != nil {
//		fmt.Println(errFile1)
//		return
//	}
//	_ = writer.WriteField("upload_preset", "oexwzbcj")
//	_ = writer.WriteField("api_key", "289463472338254")
//	err := writer.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, payload)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//}
