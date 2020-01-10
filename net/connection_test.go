package net

import (
	"net"
	"reflect"
	"testing"
)

var (
	inputData      = FakeData{"foo"}
	ciphertextData = []byte{0x00, 0x00, 0x00, 0x0D, 0xD0, 0xF2, 0xB4, 0xDB, 0xB4, 0x96, 0xAC, 0x8E, 0xE8, 0x87, 0xE8, 0xCA, 0xB7}
)

type FakeData struct {
	Foo string
}

type FakeServer struct {
	listener *net.TCPListener
}

func NewFakeServer(callback func([]byte) []byte) (*FakeServer, error) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv6loopback})
	if err != nil {
		return nil, err
	}

	go func(cb func([]byte) []byte) {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			buffer := make([]byte, 256)
			n, err := conn.Read(buffer)
			if err != nil {
				return
			}
			conn.Write(cb(buffer[:n]))
			conn.Close()
		}
	}(callback)

	return &FakeServer{
		listener: listener,
	}, nil
}

func (f *FakeServer) Close() error {
	return f.listener.Close()
}

func (f *FakeServer) Addr() net.Addr {
	return f.listener.Addr()
}

func TestNewConnection(t *testing.T) {
	server, err := NewFakeServer(nil)
	if err != nil {
		t.Error(err)
	}
	defer server.Close()

	_, err = NewConnection(server.Addr().String())
	if err != nil {
		t.Error(err)
	}
}

func TestNewConnectionErr(t *testing.T) {
	_, err := NewConnection("localhost:1")
	if err == nil {
		t.Error("No error was raised")
	}
}

func TestConnectionWrite(t *testing.T) {
	server, err := NewFakeServer(func(input []byte) []byte {
		if !reflect.DeepEqual(ciphertextData, input) {
			t.Errorf("Expected %v, got %v", ciphertextData, input)
		}
		return ciphertextData
	})
	if err != nil {
		t.Error(err)
	}
	defer server.Close()

	conn, err := NewConnection(server.Addr().String())
	if err != nil {
		t.Error(err)
	}

	err = conn.Write(FakeData{"foo"})
	if err != nil {
		t.Error(err)
	}

	var result FakeData
	err = conn.Read(&result)
	if err != nil {
		t.Error(err)
	}
	if result.Foo != "foo" {
		t.Errorf("Expected foo, got %v", result.Foo)
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}
