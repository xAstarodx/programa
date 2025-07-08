package main

import (
	"net/http"
	"fmt"

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

func getLibroPorID(c *gin.Context) {
    id := c.Param("id")
    for _, libro := range libros {
        if libro.ID == id {
            c.IndentedJSON(http.StatusOK, libro)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje": "Libro no encontrado"})
}


deleteLibroPorID(c *gin.Context) {
	id := c.Param("id")
	for i, libro := range libros {		
		if libro.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"mensaje": "Libro eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje": "Libro no encontrado"})
}	



func main() {
	router := gin.Default()
	router.GET("/", getLibros)
	router.POST("/post",postLibros)
	router.GET("/:id", getLibroPorID)
	router.DELETE("/delete:id", deleteLibroPorID)
	router.Run("localhost:8080")
	fmt.Println("Servidor corriendo en http://localhost:8080")
}