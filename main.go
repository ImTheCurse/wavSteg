package main

import (
	"github.com/ImTheCurse/wavSteg/Decode"
)

func main() {
	//fmt.Println("hello world!")
	//fp, err := os.Open("CantinaBand3.wav")
	//if err != nil {
	//return
	//}
	// encode.EncodeAudio(fp, "hello beautiful world! its getting a bit late but I believe we can do this not on our own' but together")
	Decode.Decode("enc_file.wav")
}
