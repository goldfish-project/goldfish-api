package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
	"goldfish-api/internal/core/domain"
	"net/http"
)

func init() {
	fmt.Println("Init...")
	orm.RegisterTable((*domain.WorkspaceToUser)(nil))
	fmt.Println("Done...")
}

func main() {
	// load and read config
	config := readConfigFromEnviroment()

	// setup database
	db := pg.Connect(config.DB.PGOptions())
	if err := createSchema(db); err != nil {
		panic(err)
	}

	// create http router
	router := mux.NewRouter()

	fmt.Println("Starting server on port " + config.Port)

	// start web service
	if err := http.ListenAndServe(config.Host + ":" + config.Port, router); err != nil {
		panic(err)
	}

	fmt.Println(config)
}

// createSchema creates the database schema of the applications models
func createSchema(db *pg.DB) error {
	/*orm.SetTableNameInflector(func(s string) string {
		return "goldfish_" + s
	})*/

	orm.RegisterTable((*domain.WorkspaceToUser)(nil))

	models := []interface{}{
		(*domain.User)(nil),
		(*domain.Workspace)(nil),
		(*domain.WorkspaceToUser)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:          true,
			IfNotExists:   false,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}