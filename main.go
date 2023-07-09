package main

import (
	"code.byted.org/overpass/data_life_trade_order_search/kitex_gen/data/life/trade_order_search"
	"fmt"
	"structurer/json"
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
	sb := SB{
		//World: "world",
		//SA: &SA{
		//	Hello: "hello",
		//},
		//SliceInt:    []int{1, 2, 3},
		//SliceString: []string{"s1", "s2"},
		SliceStruct: []*SB{
			{
				World: "ss world",
			},
		},
	}
	bs, _ := json.Marshal(sb)
	fmt.Println(string(bs))
}

var rawJson = "{\n    \"OrderId\": \"1003379872014106040\",\n    \"OrderDetailAccess\": {\n        \"GetDetail\": true,\n        \"GetSnapShot\": true,\n        \"GetOrderItem\": true,\n        \"GetOrderFee\": true,\n        \"GetExchangeProductSnap\": true,\n        \"GetOrderService\": true\n    },\n    \"NeedFakeItem\": false,\n    \"MasterQuery\": false,\n    \"Downgrade2MasterQuery\": false,\n    \"Base\": {\n        \"LogID\": \"\",\n        \"Caller\": \"\",\n        \"Addr\": \"\",\n        \"Client\": \"\",\n        \"TrafficEnv\": {\n            \"Open\": false,\n            \"Env\": \"\"\n        },\n        \"Extra\": {\n            \"\": \"\",\n            \"env\": \"prod\"\n        }\n    }\n}"
