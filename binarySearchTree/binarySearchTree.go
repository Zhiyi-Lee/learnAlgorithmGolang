package binarySearchTree

type Note struct {
	key   int
	value interface{}
	left  *Note
	right *Note
}

type BST struct {
	count int
	root  *Note
}

func (t *BST) Size() int {
	return t.count
}

func (t *BST) IsEmply() bool {
	return t.count == 0
}

func (t *BST) Insert(key int, value interface{}) {
	t.root = t.insert(t.root, key, value)
}

//查找键值为key的元素是否存在
func (t *BST) Contain(key int) bool {
	return t.contain(t.root, key)
}

//查找键值为key的元素的value
func (t *BST) Search(key int) *interface{} {
	return t.search(t.root, key)
}

//向以node为根的二叉搜索树中插入节点(key,value)
//返回插入新节点后的新的根node
//O(logn)
func (t *BST) insert(note *Note, key int, value interface{}) *Note {
	//找到空的插入点 则跳出递归
	if note == nil {
		t.count++
		return &Note{key, value, nil, nil}
	}

	if key == note.key {
		//key已经存在则更新value
		note.value = value
	} else if key > note.key {
		note.right = t.insert(note.right, key, value)
	} else {
		note.left = t.insert(note.left, key, value)
	}

	//返回新的根节点
	return note
}

//在note中搜索键值为key的元素
//返回bool
func (t *BST) contain(note *Note, key int) bool {
	if note == nil {
		return false
	}

	if key == note.key {
		return true
	} else if key > note.key {
		return t.contain(note.right, key)
	} else {
		return t.contain(note.left, key)
	}
}

//在note中搜索键值为key的元素的value
func (t *BST) search(note *Note, key int) *interface{} {
	if note == nil {
		return nil
	}

	if key == note.key {
		return &note.value
	} else if key > note.key {
		return t.search(note.right, key)
	} else {
		return t.search(note.left, key)
	}
}
