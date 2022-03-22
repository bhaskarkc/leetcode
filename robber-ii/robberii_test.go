package robberii_test

import (
	"testing"

	robberii "xlinkerz/leetcode/robber-ii"

	"github.com/stretchr/testify/assert"
)


func TestRob(t *testing.T) {
	t.Parallel()

	type fields struct {
        houses []int
	}

	type want struct {
        money int
    }

	testTable := map[string]struct {
		Fields fields
		Want   want
	}{
        "2_3" : {
        	Fields: fields{ houses: []int{2,3}},
        	Want: want{3},
        },
        "1_2_3" : {
            Fields: fields{houses: []int{1,2,3}},
            Want: want{3},
        },
        "3_2_3" : {
        	Fields: fields{ houses: []int{3,2,3}},
        	Want: want{3},
        },
        "2_3_2" : {
            Fields: fields{houses: []int{2,3,2}},
            Want: want{3},
        },
        "1_2_3_1" : {
            Fields: fields{houses: []int{1,2,3,1}},
            Want: want{4},
        },
        "[1,3,1,3,100]" : {
            Fields: fields{houses: []int{1,3,1,3,100}},
            Want: want{103},
        },
    }

	for name, tt := range testTable {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
            m := robberii.Rob(tt.Fields.houses)
            assert.Equal(t, m, tt.Want.money)
        })
    }
}
