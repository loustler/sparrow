package download

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filepath string) error {
	res, err := http.Get(url)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("Cannot receive status code 200 from Given URL")
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)

	if err != nil {
		return err
	}

	return nil
}
