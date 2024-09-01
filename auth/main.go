package main

import (
	"context"
	"log"
	"net"
	"time"

	//"example.com/m/auth"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

var jwtKey = []byte("my_secret_key")

type server struct {
	auth.UnimplementedAuthServiceServer
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *server) Authenticate(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	if req.Username != "exampleuser" || req.Password != "examplepassword" {
		return &auth.AuthResponse{Message: "Invalid credentials"}, nil
	}

	expirationTime := time.Now().Add(time.Minute * 15)
	claims := &Claims{
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &auth.AuthResponse{Token: tokenString, Message: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &server{})

	log.Printf("Server is listening on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
