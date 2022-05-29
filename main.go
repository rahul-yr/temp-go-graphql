package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/rahul-yr/temp-go-graphql/fakejson"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func RoutePerform(c *gin.Context) {
	// log.Println("==========================")
	// log.Println(c.Request)
	// log.Println("==========================")
	var reqObj postData
	if err := c.ShouldBindJSON(&reqObj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// perform cleanup and custom validations

	// perform business logic
	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         fakejson.Schema,
		RequestString:  reqObj.Query,
		VariableValues: reqObj.Variables,
		OperationName:  reqObj.Operation,
	})
	// log.Println("==========================")
	// log.Println(result)
	// log.Println("==========================")

	if len(result.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Errors})
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/graphql", RoutePerform)
	router.Run()
}
