package main

import (
	"fmt"
	"golang.org/x/net/html"
	"bytes"
	"strings"
	"os"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func main() {
	parseTree, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing input stream with err:", err)
		os.Exit(-1)
	}

	words, images := countWordsAndImages(parseTree)

	fmt.Printf("Found %d words and %d images\n", words, images)
}

func countWordsAndImages(p *html.Node) (int, int) {
	var words, images int

	searchAndCount(p, &words, &images)
	
	return words, images
}

func searchAndCount(p *html.Node, pwords, pimages *int) {

	if p.Type == html.TextNode {
		*pwords += len(strings.Fields(p.Data))
	} else if p.Type == html.ElementNode && p.Data == "img" {
		*pimages++
	}

	for c := p.FirstChild; c != nil; c = c.NextSibling {
		searchAndCount(c, pwords, pimages)
	}
}