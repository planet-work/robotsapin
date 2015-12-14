package sapi

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Image struct {
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Data     string `json:"data"`
}

type DisplayStatus struct {
	Status    string `json:"status"`
	ImageName string `json:"img_name"`
	ImageData string `json:"imgdata"`
}

func DisplayList() ([]*Image, error) {
	var il []*Image
	var data []byte
	logD.Println("Getting images from ", Settings.PictureDir)
	files, _ := ioutil.ReadDir(Settings.PictureDir)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".png") {
			logD.Println(">>>", f.Name())
			var i Image
			i.Name = f.Name()
			i.Filename = f.Name()
			f, err := os.Open(Settings.PictureDir + "/" + i.Filename)
			if err != nil {
				logE.Println(err)
			}
			defer f.Close()
			data, _ = ioutil.ReadAll(f)
			i.Data = base64.URLEncoding.EncodeToString(data)
			/*
				res, err := ReadId3V1Tag(Settings.MusicDir + "/" + f.Name())
				if err != nil {
					logE.Println(err)
				}*/
			il = append(il, &i)
		}
	}
	return il, nil
}

func DisplayData(data string) error {
	logD.Println("Displaying image from data")
	cmd := exec.Command("python/displayimg.py ", data)
	err := cmd.Run()
	return err
}

func DisplayImage(filename string) error {
	logD.Println("Displaying image ", filename)
	cmd := exec.Command("python/displayimg.py ", Settings.PictureDir+"/"+filename)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	err := cmd.Run()
	return err
}
