package filesystem

type Inode struct {
	Name     string
	Parent   *Inode
	Children []*Inode
	Size     int
	IsDir    bool
}

func NewInode(name string, parent *Inode, size int, isDir bool) *Inode {
	return &Inode{
		Name:   name,
		Parent: parent,
		Size:   size,
		IsDir:  isDir,
	}
}

func (i *Inode) AddChild(child *Inode) {
	// add the child inode and update the size of all parents
	i.Children = append(i.Children, child)
	if child.Size != 0 {
		i.Size += child.Size
		// update parent sizes
		parent := i.Parent
		for parent != nil {
			parent.Size += child.Size
			parent = parent.Parent
		}
	}
}
