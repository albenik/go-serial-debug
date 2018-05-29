module github.com/albenik/go-serial-debug

replace (
	go.bug.st/serial.v1 v0.0.0 => github.com/albenik/go-serial v1.0.0-my2
)

require (
	github.com/albenik/iolog v0.7.1
	go.bug.st/serial.v1 v0.0.0
)
