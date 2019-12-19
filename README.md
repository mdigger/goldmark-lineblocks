# goldmark-lineblocks

[![GoDoc](https://godoc.org/github.com/mdigger/goldmark-lineblocks?status.svg)](https://godoc.org/github.com/mdigger/goldmark-lineblocks)

This [goldmark](http://github.com/yuin/goldmark) parser extension adds support for inline blocks in markdown.

```go
var md = goldmark.New(
    goldmark.WithExtensions(lineblocks.Enable),
)
err := md.Convert(source, os.Stdout)
if err != nil {
    log.Fatal(err)
}
```

A line block is a sequence of lines beginning with a vertical bar (`|`) followed by a space. The division into lines will be preserved in the output, as will any leading spaces; otherwise, the lines will be formatted as Markdown. This is useful for verse and addresses:

    | The limerick packs laughs anatomical
    | In space that is quite economical.
    |    But the good ones I've seen
    |    So seldom are clean
    | And the clean ones so seldom are comical

    | 200 Main St.
    | Berkeley, CA 94718

The lines can be hard-wrapped if needed.

    | The Right Honorable Most Venerable and 
      Righteous Samuel L. Constable, Jr.
    | 200 Main St.
    | Berkeley, CA 94718

This syntax is borrowed from [reStructuredText](http://docutils.sourceforge.net/docs/ref/rst/introduction.html).

