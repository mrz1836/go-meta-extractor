package metaextractor

// Tags is the html/meta tags to extract and process
type Tags struct {
	Author              string `json:"author"`
	Description         string `json:"description"`
	OGAuthor            string `json:"og_author"`
	OGDescription       string `json:"og_description"`
	OGImage             string `json:"og_image"`
	OGPublisher         string `json:"og_publisher"`
	OGSiteName          string `json:"og_site_name"`
	OGTitle             string `json:"og_title"`
	Title               string `json:"title"`
	TwitterDescription  string `json:"twitter_description"`
	TwitterImage        string `json:"twitter_image"`
	TwitterCard         string `json:"twitter_card"`
	TwitterPlayer       string `json:"twitter_player"`
	TwitterPlayerHeight string `json:"twitter_player_height"`
	TwitterPlayerWidth  string `json:"twitter_player_width"`
	TwitterTitle        string `json:"twitter_title"`
}

// todo: parse the apple mobile title
// <meta name="apple-mobile-web-app-title" content="SiteTitle"/>

// Tag and Property constants for parsing
const (
	TagBody                = "body"
	TagContent             = "content"
	TagMeta                = "meta"
	TagMetaAuthor          = "author"
	TagMetaDescription     = "description"
	TagName                = "name"
	TagOGAuthor            = "og:author"
	TagOGDescription       = "og:description"
	TagOGImage             = "og:image"
	TagOGPublisher         = "og:publisher"
	TagOGSiteName          = "og:site_name"
	TagOGTitle             = "og:title"
	TagProperty            = "property"
	TagTitle               = "title"
	TagTwitterCard         = "twitter:card"
	TagTwitterDescription  = "twitter:description"
	TagTwitterImage        = "twitter:image"
	TagTwitterPlayer       = "twitter:player"
	TagTwitterPlayerHeight = "twitter:player:height"
	TagTwitterPlayerWidth  = "twitter:player:width"
	TagTwitterTitle        = "twitter:title"
)
