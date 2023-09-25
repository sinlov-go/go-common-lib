package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistanceBetweenTimestampHuman(t *testing.T) {
	// mock DistanceBetweenTimestampSecondHuman
	tests := []struct {
		name       string
		form       int64
		to         int64
		wantResult string
	}{
		{
			name:       "zero",
			form:       0,
			to:         0,
			wantResult: "0s",
		},
		{
			name:       "less",
			form:       10000,
			to:         0,
			wantResult: "N/A",
		},
		{
			name:       "second",
			form:       1695486518,
			to:         1695486528,
			wantResult: "10s",
		},
		{
			name:       "minutes",
			form:       1695386518,
			to:         1695388528,
			wantResult: "33m 30s",
		},
		{
			name:       "hours",
			form:       1695386518,
			to:         1695398528,
			wantResult: "3h 20m 10s",
		},
		{
			name:       "days",
			form:       1695386518,
			to:         1695486528,
			wantResult: "1d 3h 46m 50s",
		},
		{
			name:       "more times",
			form:       1685386518,
			to:         1695486528,
			wantResult: "116d 21h 33m 30s",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do DistanceBetweenTimestampSecondHuman
			gotResult := DistanceBetweenTimestampSecondHuman(tc.form, tc.to)

			// verify DistanceBetweenTimestampSecondHuman
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}
