package dockerhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DockerImageName struct {
	Org   string
	Image string
	Tag   string
}

func (d DockerImageName) String() string {
	if d.Tag == "" {
		return fmt.Sprintf("%s/%s", d.Org, d.Image)
	}
	return fmt.Sprintf("%s/%s:%s", d.Org, d.Image, d.Tag)
}

func (d DockerImageName) Pretty() string {
	image := d.Image
	if d.Tag != "" {
		image = fmt.Sprintf("%s:%s", image, d.Tag)
	}
	if d.Org == "library" {
		return image
	}
	return fmt.Sprintf("%s/%s", d.Org, image)
}

type DockerhubTag struct {
	Name          string           `json:"name"`
	LastUpdated   string           `json:"last_updated"`
	TagLastPushed string           `json:"tag_last_pushed"`
	Digest        string           `json:"digest"`
	Images        []DockerhubImage `json:"images"`
}

type dockerhubResponse struct {
	Count int `json:"count"`
	// Next     string         `json:"next"`
	// Previous string         `json:"previous"`
	Results []DockerhubTag `json:"results"`
}

type DockerhubImage struct {
	Digest       string `json:"digest"`
	Architecture string `json:"architecture"`
	Os           string `json:"os"`
	LastPushed   string `json:"last_pushed"`
}

func ParseDockerImage(imageName string) DockerImageName {
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

	return DockerImageName{
		Org:   org,
		Image: image,
		Tag:   tag,
	}
}

func GetDockerImageTagDetails(image DockerImageName) (DockerhubTag, error) {
	res, err := http.Get(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/%s", image, image.Tag))
	if err != nil {
		return DockerhubTag{}, err
	}
	defer res.Body.Close()

	var dResponse DockerhubTag
	err = json.NewDecoder(res.Body).Decode(&dResponse)
	if err != nil {
		return DockerhubTag{}, err
	}

	return dResponse, nil
}

func GetDockerImageTags(image DockerImageName) ([]DockerhubTag, error) {
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
