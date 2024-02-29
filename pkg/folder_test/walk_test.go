package folder_test

import (
	"github.com/sinlov-go/go-common-lib/pkg/folder"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rootLevCnt       = 3
	innerLev1JsonCnt = 5
	innerLev1txtCnt  = 2
	innerLev2JsonCnt = 4
)

func TestWalkAllByMatchPath(t *testing.T) {
	// mock WalkAllByMatchPath
	t.Logf("~> mock WalkAllByMatchPath")
	err, testDataPath := createTestFileTree(t)
	if err != nil {
		t.Fatal(err)
	}

	// do WalkAllByMatchPath
	t.Logf("~> do WalkAllByMatchPath")
	matchJsonPath, err := folder.WalkAllByMatchPath(testDataPath, `.*.json$`, true)

	assert.NotEmpty(t, matchJsonPath)
	// verify WalkAllByMatchPath
	assert.Equal(t,
		rootLevCnt+innerLev1JsonCnt+innerLev2JsonCnt,
		len(matchJsonPath))
}

func TestWalkAllByGlob(t *testing.T) {
	// mock WalkAllByGlob

	t.Logf("~> mock WalkAllByGlob")
	err, testDataPath := createTestFileTree(t)
	if err != nil {
		t.Fatal(err)
	}

	// do WalkAllByGlob
	t.Logf("~> do WalkAllByGlob")

	matchJsonPath, err := folder.WalkAllByGlob(testDataPath, `*.json`, true)
	// verify WalkAllByGlob
	assert.NotEmpty(t, matchJsonPath)
	// verify WalkAllByMatchPath
	assert.Equal(t,
		rootLevCnt,
		len(matchJsonPath))

	matchJsonPath, err = folder.WalkAllByGlob(testDataPath, `**/*.json`, true)
	// verify WalkAllByGlob
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, matchJsonPath)
	assert.Equal(t,
		innerLev1JsonCnt,
		len(matchJsonPath))

	matchJsonPath, err = folder.WalkAllByGlob(testDataPath, `**/*.txt`, true)
	// verify WalkAllByGlob
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, matchJsonPath)
	assert.Equal(t,
		innerLev1txtCnt,
		len(matchJsonPath))

	matchJsonPath, err = folder.WalkAllByGlob(testDataPath, `**/**/*.json`, true)
	// verify WalkAllByGlob
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, matchJsonPath)
	assert.Equal(t,
		innerLev2JsonCnt,
		len(matchJsonPath))

	matchJsonPath, err = folder.WalkAllByGlob(testDataPath, `**/**/*.txt`, true)
	// verify WalkAllByGlob
	if err != nil {
		t.Fatal(err)
	}
	assert.Empty(t, matchJsonPath)
}

func createTestFileTree(t *testing.T) (error, string) {
	currentFilePath, err := folder.GetCurrentFilePath()
	if err != nil {
		t.Error(err)
	}
	currentFolder := folder.PathParent(currentFilePath)
	testDataPath := folder.PathJoin(currentFolder, "testdata")

	err = folder.RmDirForce(testDataPath)
	if err != nil {
		t.Error(err)
	}
	err = folder.Mkdir(testDataPath)
	if err != nil {
		t.Error(err)
	}

	rootLevCnt := 3

	err = addTextFileByTry(testDataPath, "data", "json", rootLevCnt)
	if err != nil {
		t.Error(err)
	}

	innerLev1JsonCnt := 5
	innerLev1txtCnt := 2

	innerLev1Folder := folder.PathJoin(testDataPath, "inner_1")
	err = addTextFileByTry(innerLev1Folder, "data", "json", innerLev1JsonCnt)
	if err != nil {
		t.Error(err)
	}
	err = addTextFileByTry(innerLev1Folder, "example", "txt", innerLev1txtCnt)
	if err != nil {
		t.Error(err)
	}

	innerLev2JsonCnt := 4
	innerLev2Folder := folder.PathJoin(innerLev1Folder, "inner_2")
	err = addTextFileByTry(innerLev2Folder, "data", "json", innerLev2JsonCnt)
	if err != nil {
		t.Error(err)
	}
	return err, testDataPath
}
