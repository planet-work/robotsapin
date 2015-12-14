package sapi

import (
	"fmt"
	"github.com/hybridgroup/gobot/platforms/firmata"
	//"github.com/hybridgroup/gobot/platforms/gpio"
	"strconv"
	"time"
	//	"encoding/base64"
	//	"io/ioutil"
	//	"os"
	//	"strings"
)

type LedSequence struct {
	Id      int    `json:"id"`
	Data    string `json:"data"`
	Speed   int    `json:"speed"`
	Reverse bool   `json:"reverse"`
}

type ToperStatus struct {
	SequenceId int `json:"sequence_id"`
	Speed      int `json:"speed"`
}

var dataPin = "8"
var clockPin = "5"
var latchPin = "4"
var firmataAdaptor *firmata.FirmataAdaptor

var seq1 = []string{"00000001", "00000010", "00000100", "00001000", "00010000", "00100000", "01000000", "10000000"}
var seq2 = []string{"11111110", "11111101", "11111011", "11110111", "11101111", "11011111", "10111111", "01111111"}
var seq3 = []string{"00010001", "00100010", "01000100", "10001000", "00010001", "00100010", "01000100", "10001000"}
var seq4 = []string{"11101110", "11011101", "10111011", "01110111", "11101110", "11011101", "10111011", "01110111"}
var seq5 = []string{"00000000", "11111111", "00000000", "11111111", "00000000", "11111111", "00000000", "11111111"}

func TopperList() ([]*LedSequence, error) {
	var ls []*LedSequence
	//var data []byte
	s := LedSequence{}
	s.Id = 1
	s.Speed = 2
	s.Data = ""
	s.Reverse = false
	ls = append(ls, &s)
	return ls, nil
}

func initTopper() {
	firmataAdaptor = firmata.NewFirmataAdaptor("arduino", "/dev/ttyUSB0")
	//led := gpio.NewLedDriver(firmataAdaptor, "led", "13")
}

func updateLeds(values string) (err error) {
	timer := time.Tick(1 * time.Millisecond)
	fmt.Println("LEDS:", values)

	firmataAdaptor.DigitalWrite(latchPin, 0)
	for i := 0; i < 8; i++ {
		bit, _ := strconv.ParseBool(string(values[i]))
		if bit == true {
			firmataAdaptor.DigitalWrite(dataPin, 1)
		} else {
			firmataAdaptor.DigitalWrite(dataPin, 0)
		}
		firmataAdaptor.DigitalWrite(clockPin, 1)
		<-timer
		firmataAdaptor.DigitalWrite(clockPin, 0)
	}
	firmataAdaptor.DigitalWrite(latchPin, 1)
	return err
}

func TopperSetSequence(seqId int, speed int) error {
	logD.Println("Topper sequence ", seqId)

	return nil
}
