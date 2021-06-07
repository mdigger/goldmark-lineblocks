package lineblocks

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type transformer struct{}

// Transform is a ParagraphTransformer implementation.
func (transformer) Transform(node *ast.Paragraph, reader text.Reader, pc parser.Context) {
	source := reader.Source()
	lines := node.Lines()
	if source == nil || lines.Len() == 0 {
		return
	}

	// check line block start
	start := lines.At(0).Start
	if len(source) < start || source[start] != '|' {
		return
	}

	// create new line block
	inlineBlock := NewLineBlock()
	inlineBlock.SetBlankPreviousLines(node.HasBlankPreviousLines())

	// skip last new line and spaces symbols in previous (last) line of block
	trimRightLastLine := func() {
		last := inlineBlock.LastChild()
		if last == nil {
			return
		}

		lines := last.Lines()
		pos := lines.Len() - 1
		if pos < 0 {
			return
		}

		line := lines.At(pos)
		line = line.TrimRightSpace(source)
		lines.Set(pos, line)
	}

	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		value := line.Value(source)
		// check is new line of line block
		if len(value) > 0 && value[0] == '|' {
			trimRightLastLine() // remove spaces in previous line of block

			width, pos := util.IndentWidth(value[1:], 1)
			line.Start += 1 + pos // skip '|' and remove left spaces
			if width > 0 {
				width-- // skip one space after '|'
			}

			// set line block
			inlineLine := NewLineBlockItem(width)
			inlineLine.Lines().Append(line)

			// add new line to line block
			inlineBlock.AppendChild(inlineBlock, inlineLine)
		} else {
			// add line to exist line of block
			inlineBlock.LastChild().Lines().Append(line)
		}
	}

	// replace paragraph with line block
	node.Parent().ReplaceChild(node.Parent(), node, inlineBlock)
}

// ParagraphTransformer is a ParagraphTransformer implementation
// that parses  and transform line blocks from paragraphs.
var ParagraphTransformer parser.ParagraphTransformer = new(transformer)
