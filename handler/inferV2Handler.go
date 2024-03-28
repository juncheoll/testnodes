package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testNodes/setting"
	"time"

	"github.com/gorilla/mux"
)

var mutexs []sync.Mutex = make([]sync.Mutex, 100)

func (h *Handler) testInferV2Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]
	model := vars["model"]
	version := vars["version"]
	port := vars["tn"]

	tn, err := strconv.Atoi(string(port[2:]))
	if err != nil {
		log.Println("오류오루ㅠㅇ룽")
		return
	}

	type requestData struct {
		Port      string  `json:"port"`
		Id        string  `json:"id"`
		Model     string  `json:"model"`
		Version   string  `json:"version"`
		BurstTime float64 `json:"burstTime"`
	}

	reqData := requestData{
		Port:      port,
		Id:        provider,
		Model:     model,
		Version:   version,
		BurstTime: 0,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		log.Println("marshal 실패", err)
		return
	}

	//매니저에게 인퍼런스 인크리즈 알리기
	http.Post("http://"+setting.ManagerUrl+"/inference/start", "application/json", bytes.NewBuffer(jsonData))

	mutexs[tn].Lock()
	defer mutexs[tn].Unlock()

	//요청받으면 랜덤한 인퍼런스 타임으로 결과값 돌려주기.
	inferTime := getRandNum(500, 10000)
	time.Sleep(time.Millisecond * time.Duration(inferTime))

	//랜덤한 확률로 인퍼런스 중 fault상황
	randRate := rand.Float64()
	if randRate < 0.1 {
		os.Exit(1)
	}
	reqData.BurstTime = float64(inferTime) / 1000
	jsonData, err = json.Marshal(reqData)
	if err != nil {
		log.Println("marshal 실패", err)
		return
	}

	//정상 수행 후 응답 상황
	http.Post("http://"+setting.ManagerUrl+"/inference/start", "application/json", bytes.NewBuffer(jsonData))
}

func getRandNum(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
