package main

import (
	"client/pbs"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var gorpc_conn *grpc.ClientConn
var noderpc_conn *grpc.ClientConn

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": "helloworld",
	})
}

func GetRpcMessage(w http.ResponseWriter, r *http.Request) {

	message := pbs.Message{
		Body: "client",
	}

	go_c := pbs.NewChatServiceClient(gorpc_conn)
	node_c := pbs.NewChatServiceClient(noderpc_conn)

	go_response, err := go_c.SayHello(context.Background(), &message)
	node_response, err2 := node_c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatal("err : ", err)
	}
	if err2 != nil {
		log.Fatal("err2 : ", err2)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"GoRpc":   go_response.Body,
		"NodeRpc": node_response.Body,
	})
}

func main() {

	var err error
	var err2 error

	gorpc_conn, err = grpc.Dial("localhost:3001", grpc.WithInsecure())
	noderpc_conn, err2 = grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err : ", err)
	}
	if err2 != nil {
		log.Fatal("err2 : ", err2)
	}

	defer gorpc_conn.Close()
	defer noderpc_conn.Close()
	/*
		message := pbs.Message{
			Body: "client",
		}

		go_c := pbs.NewChatServiceClient(gorpc_conn)
		node_c := pbs.NewChatServiceClient(noderpc_conn)

		go_response, err := go_c.SayHello(context.Background(), &message)
		node_response, err2 := node_c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatal("err : ", err)
		}
		if err2 != nil {
			log.Fatal("err2 : ", err2)
		}
		log.Printf("From Go Server : %s", go_response.Body)
		log.Printf("From Node Server : %s", node_response.Body)
	*/

	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorld).Methods("GET")
	r.HandleFunc("/rpc", GetRpcMessage).Methods("GET")
	err3 := http.ListenAndServe(":3000", r)
	if err3 != nil {
		log.Fatal("err3 : ", err3)
	}
}
