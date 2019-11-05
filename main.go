package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"encoding/csv"
	"io"
	"strings"
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
		"actor": {word: "actor"},
		"black": {word: "black"},
		"clay": {word: "clay"},
		"disease": {word: "disease"},
		"adjective": {word: "adjective"},
		"blind": {word: "blind"},
		"clean": {word: "clean"},
		"doctor": {word: "doctor"},
		"adult": {word: "adult"},
		"blood": {word: "blood"},
		"dog": {word: "dog"},
		"afternoon": {word: "afternoon"},
		"blue": {word: "blue"},
		"clock": {word: "clock"},
		"dollar": {word: "dollar"},
		"air": {word: "air"},
		"boat": {word: "boat"},
		"close": {word: "close"},
		"door": {word: "door"},
		"airport": {word: "airport"},
		"body": {word: "body"},
		"clothing": {word: "clothing"},
		"dot": {word: "dot"},
		"alive": {word: "alive"},
		"bone": {word: "bone"},
		"club": {word: "club"},
		"down": {word: "down"},
		"animal": {word: "animal"},
		"book": {word: "book"},
		"coat": {word: "coat"},
		"draw": {word: "draw"},
		"apartment": {word: "apartment"},
		"bottle": {word: "bottle"},
		"coffee": {word: "coffee"},
		"dream": {word: "dream"},
		"apple": {word: "apple"},
		"bottom": {word: "bottom"},
		"cold": {word: "cold"},
		"dress": {word: "dress"},
		"April": {word: "April"},
		"box": {word: "box"},
		"color": {word: "color"},
		"drink": {word: "drink"},
		"arm": {word: "arm"},
		"boy": {word: "boy"},
		"computer": {word: "computer"},
		"drive": {word: "drive"},
		"army": {word: "army"},
		"brain": {word: "brain"},
		"consonant": {word: "consonant"},
		"drug": {word: "drug"},
		"art": {word: "art"},
		"bread": {word: "bread"},
		"contract": {word: "contract"},
		"dry": {word: "dry"},
		"artist": {word: "artist"},
		"break": {word: "break"},
		"cook": {word: "cook"},
		"dust": {word: "dust"},
		"attack": {word: "attack"},
		"breakfast": {word: "breakfast"},
		"cool": {word: "cool"},
		"ear": {word: "ear"},
		"August": {word: "August"},
		"bridge": {word: "bridge"},
		"copper": {word: "copper"},
		"Earth": {word: "Earth"},
		"author": {word: "author"},
		"brother": {word: "brother"},
		"corn": {word: "corn"},
		"east": {word: "east"},
		"baby": {word: "baby"},
		"brown": {word: "brown"},
		"corner": {word: "corner"},
		"eat": {word: "eat"},
		"back": {word: "back"},
		"build": {word: "build"},
		"count": {word: "count"},
		"edge": {word: "edge"},
		"building": {word: "building"},
		"country": {word: "country"},
		"egg": {word: "egg"},
		"bad": {word: "bad"},
		"burn": {word: "burn"},
		"court": {word: "court"},
		"eight": {word: "eight"},
		"bag": {word: "bag"},
		"bus": {word: "bus"},
		"cow": {word: "cow"},
		"eighteen": {word: "eighteen"},
		"ball": {word: "ball"},
		"buy": {word: "buy"},
		"crowd": {word: "crowd"},
		"eighty": {word: "eighty"},
		"banana": {word: "banana"},
		"cake": {word: "cake"},
		"cry": {word: "cry"},
		"election": {word: "election"},
		"band": {word: "band"},
		"call": {word: "call"},
		"cup": {word: "cup"},
		"electronics": {word: "electronics"},
		"bank": {word: "bank"},
		"camera": {word: "camera"},
		"curved": {word: "curved"},
		"eleven": {word: "eleven"},
		"bar": {word: "bar"},
		"camp": {word: "camp"},
		"cut": {word: "cut"},
		"energy": {word: "energy"},
		"bathroom": {word: "bathroom"},
		"car": {word: "car"},
		"dance": {word: "dance"},
		"engine": {word: "engine"},
		"beach": {word: "beach"},
		"card": {word: "card"},
		"dark": {word: "dark"},
		"evening": {word: "evening"},
		"beard": {word: "beard"},
		"carry": {word: "carry"},
		"date": {word: "date"},
		"exercise": {word: "exercise"},
		"beat": {word: "beat"},
		"cat": {word: "cat"},
		"daughter": {word: "daughter"},
		"expensive": {word: "expensive"},
		"beautiful": {word: "beautiful"},
		"catch": {word: "catch"},
		"day": {word: "day"},
		"explode": {word: "explode"},
		"bed": {word: "bed"},
		"ceiling": {word: "ceiling"},
		"dead": {word: "dead"},
		"eye": {word: "eye"},
		"bedroom": {word: "bedroom"},
		"cell": {word: "cell"},
		"phone": {word: "phone"},
		"deaf": {word: "deaf"},
		"face": {word: "face"},
		"beef": {word: "beef"},
		"centimeter": {word: "centimeter"},
		"death": {word: "death"},
		"beer": {word: "beer"},
		"chair": {word: "chair"},
		"December": {word: "December"},
		"fall": {word: "fall"},
		"bend": {word: "bend"},
		"cheap": {word: "cheap"},
		"deep": {word: "deep"},
		"family": {word: "family"},
		"beverage": {word: "beverage"},
		"cheese": {word: "cheese"},
		"diamond": {word: "diamond"},
		"famous": {word: "famous"},
		"bicycle": {word: "bicycle"},
		"chicken": {word: "chicken"},
		"die": {word: "die"},
		"big": {word: "big"},
		"large": {word: "large"},
		"child": {word: "child"},
		"dig": {word: "dig"},
		"fan": {word: "fan"},
		"bill": {word: "bill"},
		"church": {word: "church"},
		"dinner": {word: "dinner"},
		"farm": {word: "farm"},
		"billion": {word: "billion"},
		"circle": {word: "circle"},
		"direction": {word: "direction"},
		"fast": {word: "fast"},
		"bird": {word: "bird"},
		"city": {word: "city"},
		"dirty": {word: "dirty"},
		"father": {word: "father"},
		"February": {word: "February"},
		"hair": {word: "hair"},
		"kitchen": {word: "kitchen"},
		"million": {word: "million"},
		"feed": {word: "feed"},
		"half": {word: "half"},
		"knee": {word: "knee"},
		"minute": {word: "minute"},
		"female": {word: "female"},
		"hand": {word: "hand"},
		"knife": {word: "knife"},
		"mix": {word: "mix"},
		"stir": {word: "stir"},
		"fifteen": {word: "fifteen"},
		"hang": {word: "hang"},
		"lake": {word: "lake"},
		"Monday": {word: "Monday"},
		"fifth": {word: "fifth"},
		"happy": {word: "happy"},
		"lamp": {word: "lamp"},
		"money": {word: "money"},
		"fifty": {word: "fifty"},
		"hard": {word: "hard"},
		"laptop": {word: "laptop"},
		"month": {word: "month"},
		"fight": {word: "fight"},
		"hat": {word: "hat"},
		"laugh": {word: "laugh"},
		"moon": {word: "moon"},
		"find": {word: "find"},
		"he": {word: "he"},
		"lawyer": {word: "lawyer"},
		"morning": {word: "morning"},
		"finger": {word: "finger"},
		"head": {word: "head"},
		"leaf": {word: "leaf"},
		"mother": {word: "mother"},
		"fire": {word: "fire"},
		"healthy": {word: "healthy"},
		"learn": {word: "learn"},
		"mountain": {word: "mountain"},
		"first": {word: "first"},
		"hear": {word: "hear"},
		"left": {word: "left"},
		"mouse": {word: "mouse"},
		"fish": {word: "fish"},
		"heart": {word: "heart"},
		"leg": {word: "leg"},
		"mouth": {word: "mouth"},
		"five": {word: "five"},
		"heat": {word: "heat"},
		"lemon": {word: "lemon"},
		"movie": {word: "movie"},
		"flat": {word: "flat"},
		"heaven": {word: "heaven"},
		"letter": {word: "letter"},
		"murder": {word: "murder"},
		"floor": {word: "floor"},
		"heavy": {word: "heavy"},
		"library": {word: "library"},
		"music": {word: "music"},
		"flower": {word: "flower"},
		"hell": {word: "hell"},
		"lie": {word: "lie"},
		"narrow": {word: "narrow"},
		"fly": {word: "fly"},
		"high": {word: "high"},
		"lift": {word: "lift"},
		"nature": {word: "nature"},
		"follow": {word: "follow"},
		"hill": {word: "hill"},
		"neck": {word: "neck"},
		"food": {word: "food"},
		"hole": {word: "hole"},
		"needle": {word: "needle"},
		"foot": {word: "foot"},
		"horse": {word: "horse"},
		"light": {word: "light"},
		"neighbor": {word: "neighbor"},
		"hospital": {word: "hospital"},
		"lip": {word: "lip"},
		"network": {word: "network"},
		"forest": {word: "forest"},
		"hot": {word: "hot"},
		"listen": {word: "listen"},
		"new": {word: "new"},
		"fork": {word: "fork"},
		"hotel": {word: "hotel"},
		"location": {word: "location"},
		"newspaper": {word: "newspaper"},
		"forty": {word: "forty"},
		"hour": {word: "hour"},
		"lock": {word: "lock"},
		"nice": {word: "nice"},
		"four": {word: "four"},
		"house": {word: "house"},
		"long": {word: "long"},
		"night": {word: "night"},
		"fourteen": {word: "fourteen"},
		"human": {word: "human"},
		"loose": {word: "loose"},
		"nine": {word: "nine"},
		"fourth": {word: "fourth"},
		"hundred": {word: "hundred"},
		"lose": {word: "lose"},
		"nineteen": {word: "nineteen"},
		"Friday": {word: "Friday"},
		"husband": {word: "husband"},
		"loud": {word: "loud"},
		"ninety": {word: "ninety"},
		"friend": {word: "friend"},
		"I": {word: "I"},
		"love": {word: "love"},
		"no": {word: "no"},
		"front": {word: "front"},
		"ice": {word: "ice"},
		"low": {word: "low"},
		"north": {word: "north"},
		"game": {word: "game"},
		"image": {word: "image"},
		"lunch": {word: "lunch"},
		"nose": {word: "nose"},
		"garden": {word: "garden"},
		"inch": {word: "inch"},
		"magazine": {word: "magazine"},
		"note": {word: "note"},
		"gasoline": {word: "gasoline"},
		"injury": {word: "injury"},
		"male": {word: "male"},
		"November": {word: "November"},
		"gift": {word: "gift"},
		"inside": {word: "inside"},
		"man": {word: "man"},
		"nuclear": {word: "nuclear"},
		"girl": {word: "girl"},
		"instrument": {word: "instrument"},
		"musical": {word: "musical"},
		"manager": {word: "manager"},
		"number": {word: "number"},
		"glass": {word: "glass"},
		"island": {word: "island"},
		"map": {word: "map"},
		"ocean": {word: "ocean"},
		"go": {word: "go"},
		"it": {word: "it"},
		"March": {word: "March"},
		"October": {word: "October"},
		"God": {word: "God"},
		"January": {word: "January"},
		"market": {word: "market"},
		"office": {word: "office"},
		"gold": {word: "gold"},
		"job": {word: "job"},
		"marriage": {word: "marriage"},
		"oil": {word: "oil"},
		"good": {word: "good"},
		"juice": {word: "juice"},
		"marry": {word: "marry"},
		"old": {word: "old"},
		"grandfather": {word: "grandfather"},
		"July": {word: "July"},
		"material": {word: "material"},
		"young": {word: "young"},
		"grandmother": {word: "grandmother"},
		"jump": {word: "jump"},
		"May": {word: "May"},
		"one": {word: "one"},
		"grass": {word: "grass"},
		"June": {word: "June"},
		"mean": {word: "mean"},
		"open": {word: "open"},
		"gray": {word: "gray"},
		"key": {word: "key"},
		"medicine": {word: "medicine"},
		"green": {word: "green"},
		"kill": {word: "kill"},
		"melt": {word: "melt"},
		"orange": {word: "orange"},
		"ground": {word: "ground"},
		"kilogram": {word: "kilogram"},
		"metal": {word: "metal"},
		"outside": {word: "outside"},
		"grow": {word: "grow"},
		"king": {word: "king"},
		"meter": {word: "meter"},
		"page": {word: "page"},
		"gun": {word: "gun"},
		"kiss": {word: "kiss"},
		"milk": {word: "milk"},
		"pain": {word: "pain"},
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
	writeToFile(words)
}

func writeToFile(data map[string]EnglishWord) {
	file, err := os.Create("words.csv")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		line := []string{value.word, value.ipa, value.audioLink}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
		if value.audioLink != "" {
			errDl := downloadAudio("audio/"+value.word+".mp3", value.audioLink)
			if errDl != nil {
				panic(errDl)
			}
		}
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
	ipa := doc.Find("span.pron-g[geo='n_am'] span.phon").First()
	text := ipa.Text()
	text = strings.Replace(text, "NAmE", "", -1)
	text = strings.Replace(text, "//", "/", -1)
	return text
}

func getAudioLink(doc *goquery.Document) string {
	audio, _ := doc.Find("span.pron-g[geo='n_am'] div.pron-us").Attr("data-src-mp3")
	return audio
}

func downloadAudio(filepath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, errCopy := io.Copy(out, resp.Body)
	return errCopy
}
