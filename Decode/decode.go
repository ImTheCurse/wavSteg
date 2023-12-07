package Decode

import (
	"log"
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

	f, err := os.Create("results/dec_msg.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(decodedMessage)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func checkMarked(val int, prevVal int) bool {
	return val%10 == 0 && prevVal%10 != 0
}
