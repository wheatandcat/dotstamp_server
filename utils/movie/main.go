package movie

import (
	"dotstamp_server/utils"
	"net/http"
	"os"
	"os/exec"

	youtube "google.golang.org/api/youtube/v3"
)

// Upload アップロード
type Upload struct {
	UserContributionID string
	Title              string
	Description        string
	CategoryID         string
	VideoStatus        string
}

// getRootPath パスを取得する
func getRootPath() (string, error) {
	p, err := utils.GetAppPath()
	if err != nil {
		return "", err
	}

	return p + "/", nil
}

// Make 作成する
func Make(file string) error {
	path, err := getRootPath()
	if err != nil {
		return err
	}

	movie := path + "static/movie/input.mp4"
	sound := path + "static/files/tmp/sound/" + file + ".m4a"
	dist := path + "static/files/movie/" + file + ".mp4"

	cmd := "ffmpeg -y -i " + movie + " -i " + sound + " -map 0:0 -map 1:0 -movflags faststart -vcodec libx264 -acodec copy " + dist

	_, err = exec.Command("sh", "-c", cmd).Output()

	return err
}

// UploadToYoutube YouTubeにアップロードする
func UploadToYoutube(client *http.Client, u Upload) (string, error) {
	if utils.IsTest() {
		return "", nil
	}

	path, err := getRootPath()
	if err != nil {
		return "", err
	}

	filename := path + "static/files/movie/" + u.UserContributionID + ".mp4"

	service, err := youtube.New(client)
	if err != nil {
		return "", err
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       u.Title,
			Description: u.Description,
			CategoryId:  u.CategoryID,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: u.VideoStatus},
	}

	upload.Snippet.Tags = []string{"test", "upload", "api"}

	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return "", err
	}

	response, err := call.Media(file).Do()
	if err != nil {
		return "", err
	}

	return response.Id, nil
}
