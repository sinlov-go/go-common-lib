package string_tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Str2LineRaw(t *testing.T) {
	// mock _Str2LineRaw
	mockEnvDroneCommitMessage := "mock message commit\nmore line\nand more line\r\n"
	commitMessage := mockEnvDroneCommitMessage
	t.Logf("~> mock _Str2LineRaw")
	// do _Str2LineRaw
	t.Logf("~> do _Str2LineRaw")
	lineRaw := String2LineRaw(commitMessage)
	//commitMessage = strings.Replace(commitMessage, "\n", `\\n`, -1)
	t.Logf("lineRaw: %v", lineRaw)
	assert.Equal(t, "mock message commit\\nmore line\\nand more line\\n", lineRaw)
	// verify _Str2LineRaw
}
