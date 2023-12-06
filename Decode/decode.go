package Decode

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func Decode(fname string) {
	fp, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var decodedMessage string
	Dec := wav.NewDecoder(fp)
	buff, err := Dec.FullPCMBuffer()
	if err != nil {
		panic(err)
	}

	for i, val := range buff.Data {
		if i == 0 {
			continue
		}
		if checkMarked(val, buff.Data[i-1]) {
			decodedMessage += string(byte(buff.Data[i+1]))
		}
	}
	decodedMessage += "\n"
	fmt.Print(decodedMessage)
}

func checkMarked(val int, prevVal int) bool {
	return val%10 == 0 && prevVal%10 != 0
}
