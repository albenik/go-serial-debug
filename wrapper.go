package serialdebug

import (
	"fmt"

	"github.com/albenik/iolog"
)

type OpenFunc func() (SerialPort, error)

func Wrap(fn OpenFunc, log *iolog.IOLog) OpenFunc {
	return func() (SerialPort, error) {
		port, err := log.LogAny("open", func() (interface{}, error) {
			port, err := fn()
			port = &PortWrapper{port: port, log: log}
			return port, err
		}, false)
		return port.(SerialPort), err
	}
}

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
	port SerialPort
	log  *iolog.IOLog
}

func (pw *PortWrapper) SetReadTimeout(t int) error {
	_, err := pw.log.LogAny("set_read_timeout", func() (interface{}, error) {
		err := pw.port.SetReadTimeout(t)
		return t, err
	}, true)
	return err
}

func (pw *PortWrapper) SetReadTimeoutEx(t, i uint32) error {
	_, err := pw.log.LogAny("set_read_timeout_ex", func() (interface{}, error) {
		err := pw.port.SetReadTimeoutEx(t, i)
		return &struct{ T, I uint32 }{t, i}, err
	}, true)
	return err
}

func (pw *PortWrapper) SetFirstByteReadTimeout(t uint32) error {
	_, err := pw.log.LogAny("set_first_byte_read_timeout", func() (interface{}, error) {
		err := pw.port.SetFirstByteReadTimeout(t)
		return t, err
	}, true)
	return err
}

func (pw *PortWrapper) SetWriteTimeout(t int) error {
	_, err := pw.log.LogAny("set_write_timeout", func() (interface{}, error) {
		err := pw.port.SetWriteTimeout(t)
		return t, err
	}, true)
	return err
}

func (pw *PortWrapper) ReadyToRead() (r uint32, err error) {
	pw.log.LogAny("ready_to_read", func() (interface{}, error) {
		r, err = pw.port.ReadyToRead()
		return r, err
	}, true)
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
	}, true)
	return err
}

func (pw *PortWrapper) ResetOutputBuffer() error {
	_, err := pw.log.LogAny("reset_output_buffer", func() (interface{}, error) {
		err := pw.port.ResetOutputBuffer()
		return nil, err
	}, true)
	return err
}

func (pw *PortWrapper) SetDTR(dtr bool) error {
	_, err := pw.log.LogAny("set_dtr", func() (interface{}, error) {
		err := pw.port.SetDTR(dtr)
		return dtr, err
	}, true)
	return err
}

func (pw *PortWrapper) SetRTS(rts bool) error {
	_, err := pw.log.LogAny("set_rts", func() (interface{}, error) {
		err := pw.port.SetRTS(rts)
		return rts, err
	}, true)
	return err
}

func (pw *PortWrapper) Close() error {
	_, err := pw.log.LogAny("close", func() (interface{}, error) {
		err := pw.port.Close()
		return nil, err
	}, true)
	pw.port = nil
	return err
}

func (pw *PortWrapper) String() string {
	return pw.port.String()
}
