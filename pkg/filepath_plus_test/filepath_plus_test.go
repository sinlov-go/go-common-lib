package filepath_plus_test

import (
	"errors"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateEmptyFile(t *testing.T) {

	testDataFullPath, errGoldenPath := testGoldenKit.GetOrCreateTestDataFullPath("create_empty_file")
	assert.Nil(t, errGoldenPath)
	errRmDir := filepath_plus.RmDir(testDataFullPath, true)
	assert.Nil(t, errRmDir)

	firstFile := filepath.Join(testDataFullPath, "empty.golden")
	errCreateEmptyFile := filepath_plus.CreateEmptyFile(firstFile, os.FileMode(0o644))
	assert.Nil(t, errCreateEmptyFile)
	assert.True(t, filepath_plus.PathExistsFast(firstFile))
	errCreateEmptyFile = filepath_plus.CreateEmptyFile(firstFile, os.FileMode(0o644))
	assert.True(t, errors.Is(errCreateEmptyFile, filepath_plus.CreateEmptyFilePathExisted))

	secondFile := filepath.Join(testDataFullPath, "some_word.golden")
	errWriteSecondFile := filepath_plus.WriteFileByByte(secondFile, []byte("some word"), os.FileMode(0o644), true)
	assert.Nil(t, errWriteSecondFile)
	assert.True(t, filepath_plus.PathExistsFast(secondFile))
	errCreateSecondEmptyFile := filepath_plus.CreateEmptyFile(secondFile, os.FileMode(0o644))
	assert.True(t, errors.Is(errCreateSecondEmptyFile, filepath_plus.CreateEmptyFilePathExisted))
}
