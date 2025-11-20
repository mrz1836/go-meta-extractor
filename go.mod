module github.com/mrz1836/go-meta-extractor

go 1.24.0

require (
	github.com/stretchr/testify v1.11.1
	golang.org/x/net v0.47.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Security: Force golang.org/x/crypto to v0.45.0 to fix CVE-2025-47914 and CVE-2025-58181
replace golang.org/x/crypto => golang.org/x/crypto v0.45.0
