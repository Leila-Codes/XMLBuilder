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
	"os"
)

func main() {
	// Open a file for writing
	file, err := os.OpenFile("output.xml", os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		panic(err)
	}

	// Create a new XML document
	doc := XMLBuilder.NewDocument(file)
	
	// Write your root element
	root := doc.Element("UserAccount")
	
	// Write its children
	root.Element("FirstName").Content("Leila-Codes").Close(false)
	root.Element("Metadata").Attr("age", "5").Close(true)
	
	// Close the root node (specifying whether it's a self-closing element or not)
	root.Close(false)

	// Close the file
	file.Close()
	
	/* File Contents:
	<?xml version="1.0" encoding="UTF-8"?>
	<UserAccount>
		<FirstName>Leila-Codes</FirstName>
		<Metadata age="5"/>
	</UserAccount>
	*/
}
```
