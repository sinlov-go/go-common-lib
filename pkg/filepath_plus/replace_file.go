package filepath_plus

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ReplaceFileLineByLine
//
// read file as lines, this method will read all line, so if file is too big, will be slow
func ReplaceFileLineByLine(path string, reg string, index int, replace string) error {
	compLine, err := regexp.Compile(reg)
	if err != nil {
		return err
	}
	if !PathExistsFast(path) {
		return fmt.Errorf("read path %s not exists", path)
	}
	if PathIsDir(path) {
		return fmt.Errorf("read path %s is dir", path)
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("AlterFile get file info at path: %v, err: %v", path, err)
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("read path %s error %s", path, err)
	}
	defer func(file *os.File) {
		errFileClose := file.Close()
		if errFileClose != nil {
			fmt.Printf("read close file err: %v\n", errFileClose)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var readLine []string
	for scanner.Scan() {
		lineText := scanner.Text()
		findGroup := compLine.FindStringSubmatch(lineText)
		if len(findGroup) < index+1 {
			readLine = append(readLine, lineText)
			continue
		}
		findGroup[index] = replace
		readLine = append(readLine, strings.Join(findGroup[1:], ""))
	}
	joinData := strings.Join(readLine, "\n")

	err = os.WriteFile(path, []byte(joinData), fileInfo.Mode())
	if err != nil {
		return fmt.Errorf("ReplaceFileLineByLine write data at path: %v, err: %v", path, err)
	}

	return nil
}
