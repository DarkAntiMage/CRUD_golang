package main

import (
	"echofw/config"
	"echofw/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

//----------
// Handlers
//----------

func main() {
	config.ConnectDB()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	//e.GET("/users", getAllUsers)
	e.POST("/users", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		response := new(Response)
		fmt.Println(user)
		if user.CreateUser() != nil {
			response.ErrorCode = 10
			response.Message = "Gagal create data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses create data user"
			response.Data = *user
		}
		fmt.Println(user)
		return c.JSON(http.StatusOK, response)
	})
	e.GET("/users", func(c echo.Context) error {
		response := new(Response)
		users, err := model.GetAllUsers() // method get all
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Gagal melihat data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses melihat data user"
			response.Data = users
		}
		return c.JSON(http.StatusOK, response)
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		response := new(Response)
		id, _ := strconv.Atoi(c.Param("id"))
		if user.UpdateUser(id) != nil { // method update user
			response.ErrorCode = 10
			response.Message = "Gagal update data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses update data user"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user, _ := model.GetUser(id) // method get by email
		response := new(Response)

		if user.DeleteUser() != nil { // method update user
			response.ErrorCode = 10
			response.Message = "Gagal menghapus data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses menghapus data user"
		}
		return c.JSON(http.StatusOK, response)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
