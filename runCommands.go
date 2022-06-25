package main

import (
	"2048-Game/models"
	"2048-Game/service"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runCommands(manager *service.Grid){
	printGrid(manager.GetGrid())
	for {
		switch readInput() {
		case models.Left:
			manager.GoLeft()
		case models.Right:
			manager.GoRight()
		case models.Top:
			manager.GoTop()
		case models.Bottom:
			manager.GoDown()
		}
		if addValueAndCheckSuccess(manager) {
			return
		}
	}
}

func readInput()string{
	reader :=bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	return strings.Split(val,"\n")[0]
}

func printGrid(str []string){
	fmt.Println(str)
}

func addValueAndCheckSuccess(manager *service.Grid)(exit bool){
	if !manager.AddRandValue(){
		if manager.CheckForSuccess() {
			fmt.Println("Congratulations")
		} else {
			fmt.Println("Game Over")
		}
		exit = true
	}
	printGrid(manager.GetGrid())
	return
}