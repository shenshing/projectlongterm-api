package blogs

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/projectlongterm-api/tables"
	"example.com/projectlongterm-api/utils"
	"github.com/gin-gonic/gin"
)

type ArticleWebFormat struct {
	Title       string
	Excerpt     string
	Author      string
	Avatar      string
	Date        string
	ReadTime    string
	Image       string
	Slug        string
	Tags        string
	OriginalUrl string
	AuthorUrl   string
}

func Test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Success called.")
}

func Paginate(c *gin.Context, db *sql.DB, w http.ResponseWriter) {
	utils.EnableCors(w)
	limit_as_param := c.Param("limit")   // page size
	offset_as_param := c.Param("offset") // to skip, index started from 0.

	// For example, page 1 will have articles from index 0 to 9, and page 2 will have index from 10 - 19

	fmt.Println("limit_as_param: ", limit_as_param)
	fmt.Println("offset_as_param: ", offset_as_param)

	limit, limit_convert_err := strconv.ParseInt(limit_as_param, 10, 64)

	if limit_convert_err != nil {
		fmt.Println("Failed to convert 'limit' value!!!")
	}
	offset, offset_convert_err := strconv.ParseInt(offset_as_param, 10, 64)
	if offset_convert_err != nil {
		fmt.Println("Failed to convert 'offset' value!!!")
	}

	if limit == 0 {
		limit = 10
	}

	if offset == 0 {
		offset = 1
	}
	offset = (offset - 1) * limit

	fmt.Printf("---> Limit is %d and offset is %d", limit, offset)

	query := `SELECT * FROM articles WHERE type = 'blog' ORDER BY created_at DESC LIMIT ? OFFSET ?;`
	rows, query_err := db.Query(query, limit, offset)

	if query_err != nil {
		fmt.Println("Failed to query!!!")
	}

	articles := []tables.Article{}
	for rows.Next() {
		var article tables.Article
		if err := rows.Scan(&article.Id, &article.Lang, &article.Title, &article.Body, &article.Html, &article.Author, &article.Created_at, &article.Updated_at, &article.Read_time, &article.Author_profile_url, &article.Thumbnail_url, &article.Original_url, &article.Article_tags, &article.Type, &article.Excerpt, &article.Avatar, &article.Slug); err != nil {
			log.Fatal("Something went wrong when Scan row!!!", err)
		}
		articles = append(articles, article)
	}

	c.IndentedJSON(http.StatusOK, articles)
}

func Trending(c *gin.Context, db *sql.DB, w http.ResponseWriter) {
	utils.EnableCors(w)
	const max int = 4 // Homepage should only display 6 articles, else could be messy.

	fmt.Printf("Start query trending article. Maximum %d articles.", max)

	query := "SELECT * FROM articles WHERE type = 'blog' ORDER BY created_at DESC LIMIT ?;"
	rows, query_err := db.Query(query, max)

	if query_err != nil {
		log.Fatal("[Trending-Article] Failed to query articles!!!")
	}

	articles := []tables.Article{}
	for rows.Next() {
		var article tables.Article

		if err := rows.Scan(&article.Id, &article.Lang, &article.Title, &article.Body, &article.Html, &article.Author, &article.Created_at, &article.Updated_at, &article.Read_time, &article.Author_profile_url, &article.Thumbnail_url, &article.Original_url, &article.Article_tags, &article.Type, &article.Excerpt, &article.Avatar, &article.Slug, &article.Author_name); err != nil {
			log.Fatal("[Trending-Article] Something went wrong when Scan row!!!", err)
		}
		// fmt.Println("here it is: ", article.Created_at)
		articles = append(articles, article)
		// fmt.Println(articles)
	}

	c.IndentedJSON(http.StatusOK, articles)
}

func GetArticleBySlug(c *gin.Context, db *sql.DB, w http.ResponseWriter) {
	utils.EnableCors(w)
	slug := c.Param("slug")

	fmt.Printf("[Get-Article-By-Slug] Get article by slug %s\n.", slug)
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slug parameter is required"})
		return
	}

	query := `SELECT * FROM articles WHERE type = 'blog' AND slug = ?;`
	rows, query_err := db.Query(query, slug)

	if query_err != nil {
		fmt.Println("[Get-Article-By-Slug] Failed to query. Error: ", query_err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Check log on the backend for more information."})
	}

	articleResponse := tables.Article{}

	for rows.Next() {
		var article tables.Article
		if err := rows.Scan(&article.Id, &article.Lang, &article.Title, &article.Body, &article.Html, &article.Author, &article.Created_at, &article.Updated_at, &article.Read_time, &article.Author_profile_url, &article.Thumbnail_url, &article.Original_url, &article.Article_tags, &article.Type, &article.Excerpt, &article.Avatar, &article.Slug, &article.Author_name); err != nil {
			fmt.Println("[Get-Article-By-Slug] Something went wrong when extract query result. Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Check log on the backend for more information."})
		}
		// fmt.Println(article)
		articleResponse = article
	}

	c.IndentedJSON(http.StatusOK, articleResponse)
}
