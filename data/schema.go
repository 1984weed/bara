package data

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var userType *graphql.Object
var widgetType *graphql.Object
var questionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Question",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SubmitCodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SubmitCodeType",
		Fields: graphql.Fields{
			"typedCode": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lang": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

type SubmitCode struct {
	TypedCode string
	Result    string
	Lang      string
}

var nodeDefinitions *relay.NodeDefinitions
var widgetConnection *relay.GraphQLConnectionDefinitions

var Schema graphql.Schema

type Question struct {
	ID          int64
	Title       string
	Description string
}

func init() {

	/**
	 * We get the node interface and field from the Relay library.
	 *
	 * The first method defines the way we resolve an ID to its object.
	 * The second defines the way we resolve an object to its GraphQL type.
	 */
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ct context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "User" {
				return GetUser(resolvedID.ID), nil
			}
			if resolvedID.Type == "Widget" {
				return GetWidget(resolvedID.ID), nil
			}
			return nil, nil
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *User:
				return userType
			case *Widget:
				return widgetType
			}
			return nil
		},
	})

	/**
	 * Define your own types here
	 */
	widgetType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Widget",
		Description: "A shiny widget'",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Widget", nil),
			"name": &graphql.Field{
				Description: "The name of the widget",
				Type:        graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})
	widgetConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "WidgetConnection",
		NodeType: widgetType,
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A person who uses our app",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"widgets": &graphql.Field{
				Type:        widgetConnection.ConnectionType,
				Description: "A person's collection of widgets",
				Args:        relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					dataSlice := WidgetsToInterfaceSlice(GetWidgets()...)
					return relay.ConnectionFromArray(dataSlice, args), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,
			// Add you own root fields here
			"viewer": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(), nil
				},
			},
		},
	})
	/**
	 * This is the type that will be the root of our mutations,
	 * and the entry point into performing writes in our schema.
	 */
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"codesSubmit": &graphql.Field{
				Type: SubmitCodeType,
				Args: graphql.FieldConfigArgument{
					"typedCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lang": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Println("ki------------ta")
					// nodeClient := remoteNewNodeJsClient()
					result := SubmitCode{
						TypedCode: p.Args["typedCode"].(string),
						Lang:      p.Args["lang"].(string),
						Result:    "OK",
					}

					return result, nil
				},
			},
			// Add you own mutations here
			"createQuestion": &graphql.Field{
				Type:        questionType,
				Description: "Create new questions",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"funcName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"args": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"testcases": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					question := Question{
						ID:          int64(rand.Intn(100000)),
						Title:       p.Args["title"].(string),
						Description: p.Args["description"].(string),
					}

					return question, nil
				},
			},
		},
	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		panic(err)
	}
}
