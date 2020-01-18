package main

import (
	"encoding/json"
	"flag"
	"github.com/phillipahereza/go_microservices/gelftail/aggregator"
	"github.com/phillipahereza/go_microservices/gelftail/transformer"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"os"
	"sync"
)

var authToken = ""
var port *string

func init() {
	data, err := ioutil.ReadFile("token.txt")
	if err != nil {
		msg := "Can not find token.txt that should contain the loggly token"
		logrus.Errorln(msg)
	}
	authToken = string(data)

	port = flag.String("port", "12202", "UDP port for the geftail")
	flag.Parse()
}

func main() {
	logrus.Println("Starting Gelf-Tail server .....")
	serverConn := startUDPServer(*port)
	defer serverConn.Close()

	var bulkQueue = make(chan []byte, 1)

	go aggregator.Start(bulkQueue, authToken)
	go listenForLogStatements(serverConn, bulkQueue)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func listenForLogStatements(conn *net.UDPConn, queue chan []byte) {
	buf := make([]byte, 8192) // buffer that can store UDP payload of 8kb
	var item map[string]interface{}
	for {
		n, _, err := conn.ReadFromUDP(buf)
		logrus.Debugf("Buffer contains: %s", string(buf[0:n]))
		if err != nil {
			logrus.Errorf("Problem readin UDP message into buffer: %v\n", err.Error())
			continue
		}
		err = json.Unmarshal(buf[0:n], &item)
		if err != nil {
			logrus.Errorln("Error unmarshalling log into JSON: " + err.Error())
			item = nil
			continue
		}

		processedLogMessage, err := transformer.ProcessLogStatement(item)
		if err != nil {
			logrus.Printf("Problem parsing message: %v", string(buf[0:n]))
		} else {
			queue <- processedLogMessage
		}
		item = nil

	}
}

func startUDPServer(port string) *net.UDPConn {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	checkError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)

	return ServerConn
}

func checkError(err error) {
	if err != nil {
		logrus.Println("Error: ", err)
		os.Exit(0)
	}
}
