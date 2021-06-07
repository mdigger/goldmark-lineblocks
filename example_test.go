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
	// <div class="line-block">
	// <div class="line">The limerick packs laughs anatomical</div>
	// <div class="line">In space that is quite economical.</div>
	// <div class="line">&nbsp;&nbsp;&nbsp;But the good ones I've seen</div>
	// <div class="line">&nbsp;&nbsp;&nbsp;So seldom are clean</div>
	// <div class="line">And the clean ones so seldom are comical</div>
	// </div>
}
