package XMLBuilder

import "strings"

const (
	StartElement = "<"
	EndElement   = ">"
	Quote        = "\""
	Apostrophe   = "'"
	Ampersand    = "&"
	Space        = " "
	Equals       = "="
)

var xmlEncoder = strings.NewReplacer(
	StartElement, "&lt;",
	EndElement, "&gt;",
	Quote, "&quot;",
	Apostrophe, "&apos;",
	Ampersand, "&amp;",
)

var xmlDecoder = strings.NewReplacer(
	"&lt;", StartElement,
	"&gt;", EndElement,
	"&quot;", Quote,
	"&apos;", Apostrophe,
	"&amp;", Ampersand,
)

func quote(s string) string {
	return Quote + xmlEncoder.Replace(s) + Quote
}
