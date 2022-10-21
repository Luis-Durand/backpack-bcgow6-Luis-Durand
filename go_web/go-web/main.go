/* ,
   "Usuario2":{
       "Id":2,
        "Nombre":"Pepe",
        "Apellido":"Diaz",
        "Email":"Pepe@gmail.com",
        "Altura":1.75,
        "Activo":false,
        "Creacion":"20/04"

   } */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* type Users struct {
	Users []user
} */

type user struct {
	Id       int
	Nombre   string
	Apellido string
	Email    string
	Edad     int
	Altura   float32
	Activo   bool
	Creacion string
}

func GetAll(c *gin.Context) {

	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		errors.New("fallo en la lectura del json")
		return
	}

	var p []user
	var filter []user
	if err := json.Unmarshal(data, &p); err != nil {
		errors.New("error en el parseo")
	}
	numEdad, _ := strconv.Atoi(c.Query("Edad"))
	numId, err := strconv.Atoi(c.Query("Id"))

	for _, users := range p {
		if numId == users.Id || c.Query("Nombre") == (users.Nombre) || c.Query("Apellido") == users.Nombre || c.Query("Email") == users.Nombre || numEdad == users.Edad || c.Query("Altura") == users.Nombre || c.Query("Activo") == users.Nombre || c.Query("Creacion") == users.Nombre {
			filter = append(filter, users)

		} else {
			fmt.Println("No encontro a nadie")
		}

	}

	c.JSON(200, gin.H{"users": filter})
}

func GetOne(c *gin.Context) {
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		errors.New("fallo en la lectura del json")
		return
	}

	var p []user

	if err := json.Unmarshal(data, &p); err != nil {
		errors.New("error en el parseo")
	}

	if err != nil {
		errors.New("fallo de conversion de string a int")
		return
	}
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.New("Fallo en conversion")

	}
	var userFound user

	for _, user := range p {
		if user.Id == idParam {
			userFound = user
			break
		}
	}
	if userFound.Id > 0 {
		c.String(200, "Aca esta el user %#v", userFound)
	} else {
		c.String(404, "No se encontro el user")
	}

}

func main() {

	/* p:=User{
		Id:
	} */

	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Luis",
		})
	})

	router.GET("/users", GetAll)
	router.GET("/oneuser/:id", GetOne)
	router.Run()

}
