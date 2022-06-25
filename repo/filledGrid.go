package repo

type FilledGrid struct {
	mp map[int]int
}

func NewFilledGrid()*FilledGrid{
	return &FilledGrid{
		mp : make(map[int]int, 0),
	}
}

func(f *FilledGrid)AddPos(idx int, val int){
	f.mp[idx] = val
}

func(f *FilledGrid)RemovePos(idx int){
	delete(f.mp,idx)
}

func(f *FilledGrid)GetVal(idx int) (int, bool){
	val ,ok := f.mp[idx]
	if ok {
		return val, true
	}

	return 0, false
}