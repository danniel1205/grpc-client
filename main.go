package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/danniel1205/grpc-service/helloservice"
	"google.golang.org/grpc"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	fmt.Println(fmt.Sprintf("Server address is %s", *serverAddr))
	var opts []grpc.DialOption
	//if *tls {
	//	if *caFile == "" {
	//		*caFile = data.Path("x509/ca_cert.pem")
	//	}
	//	creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
	//	if err != nil {
	//		log.Fatalf("Failed to create TLS credentials %v", err)
	//	}
	//	opts = append(opts, grpc.WithTransportCredentials(creds))
	//} else {
	//	opts = append(opts, grpc.WithInsecure())
	//}
	opts = append(opts, grpc.WithInsecure())

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := helloservice.NewHelloServiceClient(conn)

	for {
		resp, _ := client.SayHello(context.Background(), &helloservice.Request{
			Name: "Daniel",
			From: "Beijing",
		})

		if resp != nil {
			fmt.Println(resp.GetMessage())
		} else {
			fmt.Println("Server returns nil")
		}

		time.Sleep(5 * time.Second)
	}
}
