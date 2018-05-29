package serialdebug

import (
	"fmt"

	"github.com/albenik/iolog"
)

type OpenFunc func() (SerialPort, error)

func NewWrappedOpener(open OpenFunc, log *iolog.IOLog) OpenFunc {
	return func() (port SerialPort, err error) {
		log.LogAny("open", func() (interface{}, error) {
			if port, err = open(); err == nil {
				port = &PortWrapper{port: port, log: log}
			}
			return port, err
		})
		return
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

func (pw *PortWrapper) String() string {
	return pw.port.String()
}

func (pw *PortWrapper) SetReadTimeout(t int) error {
	return pw.log.LogAny("set_read_timeout", func() (interface{}, error) {
		err := pw.port.SetReadTimeout(t)
		return t, err
	})
}

func (pw *PortWrapper) SetReadTimeoutEx(t, i uint32) error {
	return pw.log.LogAny("set_read_timeout_ex", func() (interface{}, error) {
		err := pw.port.SetReadTimeoutEx(t, i)
		return &struct{ T, I uint32 }{t, i}, err
	})
}

func (pw *PortWrapper) SetFirstByteReadTimeout(t uint32) error {
	return pw.log.LogAny("set_first_byte_read_timeout", func() (interface{}, error) {
		err := pw.port.SetFirstByteReadTimeout(t)
		return t, err
	})
}

func (pw *PortWrapper) SetWriteTimeout(t int) error {
	return pw.log.LogAny("set_write_timeout", func() (interface{}, error) {
		err := pw.port.SetWriteTimeout(t)
		return t, err
	})
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
	return pw.log.LogAny("reset_input_buffer", func() (interface{}, error) {
		err := pw.port.ResetInputBuffer()
		return nil, err
	})
}

func (pw *PortWrapper) ResetOutputBuffer() error {
	return pw.log.LogAny("reset_output_buffer", func() (interface{}, error) {
		err := pw.port.ResetOutputBuffer()
		return nil, err
	})
}

func (pw *PortWrapper) SetDTR(dtr bool) error {
	return pw.log.LogAny("set_dtr", func() (interface{}, error) {
		err := pw.port.SetDTR(dtr)
		return dtr, err
	})
}

func (pw *PortWrapper) SetRTS(rts bool) error {
	return pw.log.LogAny("set_rts", func() (interface{}, error) {
		err := pw.port.SetRTS(rts)
		return rts, err
	})
}

func (pw *PortWrapper) Close() error {
	return pw.log.LogAny("close", func() (interface{}, error) {
		err := pw.port.Close()
		return nil, err
	})
}

func (pw *PortWrapper) StartLoggging() {
	pw.log.Start()
}

func (pw *PortWrapper) StopLogging() []*iolog.Record {
	return pw.log.Stop()
}
