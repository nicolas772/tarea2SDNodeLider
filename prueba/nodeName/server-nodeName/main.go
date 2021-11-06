package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"io/ioutil"
	"strconv"
	pb "sisDistribuidos/nameNode-grpc/nodeName/protonodename"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewInformarJugadasServer() *InformarJugadasServer {
	return &InformarJugadasServer{
		registro: &pb.Registro{},
	}
}

type InformarJugadasServer struct {
	pb.UnimplementedInformarJugadasServer
	registro *pb.Registro
}

func (server *InformarJugadasServer) Run() error{
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInformarJugadasServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} 
	return s.Serve(lis)
}

func (s *InformarJugadasServer) InfoJugadas(ctx context.Context, in *pb.InfoJugada) (*pb.Registro, error) {
	var ips  = [3] string{"1.1","1.2","1.3"}
	var ipAsign = ips[rand.Intn(3)]
	var path = "nodeName/server-nodeName/registroGeneral.txt"
	var mensaje string = "Jugador_"+ strconv.Itoa(int(in.GetNumeroJugador())) + " Ronda_"+ strconv.Itoa(int(in.GetRonda())) + " " + string(ipAsign) + "\n"
	
	contenido, _ := ioutil.ReadFile(path)
	
	contenido  = append(contenido, []byte(mensaje)...)

    err := ioutil.WriteFile(path, contenido, 0644)
	
    if err != nil {
        log.Fatal(err)
    }
	
	var pointer *int32
	var entero = int32(in.GetJugada())
	pointer = &entero
	for i := 1; i < 6; i++ {
		var entero2 = rand.Intn(10)
		entero = int32(entero2)
		s.registro.Jugadas = append(s.registro.Jugadas, *pointer)
	}
	return s.registro, nil
}

func main() {
	var info_jugada *InformarJugadasServer = NewInformarJugadasServer()
	if err := info_jugada.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}