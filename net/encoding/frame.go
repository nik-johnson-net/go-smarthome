package encoding

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	frameHeaderSize = 4
	maxFrameSize    = 1 << 30 // 1 GB
)

var (
	// ErrFrameTooLarge means an incoming frame exceeds limits
	ErrFrameTooLarge = errors.New("recv frame exceeds max frame size of 1GB")
)

// Frame prepends data with the size of data and returns the new data byte slice.
func Frame(data []byte) []byte {
	frameSize := len(data)
	frameBuffer := make([]byte, frameHeaderSize, frameSize+frameHeaderSize)
	binary.BigEndian.PutUint32(frameBuffer, uint32(frameSize))
	return append(frameBuffer, data...)
}

// ReadFrame reads a full frame from r and returns the payload.
func ReadFrame(r io.Reader) ([]byte, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, err
	}

	frameSize := binary.BigEndian.Uint32(header)
	if frameSize > maxFrameSize {
		return nil, ErrFrameTooLarge
	}
	frameBuffer := make([]byte, frameSize)

	_, err := io.ReadFull(r, frameBuffer)
	return frameBuffer, err
}
