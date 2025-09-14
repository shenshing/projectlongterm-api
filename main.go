package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"example.com/projectlongterm-api/endpoints/blogs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// * Load .env
	err := godotenv.Load("./env/local.env")
	if err != nil {
		log.Fatal("Error load .env file.")
	} else {
		fmt.Println("Success load .env file.")
	}

	// * Connect to MySQL.
	mysql_username := os.Getenv("MYSQL_USERNAME")
	mysql_password := os.Getenv("MYSQL_PASSWORD")
	mysql_database := os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql", mysql_username+":"+mysql_password+"@/"+mysql_database)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("[Connect-MySQL] Successfully connecting to MySQL.")
	}
	defer db.Close()

	// * List all routes.
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.GET("albums/:id", getAlbumByID)
	// router.POST("albums", postAlbums);

	// Articles -- pagination
	router.GET("/blogs/paginate/:limit/:offset", func(c *gin.Context) {
		blogs.Paginate(c, db, c.Writer)
	})
	// Article -- trending
	router.GET("/blogs/paginate/trending", func(c *gin.Context) {
		blogs.Trending(c, db, c.Writer)
	})
	// Article -- get-article-by-slug
	router.GET("/blogs/get-article-by-slug/:slug", func(c *gin.Context) {
		blogs.GetArticleBySlug(c, db, c.Writer)
	})

	// * Create table
	// for index, create_table_query := range tables.Query_to_create_table {
	// 	fmt.Println(index, "Create table")
	// 	_, err := db.Exec(create_table_query)
	// 	if err != nil {
	// 		log.Printf("Error creating table at index %d: %v", index, err)
	// 	} else {
	// 		fmt.Printf("Successfully executed table creation query at index %d\n", index)
	// 	}
	// }

	// * ALter table (add, remove columns)
	// for index, alter_table_query := range tables.Update_table_string {
	// 	fmt.Println(index, " Alter table ")
	// 	_, err := db.Exec(alter_table_query)
	// 	if err != nil {
	// 		log.Printf("Error altering table at index %d, %v", index, err)
	// 	} else {
	// 		fmt.Printf("Successfully altering table query at index %d\n", index)
	// 	}
	// }

	// * Insert default article.
	// tables.Insert_articles(db)

	// port := os.Getenv("PORT");
	// fmt.Println("PORT is ", port)

	// fmt.Println("Main application running.")
	// utils.HelperFunction()

	// * Run web server.
	router.Run("localhost:9090")

}
