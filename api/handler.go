package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handlerGetPostList(c echo.Context) error {
	var posts []Post
	limit := c.QueryParam("limit")

	query := `
		SELECT 
			id, title, content, create_at
		FROM
			posts
		ORDER BY
			create_at DESC
		LIMIT
			?`
	if err := db.Select(&posts, query, limit); err != nil {
		return c.String(http.StatusNotFound, "記事がありません")
	}
	return c.JSON(http.StatusOK, posts)
}

func handlerGetPost(c echo.Context) error {
	var post Post
	postID := c.Param("post_id")
	query := `
		SELECT 
			id, title, content, create_at 
		FROM
			posts
		WHERE
			id = ?`
	if err := db.Get(&post, query, postID); err != nil {
		return c.String(http.StatusNotFound, "記事がありません")
	}
	return c.JSON(http.StatusOK, post)
}

func handlerGetComments(c echo.Context) error {
	var comments []Comment
	postID := c.Param("post_id")
	limit := c.QueryParam("limit")
	query := `
		SELECT
			id, poster, content, post_id, create_at
		FROM
			comments
		WHERE
			post_id = ?
		ORDER BY
			create_at DESC
		LIMIT
			?`
	if err := db.Select(&comments, query, postID, limit); err != nil {
		return c.String(http.StatusNotFound, "コメントがありません")
	}
	return c.JSON(http.StatusOK, comments)
}

func handlerCreatePost(c echo.Context) error {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	query := `
		INSERT INTO
			posts (title, content)
		VALUES
			(:title, :content)`
	_, err := db.NamedExec(query, req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "")
}

func handlerCreateComment(c echo.Context) error {
	var req struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	postID := c.Param("post_id")
	query := `
		INSERT INTO
			comments (poster, content, post_id)
		VALUES
			(?, ?, ?)`
	_, err := db.Queryx(query, req.Name, req.Content, postID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, "")
}
