package lib

import (
	"os/exec"
	"bytes"
	"regexp"
)

func getDir(path string) ([]string, error){
	bin := "ls"
	arg1 := path

	cmd := exec.Command(bin, arg1)

	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	out :=  bytes.Split(stdout,[]byte("\n"))
	strOut := make([]string, len(out))

	for i := 0; i < len(out); i++ {
		strOut[i] = string(out[i])
		}

	strOut = strOut[:(len(strOut) - 1)]
	return strOut, nil
}

func GetAudioPlaylist(path string) ([]string, error){
	files, err := getDir(path)
	if err != nil {
		return nil, err
	}

	count := 0
	r, err := regexp.Compile(`.*\.(wav|mp3)`)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(files); i++ {
		if r.MatchString(files[i]){
			count += 1
		}
	}

	playlist := make([]string, count)
	y := 0
	for i := 0; i < len(files); i++ {
		if r.MatchString(files[i]){
			playlist[y] = files[i]
			y++
		}
	}
	return playlist, nil
}
