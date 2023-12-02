package main

import (
	"os"

	"github.com/ImTheCurse/wavSteg/encode"
)

func main() {

	fp, err := os.Open("CantinaBand3.wav")

	if err != nil {
		return
	}
	encode.EncodeAudio(fp, "hello beautiful world! its getting a bit late but I believe we can do this not on our own' but together")
}
