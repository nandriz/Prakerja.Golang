package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CoffeeProduct struct {
	Foto  string `json:"foto"`
	Nama  string `json:"nama"`
	Size  string `json:"size"`
	Harga uint   `json:"harga"`
	Link  string `json:"link"`
}

// not working
type CoffeeProducts struct {
	Arabica []CoffeeProduct `gorm:"arabica"`
	Robusta []CoffeeProduct `json:"robusta"`
	Nonkopi []CoffeeProduct `json:"nonkopi"`
}

// not working -end

type User struct {
	Id    int    `json:"id" gorm:"primaryKey autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BaseResponse struct {
	Status  bool        `json:"ststus"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	InitDatabase()
	e := echo.New()

	// Routing
	e.GET("/users", GetUsersController)
	e.POST("/users", AddUsersController)
	e.GET("/CoffeeProduct", GetCoffeeProduct)
	e.POST("/CoffeeProduct", AddCoffeeProduct)
	// e.GET("/CoffeeProducts", GetCoffeeProducts)
	// e.POST("//add-coffee-product", AddCoffeeProductHandler)
	e.GET("/users/:id", GetUsersDetailController)
	e.POST("/login", LoginController)
	e.Start(":8000")

}

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/prakerja10?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal mengambil data")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&CoffeeProduct{})
	// DB.AutoMigrate(&CoffeeProducts{})
}

// not working

func AddCoffeeProductHandler(c echo.Context) error {
	var data map[string][]CoffeeProduct
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  false,
			Message: "Invalid JSON data",
			Data:    nil,
		})
	}

	// Extract the "arabica" data from the JSON and insert it into the database
	arabicaProducts := data["arabica"]

	// Insert 'arabicaProducts' into the MySQL database using GORM or your preferred database library

	// Return a response indicating success
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Coffee products added successfully",
		Data:    arabicaProducts,
	})
}

// not working end

func AddCoffeeProduct(c echo.Context) error {
	// Parse the JSON data from the request body into a struct
	var coffeeProduct CoffeeProduct
	c.Bind(&coffeeProduct)
	result := DB.Create(&coffeeProduct)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  false,
			Message: "Invalid JSON data",
			Data:    nil,
		})
	}

	// Process the coffee product data as needed
	// ...

	// Return a response indicating success
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Coffee product added successfully",
		Data:    coffeeProduct,
	})
}

func GetCoffeeProduct(c echo.Context) error {
	// country := c.QueryParam("country")

	var coffeeProduct []CoffeeProduct
	result := DB.Find(&coffeeProduct)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  false,
			Message: "Gagal",
			Data:    coffeeProduct,
		})
	}

	// If there are no errors, you can proceed to use the 'users' slice here

	// users = append(users, User{1, country, "A"})
	// users = append(users, User{1, "B", "B"})
	// users = append(users, User{1, "C", "C"})

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Berhasil Diambil",
		Data:    coffeeProduct,
	})
}

// not working
func AddCoffeeProducts(c echo.Context) error {
	// Parse the JSON data from the request body into a struct
	var coffeeProducts CoffeeProducts
	c.Bind(&coffeeProducts)
	result := DB.Create(&coffeeProducts)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  false,
			Message: "Invalid JSON data",
			Data:    nil,
		})
	}

	// Process the coffee product data as needed
	// ...

	// Return a response indicating success
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Coffee product added successfully",
		Data:    coffeeProducts,
	})
}

func GetCoffeeProducts(c echo.Context) error {
	// country := c.QueryParam("country")

	var coffeeProducts []CoffeeProducts
	result := DB.Find(&coffeeProducts)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  false,
			Message: "Gagal",
			Data:    coffeeProducts,
		})
	}

	// If there are no errors, you can proceed to use the 'users' slice here

	// users = append(users, User{1, country, "A"})
	// users = append(users, User{1, "B", "B"})
	// users = append(users, User{1, "C", "C"})

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Berhasil Diambil",
		Data:    coffeeProducts,
	})
}

// not working end

func AddUsersController(c echo.Context) error {

	var user User
	c.Bind(&user)

	result := DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  false,
			Message: "Gagal Menambah Data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Data Ditambahkan",
		Data:    user,
	})

}

func LoginController(c echo.Context) error {

	var userLogin UserLogin
	c.Bind(&userLogin)

	//email := c.FormValue("email")
	//password := c.FormValue("password")

	if userLogin.Email == "admin@najasa.id" && userLogin.Password == "123QWE" {
		// users := User{1, email, password}
		return c.JSON(http.StatusOK, BaseResponse{
			Status:  true,
			Message: "Berhasil Diambil",
			Data:    userLogin,
		})
	}
	return c.JSON(http.StatusUnauthorized, BaseResponse{
		Status:  false,
		Message: "Gagal Diambil",
		Data:    nil,
	})
}

// fungsi Param
func GetUsersDetailController(c echo.Context) error {
	id := c.Param("id")
	var users User = User{1, id, "A"}

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Berhasil Diambil",
		Data:    users,
	})
}

// end fungsi Param

func GetUsersController(c echo.Context) error {
	// country := c.QueryParam("country")

	var users []User
	result := DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  false,
			Message: "Gagal",
			Data:    users,
		})
	}

	// If there are no errors, you can proceed to use the 'users' slice here

	// users = append(users, User{1, country, "A"})
	// users = append(users, User{1, "B", "B"})
	// users = append(users, User{1, "C", "C"})

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Berhasil Diambil",
		Data:    users,
	})
}
