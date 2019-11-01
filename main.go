package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

/*
	words = {...}
	foreach word do:
		- get html page;
		- extract IPA;
		- download audiofile;
	save to csv;
*/

const baseurl = `https://www.oxfordlearnersdictionaries.com/us/definition/english/`

type EnglishWord struct {
	word, ipa, audioLink string
}

func main() {
	words := map[string]EnglishWord{
		"bad": {word:"bad"},
		"happen": {word:"happen"},
	}

	for key, item := range words {
		wordPage, err := GetLatestBlogTitles(baseurl + key)
		if err != nil {
			log.Println(err)
		}
		ipa := getIPA(wordPage)
		audioLink := getAudioLink(wordPage)
		words[key] = struct{ word, ipa, audioLink string }{word: item.word, ipa: ipa, audioLink: audioLink}
	}

	for _, item := range words {
		fmt.Println("word: ", item.word)
		fmt.Println("ipa: ", item.ipa)
		fmt.Println("audio: ", item.audioLink)
		fmt.Println("--")
	}
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestBlogTitles(url string) (*goquery.Document, error) {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, err
}

func getIPA(doc *goquery.Document) string {
	ipa := doc.Find("span.pron-g[geo='n_am']").First()
	return ipa.Text()
}

func getAudioLink(doc *goquery.Document) string {
	audio, _ := doc.Find("span.pron-g[geo='n_am'] div.pron-us").Attr("data-src-mp3")
	return audio
}

func downloadAudio() {

}
