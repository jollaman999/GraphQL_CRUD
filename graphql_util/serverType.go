package graphql_util

import "github.com/graphql-go/graphql"

var serverType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Server",
		Fields: graphql.Fields {
			"uuid": &graphql.Field {
				Type: graphql.String,
			},
			"server_name": &graphql.Field {
				Type: graphql.String,
			},
			"server_disc": &graphql.Field {
				Type: graphql.String,
			},
			"cpu": &graphql.Field {
				Type: graphql.Int,
			},
			"memory": &graphql.Field {
				Type: graphql.Int,
			},
			"disk_size": &graphql.Field {
				Type: graphql.Int,
			},
			"created": &graphql.Field {
				Type: graphql.DateTime,
			},
		},
	},
)