package gin

import (
	"github.com/gin-gonic/gin"

	govuktemplate "github.com/0xnu/govuk-frontend-go/pkg/govuk/template"
)

func (a *Adapter) MountTemplates(r *gin.Engine) error {
	t, err := govuktemplate.Parse()
	if err != nil {
		return err
	}
	r.SetHTMLTemplate(t)
	return nil
}
