package usc

type USC struct {
    lists map[string] *List
    refreshNeeded bool
    handler Handler
    network Network
}

func MakeUSC(handler Handler) *USC {
    result := &USC{lists:map[string] *List {}, handler:handler}
    result.network = Network{usc:result}

    return result
}

func (u *USC) NeedRefresh() {
    u.refreshNeeded = true
}

func (u *USC) refresh() bool {
    result := u.refreshNeeded
    u.refreshNeeded = false
    return result
}

func (u *USC) toJson() map[string] interface{} {
    result := make(map[string] interface{})

    for name,list := range u.lists {
        result[name] = list.toJson()
    }

    return result
}

func (u *USC) handleCommand(params []Param) string {
    if params == nil || len(params) == 0 {
        return "Empty command"
    }

    return u.handler(params)
}

func (u *USC) AddList(name string) *List {
    if _,ok := u.lists[name]; !ok {
        u.lists[name] = &List{nodes: []Node{}}
    }
    return u.lists[name]
}

func (u *USC) Start(port int) {
    u.network.Start(port)
}

func (u *USC) Stop() {
    u.network.Stop()
}
