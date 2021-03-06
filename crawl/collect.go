package crawl

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// All takes a reader object (like the one returned from http.Get())
// a tag name and attribute name and extracts the attribute value from all
// specified tags on the page.
// It does not close the reader passed to it.
func All(httpBody io.Reader, tagName, attributeName string) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if err := page.Err(); err != nil {
			fmt.Println(err)
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == tagName {
			for _, attr := range token.Attr {
				if attr.Key == attributeName {
					tl := trimHash(attr.Val)
					_, err := url.Parse(tl)
					if err != nil {
						continue
					}
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}
