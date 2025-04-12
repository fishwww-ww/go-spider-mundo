package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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
	gorm.Model
	Name         string `json:"Name" gorm:"column:name"`
	Introduction string `json:"Introduction" gorm:"column:introduction"`
	Require      string `json:"Require" gorm:"column:require"`
	Number       string `json:"Number" gorm:"column:number"`
	Publisher    string `json:"Publisher" gorm:"column:publisher"`
	Contact      string `json:"Contact" gorm:"column:contact"`
	ID           int    `json:"ID" gorm:"column:id"`
}

var DB *gorm.DB

func main() {
	InitDB()
	mundo()
}

func mundo() {
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
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	// 使用 GORM 插入数据
	result := DB.Create(&response.Data.Team.Content)
	if result.Error != nil {
		log.Fatal("插入数据失败:", result.Error)
	}

	fmt.Printf("成功插入 %d 条数据\n", result.RowsAffected)
}
