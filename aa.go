package main

import (
	"encoding/json"
	"fmt"
)

// Q1 123
func Q1() {
	var Data APIResult
	_ = json.Unmarshal(getAPIResult(), &Data)

	fmt.Printf("Q1: %+v \n", Data.Ret[0])
}

// Q2 123
func Q2() {
	/* Q2-1
	IP:192.17.55.3	Host:docs.google.com 		Port:80
	IP:192.17.33.17 Host:ts-phpadmin.vir999.com Port:80
	IP:192.17.99.52 Host:jsonviewer.stack.hu 	Port:7711
	*/
	var setList []*conn
	set1 := conn{IP: "192.17.55.3", Host: "docs.google.com", Port: "80"}
	set2 := conn{IP: "192.17.33.17", Host: "ts-phpadmin.vir999.com", Port: "80"}
	set3 := conn{IP: "192.17.99.52", Host: "jsonviewer.stack.hu", Port: "7711"}

	setList = append(setList, &set1)
	setList = append(setList, &set2)
	setList = append(setList, &set3)

	fmt.Printf("Q2-1: ")
	printResult(setList)

	/* Q2-2
	IP:177.18.2.33 Host:github.com Port:80
	*/
	set4 := conn{IP: "177.18.2.33", Host: "github.com", Port: "80"}
	setList = append(setList, &set4)

	fmt.Printf("Q2-2: ")
	printResult(setList)

	// Q2-3
	updateSetting(setList)

	fmt.Printf("Q2-3: ")
	printResult(setList)
}

// getAPIResult 321
func getAPIResult() []byte {
	userData := `{
			"ret":[{
					"id":148500286,
					"parent_id":1414663,
					"parent":1414663,
					"all_parents":[1414663,59735,58971,24367,6],
					"domain":6,
					"last_login":"2018-06-12T09:17:51+0800",
					"currency":"CNY",
					"password_expire_at":"2020-09-04T18:19:02+0800",
					"password_reset":false,
					"last_bank":54,
					"username":"wesley01",
					"block":false,
					"bankrupt":false,
					"cash":{
						"id":71814701,
						"user_id":148500286,
						"balance":4923642.0545,
						"pre_sub":0,
						"pre_add":0,
						"currency":"CNY",
						"last_entry_at":"20180612155900"
					},
					"cash_fake":null,
					"card":null,
					"enabled_card":null
			}],
			"result":"ok",
			"profile":{
					"execution_time":14,
					"server_name":"ipl-web01.rd5.qa"
			}
		}`

	return []byte(userData)
}

func updateSetting(setList []*conn) {
	for i, set := range setList {
		if set.Host == "jsonviewer.stack.hu" {
			setList[i].Port = "80"
		}
	}
}

func printResult(result []*conn) {
	for _, set := range result {
		fmt.Printf("%+v ", *set)
	}

	fmt.Println()
}
