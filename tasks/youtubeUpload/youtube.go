// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	youtube "google.golang.org/api/youtube/v3"
)

func init() {
	registerDemo("youtube", youtube.YoutubeUploadScope, youtubeMain)
}

// Flags
var (
	title              = flag.String("title", "", "Youtube upload title")
	description        = flag.String("description", "", "Youtube upload description")
	categoryID         = flag.String("categoryId", "", "Youtube upload categoryId")
	userContributionID = flag.String("userContributionID", "", "Youtube upload userContributionID")
	videoStatus        = flag.String("videoStatus", "", "Youtube upload VideoStatus")
)

func youtubeMain(client *http.Client, argv []string) {
	id, _ := strconv.Atoi(*userContributionID)
	t := models.MovieTypeYoutube

	if len(argv) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: youtube filename")
		contributions.AddOrSaveMovie(id, "", t, models.StatusError)
		return
	}
	filename := argv[0]

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       *title,
			Description: *description, // can not use non-alpha-numeric characters
			CategoryId:  *categoryID,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: *videoStatus},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	upload.Snippet.Tags = []string{"test", "upload", "api"}

	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalf("Error opening %v: %v", filename, err)
		contributions.AddOrSaveMovie(id, "", t, models.StatusError)
		return
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
		contributions.AddOrSaveMovie(id, "", t, models.StatusError)
		return
	}

	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)

	contributions.AddOrSaveMovie(id, response.Id, t, models.StatusPublic)
}
