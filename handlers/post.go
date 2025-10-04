package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/erfanfs10/MrRobot-BackEnd/db"
	"github.com/erfanfs10/MrRobot-BackEnd/models"
	"github.com/erfanfs10/MrRobot-BackEnd/queries"
	"github.com/labstack/echo/v4"
)

func PostListCategory(c echo.Context) error {
	category_title := c.Param("title")
	posts := []models.Post{}
	query := queries.BuildPostListQuery("category", category_title)
	err := db.DB.Select(&posts, query)
	if err != nil {
		errText := fmt.Sprintf("can not get post wish category title: %v", category_title)
		c.Set("err", errText)
	}
	return c.JSON(http.StatusOK, posts)

}

func PostListTag(c echo.Context) error {
	tag_title := c.Param("title")
	posts := []models.Post{}
	query := queries.BuildPostListQuery("tag", tag_title)
	err := db.DB.Select(&posts, query)
	if err != nil {
		errText := fmt.Sprintf("can not get post wish tag title: %v", tag_title)
		c.Set("err", errText)
	}
	return c.JSON(http.StatusOK, posts)

}

func PostDetail(c echo.Context) error {
	post_slug := c.Param("slug")

	post := models.Post{}
	relatedPosts := []models.Post{}

	q := queries.BuildPostListQuery("detail", post_slug)
	err := db.DB.Get(&post, q)
	if err != nil {
		errText := fmt.Sprintf("can not get post wish slug: %v", post_slug)
		c.Set("err", errText)
	}

	// convert raw json message to postCategory struct to
	// extract the first category title item
	var postCategories []models.PostCategory
	err = json.Unmarshal(*post.Categories, &postCategories)
	if err != nil {
		errText := fmt.Sprintf("can not get first category of current post: %s", post_slug)
		c.Set("err", errText)
	}

	// only get related post if current post has a category
	if len(postCategories) >= 1 {
		query := queries.BuildPostListQuery("category", *postCategories[0].Title)
		err = db.DB.Select(&relatedPosts, query)
		if err != nil {
			errText := fmt.Sprintf("can not get related posts wish slug: %v", post_slug)
			c.Set("err", errText)
		}
	}

	response := struct {
		Post         models.Post   `json:"post"`
		RelatedPosts []models.Post `json:"related_posts"`
	}{}

	response.Post = post
	response.RelatedPosts = relatedPosts

	return c.JSON(http.StatusOK, response)

}

func PostRandom(c echo.Context) error {
	posts := []models.Post{}
	q := queries.BuildPostListQuery("random", "")
	err := db.DB.Select(&posts, q)
	if err != nil {
		c.Set("err", "can not get random posts")
	}
	return c.JSON(http.StatusOK, posts)

}
