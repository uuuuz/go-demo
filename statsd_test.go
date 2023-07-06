package main

import (
	"context"
	"fmt"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/influxdata/influxdb-client-go/v2"
	"log"
	"testing"
	"time"
)

// influxdb demo

// Create a Database with a query
func createDatabase(conn client.Client) {
	q := client.NewQuery("CREATE DATABASE test", "", "")
	if response, err := conn.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}

func connInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://47.103.121.55:8086",
		Username: "admin",
		Password: "Welcome123",
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

// query
func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// insert
func writesPoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "telegraf",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   221.1,
		"system": 44.3,
		"user":   16.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}

func TestExecStatsd(t *testing.T) {
	conn := connInflux()
	fmt.Println(conn)

	//createDatabase(conn)
	// insert
	writesPoints(conn)
	//
	// 获取10条数据并展示
	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, err := queryDB(conn, qs)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range res[0].Series[0].Values {
		for j, value := range row {
			log.Printf("j:%d value:%v\n", j, value)
		}
	}
}

func TestStatsd(t *testing.T) {
	c := influxdb2.NewClient("http://localhost:8087", "a1mRtYjUrNdD4OyOuFjkEGbj3VVYQLhj0j1wvXU9QtyxS_7i6Oew__a_0dOHK4LcpA9jQ4pCPjB0Kw_4DeMbXw==")
	defer c.Close()

	// 可复用
	// c.WriteAPI()
	// q := client.NewQuery("SELECT count(value) FROM test", "test", "n")
	res, err := c.QueryAPI("wxm").Query(context.TODO(), `from(bucket:"test")|> range(start: -1h)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	//if response.Error() != nil {
	//	fmt.Println(response.Error())
	//}
	for res.Next() {
		fmt.Println(res.Record().Values())
	}
}
