package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type libro struct {
	ID     string `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Año    int    `json:"año"`
}

var libros = []libro{
	{ID: "1", Titulo: "Cien años de soledad", Autor: "Gabriel garcia marquez", Año: 1972},
	{ID: "2", Titulo: "Don quijote de la mancha", Autor: "Miguel de cervantes", Año: 1605},
	{ID: "3", Titulo: "La casa de los espiritus", Autor: "Isabel allende", Año: 1982},
}

func getLibros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, libros)
}

func postLibros(c *gin.Context){
	var nuevoLibro libro
	c.BindJSON(&nuevoLibro)
	libros=append(libros,nuevoLibro)
	c.IndentedJSON(http.StatusCreated,libros)
}







func main() {
	router := gin.Default()
	router.GET("/", getLibros)
	router.POST("/",postLibros)
	router.Run("localhost:8080")
}
