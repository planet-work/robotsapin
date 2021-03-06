package sapi

import (
	"bufio"
	"io"
	//"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Mp3Player struct {
	Status     string
	LoadedFile string
	Position   float64
	Duration   float64
	Stdin      io.WriteCloser
	Stdout     *bufio.Reader
	Stderr     *bufio.Reader
	Pid        int
}

func (player *Mp3Player) Spawn() {
	player.Status = "init2"
	c := make(chan int)
	if player.Pid > 0 {
		syscall.Kill(player.Pid, 9)
	}
	go func() {
		var err error
		logD.Println("mpg321 -R foo")
		cmd := exec.Command("mpg321", "-R", "foo")
		//cmd.Stderr = os.Stderr
		player.Stdin, err = cmd.StdinPipe()
		if err != nil {
			logE.Println(err)
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			logE.Println("ERROR Getting stdout", err)
		}
		player.Stdout = bufio.NewReader(stdout)

		stderr, err := cmd.StderrPipe()
		if err != nil {
			logE.Println("ERROR Getting stderr", err)
		}
		player.Stderr = bufio.NewReader(stderr)

		player.Status = "started"
		cmd.Start()
		c <- 1
		player.Pid = cmd.Process.Pid
		logD.Println("Started backgroup MP3 Player PID:", player.Pid)
		defer player.Kill()
		if err := cmd.Wait(); err != nil {
			syscall.Kill(player.Pid, 9)
			logE.Println("ERROR", err)
		}
	}()

	go func() {
		_ = <-c
		var frame string
		logD.Println("MP3PLAYER : starting stderr reader")

		scanner := bufio.NewScanner(player.Stderr)
		for scanner.Scan() {
			// @F 4870 23 127.21 0.62
			frame = scanner.Text()
			status := strings.Split(string(frame), " ")
			if status[0] == "@F" {
				Status.Music.Position, _ = strconv.ParseFloat(status[3], 32)
				remaining, _ := strconv.ParseFloat(status[4], 32)
				if Status.Music.Duration == 0 {
					Status.Music.Duration = remaining
				}
				if remaining < 1 {
					Status.Music.Status = "stopped"
				}
			}
		}

		logD.Println("MP3PLAYER : stderr reader stopper")
	}()

}

func (player *Mp3Player) Kill() {
	logD.Println("Killing MP3 player with PID", player.Pid)
	syscall.Kill(player.Pid, 9)
}

func (player *Mp3Player) Play(filename string) {
	player.LoadedFile = filename
	player.Position = 0
	Status.Music.Duration = 0
	Status.Music.Position = 0
	logD.Println("Playing ", Settings.MusicDir+"/"+filename)
	player.Stdin.Write([]byte("LOAD " + Settings.MusicDir + "/" + filename + "\n"))
}

func NewMp3Player() *Mp3Player {
	return &Mp3Player{Status: "init", Pid: 0}
}

func (player *Mp3Player) Pause() {
	logD.Println("Pausing")
	player.Stdin.Write([]byte("PAUSE\n"))
}

func (player *Mp3Player) Stop() {
	logD.Println("Stop")
	player.Stdin.Write([]byte("STOP\n"))
}

func (player *Mp3Player) Gain(gain int) {
	logD.Println("Gain" + strconv.Itoa(gain))
	player.Stdin.Write([]byte("GAIN " + strconv.Itoa(gain) + "\n"))
}

func (player *Mp3Player) Quit() {
	logD.Println("QUIT")
	player.Stdin.Write([]byte("QUIT\n"))
}
