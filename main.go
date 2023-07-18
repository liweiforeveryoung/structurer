package main

import (
	"bytes"
	"fmt"
	"structurer/st"
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
	obj := SB{
		World: "world",
		SA: &SA{
			Hello: "hello",
		},
		SliceInt: []int{
			1,
			2,
			3,
		},
		SliceString: []string{
			"s1",
			"s2",
			"s3",
		},
		SliceStruct: []*SB{
			nil,
			{
				World: "world2",
			},
		},
	}
	pathMap, bs, _ := st.Marshal(obj)
	fmt.Println("=======   pathMap:   ", GenerateImportPkgPath(pathMap))
	fmt.Println("=======   bs1:   ", string(bs))
}

func GenerateImportPkgPath(pathMap map[string]string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("import (\n")
	for shortcut, completePath := range pathMap {
		buf.WriteString(shortcut)
		buf.WriteString(" \"")
		buf.WriteString(completePath)
		buf.WriteByte('"')
		buf.WriteString("\n")
	}
	buf.WriteString(")")
	return buf.String()
}
