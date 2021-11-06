package main

import (
	"context"
	"log"
	"math/rand"
	pb "sisDistribuidos/nameNode-grpc/nodeName/protonodename"

	"google.golang.org/grpc"
)

const (
	address = "10.6.40.196:40051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewInformarJugadasClient(conn)

	for i := 1; i < 11; i++ {

		var NumeroJugador int = rand.Intn(14)
		var Ronda int = rand.Intn(3)
		var jugada int = rand.Intn(3)

		r, err := serviceClient.InfoJugadas(context.Background(), &pb.InfoJugada{NumeroJugador: int32(NumeroJugador), Ronda: int32(Ronda), Jugada: int32(jugada)})

		if err != nil {
			log.Fatalf("could not informar jugada: %v", err)
		}
		log.Printf("Ronda %v ,Jugadas: %v "+"del jugador %v", Ronda, r.GetJugadas(), NumeroJugador)
	}
}
