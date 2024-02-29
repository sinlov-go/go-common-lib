package string_tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrInArr(t *testing.T) {
	// mock StringInArr
	var testArr = []string{
		"a",
		"b",
		"c",
	}
	t.Logf("~> mock StringInArr")
	// do StringInArr
	t.Logf("~> do StringInArr")
	// verify StringInArr
	assert.True(t, StringInArr("c", testArr))
	assert.False(t, StringInArr("d", testArr))
}

func BenchmarkStrInArr(b *testing.B) {
	var testArr = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
	}
	for i := 0; i < b.N; i++ {
		assert.True(b, StringInArr("f", testArr))
	}
}

func TestStrArrRemoveDuplicates(t *testing.T) {
	// mock StringArrRemoveDuplicates

	t.Logf("~> mock StringArrRemoveDuplicates")
	var fooArr = []string{
		"a", "b", "c", "b", "a", "d", "f",
	}
	// do StringArrRemoveDuplicates
	t.Logf("~> do StringArrRemoveDuplicates")
	rdFooArr := StringArrRemoveDuplicates(fooArr)
	// verify StringArrRemoveDuplicates
	assert.Equal(t, 5, len(rdFooArr))

	var barArr []string
	for i := 0; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 1000; i < 2000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 4000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}

	rdBarArr := StringArrRemoveDuplicates(barArr)
	// verify StringArrRemoveDuplicates
	assert.Equal(t, 5000, len(rdBarArr))
}

func BenchmarkStrArrRemoveDuplicates(b *testing.B) {
	var fooArr = []string{
		"a", "b", "c", "b", "a", "d", "f",
	}
	var barArr []string
	for i := 0; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 1000; i < 2000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 4000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 0; i < b.N; i++ {
		rdFooArr := StringArrRemoveDuplicates(fooArr)
		// verify StringArrRemoveDuplicates
		assert.Equal(b, 5, len(rdFooArr))

		rdBarArr := StringArrRemoveDuplicates(barArr)
		// verify StringArrRemoveDuplicates
		assert.Equal(b, 5000, len(rdBarArr))
	}
}
