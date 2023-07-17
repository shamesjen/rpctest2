package main

import (
	"context"
	"log"
	calculator "rpctest2/kitex_gen/calculator"

	"github.com/cloudwego/kitex/client"
)

func main() {
    cli, err := calculator.NewClient("calculator", client.WithHostPorts("0.0.0.0:8888"))
    if err != nil {
        panic(err)
    }
    req := calculator.NewRequest()
    req.Num1 = 50
	req.Num2 = 100
	resp, err := cli.Add(context.Background(), req, req)
    if err != nil {
        panic(err)
    }
    // resp.Msg == "world"
	log.Println(resp)
}