package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Query struct {
	Keyword string `query:"keyword"`
	State   string `query:"state"`
	Id      string `query:"id"`
}

type Result struct {
	Referencenumber string `xml:"referencenumber" json:"referencenumber,string"`
	Date            string `xml:"date" json:"date,string"`
	Url             string `xml:"url" json:"url,string"`
	Title           string `xml:"title" json:"title,string"`
	Description     string `xml:"description" json:"description,string"`
	State           string `xml:"state" json:"state,string"`
	City            string `xml:"city" json:"city,string"`
	Country         string `xml:"country" json:"country,string"`
	Station         string `xml:"station" json:"station,string"`
	Jobtype         string `xml:"jobtype" json:"jobtype,string"`
	Salary          string `xml:"salary" json:"salary,string"`
	Category        string `xml:"category" json:"category,string"`
	ImageUrls       string `xml:"imageUrls" json:"imageurls,string"`
	Timeshift       string `xml:"timeshift" json:"timeshift,string"`
	Subwayaccess    string `xml:"subwayaccess" json:"subwayaccess,string"`
	Keywords        string `xml:"keywords" json:"keywords,string"`
}
type Response struct {
	Message string `json:"message"`
	Results []Result
}

func HRSearch(c echo.Context) (err error) {
	// クライアントからのパラメーターを取得
	q := new(Query)
	if err = c.Bind(q); err != nil {
		return
	}

	res := new(Response)
	var (
		b   map[string]interface{}
		buf bytes.Buffer
	)

	// elasticsearch へのクエリを作成
	query := CreateQuery(q)

	json.NewEncoder(&buf).Encode(query)

	fmt.Printf(buf.String())

	// elasticsearch へ接続
	es, err := ConnectElasticsearch()
	if err != nil {
		c.Error(err)
	}

	// elasticsearch へクエリ
	r, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("baito"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		c.Error(err)
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		c.Error(err)
	}

	// クエリの結果を Responce.Results に格納
	for _, hit := range b["hits"].(map[string]interface{})["hits"].([]interface{}) {
		result := new(Result)
		doc := hit.(map[string]interface{})

		fmt.Printf(result.Title)

		result.Referencenumber = doc["_source"].(map[string]interface{})["referencenumber"].(string)
		result.Date = doc["_source"].(map[string]interface{})["date"].(string)
		result.Url = doc["_source"].(map[string]interface{})["url"].(string)
		result.Title = doc["_source"].(map[string]interface{})["title"].(string)
		result.State = doc["_source"].(map[string]interface{})["state"].(string)
		result.Category = doc["_source"].(map[string]interface{})["category"].(string)
		result.Description = doc["_source"].(map[string]interface{})["description"].(string)
		result.City = doc["_source"].(map[string]interface{})["city"].(string)
		result.Country = doc["_source"].(map[string]interface{})["country"].(string)
		result.Station = doc["_source"].(map[string]interface{})["station"].(string)
		result.Jobtype = doc["_source"].(map[string]interface{})["jobtype"].(string)
		result.Salary = doc["_source"].(map[string]interface{})["salary"].(string)
		result.ImageUrls = doc["_source"].(map[string]interface{})["imageurls"].(string)
		result.Timeshift = doc["_source"].(map[string]interface{})["timeshift"].(string)
		result.Subwayaccess = doc["_source"].(map[string]interface{})["subwayaccess"].(string)
		result.Keywords = doc["_source"].(map[string]interface{})["keywords"].(string)

		res.Results = append(res.Results, *result)
	}

	res.Message = "検索に成功しました。"

	return c.JSON(http.StatusOK, res)
}