package main

import (
  "fmt"
  "os"
  "golang.org/x/net/html"
)

func main() {
  doc, err := html.Parse(os.Stdin)

  if err != nil {
  fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
    os.Exit(1)
  }

  // links := []string{"prova"}
  // fmt.Println(visit(nil, doc, 0))

  // for _, link := range visit(nil, doc, 0) {
  //   fmt.Println(link)
  // }

  outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
  if n.Type == html.ElementNode {
    stack = append(stack, n.Data) // push tag
    fmt.Println(fmt.Sprintf("stack %s, tag: %s, attrs %s", stack, n.Data, n.Attr))
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    outline(stack, c)
  }

	// for _, child := range n.children {
	// 	outline(stack, child)
	// }
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node, level uint8) []string {

  // fmt.Println(fmt.Sprintf("Level: %d, tag %s, attrs %s", level, n.Data, n.Attr))
  
  if n.Type == html.ElementNode && n.Data == "a" {
    for _, a := range n.Attr {
      if a.Key == "href" {
        links = append(links, a.Val)
      }
    }
  }

  level++

  // for initialization; condition; update
  for c := n.FirstChild; c != nil; c = c.NextSibling {
    links = visit(links, c, level)
  }

  return links
}