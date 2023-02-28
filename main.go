package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type dockerImageName struct {
	Org   string
	Image string
}

func (d dockerImageName) String() string {
	return fmt.Sprintf("%s/%s", d.Org, d.Image)
}

func (d dockerImageName) Pretty() string {
	if d.Org == "library" {
		return d.Image
	}
	return fmt.Sprintf("%s/%s", d.Org, d.Image)
}

type MyEvent struct {
	Name string `json:"name"`
}

func tagsHandler(w http.ResponseWriter, r *http.Request) {
	imageName := strings.TrimPrefix(r.URL.Path, "/tags/")
	image := parseDockerImage(imageName)

	tags, err := getDockerImageTags(image)
	if err != nil {
		http.Error(w, "no tags found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/atom+xml")

	atomFeed := generateAtomFeed(image, tags)
	fmt.Fprint(w, atomFeed)
}

func main() {
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