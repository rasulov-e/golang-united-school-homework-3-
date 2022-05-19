package homework

import (
	"fmt"
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	tC := []struct{
		input  []int64
		output []int64
	}{
		{
			input: []int64{10,9,8,7,6,5,4,3,2,1},
			output: []int64{10,9,8,7,6,5,4,3,2,1},
		},
	}
	for i, v := range tC {
		t.Run(fmt.Sprint("reverse-", i), func(t *testing.T) {
			res := reverse(v.input)
			log.Println("what we got",res)
			log.Println("what we want",v.output)
			t.Error("incorrect len")
			if len(res) != len(v.output) {
			}
		})
	}

}