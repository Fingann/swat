package main

import (
	"embed"
	"fmt"
	"github.com/fingann/swat/pkg/flags"
	"github.com/fingann/swat/tasks/cmonster"
	"github.com/fingann/swat/tasks/security"
	"io"
	"net/http"
	"os"

	"github.com/fingann/swat/pkg/renderer"
	"github.com/fingann/swat/public"
	"github.com/fingann/swat/tasks/hardcoded"
	"github.com/fingann/swat/tasks/injection"
	"github.com/fingann/swat/tasks/traversal"
	"github.com/fingann/swat/tasks/validation"

	"sync"

	"github.com/fingann/swat/tasks/xss"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
)

//go:embed dist
var distFolder embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	if _, ok := os.LookupEnv("DEBUG"); ok {
		r.Use(gin.Logger())
	} else {
		r.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: io.Discard}))
	}

	r.HTMLRender = &renderer.Templ{}

	r.StaticFS("/static/", http.FS(distFolder))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", public.Base(public.IndexPage("")))
	})

	r.GET("/index", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(public.IndexPage("")))
			return
		}
		c.HTML(http.StatusOK, "", public.IndexPage(""))
	})

	r.POST("/flag/submit", func(c *gin.Context) {
		flagValue := c.PostForm("flag")
		taskName, ok := flags.CheckFlag(flagValue)
		if ok {
			c.JSON(http.StatusOK, gin.H{"name": taskName})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Flag"})
	})

	if err := injection.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}
	if err := xss.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err := traversal.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err := hardcoded.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err := validation.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err := cmonster.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	if err := security.RegisterRoutes(r); err != nil {
		fmt.Printf("failed to register routes: %v", err)
		return
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, r *gin.Engine) {
		defer wg.Done()
		fmt.Printf("running server on :8081\n")
		if err := r.Run(":8081"); err != nil {
			fmt.Printf("failed to run server: %v", err)
			return
		}
	}(wg, r)
	go xss.VisitWebsiteWithRod("http://localhost:8081/xss")

	browser.OpenURL("http://localhost:8081")

	wg.Wait()
}
