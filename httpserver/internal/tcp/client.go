package tcp

import (
	"log"
	"net"
)

type HTTPMessage struct {
	Method  string
	Path    string
	Version string
	Body    string
}

type TCPClient struct {
	ListenAddr  string
	SendChan    chan string
	ReceiveChan chan HTTPMessage
	Ln          net.Listener
}

/**
 * Creates and initializes a new TCP client, sets up send and receive channels
 * for message handling and starts a goroutine to accept incoming TCP connections.
 *
 * @param {string} listenAddr - The address and port the client should listen.
 * @returns {TCPClient} - A TCPClient instance.
 * @throws Returns an error if the TCP listener fails.
 */
func CreateTcpClient(listenAddr string, instance string) (*TCPClient, error) {
	tcpClient := &TCPClient{
		ListenAddr:  listenAddr,
		SendChan:    make(chan string, 10),
		ReceiveChan: make(chan HTTPMessage, 10),
	}

	listener, error := net.Listen("tcp", listenAddr)
	if error != nil {
		return nil, error
	}
	tcpClient.Ln = listener

	go tcpClient.acceptLoop(instance)
	return tcpClient, nil
}

/**
 * Accepts incoming TCP connections on the listener with a loop.
 * For each new connection, a goroutine is created to handle it concurrently.
 * Closes the listener when the loop ends.
 */
func (client *TCPClient) acceptLoop(instance string) {
	defer client.Ln.Close()

	log.Printf("[%s] HTTP Server listening on %s\n", instance, client.ListenAddr)

	for {
		connection, error := client.Ln.Accept()
		if error != nil {
			log.Printf("[%s] Error accepting connection:", instance)
			continue
		}
		go client.handleConnection(connection)
	}
}
