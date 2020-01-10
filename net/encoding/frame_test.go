package encoding

import (
	"bytes"
	"io"
	"testing"
)

var (
	unframed   = []byte{0x1, 0x2, 0x3, 0x4}
	framed     = []byte{0x0, 0x0, 0x0, 0x4, 0x1, 0x2, 0x3, 0x4}
	largeFrame = []byte{0xAC, 0x0, 0x0, 0x0, 0x1, 0x2}
)

func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error(err)
	}
}

func TestFrame(t *testing.T) {
	result := Frame(unframed)
	assertEqual(t, framed, result)
}

func TestReadFrame(t *testing.T) {
	var buffer bytes.Buffer
	buffer.Write(framed)
	result, err := ReadFrame(&buffer)
	assertNoErr(t, err)
	assertEqual(t, unframed, result)
}

func TestReadFrameLarge(t *testing.T) {
	var buffer bytes.Buffer
	buffer.Write(largeFrame)
	result, err := ReadFrame(&buffer)
	assertEqual(t, ErrFrameTooLarge, err)
	assertEqual(t, []byte(nil), result)
}

func TestReadFrameErr(t *testing.T) {
	var buffer bytes.Buffer
	result, err := ReadFrame(&buffer)
	assertEqual(t, io.EOF, err)
	assertEqual(t, []byte(nil), result)
}
