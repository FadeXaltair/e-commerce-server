package api

import (
	"e-commerce-backend/config"
	"e-commerce-backend/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductKey struct {
	ProductId int `json:"product_id"`
}

type CartKey struct {
	ProductCartId int `json:"product_cart_id"`
}

// Adding product to the cart
func AddToCart(c *gin.Context) {
	var productid ProductKey
	id, _ := c.GetQuery("user-id")
	userid, _ := strconv.Atoi(id)
	err := c.Bind(&productid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "error while getting data from body",
		})
		return
	}
	err = AddProductsToCart(userid, productid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "error while adding data in database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "order added to cart",
	})
}

// give details of the products added to cart
func CartData(c *gin.Context) {
	id, _ := c.GetQuery("user-id")
	userid, _ := strconv.Atoi(id)
	data, err := GetCartData(userid)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   false,
			"message": "error while fetching values from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":        false,
		"cart-details": data,
	})
}

// Placing the order
func OrderPlaced(c *gin.Context) {
	var productid CartKey
	id, _ := c.GetQuery("user-id")
	userid, _ := strconv.Atoi(id)
	err := c.Bind(&productid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "error while getting data from body",
		})
		return
	}

	err = AddOrder(userid, productid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "error while placing order",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "order placed successfully",
	})
}

// Listing all the products aailiable
func Products(c *gin.Context) {
	data, err := ListProducts()
	if err != nil {
		c.JSON(400, gin.H{
			"error":   false,
			"message": "error while fetching values from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}

// GIves the information of all the orders user placed
func Orders(c *gin.Context) {
	id, _ := c.GetQuery("user-id")
	userid, _ := strconv.Atoi(id)
	data, err := OrderData(userid)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   false,
			"message": "error while fetching values from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":         false,
		"orders-placed": data,
	})

}

//////----------------------------- Database queries ---------------------------//////

func ListProducts() ([]database.Products, error) {
	var data []database.Products
	err := config.DB.Raw(`select * from products.all_products`).Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func AddProductsToCart(userid int, data ProductKey) error {
	err := config.DB.Exec(`insert into products.carts (user_id, product_id)
	values (?,?)`, userid, data.ProductId).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func AddOrder(userid int, data CartKey) error {
	err := config.DB.Exec(`insert into products.order_purchased (user_id, order_in_cart_id)
	values (?,?)`, userid, data.ProductCartId).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func OrderData(userid int) ([]database.Order, error) {
	var data []database.Order
	err := config.DB.Raw(`SELECT x.id as order_id, c.id as cart_id,u."name" ,u.email, ap.id as product_id, ap.product_name ,ap.product_description ,ap.product_type ,ap.price  FROM products.order_purchased x
	left join products.carts c 
	on x.order_in_cart_id  = c.id 
	left  join products.users u 
	on x.user_id = u.id 
	left join products.all_products ap 
	on c.product_id = ap.id 
	where x.user_id  =?`, userid).Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func GetCartData(userid int) ([]database.Products, error) {
	var data []database.Products
	err := config.DB.Raw(`SELECT x.product_id as id ,ap.product_name ,ap.product_description ,ap.product_type ,ap.price  FROM products.carts x
	left join products.all_products ap 
	on x.product_id = ap.id 
	where x.user_id =?`, userid).Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}
