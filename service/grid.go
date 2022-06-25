package service

import (
	"2048-Game/models"
	"2048-Game/repo"
	"fmt"
	"math"
	"math/rand"
)

type Grid struct {
	freeRepo *repo.FreeGrid
	size int
	filledRepo *repo.FilledGrid
}

func NewGrid(n int, freeGrid *repo.FreeGrid, filledGrid *repo.FilledGrid)*Grid{
	return &Grid{
		filledRepo: filledGrid,
		freeRepo: freeGrid,
		size: n,
	}
}
func (g *Grid)GoTop(){
	n := g.size
	for col:=0;col<n;col++ {
		temp := 0
		for row:=1;row<n;row++ {
			prevIdx := g.getEncIndex(row-1,col)
			currentIdx := g.getEncIndex(row, col)
			addIdx := g.getEncIndex(temp, col)
			r , t := g.checkForAdjacent(prevIdx, currentIdx, addIdx)
			row += r
			temp += t
		}
	}
}

func (g *Grid)GoDown(){
	n := g.size
	for col:=0;col<n;col++ {
		temp := n-1
		for row:=n-2;row>-1;row-- {
			nextIdx := g.getEncIndex(row+1,col)
			currentIdx := g.getEncIndex(row, col)
			addIdx := g.getEncIndex(temp, col)
			r , t := g.checkForAdjacent(nextIdx, currentIdx, addIdx)
			row -= r
			temp -= t
		}
	}
}

func (g *Grid)GoLeft(){
	n := g.size
	for row:=0;row<n;row++ {
		temp := 0
		for col:=1;col<n;col++ {
			prevIdx := g.getEncIndex(row,col-1)
			currentIdx := g.getEncIndex(row, col)
			addIdx := g.getEncIndex(row, temp)
			c , t := g.checkForAdjacent(prevIdx, currentIdx, addIdx)
			col += c
			temp += t
		}
	}
}

func (g *Grid)GoRight(){
	n := g.size
	for row:=0;row<n;row++ {
		temp := n-1
		for col:=n-2;col>-1;col-- {
			nextIdx := g.getEncIndex(row,col+1)
			currentIdx := g.getEncIndex(row, col)
			addIdx := g.getEncIndex(row, temp)
			c , t := g.checkForAdjacent(nextIdx, currentIdx, addIdx)
			col -= c
			temp -= t
		}
	}
}

func (g *Grid)AddRandValue()bool{
	pos,ok :=  g.freeRepo.GiveRandomPos()
	if !ok {
		return false
	}
	g.addValueToPos(pos,g.getRandVal())
	return true
}

func (g *Grid)CheckForSuccess()bool{
	for i:=0;i<g.size;i++{
		for j:=0;j<g.size;j++ {
			val, ok := g.filledRepo.GetVal(g.getEncIndex(i,j))
			if ok && val == models.Success {
				return true
			}
		}
	}
	return false
}


func (g *Grid)GetGrid()[]string{
	ret := make([]string,g.size)
	for i:=0;i<g.size;i++{
		str := ""
		for j:=0;j<g.size;j++ {
			val,ok := g.filledRepo.GetVal(g.getEncIndex(i,j))
			if !ok {
				str += " _"
			} else {
				str = fmt.Sprintf("%v %v", str, val)
			}
		}
		str += "\n"
		ret[i] = str
	}
	return ret
}

func (g *Grid)checkForAdjacent(lastIdx int , currentIdx int, addIdx int)(int,int){
	l,okl := g.filledRepo.GetVal(lastIdx)
	c,okc := g.filledRepo.GetVal(currentIdx)
	_,okt := g.filledRepo.GetVal(addIdx)
	if okl && okc && l == c {
		g.removeValueFromPos(lastIdx)
		g.removeValueFromPos(currentIdx)
		g.addValueToPos(addIdx, l+c)
		return 1,1
	}
	if !okt && okc {
			g.removeValueFromPos(currentIdx)
			g.addValueToPos(addIdx, c)
			return 0,1
	}
	if okt {
		return 0,1
	}
	return 0,0
}

func(g *Grid)removeValueFromPos(idx int){
	g.filledRepo.RemovePos(idx)
	g.freeRepo.AddPos(idx)
}

func(g *Grid)addValueToPos(idx int, val int){
	g.freeRepo.RemovePos(idx)
	g.filledRepo.AddPos(idx, val)
}

func(g *Grid)getEncIndex(a, b int ) int {
	return a*g.size+b
}

func(g *Grid)getRandVal() int {
	pow := rand.Int()%11 + 1
	return int(math.Pow(2,float64(pow)))
}