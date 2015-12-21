package sapi

import (
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"strconv"
	"time"
	//	"encoding/base64"
	//	"io/ioutil"
	//	"os"
	//	"strings"
)

type LedSequence struct {
	Id   int      `json:"id"`
	Data []string `json:"data"`
}

type ToperStatus struct {
	Status           string `json:"status"`
	SequenceId       int    `json:"sequence_id"`
	Speed            int    `json:"speed"`
	SequenceData     string `json:"sequence_data"`
	SequencePosition int    `json:"sequence_position"`
	Reverse          bool   `json:"reverse"`
}

var dataPin = "8"
var clockPin = "5"
var latchPin = "4"

var firmataAdaptor *firmata.FirmataAdaptor
var led *gpio.LedDriver

var sequences = [][]string{}
var seq0 = []string{"11111111"}
var seq1 = []string{"00000001", "00000010", "00000100", "00001000", "00010000", "00100000", "01000000", "10000000"}
var seq2 = []string{"11111110", "11111101", "11111011", "11110111", "11101111", "11011111", "10111111", "01111111"}
var seq3 = []string{"00010001", "00100010", "01000100", "10001000", "00010001", "00100010", "01000100", "10001000"}
var seq4 = []string{"11101110", "11011101", "10111011", "01110111", "11101110", "11011101", "10111011", "01110111"}
var seq5 = []string{"00000000", "11111111", "00000000", "11111111", "00000000", "11111111", "00000000", "11111111"}

func TopperList() ([]*LedSequence, error) {
	var ls []*LedSequence
	for id := range sequences {
		s := LedSequence{}
		s.Id = id
		s.Data = sequences[id]
		ls = append(ls, &s)
	}
	return ls, nil
}

func TopperInit() {
	firmataAdaptor = firmata.NewFirmataAdaptor("arduino", "/dev/ttyUSB0")
	firmataAdaptor.Connect()
	led = gpio.NewLedDriver(firmataAdaptor, "led", "13")
	Status.Topper.Status = "stopped"
	Status.Topper.Speed = 100
	Status.Topper.Reverse = false
	sequences = append(sequences, seq0)
	sequences = append(sequences, seq1)
	sequences = append(sequences, seq2)
	sequences = append(sequences, seq3)
	sequences = append(sequences, seq4)
	sequences = append(sequences, seq5)
	logD.Println("Topper INIT ... dome")
}

func updateLeds(values string) (err error) {
	timer := time.Tick(1 * time.Millisecond)
	//fmt.Println("LEDS:", values)

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

func RunSequence(seqId int, speed int) {
	Status.Topper.Status = "running"
	for true {
		next := sequences[seqId][Status.Topper.SequencePosition]
		//logD.Println(next)
		updateLeds(next)
		led.Toggle()
		Status.Topper.SequencePosition += 1
		if Status.Topper.SequencePosition > 7 {
			Status.Topper.SequencePosition = 0
		}
		//time.Sleep(Status.Topper.Speed * time.Millisecond)
		time.Sleep(300 * time.Millisecond)
	}
}

func TopperSetSequence(seqId int, speed int) error {
	logD.Println("Topper sequence ", seqId)
	Status.Topper.SequenceId = seqId
	Status.Topper.Speed = speed
	Status.Topper.SequenceData = ""
	go RunSequence(seqId, speed)
	return nil
}
