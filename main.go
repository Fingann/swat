package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/fingann/swat/pkg/renderer"
	"github.com/fingann/swat/public"
	"github.com/fingann/swat/tasks/hardcoded"
	"github.com/fingann/swat/tasks/injection"
	"github.com/fingann/swat/tasks/traversal"
	"github.com/fingann/swat/tasks/validation"

	"github.com/fingann/swat/tasks/xss"
	"github.com/gin-gonic/gin"
)

//go:embed dist
var distFolder embed.FS

func main() {
	r := gin.Default()
	r.HTMLRender = &renderer.Templ{}

	r.StaticFS("/static/", http.FS(distFolder))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", public.Index("") )
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", public.IndexPage("") )
	})
	if err:= injection.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}
	if err:= xss.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err:= traversal.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err:= hardcoded.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err:= validation.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	r.Run(":8081")
}
