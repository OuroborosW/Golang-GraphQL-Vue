// main.go
package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Printf("Hello")
	initDB()
	defer client.Disconnect(context.Background())

	r := gin.Default()
	r.Use(cors.Default())
	// GraphQL schema 和解析器配置可以直接来自上面的代码
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: RootQuery,
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	r.POST("/graphql", func(c *gin.Context) {
		var request struct {
			Query string `json:"query"`
		}
		if c.ShouldBindJSON(&request) == nil {
			result := graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: request.Query,
			})
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		}
	})

	r.Run()
}
