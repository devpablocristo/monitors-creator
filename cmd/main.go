package main

import (
	"fmt"

	"monitors-creator/cmd/config"
	"monitors-creator/cmd/handlers"
	usecase "monitors-creator/internal/monitors-creator"
	"monitors-creator/internal/monitors-creator/monitor"
	"monitors-creator/internal/platform/memdb"
)

func init() {
	config.LoadMonitorCreatorConfig()
}

func main() {
	//TODO: CRIAR COROUTINE QUE PINGA O DB E SE FALHAR, REABRE E ATUALIZA REFERENCIA DO DB NO REPO (SINGLETON)
	fmt.Println("initializing...")

	db := memdb.NewDB()
	r := monitor.NewMemoryRepo(db)
	u := usecase.NewMonitorUsecase(r)
	h := handlers.NewMonitorHandler(u)

	handlers.NewFuryApplication(h)
	// if err := handlers.NewHTTPServer(*h); err != nil {
	// 	panic(errors.Errorf("error starting HTTP server: %v", err))
	// }
}
