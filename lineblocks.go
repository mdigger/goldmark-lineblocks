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

type transformer struct{}

var defaultTransformer parser.ASTTransformer = new(transformer)

// // NewTransformer return inline blocks transformer.
// func NewTransformer() parser.ASTTransformer {
// 	return defaultTransformer
// }

var nbsp = []byte("&nbsp;")

// Transform implement parser.ASTTransformer inerface.
func (lb *transformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	var source = reader.Source()
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering || node.Kind() != ast.KindText {
			return ast.WalkContinue, nil
		}
		var text = node.Text(source)
		// check line block start marker
		if len(text) < 1 || text[0] != '|' {
			return ast.WalkSkipChildren, nil
		}
		// add line break
		if prev, ok := node.PreviousSibling().(*ast.Text); ok {
			if !prev.SoftLineBreak() {
				return ast.WalkSkipChildren, nil // not line start
			}
			prev.SetHardLineBreak(true)
		}
		// add spaces prefix
		var spaces = util.TrimLeftSpaceLength(text[1:])
		if spaces > 2 {
			indent := ast.NewString(bytes.Repeat(nbsp, spaces-1))
			indent.SetCode(true)
			node.Parent().InsertBefore(node.Parent(), node, indent)
		}
		// remove line block prefix and spaces
		node.(*ast.Text).Segment.Start += spaces + 1
		return ast.WalkSkipChildren, nil
	})
}

// A extension is goldmark extension for line extension in markdown.
type extension struct{}

// Extend implement goldmark.Extender interface.
func (lb *extension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(defaultTransformer, 500),
	))
}

// Extension is a initialized goldmark extension for line blocks support.
var Extension goldmark.Extender = new(extension)

// Enable is goldmark.Enable for line blocks extension.
var Enable = goldmark.WithExtensions(Extension)
