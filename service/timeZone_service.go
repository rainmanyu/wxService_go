package service

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type TimeInfo struct {
	Abbreviation string `json:"abbreviation"`
	ClientIp string `json:"client_ip"`
	Datetime string `json:"datetime"`
	DayOfWeek int `json:"day_of_week"`
	DayOfYear int `json:"day_of_year"`
	Dst bool `json:"dst"`
	DstFrom string `json:"dst_from"`
	DstOffset int  `json:"dst_offset"`
	DstUntil string `json:"dst_until"`
	RawOffset int `json:"raw_offset"`
	Timezone string `json:"timezone"`
	Unixtime int `json:"unixtime"`
	UtcDatetime string `json:"utc_datetime"`
	UtcOffset string `json:"utc_offset"`
	WeekNumber int `json:"week_number"`
  }

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://worldtimeapi.org/api/timezone/Asia/Shanghai")
	if err != nil {
		fmt.Fprint(w, "error")
	} else {
		fmt.Println("Hello world")
		fmt.Println(resp)
		
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		bodyStr := string(body)
		fmt.Println(bodyStr)

		timeObject := TimeInfo{}
		errRtn := json.Unmarshal([]byte(bodyStr), &timeObject)
		if errRtn != nil {
			fmt.Println(errRtn)
		}

		// fmt.Println(timeObject)
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(timeObject)		

		js, err := json.Marshal(timeObject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

