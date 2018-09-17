package tempconv

//二叉树实现 插入排序
type tree struct {
	value       int
	left, right *tree
}

//就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root) //传入一个空的切片，然后在后插入排序好的数据
}

func appendValues(values []int, t *tree) []int {
	if t != nil { //中序遍历
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

//创建二叉树  并排序
func add(t *tree, value int) *tree {
	if t == nil {
		//如果还没创建根节点
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
