package validate_nudity

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func DecodeImage(filePath string) (img image.Image, err error) {
	return decodeImage(filePath)
}

func decodeImage(filePath string) (img image.Image, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(reader *os.File) {
		err := reader.Close()
		if err != nil {
			return
		}
	}(reader)

	img, _, err = image.Decode(reader)
	return
}
