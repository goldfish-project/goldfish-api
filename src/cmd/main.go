package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
	"goldfish-api/internal/core/services"
	"goldfish-api/internal/handlers"
	"goldfish-api/internal/respositories/postgres"
	"net/http"
)

func init() {
	fmt.Println("Init...")
	orm.RegisterTable((*postgres.WorkspaceToUser)(nil))
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
	router = router.PathPrefix("/api").Subrouter()

	// create repositories
	userRepository := postgres.NewUserRepository(db)

	// create services
	userService := services.NewService(userRepository)

	// init http handler
	handlers.NewHTTPUserHandler(router, userService)

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

	models := []interface{}{
		(*postgres.User)(nil),
		(*postgres.Variable)(nil),
		(*postgres.Workspace)(nil),
		(*postgres.Collection)(nil),
		(*postgres.WorkspaceToUser)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}