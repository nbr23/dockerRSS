package atom

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/nbr23/dockerRSS/dockerhub"
)

func dockerhubTagToAtomEntry(image dockerhub.DockerhubImage) string {
	digest := strings.Replace(image.Digest, ":", "-", -1)

	guid := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(image.FullName.Pretty(), image.FullName, image.Os, image.Architecture, image.LastPushed))))
	return fmt.Sprintf(`
	<entry>
		<title>%s - %s</title>
		<guid>%s</guid>
		<updated>%s</updated>
		<link href="%s" />
		<content type="html"><![CDATA[Image %s@%s was pushed with tag %s.]]></content>
	</entry>
	`, image.FullName.Pretty(), image.Platform(), guid, image.LastPushed, image.FullName.GetImageURL(digest), image.FullName.RepoName(), image.Digest, image.FullName.Tag)

}

func GenerateAtomFeed(imageName dockerhub.DockerImageName, images []dockerhub.DockerhubImage) string {
	var entries []string
	var lastPushed time.Time
	for _, i := range images {
		entries = append(entries, dockerhubTagToAtomEntry(i))
		imagePushed, err := time.Parse("2006-01-02T15:04:05.999999Z", i.LastPushed)
		if err != nil {
			fmt.Println("Error while parsing date :", err)
			continue
		}
		if lastPushed.Before(imagePushed) {
			lastPushed = imagePushed
		}
	}

	return fmt.Sprintf(`
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>%s</title>
	<id>%s</id>
	<updated>%s</updated>
	%s
</feed>
	`, imageName.Pretty(), imageName.Pretty(), lastPushed, strings.Join(entries, ""))
}
