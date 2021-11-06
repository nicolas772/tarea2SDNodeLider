package main

import (
	"context"
	"log"
	"net"
	"io/ioutil"
	"strconv"
	"strings"
	pb "sisDistribuidos/nameNode-grpc/dataNode/protodatanode"
	"google.golang.org/grpc"
)

const port = ":50054"

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
	//buscar si ya esta el archivo
	var path = "dataNode/server-dataNode/jugador_" + strconv.Itoa(int(in.GetNumeroJugador())) + "__ronda_" + strconv.Itoa(int(in.GetRonda())) + ".txt"

	contenido, _ := ioutil.ReadFile(path)
	var mensaje string = strconv.Itoa(int(in.GetJugada()))

	contenido  = append(contenido, []byte(mensaje)...)
	err := ioutil.WriteFile(path, contenido, 0644)
	
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(contenido), " ")//todas las jugadas en orden
	var jugadas []int32
	for i := 0; i < len(split); i++ {
		numero, _ := strconv.Atoi(split[i])
		jugadas = append(jugadas, int32(numero))
	}

	s.registro.Jugadas = jugadas
	return s.registro, nil
}

func main() {
	var info_jugada *InformarJugadasServer = NewInformarJugadasServer()
	if err := info_jugada.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}