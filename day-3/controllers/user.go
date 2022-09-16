package controllers

import (
	"day-2/lib"
	"day-2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserList(c echo.Context) error {
	response := new(models.ResponsePagination)

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	users, err := lib.GetAllUser(c.QueryParam("q"), limit, offset)

	if err != nil {
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Sukses melihat data"
	response.Data = users
	response.Limit = limit
	response.Offset = offset
	return c.JSON(http.StatusOK, response)

}

func UserStore(c echo.Context) error {
	user := new(models.Users)
	c.Bind(user)

	response := new(models.Response)
	checkUser, _ := lib.GetOneUserByEmail(user.Email) // method get by email

	if checkUser.Email == user.Email {
		response.Message = "Email sudah pernah terdaftar"
		return c.JSON(http.StatusBadRequest, response)
	}

	if lib.CreateUser(user) != nil { // method create user
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "Sukses create data"
	response.Data = user
	return c.JSON(http.StatusOK, response)
}

func UserShow(c echo.Context) error {
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

func UserUpdate(c echo.Context) error {
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

func UserDelete(c echo.Context) error {
	user, _ := lib.GetOneUserById(c.Param("id"))
	response := new(models.Response)

	if lib.DeleteUser(&user) != nil {
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "Sukses menghapus data user"
	return c.JSON(http.StatusOK, response)
}
