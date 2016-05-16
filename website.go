package website

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/endeveit/guesslanguage"
	"github.com/golibri/fetch"
	"regexp"
	"strings"
)

// Website contains all relevant metadata from a HTML web page
type Website struct {
	URL         string
	Body        string
	Language    string
	Title       string
	Description string
	Image       string
	Favicon     string
	Feeds       []string
	Tags        []string
}

// Parse executes a given HTML string and transforms it into a Website{} struct
func Parse(s string) (Website, error) {
	doc := docFromString(s)
	w := Website{Body: s}
	w.Title = titleFromDoc(&doc)
	w.Description = descriptionFromDoc(&doc)
	w.Image = imageFromDoc(&doc)
	w.Favicon = faviconFromDoc(&doc)
	w.Feeds = feedsFromDoc(&doc, s)
	w.Tags = tagsFromDoc(&doc)
	w.URL = canonicalFromDoc(&doc)
	w.Language = detectLanguage(doc.Find("body").Text())
	return w, nil
}

// FromURL parses a Website directly from a given URL
func FromURL(URL string) (Website, error) {
	page, err := fetch.Get(URL)
	if err != nil {
		return Website{}, err
	}
	w, err := Parse(page.Body)
	if err != nil {
		return Website{}, err
	}
	return w, nil
}

func detectLanguage(str string) string {
	lang, err := guesslanguage.Guess(str)
	if err != nil {
		lang = "en"
	}
	return lang
}

func docFromString(str string) goquery.Document {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(str)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return goquery.Document{}
	}
	return *doc
}

func titleFromDoc(doc *goquery.Document) string {
	title := doc.Find("title,h1,h2,h3,h4,p,div,body").First().Text()
	rx := regexp.MustCompile(`\s[|-].{1}.+$`)
	if rx.MatchString(title) {
		i := rx.FindStringIndex(title)[0]
		title = title[:i]
	}
	return title
}

func descriptionFromDoc(doc *goquery.Document) string {
	sel := "meta[property='og:description']"
	sel += ", meta[name='twitter:description']"
	sel += ", meta[name='description']"
	desc, _ := doc.Find(sel).First().Attr("content")
	return desc
}

func imageFromDoc(doc *goquery.Document) string {
	sel := "meta[property='og:image']"
	sel += ", meta[name='twitter:image']"
	img, _ := doc.Find(sel).First().Attr("content")
	return img
}

func faviconFromDoc(doc *goquery.Document) string {
	sel := "link[rel='apple-touch-icon']"
	sel += ", link[rel='apple-touch-icon-precomposed']"
	sel += ", link[rel='shortcut icon']"
	sel += ", link[rel='icon']"
	fav, ok := doc.Find(sel).First().Attr("href")
	if !ok {
		sel = "meta[name='msapplication-TileImage']"
		ms, _ := doc.Find(sel).First().Attr("content")
		fav = ms
	}
	return fav
}

func feedsFromDoc(doc *goquery.Document, text string) []string {
	sel := "link[type='application/rss+xml']"
	sel += ", link[type='application/atom+xml']"
	matches := doc.Find(sel)

	if matches.Length() > 0 {
		feeds := make([]string, matches.Length())
		matches.Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			feeds[i] = url
		})
		return feeds
	}

	rx := regexp.MustCompile(`href=['"]([^'"]*(rss|atom|feed|xml)[^'"]*)['"]`)
	if rx.FindString(text) != "" {
		matches := rx.FindAllStringSubmatch(text, -1)
		feeds := make([]string, len(matches))
		for i, e := range matches {
			feeds[i] = e[1]
		}
		return feeds
	}

	return make([]string, 0)
}

func tagsFromDoc(doc *goquery.Document) []string {
	sel := "meta[name=keywords]"
	str, ok := doc.Find(sel).First().Attr("content")
	if !ok {
		return []string{}
	}
	str = strings.Replace(str, "|", ";", -1)
	str = strings.Replace(str, ",", ";", -1)
	if strings.Contains(str, ";") {
		list := strings.Split(str, ";")
		tags := make([]string, len(list))
		for i, e := range list {
			tags[i] = strings.Trim(e, " ")
		}
		return tags
	}
	tags := []string{str}
	return tags
}

func canonicalFromDoc(doc *goquery.Document) string {
	str, ok := doc.Find("link[rel=canonical]").First().Attr("href")
	if !ok {
		return ""
	}
	if len(str) > 3 {
		return str
	}
	return ""
}
