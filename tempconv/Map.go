package tempconv

var graph = make(map[string]map[string]bool)

//初始化map[string]map[string]bool的 惯用方式
func AddEdge(key1, key2 string) {
	edges := graph[key1]
	if edges == nil {
		edges = make(map[string]bool)
		graph[key1] = edges
	}
	edges[key2] = true
}

func hasEdge(key1, key2 string) bool {
	return graph[key1][key2]
}
