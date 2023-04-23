package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/nbr23/dockerRSS/atom"
	"github.com/nbr23/dockerRSS/dockerhub"
)

type MyEvent struct {
	Name string `json:"name"`
}

func tagsHandler(w http.ResponseWriter, r *http.Request) {
	imageName := dockerhub.ParseDockerImage(strings.TrimPrefix(r.URL.Path, "/tags/"))

	var images []dockerhub.DockerhubImage
	var err error

	if imageName.Tag != "" {
		images, err = dockerhub.GetDockerTagImagesDetails(imageName)
		if err != nil {
			http.Error(w, "tag not found", http.StatusNotFound)
			log.Printf("%s 404 tag not found: %s", r.URL.Path, err)
			return
		}
	} else {
		images, err = dockerhub.GetDockerTagsImages(imageName)
		if err != nil {
			http.Error(w, "no tags found", http.StatusNotFound)
			log.Printf("%s 404 tag not found: %s", r.URL.Path, err)
			return
		}
	}

	platformStr := r.URL.Query().Get("platform")
	if platformStr != "" {
		var filteredImages []dockerhub.DockerhubImage
		for _, i := range images {
			if i.IsPlatform(dockerhub.ParsePlatform(platformStr)) {
				filteredImages = append(filteredImages, i)
			}
		}
		images = filteredImages
	}

	format := r.URL.Query().Get("format")
	if format == "plain" {
		w.Header().Set("Content-Type", "text/xml")
	} else {
		w.Header().Set("Content-Type", "application/atom+xml")
	}

	atomFeed := atom.GenerateAtomFeed(imageName, images)
	log.Printf("%s 200", r.URL.Path)
	fmt.Fprint(w, atomFeed)
}

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	http.HandleFunc("/tags/", tagsHandler)

	static_dir := os.Getenv("HTTP_STATIC_DIR")
	if static_dir == "" {
		static_dir = "static"
	}
	fs := http.FileServer(http.Dir(static_dir))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("invalid port: %s", port)
	}
	fmt.Println("Listening on port", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
