package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func uid() string {
	now := time.Now()
	x := fmt.Sprint(rand.Intn(1000))

	return fmt.Sprint(now.Unix(), x)
}

var list = []todo{
	{ID: uid(), Title: "Wash the car", IsDone: true},
	{ID: uid(), Title: "feed the cat", IsDone: false},
	{ID: uid(), Title: "clean the room", IsDone: true},
}

func RemoveIndex(s []todo, index int) []todo {
	return append(s[:index], s[index+1:]...)
}

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World from sina!!!")
}

func getToDoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, list)
}

func getSingleToDoList(c *gin.Context) {
	id := c.Param("id")

	for _, a := range list {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}

func addToDo(c *gin.Context) {
	var newToDo todo

	newToDo.ID = uid()
	newToDo.IsDone = false

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newToDo); err != nil {
		return
	}

	// Add the new album to the slice.
	list = append(list, newToDo)
	c.IndentedJSON(http.StatusCreated, newToDo.ID)
}

func deleteToDo(c *gin.Context) {
	id := c.Param("id")

	for i, a := range list {

		if a.ID == id {
			newList := RemoveIndex(list, i)
			list = newList

			c.IndentedJSON(http.StatusOK, "ok")
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, "no such todo exists")
}

func markToDo(c *gin.Context) {
	id := c.Param("id")

	for i, a := range list {
		fmt.Println("*****************************")
		fmt.Println(i)
		fmt.Println(a)
		fmt.Println("*****************************")
		if a.ID == id {

			list[i].IsDone = !list[i].IsDone

			c.IndentedJSON(http.StatusOK, "lol")
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, "no such todo exists")
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/list", getToDoList)
	router.GET("/list/:id", getSingleToDoList)
	router.POST("/list/add", addToDo)
	router.POST("/list/mark/:id", markToDo)
	router.POST("/list/delete/:id", deleteToDo)

	router.Run("localhost:8080")
}
