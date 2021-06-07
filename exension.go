package lineblocks

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// A defaultExtension is goldmark extension for line blocks in markdown.
type defaultExtension struct{}

// Extend implement goldmark.Extender interface.
func (lb *defaultExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithParagraphTransformers(
		util.Prioritized(ParagraphTransformer, 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewHTMLRenderer(), 500),
	))
}

// Extension is a initialized goldmark extension for line blocks support.
var Extension goldmark.Extender = new(defaultExtension)

// Enable is goldmark.Enable for line blocks extension.
var Enable = goldmark.WithExtensions(Extension)
