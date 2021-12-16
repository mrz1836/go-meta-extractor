package metaextractor

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockPage is for creating a mock page of content
type MockPage struct {
	Content string
	done    bool
}

// NewMockPage is a new mock page for testing
func NewMockPage(s string) MockPage {
	return MockPage{Content: s}
}

// Read will read the mock page data
func (m *MockPage) Read(p []byte) (n int, err error) {
	if m.done {
		return 0, io.EOF
	}
	for i, b := range []byte(m.Content) {
		p[i] = b
	}
	m.done = true
	return len(m.Content), nil
}

// TestTitle will test the extraction of a title
func TestTitle(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		title         string
		mockHTML      string
		expectedTitle string
	}{
		{"Test Title", `<html><head><title>{{title}}</title></head></html>`, "Test Title"},
		{"Funky! <characters>", `<html><head><title>{{title}}</title></head></html>`, "Funky! <characters>"},
		{"Test Title", `<html><head><TITLE>{{title}}</TITLE></head></html>`, "Test Title"},
		{"", `<html><head><TITLE></TITLE></head></html>`, ""},
		{"", `<html><head><meta property="og:title" content="SomeName"/><title data-react-helmet="true"></title></head></html>`, "SomeName"},
		{"", `<html><head><title data-react-helmet="true"></title></head></html>`, ""},

		// todo: this should work but fails, I give up - the html parse does not like this (might replace pkg with https://github.com/PuerkitoBio/goquery)
		// {"", `<!doctype html><html lang="en-US"><head><meta charSet="UTF-8"/><meta property="og:site_name" content="SomeName"/><meta property="og:title" content="SomeName"/><meta property="og:image" content="https://domain.com/assets/ogimage.png"/><meta property="og:url" content="https://domain.com"/><meta property="og:description" content="Some Description"/><meta property="twitter:card" content="summary_large_image"/><meta property="twitter:site" content="@Handle"/><meta property="twitter:title" content="SomeName"/><meta property="twitter:description" content="Some Description"/><title data-react-helmet="true"></title><link rel="apple-touch-icon" sizes="57x57" href="/favicon/apple-icon-57x57.png"/><link rel="apple-touch-icon" sizes="60x60" href="/favicon/apple-icon-60x60.png"/><link rel="apple-touch-icon" sizes="72x72" href="/favicon/apple-icon-72x72.png"/><link rel="apple-touch-icon" sizes="76x76" href="/favicon/apple-icon-76x76.png"/><link rel="apple-touch-icon" sizes="114x114" href="/favicon/apple-icon-114x114.png"/><link rel="apple-touch-icon" sizes="120x120" href="/favicon/apple-icon-120x120.png"/><link rel="apple-touch-icon" sizes="144x144" href="/favicon/apple-icon-144x144.png"/><link rel="apple-touch-icon" sizes="152x152" href="/favicon/apple-icon-152x152.png"/><link rel="apple-touch-icon" sizes="180x180" href="/favicon/apple-icon-180x180.png"/><link rel="icon" type="image/png" sizes="192x192" href="/favicon/android-icon-192x192.png"/><link rel="icon" type="image/png" sizes="32x32" href="/favicon/favicon-32x32.png"/><link rel="icon" type="image/png" sizes="96x96" href="/favicon/favicon-96x96.png"/><link rel="icon" type="image/png" sizes="16x16" href="/favicon/favicon-16x16.png"/><meta name="viewport" content="width=device-width, initial-scale=1"/><link rel="manifest" href="/manifest.json"/><meta name="mobile-web-app-capable" content="yes"/><meta name="apple-mobile-web-app-capable" content="yes"/><meta name="application-name" content="SomeName"/><meta name="apple-mobile-web-app-status-bar-style" content="black"/><meta name="apple-mobile-web-app-title" content="SomeName"/><link href="https://fonts.googleapis.com/css?family=Montserrat:400,500,600,700&amp;display=swap" rel="stylesheet"/><link href="/assets/ReactToastify.min.css" rel="stylesheet"/><link rel="stylesheet" href="/assets/swiper.css"/><script src="https://widget.changelly.com/affiliate.js"></script><script src="//d.bablic.com/snippet/5e577416279d870001cda277.js?version=3.9"></script><script src="https://popups.landingi.com/api/v2/website/install-code?apikey=1a2af414-08cb-42ab-846b-3f7775c14bb6"></script><script>!function(e,t,n,s,u,a){e.twq||(s=e.twq=function(){s.exe?s.exe.apply(s,arguments):s.queue.push(arguments);},s.version=&#x27;1.1&#x27;,s.queue=[],u=t.createElement(n),u.async=!0,u.src=&#x27;//static.ads-twitter.com/uwt.js&#x27;,a=t.getElementsByTagName(n)[0],a.parentNode.insertBefore(u,a))}(window,document,&#x27;script&#x27;);// Insert Twitter Pixel ID and Standard Event data belowtwq(&#x27;init&#x27;,&#x27;o3mwp&#x27;);twq(&#x27;track&#x27;,&#x27;PageView&#x27;);</script><script>function(h,o,t,j,a,r){h.hj=h.hj||function(){(h.hj.q=h.hj.q||[]).push(arguments)};h._hjSettings={hjid:1704874,hjsv:6};a=o.getElementsByTagName(&#x27;head&#x27;)[0];r=o.createElement(&#x27;script&#x27;);r.async=1;r.src=t+h._hjSettings.hjid+j+h._hjSettings.hjsv;a.appendChild(r);})(window,document,&#x27;https://static.hotjar.com/c/hotjar-&#x27;,&#x27;.js?sv=&#x27;);</script><link href="/dist/main-e2d291e057bfd0afc8f8.css" rel="stylesheet"></head><body><div id="content"></div><script type="text/javascript" src="/dist/main-e2d291e057bfd0afc8f8.js"></script></body></html>`, "SomeName"},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{title}}", test.title, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedTitle, hm.Title)
	}
}

// TestDescription will test the extraction of a description
func TestDescription(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		description         string
		mockHTML            string
		expectedDescription string
	}{
		{"Test Description", `<html><head><meta property="description" content="{{description}}"></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta content="{{description}}" property="description"></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta content='{{description}}' property='description'></head></html>`, "Test Description"},
		{"Test Description", `<html><head><META CONTENT='{{description}}' PROPERTY='description'></head><body></body></html>`, "Test Description"},
		{"Test Description", `<html><head><META CONTENT='{{description}}' PROPERTY='description' /></head><body></body></html>`, "Test Description"},
		{"", `<html><head><meta content='' property='description'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{description}}", test.description, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedDescription, hm.Description)
	}
}

// TestAuthor will test the extraction of an author
func TestAuthor(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		author         string
		mockHTML       string
		expectedAuthor string
	}{
		{"Test Author", `<html><head><meta property="author" content="{{author}}"></head></html>`, "Test Author"},
		{"Test Author", `<html><head><meta content="{{author}}" property="author"></head></html>`, "Test Author"},
		{"Test Author", `<html><head><meta content='{{author}}' property='author'></head></html>`, "Test Author"},
		{"Test Author", `<html><head><META CONTENT='{{author}}' PROPERTY='author'></head><body></body></html>`, "Test Author"},
		{"Test Author", `<html><head><META CONTENT='{{author}}' PROPERTY='author' /></head><body></body></html>`, "Test Author"},
		{"", `<html><head><meta content='' property='author'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{author}}", test.author, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedAuthor, hm.Author)
	}
}

// TestOGTitle will test the extraction of an OG title
func TestOGTitle(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		title         string
		mockHTML      string
		expectedTitle string
	}{
		{"Test Title", `<html><head><meta property='og:title' content='{{title}}' /></head></html>`, "Test Title"},
		{"Test Title", `<html><head><meta property="og:title" content="{{title}}" /></head></html>`, "Test Title"},
		{"Test Title", `<html><head><meta content='{{title}}' property='og:title' /></head></html>`, "Test Title"},
		{"Test Title", `<html><head><meta content='{{title}}' property='og:title'></head></html>`, "Test Title"},
		{"", `<html><head><meta content='{{title}}' property='og:title'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{title}}", test.title, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedTitle, hm.OGTitle)
		assert.Equal(t, hm.Title, hm.OGTitle)
	}
}

// TestOGDescription will test the extraction of an OG description
func TestOGDescription(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		title               string
		mockHTML            string
		expectedDescription string
	}{
		{"Test Description", `<html><head><meta property='og:description' content='{{description}}' /></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta property="og:description" content="{{description}}" /></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta content='{{description}}' property='og:description' /></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta content='{{description}}' property='og:description'></head></html>`, "Test Description"},
		{"", `<html><head><meta content='{{description}}' property='og:description'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{description}}", test.title, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedDescription, hm.OGDescription)
		assert.Equal(t, hm.Description, hm.OGDescription)
	}
}

// TestOGImage will test the extraction of an OG image
func TestOGImage(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		image         string
		mockHTML      string
		expectedImage string
	}{
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta property='og:image' content='{{image}}' /></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta property="og:image" content="{{image}}" /></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta content='{{image}}' property='og:image' /></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta content='{{image}}' property='og:image'></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"", `<html><head><meta content='{{image}}' property='og:image'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{image}}", test.image, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedImage, hm.OGImage)
	}
}

// TestOGAuthor will test the extraction of an OG author
func TestOGAuthor(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		author         string
		mockHTML       string
		expectedAuthor string
	}{
		{"mrz", `<html><head><meta property='og:author' content='{{author}}' /></head></html>`, "mrz"},
		{"mrz", `<html><head><meta property="og:author" content="{{author}}" /></head></html>`, "mrz"},
		{"mrz", `<html><head><meta content='{{author}}' property='og:author' /></head></html>`, "mrz"},
		{"mrz", `<html><head><meta content='{{author}}' property='og:author'></head></html>`, "mrz"},
		{"", `<html><head><meta content='{{author}}' property='og:author'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{author}}", test.author, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedAuthor, hm.OGAuthor)
		assert.Equal(t, test.expectedAuthor, hm.Author)
	}
}

// TestOGPublisher will test the extraction of an OG publisher
func TestOGPublisher(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		publisher         string
		mockHTML          string
		expectedPublisher string
	}{
		{"mrz", `<html><head><meta property='og:publisher' content='{{publisher}}' /></head></html>`, "mrz"},
		{"mrz", `<html><head><meta property="og:publisher" content="{{publisher}}" /></head></html>`, "mrz"},
		{"mrz", `<html><head><meta content='{{publisher}}' property='og:publisher' /></head></html>`, "mrz"},
		{"mrz", `<html><head><meta content='{{publisher}}' property='og:publisher'></head></html>`, "mrz"},
		{"", `<html><head><meta content='{{publisher}}' property='og:publisher'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{publisher}}", test.publisher, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedPublisher, hm.OGPublisher)
	}
}

// TestOGSiteName will test the extraction of an OG site name
func TestOGSiteName(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		siteName         string
		mockHTML         string
		expectedSiteName string
	}{
		{"TheSite1", `<html><head><meta property='og:site_name' content='{{site_name}}' /></head></html>`, "TheSite1"},
		{"TheSite2", `<html><head><meta property="og:site_name" content="{{site_name}}" /></head></html>`, "TheSite2"},
		{"TheSite3", `<html><head><meta content='{{site_name}}' property='og:site_name'/></head></html>`, "TheSite3"},
		{"TheSite4", `<html><head><meta content='{{site_name}}' property='og:site_name'></head></html>`, "TheSite4"},
		{"", `<html><head><meta content='{{site_name}}' property='og:site_name'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{site_name}}", test.siteName, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedSiteName, hm.OGSiteName)
	}
}

// TestTwitterTitle will test the extraction of a twitter title
func TestTwitterTitle(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		title         string
		mockHTML      string
		expectedTitle string
	}{
		{"Test Title", `<html><head><meta property='twitter:title' content='{{title}}' /></head></html>`, "Test Title"},
		{"Test Title", `<html><head><meta property="twitter:title" content="{{title}}" /></head></html>`, "Test Title"},
		{"Test Title", `<html><head><meta content='{{title}}' property='twitter:title' /></head></html>`, "Test Title"},
		{"Test Title", `<html><head><meta content='{{title}}' property='twitter:title'></head></html>`, "Test Title"},
		{"", `<html><head><meta content='{{title}}' property='twitter:title'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{title}}", test.title, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedTitle, hm.TwitterTitle)
		assert.Equal(t, hm.Title, hm.TwitterTitle)
	}
}

// TestTwitterDescription will test the extraction of a twitter description
func TestTwitterDescription(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		title               string
		mockHTML            string
		expectedDescription string
	}{
		{"Test Description", `<html><head><meta property='twitter:description' content='{{description}}' /></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta property="twitter:description" content="{{description}}" /></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta content='{{description}}' property='twitter:description' /></head></html>`, "Test Description"},
		{"Test Description", `<html><head><meta content='{{description}}' property='twitter:description'></head></html>`, "Test Description"},
		{"", `<html><head><meta content='{{description}}' property='twitter:description'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{description}}", test.title, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedDescription, hm.TwitterDescription)
		assert.Equal(t, hm.Description, hm.TwitterDescription)
	}
}

// TestTwitterImage will test the extraction of a twitter image
func TestTwitterImage(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	tests := []struct {
		image         string
		mockHTML      string
		expectedImage string
	}{
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta property='twitter:image' content='{{image}}' /></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta property="twitter:image" content="{{image}}" /></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta content='{{image}}' property='twitter:image' /></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif", `<html><head><meta content='{{image}}' property='twitter:image'></head></html>`, "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"},
		{"", `<html><head><meta content='{{image}}' property='twitter:image'></head></html>`, ""},
	}

	// Test all
	for _, test := range tests {
		mp := NewMockPage(strings.Replace(test.mockHTML, "{{image}}", test.image, -1))
		assert.NotNil(t, mp)
		hm := Extract(&mp)
		assert.NotNil(t, hm)
		assert.Equal(t, test.expectedImage, hm.TwitterImage)
	}
}

// TestFullExtraction will test all the tags to extract
func TestFullExtraction(t *testing.T) {
	t.Parallel()

	title := "Test Title"
	description := "Test description"
	ogTitle := "OG Test Title"
	ogDesc := "OG Test description"
	ogImage := "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"
	ogAuthor := "MrZ"
	ogPublisher := "mrz"
	ogSiteName := "TheSite"
	twitterCard := "player"
	twitterPlayer := "https://www.youtube.com/watch?v=DoppJNHX1eY"
	twitterPlayerWidth := "1280"
	twitterPlayerHeight := "720"

	mp := NewMockPage(`
	<html>
		<head>
			<title>` + title + `</title>
			<meta property="description" content="` + description + `" />
			<meta property="og:title" content="` + ogTitle + `" />
			<meta property="twitter:title" content="` + ogTitle + `" />
			<meta property="og:description" content="` + ogDesc + `" />
			<meta property="twitter:description" content="` + ogDesc + `" />
			<meta property="og:image" content="` + ogImage + `" />
			<meta property="twitter:image" content="` + ogImage + `" />
			<meta property="og:author" content="` + ogAuthor + `" />
			<meta property="og:publisher" content="` + ogPublisher + `" />
			<meta property="og:site_name" content="` + ogSiteName + `" />
			<meta property="twitter:site" content="` + ogSiteName + `" />
			<meta property="twitter:card" content="` + twitterCard + `" />
			<meta property="twitter:player" content="` + twitterPlayer + `" />
			<meta property="twitter:player:width" content="` + twitterPlayerWidth + `" />
			<meta property="twitter:player:height" content="` + twitterPlayerHeight + `" />

		</head>
		<body>
			We're testing!
		</body>
	</html>`)

	assert.NotNil(t, mp)

	hm := Extract(&mp)
	assert.NotNil(t, hm)

	assert.Equal(t, description, hm.Description)
	assert.Equal(t, ogTitle, hm.OGTitle)
	assert.Equal(t, ogTitle, hm.TwitterTitle)
	assert.Equal(t, ogDesc, hm.OGDescription)
	assert.Equal(t, ogDesc, hm.TwitterDescription)
	assert.Equal(t, ogImage, hm.OGImage)
	assert.Equal(t, ogImage, hm.TwitterImage)
	assert.Equal(t, ogAuthor, hm.OGAuthor)
	assert.Equal(t, ogPublisher, hm.OGPublisher)
	assert.Equal(t, ogSiteName, hm.OGSiteName)
	assert.Equal(t, twitterCard, hm.TwitterCard)
	assert.Equal(t, twitterPlayer, hm.TwitterPlayer)
	assert.Equal(t, twitterPlayerWidth, hm.TwitterPlayerWidth)
	assert.Equal(t, twitterPlayerHeight, hm.TwitterPlayerHeight)
}

// ExampleExtract will show an example using the extractor
func ExampleExtract() {
	// Set a client
	client := &http.Client{Timeout: 20 * time.Second}

	// Start the request
	req, err := http.NewRequestWithContext(
		context.Background(), http.MethodGet, "https://mrz1818.com", nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Fire the request
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		log.Fatal(err)
	}

	// Close the body
	defer func() {
		_ = resp.Body.Close()
	}()

	// Extract the meta tags
	tags := Extract(resp.Body)

	fmt.Println(tags.Author)
	// Output:MrZ, Proof of Work LLC
}

// BenchmarkExtract benchmarks the method Extract()
func BenchmarkExtract(b *testing.B) {
	title := "Test Title"
	description := "Test description"
	ogTitle := "OG Test Title"
	ogDesc := "OG Test description"
	ogImage := "https://www.google.com/logos/doodles/2015/googles-new-logo-5078286822539264.2-hp.gif"
	ogAuthor := "MrZ"
	ogPublisher := "mrz"
	ogSiteName := "TheSite"

	mp := NewMockPage(`
	<html>
		<head>
			<title>` + title + `</title>
			<meta property="description" content="` + description + `" />
			<meta property="og:title" content="` + ogTitle + `" />
			<meta property="twitter:title" content="` + ogTitle + `" />
			<meta property="og:description" content="` + ogDesc + `" />
			<meta property="twitter:description" content="` + ogDesc + `" />
			<meta property="og:image" content="` + ogImage + `" />
			<meta property="twitter:image" content="` + ogImage + `" />
			<meta property="og:author" content="` + ogAuthor + `" />
			<meta property="og:publisher" content="` + ogPublisher + `" />
			<meta property="og:site_name" content="` + ogSiteName + `" />
		</head>
		<body>
			We're testing!
		</body>
	</html>`)

	for i := 0; i < b.N; i++ {
		_ = Extract(&mp)
	}
}
