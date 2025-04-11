package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Team Team `json:"Team"`
}

type Team struct {
	Content []Content `json:"Content"`
}

type Content struct {
	Name         string `json:"Name"`
	Introduction string `json:"Introduction"`
	Require      string `json:"Require"`
	Number       string `json:"Number"`
	Publisher    string `json:"Publisher"`
	Contact      string `json:"Contact"`
	ID           int    `json:"ID"`
}

var DB *sql.DB

func main() {
	InitDB()
	mundo()
}

func mundo() {
	//url := "https://api.bilibili.com/x/v2/reply/main?callback=jQuery17205371302484233957_164102593161&jsonp=jsonp&next=0&type=1&oid=251119469&mode=3&plat=1&_=1641025931610"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://qgdoywhgtdnh.sealosbja.site/timerme/api/allteam?service=mundo", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36 Edg/135.0.0.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(body))
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %+v\n Introduction: %+v\n Require: %+v\n Number: %+v\n", response.Data.Team.Content[0].Name, response.Data.Team.Content[0].Introduction, response.Data.Team.Content[0].Require, response.Data.Team.Content[0].Number)

}
