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
        if err := c.BindJSON(&nuevoLibro); err != nil {
                c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
                return
        }
        libros = append(libros, nuevoLibro)
        c.IndentedJSON(http.StatusCreated, nuevoLibro)
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

func deleteLibroPorID(c *gin.Context) {
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

func patchLibroPorID(c *gin.Context) {
        id := c.Param("id")
        
        
        for i, libro := range libros {
                if libro.ID == id {
                        
                        var actualizacion map[string]interface{}
                        if err := c.BindJSON(&actualizacion); err != nil {
                                c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
                                return
                        }
                        
                
                        if titulo, existe := actualizacion["titulo"]; existe {
                                if tituloStr, ok := titulo.(string); ok {
                                        libros[i].Titulo = tituloStr
                                }
                        }
                        
                        if autor, existe := actualizacion["autor"]; existe {
                                if autorStr, ok := autor.(string); ok {
                                        libros[i].Autor = autorStr
                                }
                        }
                        
                        if año, existe := actualizacion["año"]; existe {
                                if añoFloat, ok := año.(float64); ok {
                                        libros[i].Año = int(añoFloat)
                                }
                        }
                        
                        c.IndentedJSON(http.StatusOK, libros[i])
                        return
                }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje": "Libro no encontrado"})
}

func main() {
        router := gin.Default()
        router.GET("/libros", getLibros)
        router.POST("/libros", postLibros)
        router.GET("/libros/:id", getLibroPorID)
        router.DELETE("/libros/:id", deleteLibroPorID)
        router.PATCH("/libros/:id", patchLibroPorID)
        router.Run("localhost:8080")
}
