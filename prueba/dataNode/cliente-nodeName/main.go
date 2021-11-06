package main

import (
	"context"
	"log"
	"math/rand"
	pb "sisDistribuidos/nameNode-grpc/dataNode/protodatanode"

	"google.golang.org/grpc"
)

func main() {

	var ips = [3]string{"localhost:50054", "localhost:50068", "localhost:50054"}
	var ipAsign = ips[rand.Intn(3)] //random

	conn, err := grpc.Dial(ipAsign, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewInformarJugadasClient(conn)

	for i := 1; i < 30; i++ {

		var NumeroJugador int = rand.Intn(14)
		var Ronda int = rand.Intn(3)
		var jugada int = rand.Intn(3)

		r, err := serviceClient.InfoJugadas(context.Background(), &pb.InfoJugada{NumeroJugador: int32(NumeroJugador), Ronda: int32(Ronda), Jugada: int32(jugada)})
		if err != nil {
			log.Fatalf("could not informar jugada: %v", err)
		}
		log.Printf("Jugadas jugador %v %v \n", r.GetJugadas(), NumeroJugador)
	}
}
