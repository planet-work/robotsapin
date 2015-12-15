package sapi

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	titleEnd   = 30
	artistEnd  = 60
	albumEnd   = 90
	yearEnd    = 94
	commentEnd = 124

	// First three chars of an ID3 tag are static "TAG"
	tagStart = 3

	tagSize = 128
)

type Song struct {
	Filename string `json:"filename"`
	Name     string `json:"name"`
}

type MusicStatus struct {
	Status   string `json:"status"`
	SongName string `json:"song_name"`
	Filename string `json:"filename"`
	Position int    `json:"position"`
}

func byteString(b []byte) string {
	pos := bytes.IndexByte(b, 0)

	if pos == -1 {
		pos = len(b)
	}

	return string(b[0:pos])
}

func ReadId3V1Tag(filename string) (map[string]string, error) {
	buff_ := make([]byte, tagSize)

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	// Read last 128 bytes of file to see ID3 tag
	f.Seek(-tagSize, 2)
	f.Read(buff_)

	// First 3 characters are static "TAG"
	if byteString(buff_[0:tagStart]) != "TAG" {
		return nil, errors.New("No ID3 tag found")
	}

	buff := buff_[tagStart:]

	id3tag := map[string]string{}

	id3tag["title"] = byteString(buff[0:titleEnd])
	id3tag["artist"] = byteString(buff[titleEnd:artistEnd])
	id3tag["album"] = byteString(buff[artistEnd:albumEnd])
	id3tag["year"] = byteString(buff[albumEnd:yearEnd])
	id3tag["comment"] = byteString(buff[yearEnd:commentEnd])

	// Special case. If next-to-last comment byte is zero, then the last
	// comment byte is the track number
	if buff[commentEnd-2] == 0 {
		id3tag["track"] = fmt.Sprintf("%d", buff[commentEnd-1])
	}
	genre_code := buff[commentEnd]
	id3tag["genre"] = fmt.Sprintf("%d", genre_code)
	//id3tag["genre_name"] = codeToName[genre_code]

	return id3tag, nil
}

func MusicList() ([]*Song, error) {
	var sl []*Song
	logD.Println("Getting songs from ", Settings.MusicDir)
	files, _ := ioutil.ReadDir(Settings.MusicDir)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".mp3") {
			logD.Println(">>>", f.Name())
			var s Song
			s.Name = f.Name()
			s.Filename = f.Name()
			res, err := ReadId3V1Tag(Settings.MusicDir + "/" + f.Name())
			if err != nil {
				logE.Println(err)
			}

			for k, v := range res {
				if k == "title" {
					s.Name = v
				}
				//fmt.Printf("%s => %s\n", k, v)
			}

			sl = append(sl, &s)
		}
	}
	return sl, nil
}

func MusicPost() error {
	return nil
}

func MusicPlay(filename string) error {
	logD.Println("Playing ", filename)
	Status.Music.Status = "..."
	Status.Music.Filename = filename
	Status.Music.SongName = filename
	go func() {
		logD.Println("mpg123", Settings.MusicDir+"/"+filename)
		cmd := exec.Command("mpg123", "-R") //, Settings.MusicDir+"/"+filename)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			logE.Println(err)
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			logE.Println(err)
		}
		if err := cmd.Start(); err != nil {
			logE.Println(err)
		}
		Status.Music.Status = "playing"
		stdin.Write([]byte("LOAD" + Settings.MusicDir + "/" + filename))
		x, _ := stdout.Reader.ReadString('\n')
		if err := cmd.Wait(); err != nil {
			logE.Println(err)
		}
		//Status.Music.Status = "stopped"
		logD.Println("Song terminated")
	}()
	//var out bytes.Buffer
	//cmd.Stdout = &out
	logD.Println("Return XXX")
	return nil
}
