package main

import (
	"github.com/inahym196/bomb/internal/controller"
	"github.com/inahym196/bomb/internal/interactor"
)

func main() {

	interactor := interactor.NewGameInteractor()
	controller.NewCLIController(interactor).Run()
}
