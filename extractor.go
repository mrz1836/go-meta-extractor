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
			return tags
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == TagBody {
				return tags
			}
			if t.Data == TagTitle {
				titleFound = true
			}
			if t.Data == TagMeta {

				if value, ok = extractMetaProperty(t, TagMetaDescription); ok {
					tags.Description = truncateField(value, MaxFieldLength)
				}

				if value, ok = extractMetaProperty(t, TagMetaAuthor); ok {
					tags.Author = truncateField(value, MaxFieldLength)
				}

				if value, ok = extractMetaProperty(t, TagOGTitle); ok {
					tags.OGTitle = truncateField(value, MaxFieldLength)
					if len(tags.Title) == 0 {
						tags.Title = truncateField(value, MaxFieldLength)
					}
				}

				if value, ok = extractMetaProperty(t, TagOGDescription); ok {
					tags.OGDescription = truncateField(value, MaxFieldLength)
					if len(tags.Description) == 0 {
						tags.Description = truncateField(value, MaxFieldLength)
					}
				}

				if value, ok = extractMetaProperty(t, TagOGImage); ok {
					tags.OGImage = truncateField(value, MaxFieldLength)
				}

				if value, ok = extractMetaProperty(t, TagOGSiteName); ok {
					tags.OGSiteName = truncateField(value, MaxFieldLength)
				}

				if value, ok = extractMetaProperty(t, TagOGPublisher); ok {
					tags.OGPublisher = truncateField(value, MaxFieldLength)
				}

				if value, ok = extractMetaProperty(t, TagOGAuthor); ok {
					tags.OGAuthor = truncateField(value, MaxFieldLength)
					if len(tags.Author) == 0 {
						tags.Author = truncateField(value, MaxFieldLength)
					}
				}

				// Twitter card (use if OG not found)
				if value, ok = extractMetaProperty(t, TagTwitterTitle); ok {
					tags.TwitterTitle = truncateField(value, MaxFieldLength)
					if len(tags.Title) == 0 {
						tags.Title = truncateField(value, MaxFieldLength)
					}
				}

				if value, ok = extractMetaProperty(t, TagTwitterDescription); ok {
					tags.TwitterDescription = truncateField(value, MaxFieldLength)
					if len(tags.Description) == 0 {
						tags.Description = truncateField(value, MaxFieldLength)
					}
				}

				if value, ok = extractMetaProperty(t, TagTwitterImage); ok {
					tags.TwitterImage = truncateField(value, MaxFieldLength)
					if len(tags.OGImage) == 0 {
						tags.OGImage = truncateField(value, MaxFieldLength)
					}
				}

				if value, ok = extractMetaProperty(t, TagTwitterCard); ok {
					tags.TwitterCard = truncateField(value, MaxFieldLength)
				}

				if value, ok = extractMetaProperty(t, TagTwitterPlayer); ok {
					tags.TwitterPlayer = truncateField(value, MaxFieldLength)
				}
				if value, ok = extractMetaProperty(t, TagTwitterPlayerWidth); ok {
					tags.TwitterPlayerWidth = truncateField(value, MaxFieldLength)
				}
				if value, ok = extractMetaProperty(t, TagTwitterPlayerHeight); ok {
					tags.TwitterPlayerHeight = truncateField(value, MaxFieldLength)
				}
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				tags.Title = truncateField(t.Data, MaxFieldLength)
				titleFound = false
			}
		case html.CommentToken, html.DoctypeToken, html.EndTagToken:
			continue
		}
	}
}

// truncateField truncates a string to maxLen bytes if it exceeds that limit
// It handles Unicode properly by ensuring we don't truncate in the middle of a character
func truncateField(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	// Ensure we don't cut in the middle of a UTF-8 character
	// by finding the last valid rune boundary within maxLen bytes
	for i := maxLen; i >= 0; i-- {
		if i == 0 {
			return ""
		}
		if s[i]>>6 != 2 { // This byte is not a continuation byte (0b10xxxxxx)
			return s[:i]
		}
	}

	return ""
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
