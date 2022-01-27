package myutils

import (
	"golang.org/x/net/html"
)

/* ************ Private Methods ************* */
func getAttribute(n *html.Node, key string) (string, bool){
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func traverse(n *html.Node, id string, attrKey string) *html.Node {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, attrKey)
		if ok && s == id {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, id, attrKey)
		if result != nil {
			return result
		}
	}
	return nil
}

/* ************ Public Methods ************* */
func GetElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id, "id")
}

func GetElementByType(n *html.Node, id string) *html.Node {
	return traverse(n, id, "type")
}