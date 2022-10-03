package main

import (
	"bookstore/pb"
	"context"
	"testing"
)

func Test_server_ListBooks(t *testing.T) {
	// 初始化
	db, _ := NewDB("test.db")
	s := server{bs: &bookstore{db: db}}

	// rpc请求
	req := &pb.ListBooksRequest{
		Shelf: 4,
	}
	res, err := s.ListBooks(context.Background(), req)
	if err != nil {
		t.Fatalf("s.ListBooks failed, err:%v\n", err)
	}
	t.Logf("next_page_token:%v\n", res.GetNextPageToken())
	for i, book := range res.Books {
		t.Logf("%d: %#v\n", i, book)
	}
}
