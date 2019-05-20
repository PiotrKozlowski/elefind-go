package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide path to the book.")
	}

	path := os.Args[1]

	if !strings.Contains(path, ".xhtml") {
		log.Fatal("Filename have to have xhtml extension.")
	}

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't open file at %v", path)
	}

	c := strings.NewReader(string(bs))
	doc, err := goquery.NewDocumentFromReader(c)
	if err != nil {
		log.Fatal("Error while parsing xhtml from file.")
	}

	html, err := doc.Find("body").Html()
	if err != nil {
		log.Fatal("Couldn't find element body.")
	}
	doc.Find("body").ReplaceWithHtml(withoutNamespaces(html))

	docHTML, err := doc.Html()
	if err != nil {
		log.Fatal("Couldn't remove namespaces from tag names.")
	}

	wrErr := ioutil.WriteFile(path, []byte(docHTML), 0644)
	if wrErr != nil {
		log.Fatalf("Couldn't update file at path %v", path)
	}
}

func withoutNamespaces(html string) string {
	// Replace <m:math></m:math> to <math></math>
	var re = regexp.MustCompile(`([a-zA-Z]):([a-zA-Z]+)`)
	s := re.ReplaceAllString(html, `$2`)
	return s
}
