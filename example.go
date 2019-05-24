package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type FormParams struct {
	FieldA     string   `form:"field_a"`
	FieldB     string   `form:"field_b"`
	ArrayField []string `form:"array_field"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/u/:name/:id", func(c *gin.Context) {
		var person Person

		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})

			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	r.GET("/get", func(c *gin.Context) {
		var b FormParams
		c.Bind(&b)
		c.JSON(200, gin.H{
			"field_a":     b.FieldA,
			"field_b":     b.FieldB,
			"array_field": b.ArrayField,
		})
	})

	r.Run(":8088")
}
