package lineblocks_test

import (
	"log"
	"os"

	lineblocks "github.com/mdigger/goldmark-lineblocks"
	"github.com/yuin/goldmark"
)

func Example() {
	var md = goldmark.New(
		lineblocks.Enable,
		// goldmark.WithExtensions(lineblocks.Extension),
	)
	var source = []byte(`
| The limerick packs laughs anatomical
| In space that is quite economical.
|    But the good ones I've seen
|    So seldom are clean
| And the clean ones so seldom are comical`)
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// <p>The limerick packs laughs anatomical<br>
	// In space that is quite economical.<br>
	//    But the good ones I've seen<br>
	//    So seldom are clean<br>
	// And the clean ones so seldom are comical</p>
}
