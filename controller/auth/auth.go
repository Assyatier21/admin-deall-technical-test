package auth

import (
	"admin/config"
	"admin/database"
	e "admin/entity"
	"admin/models"
	"admin/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) (err error) {
	var (
		user  e.User
		query string
		rows  *sql.Rows
		data  []interface{}
	)

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if c.FormValue("username") == "" {
		res := models.SetResponse(http.StatusBadRequest, "username can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	if c.FormValue("password") == "" {
		res := models.SetResponse(http.StatusBadRequest, "password can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.GetUserByUsernamePass, user.Username, utils.Hash_256(user.Password), user.Token, user.RoleId)
	rows, err = config.DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var temp = e.User{}
		if err := rows.Scan(&temp.Id, &temp.Username, &temp.Password, &temp.RoleId, &temp.Token); err != nil {
			log.Fatal(err)
		}
		fmt.Println(temp)

		var registeredUser = e.RegisteredUser{}

		registeredUser.Id = temp.Id
		registeredUser.Username = temp.Username
		registeredUser.RoleId = temp.RoleId
		registeredUser.Token = temp.Token

		data = append(data, registeredUser)
	}

	if len(data) > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusUnauthorized, "username or password is incorrect", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func Register(c echo.Context) (err error) {
	var (
		user           e.User
		registeredUser e.RegisteredUser
		query          string
		rows           sql.Result
		data           []interface{}
	)

	c.Bind(&user)

	if c.FormValue("username") == "" {
		res := models.SetResponse(http.StatusBadRequest, "username can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else if !utils.IsValidAlphaNumeric(c.FormValue("username")) {
		res := models.SetResponse(http.StatusBadRequest, "username must only contains alphabet or numeric", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	if c.FormValue("password") == "" {
		res := models.SetResponse(http.StatusBadRequest, "password can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
	if i, _ := strconv.Atoi(c.FormValue("role_id")); i == 0 {
		res := models.SetResponse(http.StatusBadRequest, "role_id can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.InsertUser, user.Id, user.Username, utils.Hash_256(user.Password), user.RoleId, models.AccessToken)
	rows, err = config.DB.Exec(query)

	if err != nil {
		log.Println(err.Error())
		return
	}

	id, _ := rows.LastInsertId()

	registeredUser.Id = id
	registeredUser.Username = user.Username
	registeredUser.RoleId = user.RoleId
	registeredUser.Token = models.AccessToken

	data = append(data, registeredUser)

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "failed to register account", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
