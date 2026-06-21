package gin

import (
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/0xnu/govuk-frontend-go/internal/govuk/assets"
)

var SupportedGovUKFrontendVersions = []string{"6.0.0", "6.1.0", "6.2.0"}

type Adapter struct {
	AssetsPrefix string
	AssetsRoot   string
}

func New() *Adapter {
	return &Adapter{
		AssetsPrefix: "/assets",
		AssetsRoot:   path.Join("build", "govuk-frontend", "current"),
	}
}

func (a *Adapter) SetGovUKFrontendVersion(version string) error {
	root := path.Join("build", "govuk-frontend", "v"+version)

	sub, err := assets.Subdir(root)
	if err != nil {
		return fmt.Errorf("govuk-frontend assets version %q not found: %w", version, err)
	}

	if _, err := fs.Stat(sub, "govuk-frontend.min.css"); err != nil {
		return fmt.Errorf("govuk-frontend assets version %q is missing expected files: %w", version, err)
	}

	a.AssetsRoot = root
	return nil
}

func (a *Adapter) Mount(r *gin.Engine) {
	r.GET(path.Join(a.AssetsPrefix, "/*filepath"), a.assetsHandler())
	r.GET("/favicon.ico", a.faviconHandler())
}

func (a *Adapter) assetsHandler() gin.HandlerFunc {
	sub, err := assets.Subdir(a.AssetsRoot)
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(sub))

	return func(c *gin.Context) {
		fp := strings.TrimPrefix(c.Param("filepath"), "/")
		fp = strings.TrimPrefix(path.Clean("/"+fp), "/")
		if strings.HasPrefix(fp, "..") {
			c.Status(http.StatusNotFound)
			return
		}

		servePath := fp
		if _, err := fs.Stat(sub, servePath); err != nil {
			alt := path.Join("assets", fp)
			if _, err := fs.Stat(sub, alt); err != nil {
				c.Status(http.StatusNotFound)
				return
			}
			servePath = alt
		}

		r := c.Request.Clone(c.Request.Context())
		r.URL.Path = "/" + servePath
		fileServer.ServeHTTP(c.Writer, r)
	}
}

func (a *Adapter) faviconHandler() gin.HandlerFunc {
	sub, err := assets.Subdir(a.AssetsRoot)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		for _, fp := range []string{"assets/images/favicon.ico", "images/favicon.ico"} {
			if _, err := fs.Stat(sub, fp); err == nil {
				c.FileFromFS(fp, http.FS(sub))
				return
			}
		}

		c.Status(http.StatusNotFound)
	}
}
