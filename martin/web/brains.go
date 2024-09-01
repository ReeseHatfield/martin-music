package web

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ReeseHatfield/ffmpeg"
)

const (
	baseURL = `https://musicbrainz.org/ws/2/release`
)

type Metadata struct {
	XMLName     xml.Name  `xml:"metadata"`
	ReleaseList []Release `xml:"release-list>release"`
}

type Release struct {
	ID           string       `xml:"id,attr"`
	Title        string       `xml:"title"`
	Status       string       `xml:"status"`
	ArtistCredit []NameCredit `xml:"artist-credit>name-credit"`
}

type NameCredit struct {
	Name   string `xml:"name"`
	Artist Artist `xml:"artist"`
}

type Artist struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name"`
	SortName string `xml:"sort-name"`
}

var ErrCoverArtNotFound = errors.New("Could not get cover art from ID")
var ErrNoReleasesFound = errors.New("No releases could be found in release group")

func GetCover(albumTitle, artistName string) (*ffmpeg.Image, error) {

	params := url.Values{}
	// release+group seems to get more accurate results than just `release`
	params.Add("query", fmt.Sprintf(`release+group:"%s" AND artist:"%s"`, albumTitle, artistName))
	params.Add("fmt", "xml") // Ensuring response is in XML format
	getMe := fmt.Sprintf("%s/?%s", baseURL, params.Encode())

	fmt.Printf("Query URL: %v\n", getMe)
	res, err := http.Get(getMe)
	if err != nil {
		log.Println("Error fetching release info:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %s", res.Status)
	}

	xmlData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var metadata Metadata
	err = xml.Unmarshal(xmlData, &metadata)
	if err != nil {
		log.Println("Error unmarshalling XML:", err)
		return nil, err
	}

	if len(metadata.ReleaseList) < 1 {
		log.Println("No entries found in release list")
		return nil, ErrNoReleasesFound
	}

	// use an official release if one exists
	var topRelease *Release
	for _, release := range metadata.ReleaseList {
		if release.Status == "Official" {
			topRelease = &release
			break
		}
	}

	// oops no top results, use fallback. Common with more obscure records
	if topRelease == nil {
		topRelease = &metadata.ReleaseList[0]
	}

	fmt.Printf("Top Release ID: %s, Title: %s\n", topRelease.ID, topRelease.Title)
	coverURL, err := fetchArtFromID(topRelease.ID)
	if err != nil {
		log.Println("Err fetching cover art:", err)
		return nil, err
	}

	fmt.Println("Cover Art URL:", coverURL)

	return downloadImageFromURL(coverURL)
}

func fetchArtFromID(releaseID string) (string, error) {
	coverArtURL := fmt.Sprintf("https://coverartarchive.org/release/%s/front", releaseID)

	resp, err := http.Get(coverArtURL)
	if err != nil {
		log.Println("Failed to fetch cover art:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrCoverArtNotFound
	}

	return coverArtURL, nil
}

func downloadImageFromURL(imgURL string) (*ffmpeg.Image, error) {

	res, err := http.Get(imgURL)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	fmt.Println("img url: " + imgURL)
	fmt.Println("img hash: " + hashURL(imgURL))

	path := fmt.Sprintf("../temp/%s.jpg", hashURL(imgURL))

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	img := ffmpeg.Image{
		Path: absPath,
	}

	return &img, nil
}

func hashURL(url string) string {
	hash := sha256.New()

	hash.Write([]byte(url))

	hashBytes := hash.Sum(nil)
	// will return unsupported character.
	// would probably be fine, but harder to work with without the encoding
	hashString := base64.URLEncoding.EncodeToString(hashBytes)

	return hashString
}
