package filepath_plus

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// PathExists
//
//	path exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathExistsFast
//
//	path exists fast
func PathExistsFast(path string) bool {
	exists, _ := PathExists(path)
	return exists
}

// PathIsDir
//
//	path is dir
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// RmDir
//
//	remove dir by path
//
//nolint:golint,unused
func RmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := PathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// Mkdir
// will use FetchDefaultFolderFileMode()
func Mkdir(path string) error {
	err := os.MkdirAll(path, FetchDefaultFolderFileMode())
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// CopyFile
//
//	sourcePath is the path to the file in the embedded filesystem.
//	target is the target path where the file will be copied.
//	perm is the permission mode of the target file. os.FileMode(0o644) or os.FileMode(0o666)
//	coverage true will coverage old
func CopyFile(
	sourcePath string,
	target string,
	perm os.FileMode,
	coverage bool,
) error {
	sourceFile, errOpenSource := os.Open(sourcePath)
	if errOpenSource != nil {
		return fmt.Errorf("errOpenSource at CopyFile sourcePath %v", errOpenSource)
	}

	if !coverage {
		exists, errExist := PathExists(target)
		if errExist != nil {
			return errExist
		}

		if exists {
			return fmt.Errorf("not coverage, which target path exist %v", target)
		}
	}

	parentPath := filepath.Dir(target)
	if !PathExistsFast(parentPath) {
		errMkParentPath := os.MkdirAll(parentPath, FetchDefaultFolderFileMode())
		if errMkParentPath != nil {
			return fmt.Errorf(
				"can not CopyFile at new dir, at parent path: %v, why: %v",
				parentPath,
				errMkParentPath,
			)
		}
	}

	targetFile, errOpenTarget := os.OpenFile(target, os.O_CREATE|os.O_RDWR, perm)
	if errOpenTarget != nil {
		return fmt.Errorf("errOpenSource at CopyFile target mode: %v at: %s , %v", perm, target, errOpenTarget)
	}

	_, errCopy := io.Copy(targetFile, sourceFile)
	if errCopy != nil {
		return fmt.Errorf("errCopy at CopyFile target %v", errCopy)
	}

	return nil
}

// ReadFileAsByte
//
//	read file as byte array
func ReadFileAsByte(path string) ([]byte, error) {
	exists, err := PathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	if PathIsDir(path) {
		return nil, fmt.Errorf("path is dir: %s", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// ReadFileAsLines
//
// read file as lines, this method will read all line, so if file is too big, will be slow
func ReadFileAsLines(path string) ([]string, error) {
	if !PathExistsFast(path) {
		return nil, fmt.Errorf("read path %s not exists", path)
	}
	if PathIsDir(path) {
		return nil, fmt.Errorf("read path %s is dir", path)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read path %s error %s", path, err)
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
		readLine = append(readLine, scanner.Text())
	}
	return readLine, nil
}

// ReadFileAsJson
//
//	read file as json
func ReadFileAsJson(path string, v interface{}) error {
	fileAsByte, errRead := ReadFileAsByte(path)
	if errRead != nil {
		return fmt.Errorf("path: %s , read file as err: %v", path, errRead)
	}
	err := json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
}

// WriteFileByByte
//
//	write bytes to file
//	path most use Abs Path
//	data []byte
//	fileMod os.FileMode(0o666) os.FileMode(0o644)
//	coverage true will coverage old
func WriteFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !PathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, FetchDefaultFolderFileMode())
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}

// WriteFileAsJson write json file
//
//	path most use Abs Path
//	v data
//	fileMod os.FileMode(0o666) or os.FileMode(0o644)
//	coverage true will coverage old
//	beauty will format json when write
func WriteFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if beauty {
		var str bytes.Buffer
		errJson := json.Indent(&str, data, "", "  ")
		if errJson != nil {
			return errJson
		}
		return WriteFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return WriteFileByByte(path, data, fileMod, coverage)
}

// WriteFileAsJsonBeauty
//
//	write json file as 0666 and beauty
func WriteFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return WriteFileAsJson(path, v, os.FileMode(0o666), coverage, true)
}

// FetchDefaultFolderFileMode
// if in windows, will return os.FileMode(0o766), windows not support umask
//
//	use umask to get folder file mode
//
// if not windows, will use umask, will return os.FileMode(0o777) - umask
// not support umask will use os.FileMode(0o777)
func FetchDefaultFolderFileMode() fs.FileMode {
	switch runtime.GOOS {
	case "windows":
		return os.FileMode(0o766)
	default:
		umaskCode, err := getUmask()
		if err != nil {
			return os.FileMode(0o777)
		}
		if len(umaskCode) > 3 {
			umaskCode = umaskCode[len(umaskCode)-3:]
		}
		umaskOct, errParseUmask := strconv.ParseInt(umaskCode, 8, 64)
		if errParseUmask != nil {
			return os.FileMode(0o777)
		}
		defaultFOlderCode := 0o777
		nowOct := int(defaultFOlderCode) - int(umaskOct)
		return os.FileMode(nowOct)
	}
}

func getUmask() (string, error) {
	cmd := exec.Command("sh", "-c", "umask")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(&out)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	return "", fmt.Errorf("no output from umask command")
}
