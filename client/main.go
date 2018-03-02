package main

import (
	"log"

	"context"

	pb "github.com/rongfengliang/restyapp/service/pb"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("some wrong")
	}
	echoclient := pb.NewEchoServiceClient(con)
	userclient := pb.NewUserServiceAppClient(con)
	result, err := echoclient.Echo(context.Background(), &pb.EchoMessage{
		Message: "demapppp",
	})
	if err != nil {
		log.Fatal("some wrong")
	}
	log.Println(result.Message)
	result2, err := userclient.Message(context.Background(), &pb.AppMessage{
		Message: "demappsssssspp",
	})
	log.Println(result2.Message)

}
