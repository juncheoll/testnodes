package healthPinger

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"testNodes/setting"
)

var port string
var gpuName string

// var model_info map[string]map[string]TaskInfo = make(map[string]map[string]TaskInfo)
var model_info map[string]map[string]TaskInfo

type TaskInfo struct {
	LoadedAmount         int     `json:"loaded_amount"`
	AverageInferenceTime float32 `json:"average_inference_time"`
}

func Enter() {
	model_info = make(map[string]map[string]TaskInfo)

	port = setting.ServerPort
	cmd := exec.Command("nvidia-smi", "--query-gpu=name", "--format=csv,noheader")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("failed to get GPU name: ", err)
	}

	gpuName = strings.TrimSpace(string(output))
	log.Println("GPU NAME :", gpuName)

	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 500)
		go createNode(i)
	}
}

func createNode(i int) {
	myPort := port[:len(port)-2] + fmt.Sprintf("%02d", i)

	go alivePoster(myPort)
}

func alivePoster(myPort string) {
	postAlive(myPort)
	var cnt int = 0
	go testUpdateModel(myPort)

	for {
		cnt++
		//log.Printf("* (System) Send information to the Scheduler. (It is the %dth request)\n", cnt)

		postAlive(myPort)

		time.Sleep(8 * time.Second)
	}
}

func testUpdateModel(myPort string) {
	time.Sleep(time.Millisecond * 1000)

	type testModel struct {
		Provider string
		Name     string
		Version  string
	}

	testModels := []testModel{
		{"meta", "Llama-2-7B-Chat", "1"},
	}

	for _, model := range testModels {
		UpdateModel(model.Provider, model.Name, model.Version, myPort)
	}
}
