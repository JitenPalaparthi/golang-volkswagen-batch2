package main

import (
	"context"
	"demo/database"
	"demo/models"
	"demo/proto/userpb"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (us *UserServiceServer) CreateUser(ctx context.Context, req *userpb.UserRequest) (res *userpb.User, err error) {
	u := new(models.User)
	u.Name = req.Name
	u.Email = req.Email
	u.Status = "active"
	u.LastModified = time.Now().Unix()
	//u.Id = uint(rand.Intn(1000))
	u, err = database.CreateUser(us.DB, u)
	if err != nil {
		return nil, err
	}
	return &userpb.User{Id: int64(u.Id), Name: u.Name, Email: u.Email, Status: u.Status, LastModified: u.LastModified}, nil
}

func (us *UserServiceServer) GenerateRandomNumber(emt *emptypb.Empty, out grpc.ServerStreamingServer[userpb.RandomResponse]) error {
	stop := time.After(time.Second * 20)
out:
	for {
		select {
		case <-stop:
			break out

		default:
			time.Sleep(time.Second * 1)
			err := out.Send(&userpb.RandomResponse{Num: int64(rand.Intn(1000))})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {

	DSN := os.Getenv("DB_CONN")
	if DSN == "" {
		DSN = `admin:admin@tcp(127.0.0.1:3306)/demodb?charset=utf8mb4&parseTime=True&loc=Local`
	}

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	userServer := &UserServiceServer{DB: db}
	userpb.RegisterUserServiceServer(server, userServer)

	log.Println("gRPC service is running on port 50091")
	lis, err := net.Listen("tcp", ":50091")
	if err != nil {
		log.Fatalln(err)
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

}
