package util

import (
	"net"
	"strconv"
)

// GetRandomTCPPort returns a random available TCP port.
// It creates a listener on a random port and returns the port number.
// The listener is closed before returning the port number.
// If an error occurs, it returns 0 and the error.
func GetRandomTCPPort() (int, error) {
	// Create a listener on a random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	// Get the port from the listener
	port := listener.Addr().(*net.TCPAddr).Port
	return port, nil
}

// GetRandomTCPPortString returns a random available TCP port as a string.
// It calls GetRandomTCPPort and converts the port number to a string.
func GetRandomTCPPortString() (string, error) {
	port, err := GetRandomTCPPort()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(port), nil
}

// GetRandomPort returns a random available port (TCP or UDP).
// It creates a listener on a random port and returns the port number.
// The listener is closed before returning the port number.
// If an error occurs, it returns 0 and the error.
func GetRandomUDPPort() (int, error) {
	// Create a listener on a random port
	listener, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	// Get the port from the listener
	port := listener.LocalAddr().(*net.UDPAddr).Port
	return port, nil
}

// GetRandomUDPPortString returns a random available UDP port as a string.
// It calls GetRandomUDPPort and converts the port number to a string.
func GetRandomUDPPortString() (string, error) {
	port, err := GetRandomUDPPort()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(port), nil
}

// GetRandomBothPort returns a random available port that can be used for both TCP and UDP.
// It first gets a random TCP port, then attempts to open a UDP listener on that port.
// If successful, it returns the port number; otherwise, it returns 0 and an error
func GetRandomBothPort() (int, error) {
	// First, get a random TCP Port
	tcpPort, err := GetRandomTCPPort()
	if err != nil {
		return 0, err
	}
	// Store it, then attempt to open tcpPort as a UDP Listener
	udpListener, err := net.ListenPacket("udp", ":"+strconv.Itoa(tcpPort))
	if err != nil {
		return 0, err
	}
	defer udpListener.Close()

	// If we reach this point, we have a valid TCP and UDP listener port
	return tcpPort, nil
}

// GetRandomBothPortString returns a random available port that can be used for both TCP and UDP as a string.
func GetRandomBothPortString() (string, error) {
	port, err := GetRandomBothPort()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(port), nil
}
