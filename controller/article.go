package controller

import (
	"github.com/assyatier21/admin-deall-technical-test/config"
	"github.com/assyatier21/admin-deall-technical-test/database"
	e "github.com/assyatier21/admin-deall-technical-test/entity"
	"github.com/assyatier21/admin-deall-technical-test/models"
	"github.com/assyatier21/admin-deall-technical-test/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	article e.Article
	user    e.User
	query   string
	data    []interface{}
)

func InsertArticle(c echo.Context) (err error) {
	var (
		rows       sql.Result
		title      string
		content    string
		created_by int
	)
	currentTime := time.Now()
	c.Bind(&article)

	article.CreatedAt = fmt.Sprintf("%d-%d-%d", currentTime.Year(), currentTime.Month(), currentTime.Day())

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := models.SetResponse(http.StatusBadRequest, "title can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		title = c.FormValue("title")
	}

	if c.FormValue("content") == "" {
		res := models.SetResponse(http.StatusBadRequest, "content can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		content = c.FormValue("content")
	}

	if c.FormValue("created_by") == "" {
		res := models.SetResponse(http.StatusBadRequest, "created_by can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		if i, _ := strconv.Atoi(c.FormValue("created_by")); i == 0 {
			res := models.SetResponse(http.StatusBadRequest, "created_by can't be zero", []interface{}{})
			return c.JSON(http.StatusOK, res)
		} else {
			created_by = i
		}
	}

	query = fmt.Sprintf(database.InsertArticle, article.Id, title, content, article.CreatedAt, created_by, 0)
	rows, err = config.DB.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	id, _ := rows.LastInsertId()
	article.Id = int(id)
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")
	article.CreatedBy = created_by

	data = append(data, article)

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "failed to insert article", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func GetArticleByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "article id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.GetArticleById, id)
	err = config.DB.QueryRow(query).Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt, &article.CreatedBy, &article.Points)
	if err != nil {
		log.Println(err.Error())
	}

	data = append(data, article)

	if len(data) > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "no data was found", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func UpdateArticleByID(c echo.Context) (err error) {
	var (
		id      int
		title   string
		content string
		points  int
	)

	err = c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if c.FormValue("id") == "" {
		res := models.SetResponse(http.StatusBadRequest, "id can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		id, err = strconv.Atoi(c.FormValue("id"))
		if err != nil {
			res := models.SetResponse(http.StatusBadRequest, "id must be an integer", []interface{}{})
			return c.JSON(http.StatusOK, res)
		}
	}

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := models.SetResponse(http.StatusBadRequest, "title can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		title = c.FormValue("title")
	}

	if c.FormValue("content") == "" {
		res := models.SetResponse(http.StatusBadRequest, "content can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		content = c.FormValue("content")
	}

	if c.FormValue("points") == "" {
		res := models.SetResponse(http.StatusBadRequest, "points can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		points, err = strconv.Atoi(c.FormValue("points"))
		if err != nil {
			res := models.SetResponse(http.StatusBadRequest, "points must be an integer", []interface{}{})
			return c.JSON(http.StatusOK, res)
		}
	}
	query = fmt.Sprintf(database.UpdateArticleById, title, content, points, id)
	_, err = config.DB.Exec(query)
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "failed to update article", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	var resp = e.UpdatedArticle{}
	resp.Id = id
	resp.Title = title
	resp.Content = content
	resp.Points = points

	data = append(data, resp)
	res := models.SetResponse(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, res)
}
func DeleteArticleByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "article id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.DeleteArticleById, id)
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
		res := models.SetResponse(http.StatusBadRequest, "no data was deleted", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func UpdateArticlePointById(c echo.Context) (err error) {
	var (
		articlePoints e.ArticlePoints
		article       e.Article
		query         string
		data          []interface{}
		rows          sql.Result
	)

	c.Bind(&article)

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "article id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	points, err := strconv.Atoi(c.FormValue("points"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "points must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	articlePoints.Id = id
	articlePoints.Points = points
	data = append(data, articlePoints)

	query = fmt.Sprintf(database.UpdateArticlePoints, points, id)
	rows, err = config.DB.Exec(query)
	if err != nil {
		log.Println(err.Error())
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
