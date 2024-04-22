package main

import (
	"log"

	"github.com/pkg/errors"

	"monitors-creator/cmd/handlers"
	usecase "monitors-creator/internal/monitors-creator"
	"monitors-creator/internal/monitors-creator/monitor"
	"monitors-creator/internal/platform/memdb"
)

func main() {
	db := memdb.NewDB()
	repo := monitor.NewMemoryRepo(db)
	u := usecase.NewMonitorUsecase(repo)
	h := handlers.NewMonitorHandler(u)

	if err := handlers.NewHTTPServer(*h); err != nil {
		log.Fatal(errors.Errorf("error starting HTTP server: %v", err))
	}
}

// func start() error {
// 	if err := env.LoadEnv(); err != nil {
// 		return fmt.Errorf("error loading .env file: %w", err)
// 	}

// 	dbConn, err := setupDatabase()
// 	if err != nil {
// 		return fmt.Errorf("error setting up database: %w", err)
// 	}

// 	u := taskslist.NewTaskUsecase(dbConn)
// 	h := handler.NewHandler(u)

// 	if err := gin.NewHTTPServer(*h); err != nil {
// 		return fmt.Errorf("error starting HTTP server: %w", err)
// 	}

// 	return nil
// }

// func setupDatabase() (task.MySqlRepositoryPort, error) {
// 	// MYSQL DB
// 	dbConn, err := mysql.GetConnectionDB()
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening MySQL connection: %w", err)
// 	}
// 	d := task.NewMySqlDao(dbConn)
// 	r := task.NewMysqlRepository(d)

// 	return r, nil
// }
