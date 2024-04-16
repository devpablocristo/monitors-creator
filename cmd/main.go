package main

import (
	"fmt"
	"monitors-creator/internal/monitors-factory/monitor"
)

func main() {
	// rabbitMqManager, err := hdl.NewRabbitMQManager("amqp://guest:guest@localhost:5672/", "myQueue")
	// if err != nil {
	// 	log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	// }
	// defer rabbitMqManager.Close()

	dbConn, err := mysql.GetConnectionDB()
	if err != nil {
		return nil, fmt.Errorf("error opening MySQL connection: %w", err)
	}
	repository := monitor.NewMonitorRepository()
	_ = repository
}

// func main() {
// 	if err := start(); err != nil {
// 		log.Fatalf("Application failed: %v", err)
// 	}
// }

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
