package util_test

import (
	"fmt"
	"github.com/reinanhs/golang-basic-example/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatDataPtBr(t *testing.T) {
	tests := []struct {
		name   string
		time   string
		result string
	}{
		{
			name:   "success 1",
			time:   "2014-11-12T11:45:26.371Z",
			result: "12/11/2014 11:45",
		},
		{
			name:   "success 2",
			time:   "2014-11-12T07:04:26.371Z",
			result: "12/11/2014 07:04",
		},
		{
			name:   "success 3",
			time:   "2014-06-01T07:04:26.371Z",
			result: "01/06/2014 07:04",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt, _ := time.Parse(time.RFC3339, tt.time)
			r := util.FormatDataPtBr(lt)

			assert.Equal(t, tt.result, r, "o teste falhou porque a data não corresponde ao valor esperado")
		})
	}
}

func ExampleFormatDataPtBr() {

	lt, _ := time.Parse(time.RFC3339, "2022-06-01T07:04:26.371Z")
	fmt.Println(util.FormatDataPtBr(lt))

	// Output: 01/06/2022 07:04
}
