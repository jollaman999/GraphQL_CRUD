package graphql_util

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "User",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"name": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
