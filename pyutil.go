package pyutil

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func FileExists(file string) bool {
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func ListDir(dir string) ([]string, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		files = append(files, filepath.Join(dir, info.Name()))
	}
	return files, nil
}

func ReadLines(fpath string) ([]string, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var ret []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ret = append(ret, strings.TrimSpace(line))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
