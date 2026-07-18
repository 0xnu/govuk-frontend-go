package panel

import (
	"html/template"
	"strings"
	"testing"
)

func TestRenderConfirmation(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/panel"}}<div class="govuk-panel govuk-panel--confirmation"><h1 class="govuk-panel__title">{{.TitleText}}</h1><div class="govuk-panel__body">{{.HTML}}</div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	html, err := Render(tpl, Model{
		TitleText: "Application complete",
		HTML:      template.HTML("Your reference number<br><strong>HDJ2123F</strong>"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(html), "Application complete") {
		t.Fatalf("expected title in output: %s", html)
	}
	if !strings.Contains(string(html), "govuk-panel--confirmation") {
		t.Fatalf("expected confirmation class in output: %s", html)
	}
}

func TestRenderInterruption(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/panel"}}<div class="govuk-panel govuk-panel--interruption"><h1 class="govuk-panel__title">{{.TitleText}}</h1><div class="govuk-panel__body">{{.HTML}}</div><div class="govuk-panel__actions"><div class="govuk-button-group">{{range .Actions.Items}}<button type="button" class="govuk-button govuk-button--inverse">{{.Text}}</button>{{end}}</div></div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	html, err := Render(tpl, Model{
		TitleText: "Is your age correct?",
		HTML:      template.HTML("<p>You entered <strong>109</strong>.</p>"),
		Classes:   "govuk-panel--interruption",
		Actions: &Actions{
			Items: []ActionItem{
				{Text: "Yes", Type: "button"},
				{Text: "No", Href: "#"},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(html), "Is your age correct?") {
		t.Fatalf("expected title in output: %s", html)
	}
	if !strings.Contains(string(html), "govuk-panel--interruption") {
		t.Fatalf("expected interruption class in output: %s", html)
	}
}

func TestRenderEmptyModel(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/panel"}}<div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Render(tpl, Model{})
	if err != nil {
		t.Fatal(err)
	}
}
