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

// A LineBlocks is goldmark extension for line blocks in markdown.
type LineBlocks struct{}

// Extension is a initialized goldmark extension for line blocks support.
var Extension = new(LineBlocks)

var nbsp = []byte("&nbsp;")

// Transform implement parser.ASTTransformer inerface.
func (lb *LineBlocks) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	var source = reader.Source()
	walkToTextNode(node, func(node ast.Node) {
		text := node.Text(source)
		// check line block start marker
		if len(text) < 1 || text[0] != '|' {
			return
		}
		// add line break
		if prev, ok := node.PreviousSibling().(*ast.Text); ok {
			if !prev.SoftLineBreak() {
				return // not line start
			}
			prev.SetHardLineBreak(true)
		}
		// add spaces prefix
		spaces := util.TrimLeftSpaceLength(text[1:])
		if spaces > 2 {
			node.Parent().InsertBefore(node.Parent(), node,
				ast.NewString(bytes.Repeat(nbsp, spaces-1)))
		}
		// remove line block prefix and spaces
		node.(*ast.Text).Segment.Start += spaces + 1
	})
}

func walkToTextNode(node ast.Node, f func(node ast.Node)) {
	for node = node.FirstChild(); node != nil; node = node.NextSibling() {
		if node.Kind() == ast.KindText {
			f(node)
			continue
		}
		if node.HasChildren() {
			walkToTextNode(node, f)
		}
	}
}

// Extend implement goldmark.Extender interface.
func (lb *LineBlocks) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(lb, 500),
	))
}

// Enable is goldmark.Enable for line blocks extension.
var Enable = goldmark.WithExtensions(Extension)
