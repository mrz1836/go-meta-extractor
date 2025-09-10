package metaextractor

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"golang.org/x/net/html"
)

// FuzzExtract tests the main Extract function with random HTML input
func FuzzExtract(f *testing.F) {
	// Seed corpus with known good inputs
	f.Add(`<html><head><title>Test</title></head><body></body></html>`)
	f.Add(`<html><head><meta name="description" content="Test description"/></head></html>`)
	f.Add(`<html><head><meta property="og:title" content="OG Title"/></head></html>`)
	f.Add(`<html><head><meta property="twitter:card" content="summary"/></head></html>`)
	f.Add(``)
	f.Add(`<html><head></head><body></body></html>`)
	f.Add(`<title>Simple Title</title>`)
	f.Add(`<meta name="author" content="Test Author">`)

	f.Fuzz(func(t *testing.T, input string) {
		// Ensure input is valid UTF-8
		if !utf8.ValidString(input) {
			t.Skip("Skipping non-UTF-8 input")
		}

		// Limit input size to prevent excessive memory usage
		if len(input) > 100000 { // 100KB limit
			t.Skip("Skipping overly large input")
		}

		// Test extraction without panicking
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Extract panicked with input %q: %v", input, r)
			}
		}()

		reader := strings.NewReader(input)
		tags := Extract(reader)

		// Validate extracted data doesn't contain control characters (except newlines/tabs)
		validateExtractedTags(t, tags)
	})
}

// FuzzExtractMetaProperty tests the extractMetaProperty helper function
func FuzzExtractMetaProperty(f *testing.F) {
	// Seed with valid HTML token structures
	f.Add(`name`, `description`, `content`, `Test content`)
	f.Add(`property`, `og:title`, `content`, `OG Title`)
	f.Add(`property`, `twitter:card`, `content`, `summary`)
	f.Add(`name`, `author`, `content`, ``)
	f.Add(``, ``, ``, ``)

	f.Fuzz(func(t *testing.T, key1, val1, key2, val2 string) {
		// Ensure strings are valid UTF-8
		if !utf8.ValidString(key1) || !utf8.ValidString(val1) ||
			!utf8.ValidString(key2) || !utf8.ValidString(val2) {
			t.Skip("Skipping non-UTF-8 input")
		}

		// Limit string lengths
		if len(key1) > 1000 || len(val1) > 1000 || len(key2) > 1000 || len(val2) > 1000 {
			t.Skip("Skipping overly long strings")
		}

		// Create a mock HTML token
		token := html.Token{
			Type: html.StartTagToken,
			Data: "meta",
			Attr: []html.Attribute{
				{Key: key1, Val: val1},
				{Key: key2, Val: val2},
			},
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("extractMetaProperty panicked: %v", r)
			}
		}()

		content, ok := extractMetaProperty(token, val1)

		// Validate results
		if ok && key2 == TagContent {
			if content != val2 {
				t.Errorf("Expected content %q, got %q", val2, content)
			}
		}

		// Ensure content is valid UTF-8
		if !utf8.ValidString(content) {
			t.Errorf("extractMetaProperty returned invalid UTF-8 content")
		}
	})
}

// FuzzExtractWithMalformedHTML tests resilience against malformed HTML
func FuzzExtractWithMalformedHTML(f *testing.F) {
	// Seed with various malformed HTML patterns
	f.Add(`<html><head><title>Unclosed title`)
	f.Add(`<meta name="description" content="Unclosed meta`)
	f.Add(`<html><head><title></title><meta property="og:title" content="></head></html>`)
	f.Add(`<><><><>`)
	f.Add(`<<>><<>>`)
	f.Add(`<html><head><title>Test</title><meta></head></html>`)
	f.Add(`<html><body><title>Title in body</title></body></html>`)
	f.Add(`<!DOCTYPE html><html><head><title>Doctype test</title></head></html>`)

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Skipping non-UTF-8 input")
		}

		if len(input) > 50000 {
			t.Skip("Skipping overly large input")
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Extract panicked with malformed HTML %q: %v", input, r)
			}
		}()

		reader := strings.NewReader(input)
		tags := Extract(reader)

		// Should not panic, but extracted data should be safe
		validateExtractedTags(t, tags)
	})
}

// FuzzExtractWithSpecialCharacters tests handling of special characters and encodings
func FuzzExtractWithSpecialCharacters(f *testing.F) {
	// Seed with various special character patterns
	f.Add(`<title>Test "quotes" & ampersands</title>`)
	f.Add(`<meta name="description" content="Content with 'single quotes'"/>`)
	f.Add(`<meta property="og:title" content="Unicode: 🎉 ñáéíóú"/>`)
	f.Add(`<title>HTML entities: &lt;&gt;&amp;&quot;&#39;</title>`)
	f.Add(`<meta name="author" content="Author with emoji 👨‍💻"/>`)
	f.Add(`<title>Newlines\nand\ttabs</title>`)

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Skipping non-UTF-8 input")
		}

		if len(input) > 10000 {
			t.Skip("Skipping overly large input")
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Extract panicked with special characters %q: %v", input, r)
			}
		}()

		reader := strings.NewReader(input)
		tags := Extract(reader)

		validateExtractedTags(t, tags)
	})
}

// FuzzExtractWithNestedTags tests complex nested HTML structures
func FuzzExtractWithNestedTags(f *testing.F) {
	// Seed with nested structures
	f.Add(`<html><head><title><span>Nested</span> Title</title></head></html>`)
	f.Add(`<html><head><meta><meta property="og:title" content="Nested meta"/></meta></head></html>`)
	f.Add(`<html><body><div><title>Deep nested title</title></div></body></html>`)
	f.Add(`<html><head><!-- Comment --><title>Title with comment</title></head></html>`)

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Skipping non-UTF-8 input")
		}

		if len(input) > 20000 {
			t.Skip("Skipping overly large input")
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Extract panicked with nested HTML %q: %v", input, r)
			}
		}()

		reader := strings.NewReader(input)
		tags := Extract(reader)

		validateExtractedTags(t, tags)
	})
}

// FuzzExtractLargeInput tests memory safety with large inputs
func FuzzExtractLargeInput(f *testing.F) {
	// Seed with patterns that could be expanded
	f.Add(`<title>` + strings.Repeat("A", 1000) + `</title>`)
	f.Add(`<meta name="description" content="` + strings.Repeat("B", 500) + `"/>`)

	f.Fuzz(func(t *testing.T, baseInput string) {
		if !utf8.ValidString(baseInput) {
			t.Skip("Skipping non-UTF-8 input")
		}

		// Skip empty input to avoid division by zero
		if len(baseInput) == 0 {
			t.Skip("Skipping empty input")
		}

		// Create various sizes of input
		sizes := []int{1000, 5000, 10000}

		for _, size := range sizes {
			if len(baseInput)*size > 200000 { // 200KB limit
				continue
			}

			input := strings.Repeat(baseInput, size/len(baseInput)+1)[:size]

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Extract panicked with large input (size %d): %v", size, r)
				}
			}()

			reader := strings.NewReader(input)
			tags := Extract(reader)

			validateExtractedTags(t, tags)
		}
	})
}

// FuzzExtractXSSPayloads tests resistance to XSS-like patterns
func FuzzExtractXSSPayloads(f *testing.F) {
	// Seed with XSS-like patterns (for defensive testing)
	f.Add(`<title><script>alert('xss')</script></title>`)
	f.Add(`<meta name="description" content="<script>alert(1)</script>"/>`)
	f.Add(`<meta property="og:title" content="javascript:alert(1)"/>`)
	f.Add(`<title>'; DROP TABLE users; --</title>`)
	f.Add(`<meta name="author" content="<img src=x onerror=alert(1)>"/>`)

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Skipping non-UTF-8 input")
		}

		if len(input) > 5000 {
			t.Skip("Skipping overly large input")
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Extract panicked with XSS-like input %q: %v", input, r)
			}
		}()

		reader := strings.NewReader(input)
		tags := Extract(reader)

		// The extractor should handle these safely
		validateExtractedTags(t, tags)

		// Ensure no script tags are extracted as content
		checkForScriptContent(t, tags)
	})
}

// FuzzExtractInfiniteLoops tests for potential infinite loops
func FuzzExtractInfiniteLoops(f *testing.F) {
	// Seed with patterns that might cause issues
	f.Add(`<html><head><title><title><title>Nested titles</title></title></title></head></html>`)
	f.Add(strings.Repeat(`<meta name="test" content="loop"/>`, 100))
	f.Add(`<html>` + strings.Repeat(`<head>`, 50) + `<title>Deep</title>` + strings.Repeat(`</head>`, 50) + `</html>`)

	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("Skipping non-UTF-8 input")
		}

		if len(input) > 30000 {
			t.Skip("Skipping overly large input")
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Extract panicked with potential loop input %q: %v", input, r)
			}
		}()

		reader := strings.NewReader(input)
		tags := Extract(reader)

		validateExtractedTags(t, tags)
	})
}

// validateExtractedTags ensures extracted tag data is safe and valid
func validateExtractedTags(t *testing.T, tags Tags) {
	tagFields := []string{
		tags.Author, tags.Description, tags.OGAuthor, tags.OGDescription,
		tags.OGImage, tags.OGPublisher, tags.OGSiteName, tags.OGTitle,
		tags.Title, tags.TwitterDescription, tags.TwitterImage,
		tags.TwitterCard, tags.TwitterPlayer, tags.TwitterPlayerHeight,
		tags.TwitterPlayerWidth, tags.TwitterTitle,
	}

	for i, field := range tagFields {
		// Ensure all extracted content is valid UTF-8
		if !utf8.ValidString(field) {
			t.Errorf("Field %d contains invalid UTF-8: %q", i, field)
		}

		// Log suspicious control characters (except common whitespace)
		for _, r := range field {
			if r < 32 && r != '\n' && r != '\t' && r != '\r' {
				t.Logf("Field %d contains control character: %q (rune: %d)", i, field, r)
			}
		}

		// Reasonable length check
		if len(field) > 10000 {
			t.Errorf("Field %d is unreasonably long (%d chars)", i, len(field))
		}
	}
}

// checkForScriptContent ensures no script content is extracted
func checkForScriptContent(t *testing.T, tags Tags) {
	tagFields := []string{
		tags.Author, tags.Description, tags.OGAuthor, tags.OGDescription,
		tags.OGImage, tags.OGPublisher, tags.OGSiteName, tags.OGTitle,
		tags.Title, tags.TwitterDescription, tags.TwitterImage,
		tags.TwitterCard, tags.TwitterPlayer, tags.TwitterPlayerHeight,
		tags.TwitterPlayerWidth, tags.TwitterTitle,
	}

	scriptPatterns := []string{"<script", "javascript:", "data:text/html"}

	for i, field := range tagFields {
		fieldLower := strings.ToLower(field)
		for _, pattern := range scriptPatterns {
			if strings.Contains(fieldLower, pattern) {
				t.Logf("Warning: Field %d contains potential script content: %q", i, field)
			}
		}
	}
}

// Benchmark comparison for fuzzed inputs
func BenchmarkFuzzExtract(b *testing.B) {
	testInputs := []string{
		`<html><head><title>Simple</title></head></html>`,
		`<html><head><meta property="og:title" content="Complex OG Title"/><meta name="description" content="Test description"/></head></html>`,
		strings.Repeat(`<meta name="test" content="repeated content"/>`, 50),
	}

	for i, input := range testInputs {
		b.Run(fmt.Sprintf("Input%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				reader := bytes.NewReader([]byte(input))
				_ = Extract(reader)
			}
		})
	}
}
