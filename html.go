package lineblocks

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// htmlRenderer is a renderer.NodeRenderer implementation that renders
// LineBlock nodes.
type htmlRenderer struct {
	html.Config
}

// NewHTMLRenderer returns a new sectionHTMLRenderer
func NewHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &htmlRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *htmlRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindLineBlock, r.renderLineBlock)
	reg.Register(KindLineBlockItem, r.renderLineBlockItem)
}

func (r *htmlRenderer) renderLineBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	const tag = "div"
	if entering {
		_ = w.WriteByte('<')
		_, _ = w.WriteString(tag)
		if node.Attributes() != nil {
			html.RenderAttributes(w, node, html.GlobalAttributeFilter)
		} else {
			_, _ = w.WriteString(` class="line-block"`)
		}
	} else {
		_, _ = w.WriteString("</")
		_, _ = w.WriteString(tag)
	}
	_, _ = w.WriteString(">\n")

	return ast.WalkContinue, nil
}

func (r *htmlRenderer) renderLineBlockItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString(`<div class="line">`)

		lbi := node.(*LineBlockItem)
		for i := 0; i < lbi.Padding; i++ {
			_, _ = w.WriteString(`&nbsp;`) // line indent
		}

		if lbi.Padding == 0 && lbi.ChildCount() == 0 {
			_, _ = w.WriteString(`&#8203;`) // empty line fix
		}
	} else {
		_, _ = w.WriteString("</div>\n")
	}

	return ast.WalkContinue, nil
}
