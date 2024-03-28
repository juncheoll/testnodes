package healthPinger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"testNodes/setting"
	"testNodes/src/logCtrlr"
)

type RequestData struct {
	Port       string                         `json:"port"`
	Gpuname    string                         `json:"gpuname"`
	Model_info map[string]map[string]TaskInfo `json:"model_info"`
}

func postAlive(myPort string) {
	jsonData, err := json.Marshal(RequestData{
		Port:       myPort,
		Gpuname:    gpuName,
		Model_info: model_info,
	})
	if err != nil {
		panic(err)
	}
	resp, _ := http.Post("http://"+setting.ManagerUrl+"/alive", "application/json", bytes.NewBuffer(jsonData))
	if resp == nil || resp.StatusCode != http.StatusOK {
		logCtrlr.Error(errors.New("there is no manager" + myPort))

		//os.Exit(1)
	}
}
