package encode

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/go-audio/wav"
)

type AudioData struct {
	pcm_bytes   []int
	sample_rate int
}

func EncodeAudio(file *os.File, message string) error {
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

	initMarking(&audData.pcm_bytes)
	buff, err := insertMessageToData(audData.pcm_bytes, message)
	buff = buff
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func insertMessageToData(data []int, message string) ([]int, error) {
	idxArr := make([]int, 0)

	const ASCII_UPPER_LIMIT = 128
	const ASCII_LOWER_LIMIT = 0

	//cache valid indexes
	for i, val := range data {
		if val < ASCII_UPPER_LIMIT && val > ASCII_LOWER_LIMIT && val > 0 {
			idxArr = append(idxArr, i)
		}
	}

	if len(message) > len(idxArr) {
		return nil, errors.New("message is longer than allocation capacity")
	}

	indx := idxArr[0]
	var ero error
	//Encode each character.
	for i := 0; i < len(message); i++ {
		indx, ero = findClosestValue(data, idxArr, indx, message[i])

		//fmt.Printf("Index: %d, Value: %d", indx, data[indx])
		if ero != nil {
			return data, ero
		}

		err := markValue(&data, indx-1)

		if err != nil {
			idxArr = removeIndex(idxArr, 0)
			i--
			continue
		}
		//changing PCM values to char to insert
		data[indx] = int(message[i])
		//fmt.Printf(" msg: %d\n", int(message[i]))
	}

	return data, nil
}

func markValue(data *[]int, index int) error {
	if index != 0 {
		if (*data)[index]%10 != 0 {
			(*data)[index] = (*data)[index] - (*data)[index]%10
			return nil
		}
	}
	return errors.New("invalid Index")
}

func findClosestValue(data []int, indexes []int, startIndex int, char byte) (int, error) {
	delta := math.MaxInt32
	index := math.MinInt32
	for i, idx := range indexes {

		if idx <= 0 || idx < startIndex {

			continue
		}
		//0 is marked in data[idx-1]
		val := data[idx]
		markedVal := data[idx-1]

		if math.Abs(float64(val-int(char))) < float64(delta) && markedVal%10 != 0 {
			delta = val - int(char)
			if delta < 0 {
				delta *= -1
			}

			index = indexes[i]
			//Defined close enough similarity between 2 and 0.
			if delta <= 2 && delta >= 0 {
				return index, nil
			}
		}
	}
	if index < 0 {
		return 0, errors.New("message too long to encode, supply bigger image or a smaller message")
	}
	return index, nil

}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func initMarking(data *[]int) {
	for i, val := range *data {
		if val%10 == 0 {
			(*data)[i]++
		}
	}

}
