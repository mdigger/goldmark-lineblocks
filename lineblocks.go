// Package lineblocks is a extension for the goldmark
// (http://github.com/yuin/goldmark).
//
// This extension adds support for inline blocks in markdown.
//
// A LineBlocks is a sequence of lines beginning with a vertical bar (|)
// followed by a space. The division into lines will be preserved in the output,
// as will any leading spaces; otherwise, the lines will be formatted as
// Markdown. This is useful for verse and addresses:
//  | The limerick packs laughs anatomical
//  | In space that is quite economical.
//  |    But the good ones I've seen
//  |    So seldom are clean
//  | And the clean ones so seldom are comical
//
//  | 200 Main St.
//  | Berkeley, CA 94718
// The lines can be hard-wrapped if needed.
//  | The Right Honorable Most Venerable and Righteous Samuel L.
//    Constable, Jr.
//  | 200 Main St.
//  | Berkeley, CA 94718
// This syntax is borrowed from reStructuredText.
package lineblocks

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// A LineBlocks is goldmark extension for inline blocks in markdown.
type LineBlocks struct{}

// Extension is a initialized goldmark extension for line blocks support.
var Extension = new(LineBlocks)

var nbsp = []byte("&nbsp;")

// Transform implement parser.ASTTransformer inerface.
func (lb *LineBlocks) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	textNodes := findTexts(make([]*ast.Text, 0, 1000), node)
	source := reader.Source()
	for _, textNode := range textNodes {
		text := textNode.Text(source)
		if len(text) < 1 || text[0] != '|' {
			continue
		}
		// add line break
		if prev, ok := textNode.PreviousSibling().(*ast.Text); ok {
			if !prev.SoftLineBreak() {
				continue
			}
			prev.SetHardLineBreak(true)
		}
		// add spaces prefix
		spacesLength := util.TrimLeftSpaceLength(text[1:])
		if spacesLength > 2 {
			textNode.Parent().InsertBefore(textNode.Parent(), textNode,
				ast.NewString(bytes.Repeat(nbsp, spacesLength-1)))
		}
		textNode.Segment.Start += spacesLength + 1
		// textNode.Dump(source, 0)
	}
}

func findTexts(texts []*ast.Text, node ast.Node) []*ast.Text {
	for n := node.FirstChild(); n != nil; n = n.NextSibling() {
		if n.Kind() == ast.KindText && !n.IsRaw() {
			texts = append(texts, n.(*ast.Text))
		}
		texts = findTexts(texts, n)
	}
	return texts
}

// Extend implement goldmark.Extender interface.
func (lb *LineBlocks) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(lb, 0),
	))
}

// Enable is goldmark.Option for line blocks extension.
var Enable = goldmark.WithExtensions(Extension)
