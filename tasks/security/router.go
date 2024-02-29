package security

import (
	_ "embed"
	"fmt"
	"github.com/fingann/swat/pkg/flags"
	"net/http"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/security", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(SecurityHintPage()))
			return
		}
		c.HTML(http.StatusOK, "", SecurityHintPage())
	})

	r.GET("/.well-known/security.txt", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf(`-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

Mailto:security@geodata.no
Contact: https://geodata.no
Expires: 2024-03-14T00:00:00.000Z
Acknowledgments: %s 
Preferred-Languages: en, fr, de
Canonical: https://securitytxt.org/.well-known/security.txt
-----BEGIN PGP SIGNATURE-----

iHUEARYKAB0WIQSsP2kEdoKDVFpSg6u3rK+YCkjapwUCY9qRaQAKCRC3rK+YCkja
pwALAP9LEHSYMDW4h8QRHg4MwCzUdnbjBLIvpq4QTo3dIqCUPwEA31MsEf95OKCh
MTHYHajOzjwpwlQVrjkK419igx4imgk=
=KONn
-----END PGP SIGNATURE-----`, flags.SecurityFlag)))
	})

	return nil
}
