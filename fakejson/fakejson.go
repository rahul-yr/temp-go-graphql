package fakejson

import "github.com/graphql-go/graphql"

const base_url = "https://jsonplaceholder.typicode.com/todos"

var fakeJSONObjectConf = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FakeJson",
		Fields: graphql.Fields{
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"list": &graphql.Field{
				Type:        graphql.NewList(fakeJSONObjectConf),
				Description: "List of fake json objects",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					data, err := getFakeJSONObjects()
					if err != nil {
						return nil, err
					}
					return data, nil
				},
			},
			"getbyId": &graphql.Field{
				Type:        fakeJSONObjectConf,
				Description: "Get fake json object by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					data, err := getFakeJSONObjectById(id)
					if err != nil {
						return nil, err
					}
					return data, nil
				},
			},
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"delete": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					err := deleteTodoItem(id)
					if err != nil {
						return nil, err
					}
					return true, nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)
