package myutils

import (
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/* ************ Public Methods ************* */
func GetHtmlNode(URL string) *html.Node {
	// Send GET request for URL
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Convert the body to type string
	sb := string(body)
	// Parse the html body to html node
	node, err := html.Parse(strings.NewReader(sb))
	if err != nil {
		log.Fatalln(err)
	}

	return node
}
