package sickocommon

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func WriteJson(f string, data interface{}, perm os.FileMode) error {
	bfr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f, bfr, perm)
}

func ReadJson(f string, data interface{}) error {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	return nil
}

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, 100000)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
