package atom

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/nbr23/dockerRSS/dockerhub"
)

func dockerhubTagToAtomEntry(image dockerhub.DockerImageName, tag dockerhub.DockerhubTag) string {
	// We just use the first image in the manifest list, shouldn't matter much
	if len(tag.Digest) == 0 {
		tag.Digest = tag.Images[0].Digest
	}
	digest := strings.Replace(tag.Digest, ":", "-", -1)

	guid := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(image.Pretty(), tag.Name, tag.LastUpdated))))
	return fmt.Sprintf(`
	<entry>
		<title>%s - %s</title>
		<guid>%s</guid>
		<updated>%s</updated>
		<link href="https://hub.docker.com/layers/%s/%s/images/%s" />
	</entry>
	`, image.Pretty(), tag.Name, guid, tag.LastUpdated, image, tag.Name, digest)
}

func GenerateAtomFeed(image dockerhub.DockerImageName, tags []dockerhub.DockerhubTag) string {
	var entries []string
	var lastPushed time.Time
	for _, tag := range tags {
		entries = append(entries, dockerhubTagToAtomEntry(image, tag))
		tagPushed, err := time.Parse("2006-01-02T15:04:05.999999Z", tag.TagLastPushed)
		if err != nil {
			fmt.Println("Error while parsing date :", err)
			continue
		}
		if lastPushed.Before(tagPushed) {
			lastPushed = tagPushed
		}
	}

	return fmt.Sprintf(`
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>%s</title>
	<id>%s</id>
	<updated>%s</updated>
	%s
</feed>
	`, image.Pretty(), image.Pretty(), lastPushed, strings.Join(entries, ""))
}
