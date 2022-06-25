package repo

import "math/rand"

type FreeGrid struct {
	mp  []int
}

func NewFreeGrid(n int)*FreeGrid {
	g := make([]int, n*n)
	for i:=0;i<n*n;i++ {
		g[i] = i
	}
	return &FreeGrid{
		mp: g,
	}
}

func (g *FreeGrid) GiveRandomPos() (int, bool) {
	n := len(g.mp)
	if n == 0 {
		return -1, false
	}
	return g.mp[rand.Int()%n], true
}

func (g *FreeGrid) AddPos(pos int) {
	g.mp = append(g.mp, pos)
}

func (g *FreeGrid) RemovePos(pos int) {
	temp := make([]int, len(g.mp)-1)
	i := 0
	for _, val := range g.mp {
		if val == pos {
			continue
		}
		temp[i] = val
		i++
	}
	g.mp = temp
}

