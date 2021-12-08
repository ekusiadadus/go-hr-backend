# Go言語+Elasticsearchで簡易求人検索バックエンド構築

動作環境

|ツール|バージョン|
| --- | --- |
| Go言語 | 1.17.2 |
| Docker | 1.41 |
| docker-compose | 1.29.2  |
| Elasticsearch | 7.8.1 |
| Kibana | 7.8.1 |



##　事前準備 

/etc/sysctl.conf

```
vm.max_map_count=262144
```

.env

```
MYSQL_ROOT_PASSWORD=test
MYSQL_DATABASE=test
MYSQL_USER=test
MYSQL_PASSWORD=test
MYSQL_ALLOW_EMPTY_PASSWORD=yes
XML_PATH=[File Path]
```

求人データ構造体(テストデータ作成参考用)
```
type Jobs struct{
	Source xml.Name `xml:"source"`
	Publisher xml.Name `xml:"publisher"`
	Publisherurl xml.Name `xml:"publisherurl"`
	LastBuildDate xml.Name `xml:"lastBuildDate"`
	Jobs []Job `xml:"job"`
}

type Job struct {
	Referencenumber string  `xml:"referencenumber"`
	Date string `xml:"date"`
	Url string `xml:"url"`
	Title string `xml:"title"`
	Description string `xml:"description"`
	State string `xml:"state"`
	City string `xml:"city"`
	Country string `xml:"country"`
	Station string `xml:"station"`
	Jobtype string `xml:"jobtype"`
	Salary string `xml:"salary"`
	Category string `xml:"category"`
	ImageUrls string `xml:"imageUrls"`
	Timeshift string `xml:"timeshift"`
	Subwayaccess string `xml:"subwayaccess"`
	Keywords string `xml:"keywords"`
}
```

## 初回起動

### 1.

ルートディレクトリで、`docker-compose up`を行う
コンテナが立ち上がるまで、1-2分程かかります。

### 2.

テストデータの作成をしましょう。
構造体を参考に作成すれば、テストすることができます。

### 3.

データ投入をしましょう。
./batch 配下にある、LoadData.goを動かすとデータをElasticsearchに投入してくれます。
15万件で、2-3分程度かかります。

### 4.
./search_api配下で、`go build -o .`とすると、hr_apiのバイナリファイルが作成されます。
./hr_apiで起動しましょう。

### 5.

動作確認してみましょう。

```
 #東京都の"カフェ"の求人を検索する
http://localhost:5000/search?keyword=カフェ&state=東京都

 #東京都の"Go言語"の求人を検索する
http://localhost:5000/search?keyword=Go言語&state=東京都

 #神奈川県の"アルバイト・パート"の求人を検索する
http://localhost:5000/search?keyword=アルバイト・パート&state=神奈川県

 #求人のユニークidから検索する
http://localhost:5000/search?id=test
```

http://localhost:5000/search?[keyword]&[state]

|keyword|state|
|---|---|
|title, category, descriptionにkeywordを含む|47都道府県|
