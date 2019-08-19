package graphql_util

import (
	"../exec_util"
	"../logger"
	"../mysql_util"
	"../types"
	"github.com/graphql-go/graphql"
	"time"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		////////////////////////////// user ///////////////////////////////
		/* Create new user
		http://localhost:8001/graphql?query=mutation+_{create_user(id:"test1",name:"1"){id,name}}
		*/
		"create_user": &graphql.Field{
			Type:        userType,
			Description: "Create new user",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig {
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: create_user")

				user := types.User{
					Id:    params.Args["id"].(string),
					Name:  params.Args["name"].(string),
				}

				sql := "insert into user(id, name) values (?, ?)"
				stmt, err := mysql_util.Db.Prepare(sql)
				if err != nil {
					logger.Logger.Println(err.Error())
					return nil, nil
				}
				defer stmt.Close()
				result, err2 := stmt.Exec(user.Id, user.Name)
				if err2 != nil {
					logger.Logger.Println(err2)
					return nil, nil
				}
				logger.Logger.Println(result.LastInsertId())

				return user, nil
			},
		},

		/* Update user by id
		   http://localhost:8001/graphql?query=mutation+_{update_user(id:"test1",name:"11"){id,name}}
		*/
		"update_user": &graphql.Field{
			Type:        userType,
			Description: "Update user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: update_user")

				requested_id, _ := params.Args["id"].(string)
				name, nameOk := params.Args["name"].(string)

				user := new(types.User)

				if nameOk {
					user.Id = requested_id
					user.Name = name

					sql := "update user set name = ? where id = ?"
					stmt, err := mysql_util.Db.Prepare(sql)
					if err != nil {
						logger.Logger.Println(err.Error())
						return nil, nil
					}
					defer stmt.Close()
					result, err2 := stmt.Exec(user.Name, user.Id)
					if err2 != nil {
						logger.Logger.Println(err2)
						return nil, nil
					}
					logger.Logger.Println(result.LastInsertId())

					return user, nil
				}
				return nil, nil
			},
		},

		/* Delete user by id
		   http://localhost:8001/graphql?query=mutation+_{delete_user(id:"test1"){id}}
		*/
		"delete_user": &graphql.Field{
			Type:        userType,
			Description: "Delete user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: delete_user")

				requested_id, ok := params.Args["id"].(string)
				if ok {
					sql := "delete from user where id = ?"
					stmt, err := mysql_util.Db.Prepare(sql)
					if err != nil {
						logger.Logger.Println(err.Error())
						return nil, nil
					}
					defer stmt.Close()
					result, err2 := stmt.Exec(requested_id)
					if err2 != nil {
						logger.Logger.Println(err2)
						return nil, nil
					}
					logger.Logger.Println(result.RowsAffected())

					return requested_id, nil
				}
				return nil, nil
			},
		},

		////////////////////////////// server ///////////////////////////////
		/* Create new server
		http://localhost:8001/graphql?query=mutation+_{create_server(server_name:"ish",server_disc:"ish server",cpu:12,memory:16384,disk_size:1024000){uuid,server_name,server_disc,cpu,memory,disk_size,created}}
		*/
		"create_server": &graphql.Field{
			Type:        serverType,
			Description: "Create new server",
			Args: graphql.FieldConfigArgument {
				"server_name": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
				},
				"server_disc": &graphql.ArgumentConfig {
					Type: graphql.String,
				},
				"cpu": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
				"memory": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
				"disk_size": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: create_server")

				uuid, err := exec_util.Exec_uuidgen()
				if err != nil {
					logger.Logger.Println("Failed to generate uuid!")
					return nil, nil
				}

				server := types.Server{
					Uuid: uuid,
					Server_name:  	params.Args["server_name"].(string),
					Server_disc:  	params.Args["server_disc"].(string),
					Cpu:			params.Args["cpu"].(int),
					Memory:			params.Args["memory"].(int),
					Disk_size:			params.Args["disk_size"].(int),
					Created:			time.Now(),
				}

				sql := "insert into server(uuid, server_name, server_disc, cpu, memory, disk_size, created) values (?, ?, ?, ?, ?, ?, ?)"
				stmt, err := mysql_util.Db.Prepare(sql)
				if err != nil {
					logger.Logger.Println(err.Error())
					return nil, nil
				}
				defer stmt.Close()
				result, err2 := stmt.Exec(server.Uuid, server.Server_name, server.Server_disc, server.Cpu, server.Memory, server.Disk_size, server.Created)
				if err2 != nil {
					logger.Logger.Println(err2)
					return nil, nil
				}
				logger.Logger.Println(result.LastInsertId())

				return server, nil
			},
		},

		/* Update server by uuid
		   http://localhost:8001/graphql?query=mutation+_{update_server(uuid:"[uuid]",server_name:"ish",server_disc:"ish server",cpu:12,memory:16384,disk_size:1024000){uuid,server_name,server_disc,cpu,memory,disk_size}}
		*/
		"update_server": &graphql.Field{
			Type:        serverType,
			Description: "Update server by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
				},
				"server_name": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
				},
				"server_disc": &graphql.ArgumentConfig {
					Type: graphql.String,
				},
				"cpu": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
				"memory": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
				"disk_size": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: update_server")

				requested_uuid, _ := params.Args["uuid"].(string)
				server_name, server_nameOk := params.Args["server_name"].(string)
				server_disc, server_discOk := params.Args["server_disc"].(string)
				cpu, cpuOk := params.Args["cpu"].(int)
				memory, memoryOk := params.Args["memory"].(int)
				disk_size, disk_sizeOk := params.Args["disk_size"].(int)

				server := new(types.Server)

				if server_nameOk && server_discOk && cpuOk && memoryOk && disk_sizeOk {
					server.Uuid = requested_uuid
					server.Server_name = server_name
					server.Server_disc = server_disc
					server.Cpu = cpu
					server.Memory = memory
					server.Disk_size = disk_size

					sql := "update server set server_name = ?, server_disc = ?, cpu = ?, memory = ?, disk_size = ? where uuid = ?"
					stmt, err := mysql_util.Db.Prepare(sql)
					if err != nil {
						logger.Logger.Println(err.Error())
						return nil, nil
					}
					defer stmt.Close()
					result, err2 := stmt.Exec(server.Server_name, server.Server_disc, server.Cpu, server.Memory, server.Disk_size, server.Uuid)
					if err2 != nil {
						logger.Logger.Println(err2)
						return nil, nil
					}
					logger.Logger.Println(result.LastInsertId())

					return server, nil
				}
				return nil, nil
			},
		},

		/* Delete server by uuid
		   http://localhost:8001/graphql?query=mutation+_{delete_server(uuid:"[uuid]"){uuid}}
		*/
		"delete_server": &graphql.Field{
			Type:        serverType,
			Description: "Delete user by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: delete_server")

				requested_uuid, ok := params.Args["uuid"].(string)
				if ok {
					sql := "delete from server where uuid = ?"
					stmt, err := mysql_util.Db.Prepare(sql)
					if err != nil {
						logger.Logger.Println(err.Error())
						return nil, nil
					}
					defer stmt.Close()
					result, err2 := stmt.Exec(requested_uuid)
					if err2 != nil {
						logger.Logger.Println(err2)
						return nil, nil
					}
					logger.Logger.Println(result.RowsAffected())

					return requested_uuid, nil
				}
				return nil, nil
			},
		},
	},
})
