package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// // get user by id
func GetUserController(c echo.Context) error {
	// your solution here

	id, errcnv := strconv.Atoi(c.Param("id"))
	var data int
	if errcnv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{ //any interface kosong
			"status":  "Gagal",
			"massage": "id Harus Int",
		})
	}

	for i, val := range users {
		if val.Id == id {
			data = i
		}

	}
	return c.JSON(http.StatusOK, map[string]any{ //any interface kosong
		"status":   "berhasil",
		"massage":  "baca data berhasil",
		"articles": users[data],
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, errcnv := strconv.Atoi(c.Param("id"))
	var user int
	if errcnv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{ //any interface kosong
			"status":  "Gagal",
			"massage": "id Harus Int",
		})
	}

	for i, val := range users {
		if val.Id == id {
			user = i
		}

	}
	users = removearr(users, user)
	return c.JSON(http.StatusOK, map[string]any{ //any interface kosong
		"status":  "berhasil",
		"massage": "delete data berhasil",
		"user":    users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	id, errcnv := strconv.Atoi(c.Param("id"))
	var data int //index id yang akan diupdate
	if errcnv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{ //any interface kosong
			"status":  "Gagal",
			"massage": "id Harus Int",
		})
	}
	for k, v := range users {
		if v.Id == id {
			data = k
		}

	}
	user := User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error binding data" + errBind.Error(),
		})
	}
	if user.Id == 0 {
		user.Id = users[data].Id
	}

	users[data] = user

	return c.JSON(http.StatusOK, map[string]any{
		"messages": "Update user success",
		"user":     users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
func removearr(data []User, index int) []User {
	return append(data[:index], data[index+1:]...)
} // fung
