package usc

import (
    "threewisedroids/jsonrpc"
)

type Network struct {
    usc    *USC
    server jsonrpc.JsonRpcServer
}

func (n *Network) Start(port int) {
    n.server.Start(port,
        func(method string, params interface{}) (interface{}, interface{}) {
            switch method {
            case "list":
                return n.usc.toJson(), nil
            case "call":
                return n.call(params.([]interface{}))
            }
            return nil, "Cannot find method"
        })
}

func (n *Network) Stop() {
    n.server.Stop()
}

func (n *Network) call(params []interface{}) (interface{}, interface{}) {
    p := []Param{}

    for _, arg := range params {
        a_arg := arg.([]interface{})
        param := Param{int(a_arg[1].(float64)), a_arg[0].(string)}
        p = append(p, param)
    }

    answer := n.usc.handleCommand(p)
    refresh := n.usc.refresh()

    result := make(map[string]interface{})

    result["answer"] = answer
    result["refresh"] = refresh
    if refresh {
        result["root"] = n.usc.toJson()
    }

    return result, nil
}
