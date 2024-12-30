package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-balances/internal/database"
	"github.com.br/devfullcycle/fc-ms-balances/internal/usecase/create_balance"
	"github.com.br/devfullcycle/fc-ms-balances/internal/usecase/find_account"
	"github.com.br/devfullcycle/fc-ms-balances/internal/web"
	"github.com.br/devfullcycle/fc-ms-balances/internal/web/webserver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

type BalanceEvent struct {
	Name    string      `json:"Name"`
	Payload BalanceData `json:"Payload"`
}

type BalanceData struct {
	AccountIDFrom      string  `json:"account_id_from"`
	AccountIDTo        string  `json:"account_id_to"`
	BalanceAccountFrom float64 `json:"balance_account_id_from"`
	BalanceAccountTo   float64 `json:"balance_account_id_to"`
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "balances-mysql", "3306", "balances"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS balances (id varchar(255), account_id varchar(255), balance integer, created_at date)")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO balances (id, account_id, balance, created_at) VALUES ('b8bbd006-0318-48ba-b7a4-30dc5920f4c2', '36947137-0a7c-424d-a8be-94e77cd3a88f', 100, '2024-01-27 17:13:48'), ('ab37cbe9-b212-446b-a7ca-3dc8117281d3', '53ab91e2-1c32-4641-bd3a-0df347e82c78', 1150, '2024-01-27 17:13:48')")
	if err != nil {
		panic(err)
	}

	balanceDb := database.NewBalanceDB(db)

	createBalanceUseCase := create_balance.NewCreateBalanceUseCase(balanceDb)
	findAccountUseCase := find_account.NewFindAccountUseCase(balanceDb)

	go func() {
		webserver := webserver.NewWebServer(":3003")
		balanceHandler := web.NewWebBalanceHandler(*findAccountUseCase)
		webserver.AddHandler("/balances/{account_id}", balanceHandler.FindAccount)
		fmt.Println("Server running at port 3003")
		webserver.Start()
	}()

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"client.id":         "balances",
		"group.id":          "balances",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"balances"}
	c.SubscribeTopics(topics, nil)


	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)

			var event BalanceEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				fmt.Println("Error JSON:", err)
				continue
			}

			input := create_balance.CreateBalanceInputDTO{
				AccountID: event.Payload.AccountIDFrom,
				Balance:   event.Payload.BalanceAccountFrom,
			}
			createBalanceUseCase.Execute(input)

			input = create_balance.CreateBalanceInputDTO{
				AccountID: event.Payload.AccountIDTo,
				Balance:   event.Payload.BalanceAccountTo,
			}
			createBalanceUseCase.Execute(input)

			c.CommitMessage(msg)

		}

	}
}