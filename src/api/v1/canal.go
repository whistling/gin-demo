package v1

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
	"log"
	"time"
)

var (
	redisCon *redis.Client
	canalCon *client.SimpleCanalConnector
)

func init() {
	canalCon = client.NewSimpleCanalConnector("127.0.0.1", 11111, "", "", "example", 60000, 60*60*1000)
	err := canalCon.Connect()
	checkError(err)
	// https://github.com/alibaba/canal/wiki/AdminGuide
	//mysql 数据解析关注的表，Perl正则表达式.
	//多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\)
	//
	//常见例子：
	//
	//  1.  所有表：.*   or  .*\\..*
	//	2.  canal schema下所有表： canal\\..*
	//	3.  canal下的以canal打头的表：canal\\.canal.*
	//	4.  canal schema下的一张表：canal\\.test1
	//  5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)
	//err = connector.Subscribe(".*\\..*")
	err = canalCon.Subscribe(".*\\..*")
	checkError(err)

	//redis connect
	redisCon = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

}

func run() {
	for {
		message, err := canalCon.Get(100, nil, nil)
		checkError(err)
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(3 * time.Second)
			fmt.Println("===没有数据了===")
			continue
		}
		fmt.Println("batch id is ", batchId)
		dealEntry(message.Entries)
	}
}

func dealEntry(entrys []protocol.Entry) {
	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN ||
			entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if rowChange != nil {
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == protocol.EventType_DELETE {
					deleteRedis(rowData.GetBeforeColumns())
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == protocol.EventType_INSERT {
					insertRedis(rowData.GetBeforeColumns())
					printColumn(rowData.GetAfterColumns())
				} else if eventType == protocol.EventType_UPDATE {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}
	}
}

func deleteRedis(columns []*protocol.Column) {

}

func updateRedis(columns []*protocol.Column) {

}

func insertRedis(columns []*protocol.Column) {

}

func printColumn(columns []*protocol.Column) {
	fmt.Println(columns)
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
