package main

import (
	"os"

	"github.com/ImTheCurse/wavSteg/Decode"
	"github.com/ImTheCurse/wavSteg/encode"
)

func main() {
	fp, err := os.Open("CantinaBand3.wav")
	if err != nil {
		return
	}
	encode.EncodeAudio(fp, "isn't it insane that computers can do these kind of stuff?")

	Decode.Decode("enc_file.wav")
}
