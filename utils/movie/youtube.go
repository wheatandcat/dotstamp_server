package movie

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"dotstamp_server/utils"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	demoFunc  = make(map[string]func(*http.Client, []string))
	demoScope = make(map[string]string)
)

// GetConnectURL URLを取得する
func GetConnectURL() string {
	if utils.IsTest() {
		return ""
	}

	config := &oauth2.Config{
		ClientID:     beego.AppConfig.String("youtubeClientID"),
		ClientSecret: beego.AppConfig.String("youtubeClientSecret"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.upload"},
		RedirectURL:  beego.AppConfig.String("topurl") + "movie/youtube",
	}

	return config.AuthCodeURL("st001")
}

// GetConnect 接続を取得する
func GetConnect() (h *http.Client) {
	if utils.IsTest() {
		return h
	}

	config := &oauth2.Config{
		ClientID:     beego.AppConfig.String("youtubeClientID"),
		ClientSecret: beego.AppConfig.String("youtubeClientSecret"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.upload"},
	}

	ctx := context.Background()

	return newOAuthClient(ctx, config)
}

func saveToken(file string, token *oauth2.Token) {
	f, err := os.Create(file)
	if err != nil {
		log.Printf("Warning: failed to cache oauth token: %v", err)
		return
	}

	defer f.Close()
	gob.NewEncoder(f).Encode(token)
}

func newOAuthClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile := tokenCacheFile(config)
	token, err := tokenFromFile(cacheFile, false)
	if err != nil {
		token = tokenFromWeb(ctx, config)
		saveToken(cacheFile, token)
	} else {
		log.Printf("Using cached token %#v from %q", token, cacheFile)
	}

	return config.Client(ctx, token)
}

func osUserCacheDir() string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Caches")
	case "linux", "freebsd":
		return filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return "."
}

func tokenCacheFile(config *oauth2.Config) string {
	hash := fnv.New32a()
	hash.Write([]byte(config.ClientID))
	hash.Write([]byte(config.ClientSecret))
	hash.Write([]byte(strings.Join(config.Scopes, " ")))
	fn := fmt.Sprintf("go-api-demo-tok%v", hash.Sum32())

	return filepath.Join(osUserCacheDir(), url.QueryEscape(fn))
}

func tokenFromFile(file string, cache bool) (*oauth2.Token, error) {
	if !cache {
		return nil, errors.New("--cachetoken is false")
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	t := new(oauth2.Token)
	err = gob.NewDecoder(f).Decode(t)
	return t, err
}

func tokenFromWeb(ctx context.Context, config *oauth2.Config) (token *oauth2.Token) {
	var err error
	ch := make(chan string)
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			log.Printf("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "<h1>認証しました</h1>閉じる")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}

		log.Printf("no code")
		http.Error(rw, "", 500)
	}))

	defer ts.Close()

	config.RedirectURL = ts.URL

	authURL := config.AuthCodeURL(randState)

	go openURL(authURL)
	code := <-ch

	log.Printf("Got code: %s", code)

	token, err = config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Token exchange error: %v", err)
	}

	return token
}

func openURL(url string) {
	try := []string{"xdg-open", "google-chrome", "open"}
	for _, bin := range try {
		err := exec.Command(bin, url).Run()
		if err == nil {
			return
		}
	}
	log.Printf("Error opening URL in browser.")
}

func valueOrFileContents(value string, filename string) string {
	if value != "" {
		return value
	}
	slurp, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading %q: %v", filename, err)
	}
	return strings.TrimSpace(string(slurp))
}
