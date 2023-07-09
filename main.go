package main

import (
	"fmt"
	"structurer/json"
)

type SA struct {
	Hello string
}

type SB struct {
	World       string
	SA          *SA
	SliceInt    []int
	SliceString []string
	SliceStruct []*SB
}

func main() {
	sb := SB{
		World: "world",
		SA: &SA{
			Hello: "hello",
		},
		SliceInt:    []int{1, 2, 3},
		SliceString: []string{"s1", "s2"},
		SliceStruct: []*SB{
			{
				World: "ss world",
			},
		},
	}
	bs, _ := json.Marshal(sb)
	fmt.Println(string(bs))
}
