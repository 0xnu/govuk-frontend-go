package generic_header

import "html/template"

type Model struct {
	URL              string
	LogoText         string
	LogoHTML         template.HTML
	ContainerClasses string
	Classes          string
	Attributes       template.HTMLAttr
}
