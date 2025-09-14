package main

// import "github.com/gin-gonic/gin"
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/projectlongterm-api/endpoints/blogs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, albums)
	c.IndentedJSON(http.StatusOK, albums)
}

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range albums {
// 		// fmt.Println(index)
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

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
	// database_type := "mysql";
	// fmt.Println("mysql_username ", mysql_username);
	// fmt.Println("mysql_password ", mysql_password);
	// fmt.Println("mysql_database ", mysql_database);

	// db, err := sql.Open("mysql", "root:12345678@/projectlongterm");
	db, err := sql.Open("mysql", mysql_username+":"+mysql_password+"@/"+mysql_database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// * List all routes.
	router := gin.Default()
	router.GET("/albums", getAlbums)
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
	// 	fmt.Println(index, "Create table");
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
