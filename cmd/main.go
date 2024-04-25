package main

import (
	"fmt"

	"monitors-creator/cmd/config"
	"monitors-creator/cmd/handlers"
	usecase "monitors-creator/internal/monitors-creator"
	"monitors-creator/internal/monitors-creator/monitor"
	"monitors-creator/internal/platform/memdb"
	"monitors-creator/internal/platform/restclient"
)

func init() {
	config.LoadMonitorCreatorConfig()
}

func main() {
	//TODO: CRIAR COROUTINE QUE PINGA O DB E SE FALHAR, REABRE E ATUALIZA REFERENCIA DO DB NO REPO (SINGLETON)
	fmt.Println("initializing...")

	db := memdb.NewDB()
	r := monitor.NewMemoryRepo(db)

	datadog, err := restclient.NewEndpoint(config.Get().DatadogURL)
	if err != nil {
		panic(err)
	}

	u := usecase.NewMonitorUsecase(r, datadog)
	h := handlers.NewMonitorHandler(u)

	if err := handlers.NewFuryApplication(h); err != nil {
		panic(err)
	}
}
