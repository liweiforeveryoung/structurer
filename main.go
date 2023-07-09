package main

import (
	"bytes"
	"code.byted.org/overpass/data_life_trade_order_search/kitex_gen/data/life/trade_order_search"
	"encoding/json"
	"fmt"
	"structurer/st"
)

type SA struct {
	Hello string
}

type SB struct {
	World string
	//SA          *SA
	//SliceInt    []int
	//SliceString []string
	SliceStruct []*SB
}

func main() {
	obj := trade_order_search.GetOrderResponse{}
	_ = json.Unmarshal([]byte(rawJson), &obj)
	//
	//obj := SB{
	//	//World: "world",
	//	//SA: &SA{
	//	//	Hello: "hello",
	//	//},
	//	//SliceInt:    []int{1, 2, 3},
	//	//SliceString: []string{"s1", "s2"},
	//	SliceStruct: []*SB{
	//		{
	//			World: "ss world",
	//		},
	//	},
	//}
	pathMap, bs, _ := st.Marshal(obj)
	fmt.Println("=======   pathMap:   ", GenerateImportPkgPath(pathMap))
	fmt.Println("=======   bs1:   ", string(bs))

	bs2, _ := json.Marshal(obj)
	fmt.Println("=======   bs2:   ", string(bs2))
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

func XXX() trade_order_search.GetOrderResponse {
	return trade_order_search.GetOrderResponse{}
}

var rawJson = "{     \"BaseResp\": {\n        \"StatusCode\": 0,\n        \"StatusMessage\": \"\"\n    }\n}"
