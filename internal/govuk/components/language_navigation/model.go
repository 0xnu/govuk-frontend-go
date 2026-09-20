package language_navigation

import "html/template"

type Item struct {
	Text                    string
	HTML                    template.HTML
	Lang                    string
	HrefLang                string
	Dir                     string
	Href                    string
	Current                 bool
	LanguageDescriptionText string
	Classes                 string
	Attributes              template.HTMLAttr
}

type Model struct {
	Items      []Item
	AriaLabel  string
	Classes    string
	Attributes template.HTMLAttr
}
