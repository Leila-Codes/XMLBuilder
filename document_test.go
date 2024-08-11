package XMLBuilder

import (
	"bytes"
	"os"
	"testing"
)

func TestNewDocument(t *testing.T) {
	buff := &bytes.Buffer{}
	doc := NewDocument(buff)

	doc.Element("Computer").
		Attr("id", "my-unique-computer").
		Attr("name", "My First Computer").
		Attr("ip", "#RANDOM_IP#").Content("Hello World").Close(false)

	doc.writer.Flush()

	t.Logf("Buffer: \n%s", buff.String())
}

func TestNested(t *testing.T) {
	buff := &bytes.Buffer{}
	doc := NewDocument(buff)
	doc.Element("HacknetExtension").
		Attr("author", "Leila-Codes").
		Element("Name").Content("My First Extension").Close(false).
		Close(false)

	t.Logf("Buffer: \n%s", buff.String())
}

func TestFile(t *testing.T) {
	file, err := os.OpenFile("ExtensionInfo.xml", os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		panic(err)
	}

	doc := NewDocument(file)
	doc.Element("ExtensionInfo").
		Element("Language").Content("en-us").Close(false).
		Element("Name").Content("Intro Extension").Close(false).
		Element("AllowSaves").Content("true").Close(false).
		Element("StartingVisibleNodes").Content("advExamplePC").Close(false).
		Element("StartingMission").Content("Missions/Intro/IntroMission1.xml").Close(false).
		// <StartingActions>Actions/StartingActions.xml</StartingActions>
		Element("StartingActions").Content("Actions/StartingActions.xml").Close(false).
		Element("Description").Content(` --- Intro Extension ---
This example extension will teach the basics of building a Hacknet Extension.
Descriptions can be multiple lines, so, we're learning already!`).Close(false).
		Element("Faction").Content("Factions/ExampleFaction.xml").Close(false).
		Element("Faction").Content("Factions/IntroFaction.xml").Close(false).
		Close(false)

	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestReadmeExample(t *testing.T) {
	// Open a file for writing
	file, err := os.OpenFile("output.xml", os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		panic(err)
	}

	// Create a new XML document
	doc := NewDocument(file)

	// Write your root element
	root := doc.Element("UserAccount")

	// Write its children
	root.Element("FirstName").Content("Leila-Codes").Close(false)
	root.Element("Metadata").Attr("age", "5").Close(true)

	// Close the root node (specifying whether it's a self-closing element or not)
	root.Close(false)

	// Close the file
	file.Close()
}
