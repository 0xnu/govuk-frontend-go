package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0xnu/govuk-frontend-go/internal/app/config"
	govukgin "github.com/0xnu/govuk-frontend-go/pkg/govuk/gin"
	govuktemplate "github.com/0xnu/govuk-frontend-go/pkg/govuk/template"
)

func main() {
	r := gin.Default()
	cfg := config.Load()

	t, err := govuktemplate.Parse()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	g := govukgin.New()
	if err := g.SetGovUKFrontendVersion("6.4.0"); err != nil {
		panic(err)
	}
	g.Mount(r)

	r.GET("/", handleIndex)

	if err := r.Run(cfg.Addr()); err != nil {
		panic(err)
	}
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "page", gin.H{
		"ServiceName": "Example service",
	})
}
