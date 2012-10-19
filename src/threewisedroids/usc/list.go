package usc

type List struct {
    nodes []Node
}

func (l *List) AddItem(name string, commandId int, childs ...string) *List {
    node := Node{name:name, commandId:commandId, childs:childs[:]}

    l.nodes = append(l.nodes, node)
    return l
}

func (l *List) toJson() interface{} {
    result := make(map[string] interface{})

    for _,node := range l.nodes {
        var childs = []interface{} {}
        for _,child := range node.childs {
            childs = append(childs, child)
        }

        var array = []interface{} {node.commandId, childs}
        result[node.name] = array
    }

    return result
}

func (l *List) AddSimpleItem(name string) *List {
    return l.AddItem(name, -1)
}
