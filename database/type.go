package database

// Users struct is used to store the data from the user to database
type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Products struct {
	Id                 int    `json:"id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductType        string `json:"product_type"`
	Price              string `json:"price"`
}

// Login struct is used to login in the system
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Response struct is used for the response
type Response struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Tokens struct {
	Token string `json:"token"`
}

type Order struct {
	OrderId            int    `json:"order_id"`
	CartId             int    `json:"cart_id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	ProductId          int    `json:"product_id"`
	ProductName        string `json:"product_names"`
	ProductDescription string `json:"product_description"`
	ProductType        string `json:"product_type"`
	Price              string `json:"price"`
}
