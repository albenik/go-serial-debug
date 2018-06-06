package serialdebug

import (
	"fmt"

	"errors"
	"github.com/albenik/iolog"
)

type OpenFunc func() (SerialPort, error)

type SerialPort interface {
	fmt.Stringer

	SetReadTimeout(t int) error
	SetReadTimeoutEx(t, i uint32) error
	SetFirstByteReadTimeout(t uint32) error
	SetWriteTimeout(t int) error
	ReadyToRead() (r uint32, err error)
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	ResetInputBuffer() error
	ResetOutputBuffer() error
	SetDTR(dtr bool) error
	SetRTS(rts bool) error
	Close() error
}

type PortWrapper struct {
	open OpenFunc
	port SerialPort
	log  *iolog.IOLog
}

func NewWrapper(open OpenFunc, log *iolog.IOLog) *PortWrapper {
	return &PortWrapper{open: open, log: log}
}

func (pw *PortWrapper) Open() (SerialPort, error) {
	if pw.port != nil {
		return nil, errors.New("wrapped port already opened")
	}

	port, err := pw.log.LogAny("open", func() (interface{}, error) {
		return pw.open()
	})

	if err != nil {
		return nil, err
	}

	pw.port = port.(SerialPort)
	return pw, err
}

func (pw *PortWrapper) Close() error {
	_, err := pw.log.LogAny("close", func() (interface{}, error) {
		if pw.port == nil {
			return nil, errors.New("wrapped port already closed")
		}
		err := pw.port.Close()
		return nil, err
	})
	pw.port = nil
	return err
}

func (pw *PortWrapper) String() string {
	return pw.port.String()
}

func (pw *PortWrapper) SetReadTimeout(t int) error {
	_, err := pw.log.LogAny("set_read_timeout", func() (interface{}, error) {
		err := pw.port.SetReadTimeout(t)
		return t, err
	})
	return err
}

func (pw *PortWrapper) SetReadTimeoutEx(t, i uint32) error {
	_, err := pw.log.LogAny("set_read_timeout_ex", func() (interface{}, error) {
		err := pw.port.SetReadTimeoutEx(t, i)
		return &struct{ T, I uint32 }{t, i}, err
	})
	return err
}

func (pw *PortWrapper) SetFirstByteReadTimeout(t uint32) error {
	_, err := pw.log.LogAny("set_first_byte_read_timeout", func() (interface{}, error) {
		err := pw.port.SetFirstByteReadTimeout(t)
		return t, err
	})
	return err
}

func (pw *PortWrapper) SetWriteTimeout(t int) error {
	_, err := pw.log.LogAny("set_write_timeout", func() (interface{}, error) {
		err := pw.port.SetWriteTimeout(t)
		return t, err
	})
	return err
}

func (pw *PortWrapper) ReadyToRead() (r uint32, err error) {
	pw.log.LogAny("ready_to_read", func() (interface{}, error) {
		r, err = pw.port.ReadyToRead()
		return r, err
	})
	return
}

func (pw *PortWrapper) Read(p []byte) (n int, err error) {
	return pw.log.LogIO("read", pw.port.Read, p)
}

func (pw *PortWrapper) Write(p []byte) (n int, err error) {
	return pw.log.LogIO("write", pw.port.Write, p)
}

func (pw *PortWrapper) ResetInputBuffer() error {
	_, err := pw.log.LogAny("reset_input_buffer", func() (interface{}, error) {
		err := pw.port.ResetInputBuffer()
		return nil, err
	})
	return err
}

func (pw *PortWrapper) ResetOutputBuffer() error {
	_, err := pw.log.LogAny("reset_output_buffer", func() (interface{}, error) {
		err := pw.port.ResetOutputBuffer()
		return nil, err
	})
	return err
}

func (pw *PortWrapper) SetDTR(dtr bool) error {
	_, err := pw.log.LogAny("set_dtr", func() (interface{}, error) {
		err := pw.port.SetDTR(dtr)
		return dtr, err
	})
	return err
}

func (pw *PortWrapper) SetRTS(rts bool) error {
	_, err := pw.log.LogAny("set_rts", func() (interface{}, error) {
		err := pw.port.SetRTS(rts)
		return rts, err
	})
	return err
}
