package net

import (
	"encoding/json"
	"github.com/nik-johnson-net/go-smarthome/net/encoding"
	"net"
)

// Connection is a TP-Link Smart Home protocol connection
type Connection struct {
	conn net.Conn
}

// NewConnection establishes a connection to a TP-Link Smart Home Protocol device.
func NewConnection(address string) (*Connection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Connection{
		conn: conn,
	}, nil
}

// Write encodes data as a message to the device.
func (c *Connection) Write(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	encryptedData := encoding.Encrypt(jsonData)
	framedData := encoding.Frame(encryptedData)

	_, err = c.conn.Write(framedData)
	return err
}

// Read receives a message parsed as a data structure.
func (c *Connection) Read(data interface{}) error {
	payload, err := encoding.ReadFrame(c.conn)
	if err != nil {
		return err
	}
	jsonData := encoding.Decrypt(payload)
	return json.Unmarshal(jsonData, data)
}

// Close closes the connect once done. It's vital to close connections when done as TP-Link Smart Home devices are small embedded devices that don't accept many concurrent connections.
func (c *Connection) Close() error {
	return c.conn.Close()
}
