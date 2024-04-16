package main

import (
	"log"
	"os"
	"users_example/cmd/devapi/furyapp"
	"users_example/internal/platform/environment"
)

// TODO list
// 1. agregar métricas y logs a los casos de uso y handlers.
// 2. lockear para controlar la concurrencia en el momento de hacer actualizar una task a completed.
// 3. algunos controles más sobre ciertas reglas de negocio, por ejemplo que no se pueda asignar una tarea a un dev que ya tiene una.
// 4. wrappeo de errores de parte de internal en los handlers para poder decidir que status code tirar.
// 5. más tests unitarios y de integración
func main() {
	env := environment.GetFromString(os.Getenv("GO_ENVIRONMENT"))

	dependencies, err := furyapp.BuildDependencies(env)
	if err != nil {
		log.Fatal("error at dependencies building", err)
	}

	app := furyapp.Build(dependencies)
	if err := app.Run(); err != nil {
		log.Fatal("error at furyapp startup", err)
	}
}
