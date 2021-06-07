// Package lineblocks is a extension for the goldmark
// (http://github.com/yuin/goldmark).
//
// This extension adds support for line blocks in markdown.
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
