package main

import (
	"2048-Game/repo"
	"2048-Game/service"
)

func main(){
	sizeOfBoard := 4
	freeRepo := repo.NewFreeGrid(4)
	filledRepo := repo.NewFilledGrid()
	grid := service.NewGrid(sizeOfBoard,freeRepo,filledRepo)
	//Add two Rand Values
	grid.AddRandValue()
	grid.AddRandValue()
	runCommands(grid)
}


