package main

import (
	"fmt"
	"sync"
	"time"
)

// 操作失败信息
const (
	LoopFail     = "operation doesn't exist"
	AddFail      = "user has already existed"
	DesFail      = "user doesn't exist"
	WithdrawFail = "user doesn't exist or value is bigger than balance"
	QueryFail    = "user doesn't exist"
)

// 操作失败信息映射
var FailMes = map[string]string{
	"Loop":     LoopFail,
	"AddUser":  AddFail,
	"Desposit": DesFail,
	"Withdraw": WithdrawFail,
	"Query":    QueryFail,
}

// Bank 银行
type Bank struct {
	saving map[string]int
}

// Request 银行存取操作
type Request struct {
	op    string
	name  string
	value int
	retCh chan *Result
}

// Result 执行结果
type Result struct {
	status  bool
	balance int
}

// NewBank 新建银行
func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b
}

// Loop 银行处理客户端请求
func (b *Bank) Loop(reqCh chan *Request) {
	for req := range reqCh {
		switch req.op {
		case "Desposit":
			b.Desposit(req)
		case "Withdraw":
			b.Withdraw(req)
		case "Query":
			b.Query(req)
		case "AddUser":
			b.AddUser(req)
		default:
			ret := &Result{
				status:  false,
				balance: 0,
			}
			req.retCh <- ret
		}
	}

	// 无请求时银行退出
	fmt.Println("Bank exit")
}

// AddUser 新增用户操作
func (b *Bank) AddUser(req *Request) {
	name := req.name

	var status bool

	if _, ok := b.saving[name]; !ok {
		status = true
		b.saving[name] = 0
	}

	ret := &Result{
		status:  status,
		balance: 0,
	}
	req.retCh <- ret
}

// Desposit 存款操作
func (b *Bank) Desposit(req *Request) {
	name := req.name
	value := req.value

	var (
		ok      bool
		balance int
	)

	if _, ok = b.saving[name]; ok {
		b.saving[name] += value
		balance = b.saving[name]
	}

	ret := &Result{
		status:  ok,
		balance: balance,
	}
	req.retCh <- ret
}

// Withdraw 取款操作
func (b *Bank) Withdraw(req *Request) {
	name := req.name
	value := req.value

	var (
		status  bool
		balance int
	)

	if balance, ok := b.saving[name]; ok && balance >= value {
		status = true
		b.saving[name] -= value
		balance = b.saving[name]
	}

	ret := &Result{
		status:  status,
		balance: balance,
	}
	req.retCh <- ret
}

// Query 查询余额操作
func (b *Bank) Query(req *Request) {
	name := req.name

	balance, ok := b.saving[name]

	ret := &Result{
		status:  ok,
		balance: balance,
	}
	req.retCh <- ret
}

// xiaoming 客户小明的操作
func xiaoming(wg *sync.WaitGroup, reqCh chan<- *Request) {
	name := "xiaoming"
	retCh := make(chan *Result)
	defer func() {
		close(retCh)
		wg.Done()
	}()

	addReq := &Request{
		op:    "AddUser",
		name:  name,
		retCh: retCh,
	}
	depReq := &Request{
		op:    "Desposit",
		name:  name,
		value: 100,
		retCh: retCh,
	}
	withdrawReq := &Request{
		op:    "Withdraw",
		name:  name,
		value: 110,
		retCh: retCh,
	}
	queryReq := &Request{
		op:    "Query",
		name:  name,
		retCh: retCh,
	}

	reqs := []*Request{addReq, depReq, withdrawReq, queryReq}
	for _, req := range reqs {
		reqCh <- req
		waitResp(req)
	}
}

// xiaogang 客户小刚的操作
func xiaogang(wg *sync.WaitGroup, reqCh chan<- *Request) {
	name := "xiaogang"
	retCh := make(chan *Result)
	defer func() {
		close(retCh)
		wg.Done()
	}()

	addReq := &Request{
		op:    "AddUser",
		name:  name,
		retCh: retCh,
	}
	depReq := &Request{
		op:    "Desposit",
		name:  name,
		value: 200,
		retCh: retCh,
	}
	withdrawReq := &Request{
		op:    "Withdraw",
		name:  name,
		value: 70,
		retCh: retCh,
	}
	queryReq := &Request{
		op:    "Query",
		name:  name,
		retCh: retCh,
	}

	reqs := []*Request{addReq, depReq, withdrawReq, queryReq}
	for _, req := range reqs {
		reqCh <- req
		waitResp(req)
	}
}

// waitResp 等待请求响应req，输出信息
func waitResp(req *Request) {
	ret := <-req.retCh
	if ret.status {
		if req.op == "Desposit" || req.op == "Withdraw" {
			fmt.Printf("%s %s %d success, balance = %d.\n", req.name, req.op, req.value, ret.balance)
		} else {
			fmt.Printf("%s %s success, balance = %d.\n", req.name, req.op, ret.balance)
		}
	} else {
		if req.op == "Desposit" || req.op == "Withdraw" {
			fmt.Printf("%s %s %d fail, message = %s.\n", req.name, req.op, req.value, FailMes[req.op])
		} else if req.op == "AddUser" || req.op == "Query" {
			fmt.Printf("%s %s fail, message = %s.\n", req.name, req.op, FailMes[req.op])
		} else {
			fmt.Printf("%s %s fail, message = %s.\n", req.name, req.op, LoopFail)
		}
	}
}

func main() {
	// 创建请求的通道和银行
	reqCh := make(chan *Request, 100)
	bank := NewBank()

	// 银行处理请求
	go bank.Loop(reqCh)

	// 小明和小刚2个协程同时存取钱
	var wg sync.WaitGroup
	wg.Add(2)
	go xiaoming(&wg, reqCh)
	go xiaogang(&wg, reqCh)

	// 等待小明和小刚完成
	wg.Wait()
	close(reqCh)

	// 等待看银行是否退出
	time.Sleep(time.Second)
}
