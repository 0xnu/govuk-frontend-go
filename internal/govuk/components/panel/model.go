package panel

import "html/template"

type ActionItem struct {
	Text       string
	Type       string
	Href       string
	Classes    string
	Attributes template.HTMLAttr
}

type Actions struct {
	Items      []ActionItem
	Classes    string
	Attributes template.HTMLAttr
}

type Model struct {
	TitleText    string
	TitleHTML    template.HTML
	HeadingLevel int
	Text         string
	HTML         template.HTML
	Classes      string
	Attributes   template.HTMLAttr
	Actions      *Actions
}
