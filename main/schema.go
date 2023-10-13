// schema.go
package main

import (
	"github.com/graphql-go/graphql"
)

var StateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "State",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"search": &graphql.Field{
			Type: graphql.NewList(StateType),
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: SearchResolver,
		},
	},
})
