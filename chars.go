package XMLBuilder

import "strings"

const (
	StartElement = '<'
	EndElement   = '>'
	Quote        = '"'
	Apostrophe   = '\''
	Ampersand    = '&'
	Space        = ' '
	Equals       = '='
)

var xmlEncoder = strings.NewReplacer(
	string(StartElement), "&lt;",
	string(EndElement), "&gt;",
	string(Quote), "&quot;",
	//string(Apostrophe), "&apos;",
	string(Ampersand), "&amp;",
)

var xmlDecoder = strings.NewReplacer(
	"&lt;", string(StartElement),
	"&gt;", string(EndElement),
	"&quot;", string(Quote),
	"&apos;", string(Apostrophe),
	"&amp;", string(Ampersand),
)

func quote(s string) string {
	return string(Quote) + xmlEncoder.Replace(s) + string(Quote)
}
