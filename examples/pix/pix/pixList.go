package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	gn := gerencianet.NewGerencianet(credentials)

	
	const inicio = "2021-03-01T03:01:35Z"
	const fim = "2021-03-05T22:01:35Z"


	res, err := gn.PixList(inicio, fim) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}