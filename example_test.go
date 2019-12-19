package lineblocks_test

import (
	"log"
	"os"

	lineblocks "github.com/mdigger/goldmark-lineblocks"
	"github.com/yuin/goldmark"
)

var source = []byte(`# Line blocks

A line block is a sequence of lines beginning with a vertical bar (|) followed by a space. The division into lines will be preserved in the output, as will any leading spaces; otherwise, the lines will be formatted as Markdown. This is useful for verse and addresses:

| The limerick packs laughs anatomical
| In space that is quite economical.
|    But the good ones I've seen
|    So seldom are clean
| And the clean ones so seldom are comical

| 200 Main St.
| Berkeley, CA 94718

The lines can be hard-wrapped if needed.

| The Right Honorable Most Venerable and Righteous Samuel L.
  Constable, Jr.
| 200 Main St.
| Berkeley, CA 94718

This syntax is borrowed from [reStructuredText](http://docutils.sourceforge.net/docs/ref/rst/introduction.html).
`)

func Example() {
	var md = goldmark.New(
		goldmark.WithExtensions(lineblocks.Enable),
	)
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// <h1>Line blocks</h1>
	// <p>A line block is a sequence of lines beginning with a vertical bar (|) followed by a space. The division into lines will be preserved in the output, as will any leading spaces; otherwise, the lines will be formatted as Markdown. This is useful for verse and addresses:</p>
	// <p>The limerick packs laughs anatomical<br>
	// In space that is quite economical.<br>
	//    But the good ones I've seen<br>
	//    So seldom are clean<br>
	// And the clean ones so seldom are comical</p>
	// <p>200 Main St.<br>
	// Berkeley, CA 94718</p>
	// <p>The lines can be hard-wrapped if needed.</p>
	// <p>The Right Honorable Most Venerable and Righteous Samuel L.
	// Constable, Jr.<br>
	// 200 Main St.<br>
	// Berkeley, CA 94718</p>
	// <p>This syntax is borrowed from <a href="http://docutils.sourceforge.net/docs/ref/rst/introduction.html">reStructuredText</a>.</p>
}
