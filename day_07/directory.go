package main

type Directory struct {
	Name      string
	ParentDir *Directory
	ChildDir  []*Directory
	Size      int
}

func NewDir(parent *Directory, name string) Directory {
	return Directory{ChildDir: []*Directory{}, Size: 0, ParentDir: parent, Name: name}
}

func (dir *Directory) getSizeUnderTreshold(treshold int) int {
	sum := 0
	if dir.Size < treshold {
		sum += dir.Size
	}
	for _, child := range dir.ChildDir {
		sum += child.getSizeUnderTreshold(treshold)
	}
	return sum
}

func (dir *Directory) getDirsForSpace(target int) int {
	min := 70000000

	if dir.Size >= target && dir.Size < min {
		min = dir.Size
	}
	for _, child := range dir.ChildDir {
		res := child.getDirsForSpace(target)
		if res < min {
			min = res
		}
	}
	return min
}


func (dir *Directory) addSizeToParent(size int) {
	for {
		if dir.ParentDir == nil {
			break
		}
		dir.ParentDir.Size += size
		dir = dir.ParentDir
	}
}