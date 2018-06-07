package serialdebug_test

import (
	"github.com/albenik/go-serial-debug"
	"github.com/albenik/iolog"
)

type fakeport struct{}

func (fp *fakeport) SetReadTimeout(t int) error {
	return nil
}

func (fp *fakeport) SetReadTimeoutEx(t, i uint32) error {
	return nil
}

func (fp *fakeport) SetFirstByteReadTimeout(t uint32) error {
	return nil
}

func (fp *fakeport) SetWriteTimeout(t int) error {
	return nil
}

func (fp *fakeport) ReadyToRead() (r uint32, err error) {
	return 0, nil
}

func (fp *fakeport) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (fp *fakeport) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (fp *fakeport) ResetInputBuffer() error {
	return nil
}

func (fp *fakeport) ResetOutputBuffer() error {
	return nil
}

func (fp *fakeport) SetDTR(dtr bool) error {
	return nil
}

func (fp *fakeport) SetRTS(rts bool) error {
	return nil
}

func (fp *fakeport) Close() error {
	return nil
}

func (fp *fakeport) String() string {
	return "fakeport"
}

func open() (serialdebug.SerialPort, error) {
	return new(fakeport), nil
}

type fakeDriver struct {
	open serialdebug.OpenFunc
	port serialdebug.SerialPort
}

func (fd *fakeDriver) Connect() error {
	port, err := fd.open()
	if err == nil {
		fd.port = port
	}
	return err
}

func (fd *fakeDriver) Disconnect() error {
	port := fd.port
	fd.port = nil
	return port.Close()
}

type fakeAdapter struct {
	log *iolog.IOLog
	drv *fakeDriver
}

func (fa *fakeAdapter) StartLogging() {
	fa.log.Start()
}

func (fa *fakeAdapter) StopLogging() []*iolog.Record {
	return fa.log.Stop()
}

func Example_Logging1() {
	log := iolog.New(8)
	d := &fakeDriver{
		open: serialdebug.Wrap(open, log),
	}

	log.Start()
	d.Connect()
	d.Disconnect()
	log.Stop()

	// Output:
}

func Example_Logging2() {
	log := iolog.New(8)

	a := &fakeAdapter{
		log: log,
		drv: &fakeDriver{
			open: serialdebug.Wrap(open, log),
		},
	}

	a.StartLogging()
	a.drv.Connect()
	a.drv.Disconnect()
	a.StopLogging()

	// Output:
}
