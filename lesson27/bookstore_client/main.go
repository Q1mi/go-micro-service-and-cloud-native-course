package main

import (
	"bookstore_client/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 拨号 连接
	conn, err := grpc.Dial("127.0.0.1:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.Dial 127.0.0.1:8091 failed, err:%v\n", err)
		return
	}

	defer conn.Close()

	// 创建客户端
	c := pb.NewBookstoreClient(conn)

	res, err := c.ListBooks(context.Background(), &pb.ListBooksRequest{Shelf: 4})
	if err != nil {
		fmt.Printf("c.ListBooks failed, err:%v\n", err)
		return
	}
	fmt.Printf("next_page_token:%v\n", res.NextPageToken)
	for i, book := range res.Books {
		fmt.Printf("%d: %#v\n", i, book)
	}

}
