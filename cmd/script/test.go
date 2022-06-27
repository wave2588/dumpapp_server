package main

import (
	"dumpapp_server/pkg/common/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//now := time.Now().UnixNano() / 1e6
	//res := make(map[string]struct{})
	//for i := 0; i <= 1000; i++ {
	//	id := util.MustGenerateAppCDKEY()
	//	if _, ok := res[id]; ok {
	//		fmt.Println("存在了--->: ", id)
	//		continue
	//	}
	//	res[id] = struct{}{}
	//}
	//fmt.Println(time.Now().UnixNano()/1e6 - now)

	url := fmt.Sprintf("https://itunes.apple.com/cn/lookup?id=741292507")
	//bodyJson, _ := json.Marshal(map[string]interface{}{
	//	"id": "741292507",
	//})
	resp, err := http.DefaultClient.Post(url, "application/json", nil)
	//req, err := http.NewRequest("POST", url, strings.NewReader(string("id=741292507")))
	util.PanicIf(err)
	//req.Header.Add("Cache-Control", "no-cache")
	//resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}

	util.PanicIf(json.Unmarshal(body, &data))

	ss, _ := json.Marshal(data)
	fmt.Println(string(ss))
}
