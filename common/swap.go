package common

import "encoding/json"

// json tag 进行结构体赋值
func SwapTo(request, category interface{}) (err error) {
	dataBye, err := json.Marshal(request)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataBye, category)
	return
}
