# XMLBuilder
A quick and "dirty" functional XML builder and marshaller in Golang. Most of the time I'd recommend <ins>sticking with the standard library!</ins>

However, I've recently come across some limitations with this, particularly around interchangeably using self-closing (and non self-closing) XML tags.

## How to use
1. Install the package:
```go get github.com/Leila-Codes/XMLBuilder```
2. Create your document:

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/Leila-Codes/XMLBuilder"
)

func main() {
	doc := XMLBuilder.XDocument(
		XMLBuilder.XElement(
			// Write your element's name
			"User",
			// (optional) define any children
			// <FirstName>Leila-Codes</FirstName>
			XMLBuilder.XElement("FirstName").Content("Leila-Codes"), // set text content
			// <Metadata age="5"/>
			XMLBuilder.XElement("Metadata").Attribute("age", "5").Close(true), // self-close "true"
		).Close(),
	)

	file, err := os.Open("output.xml")
	if err != nil {
		panic(err)
	}

	// Marshal the whole thing.
	XMLBuilder.Marshal(doc, file)

	// Can also marshal in-memory
	content := bytes.Buffer{}
	XMLBuilder.Marshal(doc, content)
	
	/* OUTPUT:
	<?xml version="1.0" encoding="UTF-8"?>
	<User>
	    <FirstName>Leila-Codes</FirstName>
	    <Metadata age="5"/>
	</User>
	*/
	fmt.Print(content.String())
}
```
