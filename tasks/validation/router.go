package validation

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/fingann/swat/pkg/flags"
	"net/http"
	"strconv"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

var Flag = "flag{REEVALUATE_INPUT_BACKEND}"

var AvailableCash = 800

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/validation", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(ValidatePage()))
			return
		}
		c.HTML(http.StatusOK, "", ValidatePage())
	})

	r.GET("/validation/api/available", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprint((AvailableCash)))
	})

	//@Item("id1","Flag Item", "/path/to/flag-item.jpg", "1000.00", "This flag item holds the key to victory. Can you afford its price?")
	//        @Item("id2","Lesser Item", "/path/to/lesser-item.jpg", "10.00", "This item might not seem like much, but looks can be deceiving.")

	r.GET("/validation/products", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", Products(ProductsInfo))
	})

	r.POST("/validation/api/purchase", func(c *gin.Context) {
		var products []Purchase
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid request: %w", err)})
			return
		}
		var total int
		for _, p := range products {
			for _, product := range ProductsInfo {
				if product.Id == p.Id {
					price, err := strconv.Atoi(product.Price)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid price: %w", err)})
						return
					}
					total += price * p.Quantity
				}
			}
		}
		if total > AvailableCash {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough cash available"})
			return
		}
		AvailableCash -= total
		// Check if the flag item is in the purchase
		for _, product := range products {
			if product.Id == "id1" && product.Quantity > 0 {
				c.JSON(http.StatusOK, gin.H{"message": "Purchase successful, " + flags.ValidationFlag})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "Purchase successful"})
	})
	return nil
}

type Purchase struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

//go:embed Security.webp
var TheFlagPic []byte

//go:embed smile.webp
var Smile []byte

var ProductsInfo = []ProductData{
	{
		Id:          "id1",
		Name:        "The Flag",
		Price:       "1000",
		Description: "This flag item holds the key to victory. Can you afford its price?",
		ImageBase64: base64.StdEncoding.EncodeToString(TheFlagPic),
	},
	{
		Id:          "id2",
		Name:        "The Smile",
		Price:       "10",
		Description: "This item might not seem like much, but atlest it is never negative.",
		ImageBase64: base64.StdEncoding.EncodeToString(Smile),
	},
}
