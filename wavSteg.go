package main

import (
	"flag"
	"log"
	"os"

	"github.com/ImTheCurse/wavSteg/Decode"
	"github.com/ImTheCurse/wavSteg/encode"
)

func main() {
	audioFileName := flag.String("audio", "", "Audio file name")
	isEncode := flag.Bool("encode", false, "Encode flag")
	isDecode := flag.Bool("decode", true, "Decode flag")
	textFileName := flag.String("tfile", "", "Encode message with provided text file name")
	cliMessage := flag.String("message", "", "Encode message with Command Line Interface message")

	flag.Parse()
	var message string

	fp, err := os.Open(*audioFileName)
	if err != nil {
		log.Fatal("invalid audio file name")
	}

	if *textFileName != "" {
		textBytes, err := os.ReadFile(*textFileName)
		if err != nil {
			log.Fatal("invalid text file name")
		}
		message = string(textBytes)
	}

	if *cliMessage != "" {
		message = *cliMessage
	}

	if *isEncode {
		encode.EncodeAudio(fp, message)
		*isDecode = false
	}
	if *isDecode {
		Decode.Decode("enc_file.wav")
	}
}
