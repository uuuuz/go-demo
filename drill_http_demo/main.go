package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"net/http"

	"github.com/graphql-go/graphql"
)

func getDrillData(ctx *gin.Context) {

	val, yes := ctx.GetQuery("mode")
	if yes && val == "test" {
		_, _ = ctx.Writer.WriteString(`[
{"id": 1, "name": "marisa", "gender": "girl", "age": 14},
{"id": 2, "name": "uuz", "gender": "girl", "age": 17},
{"id": 3, "name": "sakura", "gender": "girl", "age": 15}
]`)
		return
	}
	_, _ = ctx.Writer.WriteString(`[
{"id": 1, "name": "marisa", "gender": "girl", "age": 14},
{"id": 2, "name": "uuz", "gender": "girl", "age": 17},
{"id": 3, "name": "sakura", "gender": "girl", "age": 15},
{"id": 4, "name": "han", "gender": "boy", "age": 15},
{"id": 5, "name": "wang", "gender": "boy", "age": 17}
]`)
}

func main() {
	ginRou := gin.Default()

	ginRou.POST("/gql", GraphglHandler())
	ginRou.GET("/drill-demo", getDrillData)

	_ = http.ListenAndServe(":8002", ginRou)
}

var testFields = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
	"test": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			data := getData()
			return data, nil
		},
	},
}

var testQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootQuery",
	Fields: testFields,
})

type TestView struct {
	Name      string   `json:"name"`
	Alias     string   `json:"alias"`
	ActualAge int      `json:"actual_age"`
	Detail    []string `json:"detail"`
}

func getData() []TestView {
	return []TestView{
		{
			Name:      "marisa",
			Alias:     "Marisa",
			ActualAge: 14,
			Detail:    []string{"girl", "magic"},
		},
		{
			Name:      "uuz",
			Alias:     "UUZ",
			ActualAge: 14,
			Detail:    []string{"girl", "magic", "master"},
		},
	}
}

func GraphglHandler() gin.HandlerFunc {
	schemaConfig := graphql.SchemaConfig{Query: testQuery}

	schema, _ := graphql.NewSchema(schemaConfig)

	hand := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(ctx *gin.Context) {
		hand.ContextHandler(ctx, ctx.Writer, ctx.Request)
	}
}
