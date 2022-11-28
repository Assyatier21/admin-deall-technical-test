package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/assyatier21/admin-deall-technical-test/config"
	"github.com/assyatier21/admin-deall-technical-test/database"
	e "github.com/assyatier21/admin-deall-technical-test/entity"
	"github.com/assyatier21/admin-deall-technical-test/models"
	"github.com/assyatier21/admin-deall-technical-test/utils"

	"github.com/labstack/echo/v4"
)

func CreateRegisteredUser(c echo.Context) (err error) {
	var (
		registeredUser e.RegisteredUser
		rows           sql.Result
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

	query = fmt.Sprintf(database.InsertUser, user.Id, user.Username, utils.Hash_256(user.Password), user.RoleId, user.Token)
	rows, err = config.DB.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	id, _ := rows.LastInsertId()

	registeredUser.Id = id
	registeredUser.Username = user.Username
	registeredUser.RoleId = user.RoleId

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
func GetRegisteredUser(c echo.Context) (err error) {
	var (
		rows *sql.Rows
	)

	query = fmt.Sprintf(database.GetRegisteredUser)
	rows, err = config.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		var temp = e.RegisteredUser{}
		if err := rows.Scan(&temp.Id, &temp.Username, &temp.RoleId); err != nil {
			log.Fatal(err)
		}
		data = append(data, temp)
	}

	if len(data) > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusOK, "data not found", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func UpdateRegisteredUser(c echo.Context) (err error) {
	var (
		id       int
		username string
		password string
		role_id  int
		token    string
	)

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err = strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	username = c.FormValue("username")
	password = c.FormValue("password")
	token = c.FormValue("token")

	role_id, err = strconv.Atoi(c.FormValue("role_id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "role_id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	if username == "" {
		res := models.SetResponse(http.StatusBadRequest, "username can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	if !utils.IsValidAlphaNumeric(username) {
		res := models.SetResponse(http.StatusBadRequest, "username can only accepted alphabet and numeric", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	if password == "" {
		res := models.SetResponse(http.StatusBadRequest, "password can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	if token == "" {
		res := models.SetResponse(http.StatusBadRequest, "token can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.UpdateRegisteredUser, username, utils.Hash_256(password), role_id, token, id)
	rows, err := config.DB.Exec(query)
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "failed to update registered user", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "no data was updated", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func DeleteRegisteredUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "article id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.DeleteRegisteredUser, id)
	rows, err := config.DB.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "no user was deleted", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func GetUserPoints(c echo.Context) (err error) {
	var (
		rows *sql.Rows
	)
	query = fmt.Sprintf(database.GetUsersPoints)
	rows, err = config.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		var temp = e.UserPoints{}
		if err := rows.Scan(&temp.Id, &temp.Points); err != nil {
			log.Fatal(err)
		}
		data = append(data, temp)
	}

	if len(data) > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusOK, "data not Found", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func ResetUserPoints(c echo.Context) (err error) {
	var (
		rows sql.Result
		id   int
	)
	c.Bind(&user)

	if c.FormValue("id") == "" {
		res := models.SetResponse(http.StatusBadRequest, "user id can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		id, err = strconv.Atoi(c.FormValue("id"))
		if err != nil {
			res := models.SetResponse(http.StatusBadRequest, "user id must be an integer", []interface{}{})
			return c.JSON(http.StatusOK, res)
		}
	}

	query = fmt.Sprintf(database.ResetUserPoints, id)
	rows, err = config.DB.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "user id not exist", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
