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
	imageName := strings.TrimPrefix(r.URL.Path, "/tags/")
	image := dockerhub.ParseDockerImage(imageName)

	var tags []dockerhub.DockerhubTag
	var err error

	if image.Tag != "" {
		t, err := dockerhub.GetDockerImageTagDetails(image)
		if err != nil {
			http.Error(w, "tag not found", http.StatusNotFound)
			log.Printf("%s 404 tag not found: %s", r.URL.Path, err)
			return
		}
		tags = append(tags, t)
	} else {
		tags, err = dockerhub.GetDockerImageTags(image)
		if err != nil {
			http.Error(w, "no tags found", http.StatusNotFound)
			log.Printf("%s 404 tag not found: %s", r.URL.Path, err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/atom+xml")

	atomFeed := atom.GenerateAtomFeed(image, tags)
	log.Printf("%s 200", r.URL.Path)
	fmt.Fprint(w, atomFeed)
}

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	http.HandleFunc("/tags/", tagsHandler)

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
