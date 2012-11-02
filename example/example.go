package main

import (
    "time"
    "fmt"
    "github.com/Gyscos/gotools/usc"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(4)

    server := usc.MakeUSC(func(params []usc.Param) string {
        method := params[0].Content
        methodId := params[0].CommandId
        fmt.Println("Method : " + method)

        switch methodId {

        case 0:
            return "Starting !"
        case 1:
            return "Stopping !"
        }
        return "Not understood..."
    })

    server.AddList("entry").
        AddItem("start", 0).
        AddItem("stop", 1, "number")
    server.AddList("number")

    go server.Start(3333)

    time.Sleep(10 * time.Second)

    server.Stop()
}
