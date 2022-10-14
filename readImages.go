package validate_nudity

import (
	"fmt"
	"image"
	"io"
	"net/http"
)

func loadImageFromURL(URL string) (img image.Image, err error) {
	response, err := http.Get(URL)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode != 200 {
		_, err = fmt.Printf("received non 200 response code %v, in loadImageFromURL: %v", response.StatusCode, "load img")
		return
	}

	img, _, err = image.Decode(response.Body)
	if err != nil {
		return
	}

	return
}
