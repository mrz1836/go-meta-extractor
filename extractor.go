// Package metaextractor will extract the title, description, OG & meta tags from HTML
//
// If you have any suggestions or comments, please feel free to open an issue on
// this GitHub repository!
//
// By MrZ (https://github.com/mrz1836)
package metaextractor

import (
	"io"

	"golang.org/x/net/html"
)

// Extract is the method used to extract HTML tags
func Extract(resp io.Reader) (tags Tags) {

	// Tokenize the response
	z := html.NewTokenizer(resp)

	// Set the values
	var value string
	var ok bool
	titleFound := false

	// Loop elements
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == TagBody {
				return
			}
			if t.Data == TagTitle {
				titleFound = true
			}
			if t.Data == TagMeta {

				if value, ok = extractMetaProperty(t, TagMetaDescription); ok {
					tags.Description = value
				}

				if value, ok = extractMetaProperty(t, TagMetaAuthor); ok {
					tags.Author = value
				}

				if value, ok = extractMetaProperty(t, TagOGTitle); ok {
					tags.OGTitle = value
					if len(tags.Title) == 0 {
						tags.Title = value
					}
				}

				if value, ok = extractMetaProperty(t, TagOGDescription); ok {
					tags.OGDescription = value
					if len(tags.Description) == 0 {
						tags.Description = value
					}
				}

				if value, ok = extractMetaProperty(t, TagOGImage); ok {
					tags.OGImage = value
				}

				if value, ok = extractMetaProperty(t, TagOGSiteName); ok {
					tags.OGSiteName = value
				}

				if value, ok = extractMetaProperty(t, TagOGPublisher); ok {
					tags.OGPublisher = value
				}

				if value, ok = extractMetaProperty(t, TagOGAuthor); ok {
					tags.OGAuthor = value
					if len(tags.Author) == 0 {
						tags.Author = value
					}
				}

				// Twitter card (use if OG not found)
				if value, ok = extractMetaProperty(t, TagTwitterTitle); ok {
					tags.TwitterTitle = value
					if len(tags.Title) == 0 {
						tags.Title = value
					}
				}

				if value, ok = extractMetaProperty(t, TagTwitterDescription); ok {
					tags.TwitterDescription = value
					if len(tags.Description) == 0 {
						tags.Description = value
					}
				}

				if value, ok = extractMetaProperty(t, TagTwitterImage); ok {
					tags.TwitterImage = value
					if len(tags.OGImage) == 0 {
						tags.OGImage = value
					}
				}
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				tags.Title = t.Data
				titleFound = false
			}
		case html.CommentToken, html.DoctypeToken, html.EndTagToken:
			continue
		}
	}
}

// extractMetaProperty will extract meta properties from HTML
func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if (attr.Key == TagProperty && attr.Val == prop) ||
			(attr.Key == TagName && attr.Val == prop) {
			ok = true
		}

		if attr.Key == TagContent {
			content = attr.Val
		}
	}

	return
}
