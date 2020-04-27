package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success: %s", person)
}

// curl -v "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"
// curl -v -XPOST -F 'name=appleboy' -F 'address=xyz' -F 'birthday=1992-03-15' "localhost:8085/testing"

// use httpie
// http -v "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"
// http -v "localhost:8085/testing" 'name==appleboy' 'address==xyz' 'birthday==1992-03-15'
// http -v POST "localhost:8085/testing" 'name=appleboy' 'address=xyz' 'birthday=1992-03-15'
// http -v -f POST "localhost:8085/testing" 'name=appleboy' 'address=xyz' 'birthday=1992-03-15'
