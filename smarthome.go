package smarthome

import (
	"github.com/nik-johnson-net/go-smarthome/net"
	stdnet "net"
)

// Client represents the API to a TP-Link Smart Home device.
type Client struct {
	destination string
}

// NewClient creates a Client bound to the device at destination.
func NewClient(destination string) *Client {
	return &Client{
		destination: stdnet.JoinHostPort(destination, "9999"),
	}
}

func (c *Client) sendCommand(data interface{}, response interface{}) error {
	conn, err := net.NewConnection(c.destination)
	if err != nil {
		return err
	}

	if err = conn.Write(data); err != nil {
		return err
	}

	if err = conn.Read(response); err != nil {
		return err
	}

	return conn.Close()
}

// SysInfo returns the device info.
func (c *Client) SysInfo() (SysInfoResponse, error) {
	request := GenericRequest{
		System: &SystemRequests{
			GetSysinfo: &GetSysinfoRequest{},
		},
	}
	var response GenericResponse
	if err := c.sendCommand(request, &response); err != nil {
		return SysInfoResponse{}, err
	}
	return *response.System.SysInfo, nil
}

// EMeter returns power information for Smart Plug devices.
func (c *Client) EMeter(childID ...string) (GetRealtimeResponse, error) {
	request := GenericRequest{
		Context: &Context{
			ChildIds: childID,
		},
		EMeter: &EMeterRequests{
			GetRealtime: &GetRealtimeRequest{},
		},
	}

	var response GenericResponse
	if err := c.sendCommand(request, &response); err != nil {
		return GetRealtimeResponse{}, err
	}

	emeter := *response.EMeter.GetRealtime
	normalizeEmeterResponse(&emeter)
	return emeter, nil
}
