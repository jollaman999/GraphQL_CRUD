package graphql_util

import (
	"../logger"
	"../mysql_util"
	"../types"
	"github.com/graphql-go/graphql"
	"time"
)

var queryTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			////////////////////////////// user ///////////////////////////////
			/* Get (read) single user by id
			   http://localhost:8001/graphql?query={user(id:"test1"){id,name}}
			*/
			"user": &graphql.Field {
				Type:        userType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: user")

					requested_id, ok := p.Args["id"].(string)
					if ok {
						user := new(types.User)

						var id string
						var name string

						sql := "select * from user where id = ?"
						err := mysql_util.Db.QueryRow(sql, requested_id).Scan(&id, &name)
						if err != nil {
							logger.Logger.Println(err)
							return nil, nil
						}

						user.Id = id
						user.Name = name

						return user, nil
					}
					return nil, nil
				},
			},

			/* Get (read) user list
			   http://localhost:8001/graphql?query={list_user{id,name}}
			*/
			"list_user": &graphql.Field {
				Type:        graphql.NewList(userType),
				Description: "Get user list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: list_user")

					var users []types.User
					var id string;
					var name string;

					sql := "select * from user"
					stmt, err := mysql_util.Db.Query(sql)
					if err != nil {
						logger.Logger.Println(err)
						return nil, nil
					}
					defer stmt.Close()

					for stmt.Next() {
						err := stmt.Scan(&id, &name)
						if err != nil {
							logger.Logger.Println(err)
						}

						user := types.User{Id: id, Name: name}

						logger.Logger.Println(user)
						users = append(users, user)
					}

					return users, nil
				},
			},

			////////////////////////////// server ///////////////////////////////
			/* Get (read) single server by uuid
			   http://localhost:8001/graphql?query={server(uuid:"[uuid]"){uuid,server_name,server_disc,cpu,memory,disk_size,created}}
			*/
			"server": &graphql.Field {
				Type:        serverType,
				Description: "Get server by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: server")

					requested_uuid, ok := p.Args["uuid"].(string)
					if ok {
						server := new(types.Server)

						var uuid string
						var server_name string
						var server_disc string
						var cpu int
						var memory int
						var disk_size int
						var created time.Time

						sql := "select * from server where uuid = ?"
						err := mysql_util.Db.QueryRow(sql, requested_uuid).Scan(&uuid, &server_name, &server_disc, &cpu, &memory, &disk_size, &created)
						if err != nil {
							logger.Logger.Println(err)
							return nil, nil
						}

						server.Uuid = uuid
						server.Server_name = server_name
						server.Server_disc = server_disc
						server.Cpu = cpu
						server.Memory = memory
						server.Disk_size = disk_size
						server.Created = created

						return server, nil
					}
					return nil, nil
				},
			},

			/* Get (read) server list
			   http://localhost:8001/graphql?query={list_server{uuid,server_name,server_disc,cpu,memory,disk_size,created}}
			*/
			"list_server": &graphql.Field {
				Type:        graphql.NewList(serverType),
				Description: "Get server list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: list_server")

					var servers []types.Server

					var uuid string
					var server_name string
					var server_disc string
					var cpu int
					var memory int
					var disk_size int
					var created time.Time

					sql := "select * from server"
					stmt, err := mysql_util.Db.Query(sql)
					if err != nil {
						logger.Logger.Println(err)
						return nil, nil
					}
					defer stmt.Close()

					for stmt.Next() {
						err := stmt.Scan(&uuid, &server_name, &server_disc, &cpu, &memory, &disk_size, &created)
						if err != nil {
							logger.Logger.Println(err)
						}

						server := types.Server{Uuid:uuid, Server_name:server_name, Server_disc:server_disc, Cpu:cpu, Memory:memory, Disk_size:disk_size, Created:created}

						logger.Logger.Println(server)
						servers = append(servers, server)
					}

					return servers, nil
				},
			},
		},
	})
