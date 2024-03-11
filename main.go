package main

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID 	string `json:"id"`
	Title	string `json:"title"`
	Artist	string `json:"artist"`
	Price	float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John", Price:55.90 },
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price:60.90 },
	{ID: "3", Title: "Sarah Varigner", Artist: "Manta", Price:45.90 },
	{ID: "4", Title: "Sex Desires", Artist: "Guz", Price:75.90 },
}

//getAlbuns responds with list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,albums)
}

//postAlmbum adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	//call BinsJSON to bind received JSON to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	//Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID locate the album whose ID value matchs the ID
//parameter sent by the client, then returns that album as a response

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums, looking for an album whose ID values matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("albums", postAlbums)

	router.Run("localhost:8080")
}
