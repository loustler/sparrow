package download

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func Download(url, filepath string) error {
	res, err := http.Get(url)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("cannot receive status code 200 from given url")
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
