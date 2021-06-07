package lineblocks

import (
	"strconv"

	"github.com/yuin/goldmark/ast"
)

// A LineBlock struct represents a list of inline text block.
type LineBlock struct {
	ast.BaseBlock
}

// NewLineBlock return new initialized LineBlock.
func NewLineBlock() *LineBlock {
	return new(LineBlock)
}

// KindLineBlock is a NodeKind of the LineBlock node.
var KindLineBlock = ast.NewNodeKind("LineBlock")

// Kind implements Node.Kind.
func (*LineBlock) Kind() ast.NodeKind {
	return KindLineBlock
}

// Dump implements Node.Dump.
func (b *LineBlock) Dump(source []byte, level int) {
	ast.DumpHelper(b, source, level, nil, nil)
}

// A LineBlockItem struct represents a list of inline text block.
type LineBlockItem struct {
	ast.BaseBlock
	Padding int
}

// NewLineBlockItem return new initialized LineBlockItem.
func NewLineBlockItem(padding int) *LineBlockItem {
	return &LineBlockItem{
		Padding: padding,
	}
}

// KindLineBlock is a NodeKind of the LineBlock node.
var KindLineBlockItem = ast.NewNodeKind("LineBlockItem")

// Kind implements Node.Kind.
func (*LineBlockItem) Kind() ast.NodeKind {
	return KindLineBlockItem
}

// Dump implements Node.Dump.
func (l *LineBlockItem) Dump(source []byte, level int) {
	m := map[string]string{
		"Padding": strconv.Itoa(l.Padding),
	}
	ast.DumpHelper(l, source, level, m, nil)
}
