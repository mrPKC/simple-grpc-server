package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/Pratyush/User/user/protos"
	// main "github.com/Pratyush/User"
)

func TestUserService_GetUsers(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	svc := server{}
	pb.RegisterUsersServer(srv, &svc)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}

	client := pb.NewUsersClient(conn)
	res, err := client.GetUsers(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		t.Fatalf("client.GetUser %v", err)
	}

	if res == nil {
		t.Fatalf("Unexpected values ")
	}
}

func TestUserService_GetUsersById(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	svc := server{}
	pb.RegisterUsersServer(srv, &svc)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}

	client := pb.NewUsersClient(conn)
	res, err := client.GetUserById(context.Background(), &pb.GetUserByIdRequest{Id: []int32{1, 2, 3}})
	if err != nil {
		t.Fatalf("client.GetUser %v", err)
	}

	if res == nil || res.User[0].Id != 1 {
		t.Fatalf("Unexpected values ")
	}
}
