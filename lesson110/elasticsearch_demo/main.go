package main

// go-elasticsearch demo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {

	//  连接ES
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // 你的ES服务地址
		},
	}
	// elasticsearch.NewClient()
	// elasticsearch.NewDefaultClient()  // 默认的 连接127
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("NewTypedClient failed, err:%v\n", err)
		return
	}

	// 连接成功
	fmt.Println(client)

	// 创建index
	// createIndex(client)

	// 创建document
	// indexDocument(client)
	// indexDocument2(client)

	// 查询document
	// getdocumentByID(client, "2")

	// 检索document
	searchDocument(client)
	// searchDocument2(client)

	// agg
	aggregationDemo(client)

	// update document
	// updateDocument(client)
	// updateDocument2(client)

	// delete document
	deleteDocument(client)

}

func createIndex(client *elasticsearch.TypedClient) {
	resp, err := client.Indices.
		Create("my-review-1").
		Do(context.Background())
	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return
	}
	fmt.Printf("Acknowledged:%v\n", resp.Acknowledged)
}

type Review struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userID"`
	Score       uint8     `json:"score"`
	Content     string    `json:"content"`
	Tags        []Tag     `json:"tags"`
	Status      int       `json:"status"`
	PublishTime time.Time `json:"publishTime"`
}

type Tag struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}

// indexDocument 索引文档
func indexDocument(client *elasticsearch.TypedClient) {
	// 定义 document 结构体对象
	d1 := Review{
		ID:      1,
		UserID:  147982601,
		Score:   5,
		Content: "这是一个好评！",
		Tags: []Tag{
			{1000, "好评"},
			{1100, "物超所值"},
			{9000, "有图"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	// 添加文档
	resp, err := client.Index("my-review-1").
		Id(strconv.FormatInt(d1.ID, 10)).
		Document(d1).
		Do(context.Background())
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}

// indexDocument2 索引文档
func indexDocument2(client *elasticsearch.TypedClient) {
	// 定义 document 结构体对象
	d1 := Review{
		ID:      2,
		UserID:  147982601,
		Score:   1,
		Content: "这是一个差评！",
		Tags: []Tag{
			{2000, "差评"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	// 添加文档
	resp, err := client.Index("my-review-1").
		Id(strconv.FormatInt(d1.ID, 10)).
		Document(d1).
		Do(context.Background())
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}

func getdocumentByID(client *elasticsearch.TypedClient, id string) {
	resp, err := client.Get("my-review-1", id).
		Do(context.Background())
	if err != nil {
		fmt.Printf("get document by id failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%s\n", resp.Source_)
}

// searchDocument 搜索所有文档
func searchDocument(client *elasticsearch.TypedClient) {
	// 搜索文档
	resp, err := client.Search().
		Index("my-review-1").
		Request(&search.Request{
			Query: &types.Query{
				MatchAll: &types.MatchAllQuery{},
			},
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("search document failed, err:%v\n", err)
		return
	}
	fmt.Printf("total: %d\n", resp.Hits.Total.Value)
	// 遍历所有结果
	for _, hit := range resp.Hits.Hits {
		fmt.Printf("%s\n", hit.Source_)
	}
}

// searchDocument2 指定条件搜索文档
func searchDocument2(client *elasticsearch.TypedClient) {
	// 搜索content中包含好评的文档
	resp, err := client.Search().
		Index("my-review-1").
		Request(&search.Request{ // "github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
			Query: &types.Query{
				MatchPhrase: map[string]types.MatchPhraseQuery{
					"content": {Query: "好评"},
				},
			},
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("search document failed, err:%v\n", err)
		return
	}
	fmt.Printf("total: %d\n", resp.Hits.Total.Value)
	// 遍历所有结果
	for _, hit := range resp.Hits.Hits {
		fmt.Printf("%s\n", hit.Source_)
	}
}

// aggregationDemo 聚合
func aggregationDemo(client *elasticsearch.TypedClient) {
	avgScoreAgg, err := client.Search().
		Index("my-review-1").
		Request(
			&search.Request{
				Size: some.Int(0),
				Aggregations: map[string]types.Aggregations{
					"avg_score": { // 将所有文档的 score 的平均值聚合为 avg_score
						Avg: &types.AverageAggregation{
							Field: some.String("score"),
						},
					},
				},
			},
		).Do(context.Background())
	if err != nil {
		fmt.Printf("aggregation failed, err:%v\n", err)
		return
	}
	fmt.Printf("avgScore:%#v\n", avgScoreAgg.Aggregations["avg_score"])
}

// updateDocument 更新文档
func updateDocument(client *elasticsearch.TypedClient) {
	// 修改后的结构体变量
	d1 := Review{
		ID:      1,
		UserID:  147982601,
		Score:   5,
		Content: "这是一个修改后的好评！", // 有修改
		Tags: []Tag{ // 有修改
			{1000, "好评"},
			{9000, "有图"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	resp, err := client.Update("my-review-1", "1").
		Doc(d1). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		fmt.Printf("update document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}

// updateDocument2 更新文档
func updateDocument2(client *elasticsearch.TypedClient) {
	// 修改后的json数据
	jsonStr := `{
		"id":1,
		"userID":147982601,
		"score":5,
		"content":"这是一个二次修改后的好评！",
		"tags":[
			{
				"code":1000,
				"title":"好评"
			},
			{
				"code":9000,
				"title":"有图"
			}
		],
		"status":2,
		"publishDate":"2023-12-10T15:27:18.219385+08:00"
	}`
	resp, err := client.Update("my-review-1", "1").
		Request(&update.Request{
			Doc: json.RawMessage(jsonStr),
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("update document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}

// deleteDocument 删除 document
func deleteDocument(client *elasticsearch.TypedClient) {
	resp, err := client.Delete("my-review-1", "1").
		Do(context.Background())
	if err != nil {
		fmt.Printf("delete document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}
