package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type dockerhubTag struct {
	Name          string           `json:"name"`
	LastUpdated   string           `json:"last_updated"`
	TagLastPushed string           `json:"tag_last_pushed"`
	Digest        string           `json:"digest"`
	Images        []dockerhubImage `json:"images"`
}

type dockerhubResponse struct {
	Count int `json:"count"`
	// Next     string         `json:"next"`
	// Previous string         `json:"previous"`
	Results []dockerhubTag `json:"results"`
}

type dockerhubImage struct {
	Digest       string `json:"digest"`
	Architecture string `json:"architecture"`
	Os           string `json:"os"`
	LastPushed   string `json:"last_pushed"`
}

func parseDockerImage(imageName string) dockerImageName {
	var org, image, tag string

	if strings.Contains(imageName, ":") {
		split := strings.Split(imageName, ":")
		imageName = split[0]
		tag = split[1]
	}

	// default org is "library"
	if strings.Contains(imageName, "/") {
		split := strings.Split(imageName, "/")
		org = split[0]
		image = split[1]
	} else {
		org = "library"
		image = imageName
	}

	return dockerImageName{
		Org:   org,
		Image: image,
		Tag:   tag,
	}
}

func getDockerImageTagDetails(image dockerImageName) (dockerhubTag, error) {
	res, err := http.Get(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/%s", image, image.Tag))
	if err != nil {
		return dockerhubTag{}, err
	}
	defer res.Body.Close()

	var dResponse dockerhubTag
	err = json.NewDecoder(res.Body).Decode(&dResponse)
	if err != nil {
		return dockerhubTag{}, err
	}

	return dResponse, nil
}

func getDockerImageTags(image dockerImageName) ([]dockerhubTag, error) {
	res, err := http.Get(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/?page_size=25&page=1&ordering=last_updated", image))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var dResponse dockerhubResponse
	err = json.NewDecoder(res.Body).Decode(&dResponse)
	if err != nil {
		return nil, err
	}

	if dResponse.Count == 0 || dResponse.Results == nil {
		return nil, fmt.Errorf("no tags found")
	}

	return dResponse.Results, nil
}
