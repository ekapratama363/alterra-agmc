package controllers

import (
	"day-2/lib"
	"day-2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BookList(c echo.Context) error {
	response := new(models.Response)

	books := lib.GetAllBook()

	response.Message = "Sukses melihat data"
	response.Data = books
	return c.JSON(http.StatusOK, response)
}

func BookStore(c echo.Context) error {
	book := new(models.Books)
	c.Bind(book)

	response := new(models.Response)

	if lib.CreateBook(book) != nil { // method create book
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "Sukses create data"
	response.Data = book
	return c.JSON(http.StatusOK, response)
}

func BookShow(c echo.Context) error {
	user, err := lib.GetOneUserById(c.Param("id"))
	response := new(models.Response)

	if err != nil {
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "Sukses melihat data"
	response.Data = user
	return c.JSON(http.StatusOK, response)
}

func BookUpdate(c echo.Context) error {
	user := new(models.Users)
	c.Bind(user)
	response := new(models.Response)

	if user.Email != "" {
		response.Message = "Email tidak boleh diupdate"
		return c.JSON(http.StatusInternalServerError, response)
	}

	if lib.UpdateUser(c.Param("id"), user) != nil { // method update user
		response.Message = "Gagal update data"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "Sukses update data"
	response.Data = *user
	return c.JSON(http.StatusOK, response)
}

func BookDelete(c echo.Context) error {
	user, _ := lib.GetOneUserById(c.Param("id"))
	response := new(models.Response)

	if lib.DeleteUser(&user) != nil {
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "Sukses menghapus data user"
	return c.JSON(http.StatusOK, response)
}
