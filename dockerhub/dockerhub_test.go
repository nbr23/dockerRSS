package dockerhub

import "testing"

func TestParseDockerImage(t *testing.T) {

	testCases := []struct {
		name  string
		image DockerImageName
	}{
		{"alpine", DockerImageName{Org: "library", Image: "alpine", Tag: ""}},
		{"alpine:latest", DockerImageName{Org: "library", Image: "alpine", Tag: "latest"}},
		{"nbr23/dockerrss", DockerImageName{Org: "nbr23", Image: "dockerrss", Tag: ""}},
		{"nbr23/dockerrss:latest", DockerImageName{Org: "nbr23", Image: "dockerrss", Tag: "latest"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			image := ParseDockerImage(tc.name)
			if image != tc.image {
				t.Errorf("got %q, wanted %q", image, tc.image)
			}
		})
	}
}

func TestDockerImageString(t *testing.T) {

	testCases := []struct {
		name     string
		fullName string
	}{
		{"alpine", "library/alpine"},
		{"alpine:latest", "library/alpine"},
		{"nbr23/dockerrss", "nbr23/dockerrss"},
		{"nbr23/dockerrss:latest", "nbr23/dockerrss"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			image := ParseDockerImage(tc.name).String()
			if image != tc.fullName {
				t.Errorf("got %q, wanted %q", image, tc.fullName)
			}
		})
	}
}

func TestDockerImagePretty(t *testing.T) {

	testCases := []string{
		"alpine",
		"alpine:latest",
		"nbr23/dockerrss",
		"nbr23/dockerrss:latest",
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			image := ParseDockerImage(tc).Pretty()
			if image != tc {
				t.Errorf("got %q, wanted %q", image, tc)
			}
		})
	}
}
