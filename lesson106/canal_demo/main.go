package main

import (
	"fmt"
	"time"

	pbe "github.com/Q1mi/canal-go/protocol/entry"

	"github.com/Q1mi/canal-go/client"
	"google.golang.org/protobuf/proto"
)

// canal-go client demo

func main() {
	// 连接canal-server
	// 请修改为你的 canal-server 配置
	connector := client.NewSimpleCanalConnector(
		"127.0.0.1", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		panic(err)
	}

	// mysql 数据解析关注的表，Perl正则表达式.
	err = connector.Subscribe(".*\\..*")
	if err != nil {
		fmt.Printf("connector.Subscribe failed, err:%v\n", err)
		panic(err)
	}

	// 消费消息
	for {
		message, err := connector.Get(100, nil, nil)
		if err != nil {
			fmt.Printf("connector.Get failed, err:%v\n", err)
			continue
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(time.Second)
			fmt.Println("===暂无数据===")
			continue
		}
		printEntry(message.Entries)
	}
}

func printEntry(entries []*pbe.Entry) {
	for _, entry := range entries {
		// 忽略事务开启和事务关闭类型
		if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN ||
			entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
			continue
		}
		// RowChange对象，包含了一行数据变化的所有特征
		rowChange := new(pbe.RowChange)
		// protobuf解析
		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		if err != nil {
			fmt.Printf("proto.Unmarshal failed, err:%v\n", err)
		}
		if rowChange == nil {
			continue
		}
		// 获取并打印Header信息
		header := entry.GetHeader()
		fmt.Printf("binlog[%s : %d], name[%s,%s], eventType: %s\n",
			header.GetLogfileName(),
			header.GetLogfileOffset(),
			header.GetSchemaName(), // 数据库
			header.GetTableName(),  // 数据表
			header.GetEventType(),  // 变更类型
		)
		//判断是否为DDL语句
		if rowChange.GetIsDdl() {
			fmt.Printf("isDdl:true, sql:%v\n", rowChange.GetSql())
		}

		// 获取操作类型：insert/update/delete等
		eventType := rowChange.GetEventType()
		for _, rowData := range rowChange.GetRowDatas() {
			if eventType == pbe.EventType_DELETE {
				printColumn(rowData.GetBeforeColumns())
			} else if eventType == pbe.EventType_INSERT || eventType == pbe.EventType_UPDATE {
				printColumn(rowData.GetAfterColumns())
			} else {
				fmt.Println("---before---")
				printColumn(rowData.GetBeforeColumns())
				fmt.Println("---after---")
				printColumn(rowData.GetAfterColumns())
			}
		}
	}
}

func printColumn(columns []*pbe.Column) {
	for _, col := range columns {
		fmt.Printf("%s:%s  updated=%v\n", col.GetName(), col.GetValue(), col.GetUpdated())
	}
}
