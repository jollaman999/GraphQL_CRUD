package graphql_util

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryTypes,
		Mutation: mutationTypes,
	},
)

var Graphql_handler = handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
		GraphiQL: true,
})