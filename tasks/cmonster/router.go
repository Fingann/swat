package cmonster

import (
	_ "embed"
	"github.com/fingann/swat/pkg/flags"
	"net/http"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

func SetCookieIfNotSet(c *gin.Context, name, value string) {
	_, err := c.Cookie(name)
	if err != nil {
		c.SetCookie(name, value, 0, "", "", false, true)
	}
}

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/cmonster", func(c *gin.Context) {
		SetCookieIfNotSet(c, "username", "cato")
		SetCookieIfNotSet(c, "admin", "0")
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(CmonsterPage()))
			return
		}
		c.HTML(http.StatusOK, "", CmonsterPage())
	})

	r.GET("/cmonster/admin", func(c *gin.Context) {
		admin, err := c.Cookie("admin")
		if err != nil {
			c.HTML(http.StatusOK, "", public.Base(CmonsterAdminInfo("")))
			return
		}
		if admin == "1" {
			c.HTML(http.StatusOK, "", CmonsterAdminInfo(flags.CmonsterFlag))
			return
		}
		c.HTML(http.StatusOK, "", CmonsterAdminInfo(""))
	})

	return nil
}
