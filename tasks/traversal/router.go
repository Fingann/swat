package xss

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
	
)
var Flag ="flag{try_the_admin_user}"

func RegisterRoutes(r *gin.Engine) error {
	db,err:= SetupDb()
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	r.GET("/traversal", func(c *gin.Context) {
		// Refresh the whole page if the header is not set
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Page()))
			return
		}

		c.HTML(http.StatusOK, "", Page() )
	})
	r.GET("/traversal/comments", func(c *gin.Context) {
		
		comments, err := GetComments(db)
		if err != nil {
			return
		}

		c.HTML(http.StatusOK, "", CommentRows(comments) )
	})

	

	r.POST("/task2/comments", func(c *gin.Context) {
		user,_ := c.GetPostForm("user")
		comment,_ := c.GetPostForm("comment")
		if err:= CreateComment(db, comment, user); err != nil {
			// ignore and fetch comments
		}
			comments,err:= GetComments(db)
			if err != nil {
				return
			}
			c.HTML(http.StatusOK, "", CommentRows(comments))
			go func (userAgent string)  {
				time.Sleep(4 * time.Second)	
				if !strings.Contains(userAgent, "surf") {
					VisitWebsiteWithV8("http://localhost:8081/task2?admin=true")
					VisitWebsite("http://localhost:8081/task2?admin=true")
				}
			}(c.Request.UserAgent())
					
	})
	return nil
}

//go:embed fs
var RootFs embed.FS