package skiplist

type Node struct {
	key int
	value interface{}
	levels []*Node
}

type SkipList struct{
	step int
	head *Node
	// 从0开始
	topLevel int
	count int
}

func NewSkipList(step int) *SkipList{
	var list SkipList
	list.step = step
	list.topLevel = 0
	list.count = 0
	return &list
}

func ( list *SkipList) Add(key int, value interface{}){

}

func ( list *SkipList) Delete(key int) bool{

}

func ( list *SkipList) Find(key int) (*Node, bool){

}