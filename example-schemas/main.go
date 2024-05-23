package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wundergraph/cosmo/composition-go"
)

func main() {
	config, err := composition.BuildRouterConfiguration(&composition.Subgraph{
		Name: "Account",
		Schema: `
		type Query {
			account: Account
		}
		type Account {
			user: User
		}
		type User @key(fields: "id", resolvable: false) {
			id: Int
		}
			`,
		}, &composition.Subgraph{
		Name: "Team",
		Schema: `
		type Query {
			team: Team
		}
		type Team {
			user: User
		}
		type User @key(fields: "age", resolvable: false) {
			age: Int
		}`,
	  }, &composition.Subgraph{
			Name: "User",
			Schema: `
			type Query {
				user: User
			}
			type User @key(fields: "id") @key(fields: "age") {
				id: Int
				age: Int
				salary: Int
			 }
	 `,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
	os.WriteFile("schema.json", []byte(config), 0644)
}
