# CLAUDE.md

Quick reference for Claude Code when working with **go-meta-extractor** - a Go library that extracts HTML metadata (title, description, OG tags, Twitter Cards).

## 🎯 Core Function
```go
func Extract(resp io.Reader) (tags Tags) // Extracts all meta tags from HTML
```

## 📋 Key Data Structure
```go
type Tags struct {
    Title               string // HTML <title> or fallback from OG/Twitter
    Description         string // meta description or OG/Twitter fallback
    Author              string // meta author or OG author
    OGTitle            string // og:title
    OGDescription      string // og:description
    OGImage            string // og:image
    OGSiteName         string // og:site_name
    OGPublisher        string // og:publisher
    OGAuthor           string // og:author
    TwitterTitle       string // twitter:title
    TwitterDescription string // twitter:description
    TwitterImage       string // twitter:image
    TwitterCard        string // twitter:card
    TwitterPlayer      string // twitter:player
    TwitterPlayerWidth string // twitter:player:width
    TwitterPlayerHeight string // twitter:player:height
}
```

## 🛠️ Essential Commands
```bash
magex test           # Run all tests (required before commits)
magex test:race      # Run tests with race detector
magex bench          # Run benchmarks
magex lint           # Run linter (if available)
magex help           # List all available commands
```

## ⚡ Development Rules
- **Dependencies**: Keep minimal - only uses `golang.org/x/net/html` + testify for tests
- **Field Limits**: All extracted fields truncated at `MaxFieldLength = 10000` bytes
- **UTF-8 Safety**: `truncateField()` handles Unicode boundaries properly
- **Testing**: Comprehensive unit tests + fuzz tests in `extractor_fuzz_test.go`
- **Performance**: Stops parsing at `<body>` tag for efficiency

## ✅ Pre-Commit Checklist
1. `magex test` passes ✓
2. `magex test:race` passes ✓
3. No new dependencies added without justification ✓
4. Fuzz tests pass if touching extraction logic ✓

## 🔍 Common Tasks

### Adding New Meta Tag Support
1. Add constant to `definitions.go` (e.g., `TagNewMeta = "new:meta"`)
2. Add field to `Tags` struct with JSON tag
3. Add extraction logic in `extractor.go` Extract() function
4. Add test case in `extractor_test.go`
5. Update fuzz test seeds if needed

### Testing Changes
```go
// Use MockPage for testing
mp := NewMockPage(`<html><head><title>Test</title></head></html>`)
tags := Extract(&mp)
assert.Equal(t, "Test", tags.Title)
```

### Example Usage
```go
resp, _ := http.Get("https://example.com")
defer resp.Body.Close()
tags := metaextractor.Extract(resp.Body)
fmt.Printf("Title: %s\nDescription: %s\n", tags.Title, tags.Description)
```

## 📁 Key Files
- `extractor.go` - Main extraction logic
- `definitions.go` - Constants and Tags struct
- `extractor_test.go` - Unit tests
- `extractor_fuzz_test.go` - Fuzz tests
- `examples/example.go` - Usage example

---
*For detailed conventions, see [AGENTS.md](AGENTS.md) and [tech-conventions/](tech-conventions)*
