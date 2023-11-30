package encode

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/go-audio/wav"
)

type AudioData struct {
	pcm_bytes   []int
	sample_rate int
}

type cachedIndexes struct {
	cache []int
}

func EncodeAudio(file *os.File, message string) error {
	fmt.Println("hello")
	decWav := wav.NewDecoder(file)

	if decWav == nil {
		log.Fatal("error in decoder")
		return nil
	}

	audBuff, err := decWav.FullPCMBuffer()

	if err != nil {
		log.Fatal("Error getting buffer.")
		return nil
	}
	audData := AudioData{pcm_bytes: audBuff.Data, sample_rate: audBuff.Format.SampleRate}

	return nil
}

func insertMessageToData(data []int, message string) []int {
	idxArr := make([]int, 0)

	const ASCII_UPPER_LIMIT = 128
	const ASCII_LOWER_LIMIT = 0

	for i, val := range data {
		if val < ASCII_UPPER_LIMIT && val > ASCII_LOWER_LIMIT {
			idxArr = append(idxArr, i)
		}
	}

	slices.Sort(idxArr)

	for i := 0; i < len(message); i++ {
		//data[idxArr[i]] = int(message[i])
		//find closest value.
	}
	return data
}

func markPrevValue(data []int, currentValIndex int) error {
	if currentValIndex != 0 {
		if data[currentValIndex]%10 != 0 {
			data[currentValIndex] = data[currentValIndex] - data[currentValIndex]%10
			return nil
		}
	}
	return errors.New("Invalid Index")
}
