package XMLBuilder

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	doc := &Document{Root: Element{
		Attribute: &Attribute{
			Name:  "ExtensionInfo",
			Value: "",
		},
		Children: []Element{
			{
				Attribute: &Attribute{
					Name:  "Name",
					Value: "My First Extension",
				},
			},
		},
	}}

	expects := `<?xml version="1.0" encoding="UTF-8"?>
<ExtensionInfo>
	<Name>My First Extension</Name>
</ExtensionInfo>`

	buff := &bytes.Buffer{}
	err := Marshal(doc, buff)

	assert.Nil(t, err)
	assert.Equal(t, expects, buff.String())
}
