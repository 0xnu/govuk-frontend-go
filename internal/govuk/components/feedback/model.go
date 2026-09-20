package feedback

import "html/template"

type Model struct {
	TitleText    string
	TitleHTML    template.HTML
	HeadingLevel int
	Text         string
	HTML         template.HTML
	Classes      string
	Attributes   template.HTMLAttr
}
