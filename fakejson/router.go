package fakejson

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func RoutePerform(c *gin.Context) {
	var reqObj postData
	if err := c.ShouldBindJSON(&reqObj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// perform cleanup and custom validations
	// if err := reqObj.validateParams(); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// perform business logic
	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         Schema,
		RequestString:  reqObj.Query,
		VariableValues: reqObj.Variables,
		OperationName:  reqObj.Operation,
	})
	if len(result.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": result.Errors})
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
}
