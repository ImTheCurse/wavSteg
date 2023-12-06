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
		if checkMarked(val) {
			decodedMessage += string(buff.Data[i+1])
		}
	}
	fmt.Printf("%s", decodedMessage)
}

func checkMarked(val int) bool {
	return val%10 == 0
}
