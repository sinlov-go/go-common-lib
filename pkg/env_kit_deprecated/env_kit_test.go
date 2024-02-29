package env_kit_deprecated

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvKeys(t *testing.T) {
	// mock EnvKeys
	t.Logf("~> mock EnvKeys")

	err := SetEnvBool(keyEnvDebug, true)
	if err != nil {
		t.Fatal(err)
	}

	err = SetEnvInt64(keyEnvCiNum, 2)
	if err != nil {
		t.Fatal(err)
	}

	err = SetEnvStr(keyEnvCiKey, "foo")
	if err != nil {
		t.Fatal(err)
	}

	// do EnvKeys
	t.Logf("~> do EnvKeys")

	// verify EnvKeys

	assert.True(t, FetchOsEnvBool(keyEnvDebug, false))
	err = SetEnvBool(keyEnvDebug, false)
	if err != nil {
		t.Fatal(err)
	}
	assert.False(t, FetchOsEnvBool(keyEnvDebug, false))
	assert.Equal(t, 2, FetchOsEnvInt(keyEnvCiNum, 0))
	assert.Equal(t, "foo", FetchOsEnvStr(keyEnvCiKey, ""))
	envArray := FetchOsEnvArray(keyEnvCiKeys)
	assert.Nil(t, envArray)

	err = SetEnvStr(keyEnvCiKeys, "foo, bar,My ")
	if err != nil {
		t.Fatal(err)
	}

	envArray = FetchOsEnvArray(keyEnvCiKeys)

	assert.NotNil(t, envArray)
	assert.Equal(t, "foo", envArray[0])
	assert.Equal(t, "bar", envArray[1])
	assert.Equal(t, "My", envArray[2])

	err = SetEnvU64(keyEnvCiNum, 3)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, FetchOsEnvInt(keyEnvCiNum, 0))
}

func TestEnvKeyDefaultVal(t *testing.T) {
	t.Logf("~> do EnvKeyDefaultVal")
	// do EnvKeyFailDefaultVal
	assert.Equal(t, false, FetchOsEnvBool("foo", false))
	assert.Equal(t, 1, FetchOsEnvInt("one", 1))
	err := os.Setenv("two", "two")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, FetchOsEnvInt("two", 1))

	assert.Equal(t, "", FetchOsEnvStr("bar", ""))
	assert.Equal(t, []string(nil), FetchOsEnvArray("bar"))
	err = os.Setenv("bar", ".")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []string{"."}, FetchOsEnvArray("bar"))
}
