package jsonrpc

import (
    "encoding/json"
    "fmt"
    "net"
    "strconv"
)

type JsonRpcHandler func(method string, params interface{}) (interface{}, interface{})

type JsonRpcServer struct {
    listener net.Listener
    control  chan int
}

type jsonRequest struct {
    Method string      `json:"method"`
    Params interface{} `json:"params"`
    Id     int         `json:"id"`
}

type jsonResponse struct {
    Id     int         `json:"id"`
    Result interface{} `json:"result"`
    Error  interface{} `json:"error"`
}

func handleClient(conn net.Conn, handler JsonRpcHandler) {
    defer conn.Close()

    var req jsonRequest

    decoder := json.NewDecoder(conn)
    decoder.Decode(&req)

    result, err := handler(req.Method, req.Params)

    var response = jsonResponse{Result: result, Error: err, Id: req.Id}

    encoder := json.NewEncoder(conn)
    encoder.Encode(response)
}

func (s *JsonRpcServer) run(handler JsonRpcHandler) {
    for {
        select {

        case <-s.control:
            s.control <- 2
            // fmt.Println("Routine terminated")
            return
        default:
            // fmt.Println("Accepting...")
            conn, err := s.listener.Accept()
            if err != nil {
                // fmt.Printf("Error happened : %v\n", err)
                continue
            }
            go handleClient(conn, handler)
        }
    }
}

func (s *JsonRpcServer) Start(port int, handler JsonRpcHandler) {
    s.control = make(chan int)

    host := "localhost"
    addr := host + ":" + strconv.Itoa(port)

    l, err := net.Listen("tcp", addr)

    if err != nil {
        fmt.Printf("Error : %v\n", err)
        return
    }

    s.listener = l

    go s.run(handler)
}

func (s *JsonRpcServer) Stop() {
    s.listener.Close()
    s.control <- 2
}

func RpcQuery(host string, port int, method string, params interface{}) (interface{}, interface{}) {
    addr := host + ":" + strconv.Itoa(port)

    conn, err := net.Dial("tcp", addr)

    if err != nil {
        return nil, err
    }

    var req = jsonRequest{Id: 42, Method: method, Params: params}

    encoder := json.NewEncoder(conn)
    encoder.Encode(req)

    var response jsonResponse

    decoder := json.NewDecoder(conn)
    decoder.Decode(&response)

    return response.Result, response.Error
}
