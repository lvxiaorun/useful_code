package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
)

type TestRedis struct {
	Name string
	Age int
}
func GetCon() redis.Conn{
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("connect to redis err:",err)
	}
	return c
}

func main(){
	c := GetCon()
	name,err := redis.Bytes(c.Do("get","name"))
	if err != nil{
		fmt.Println("get key name err:",err)
	}
	fmt.Println(name)
	age,err1 := redis.Int(c.Do("get","age"))
	if err1 != nil{
		fmt.Println("get key age err:",err1)
	}
	_,err2 := c.Do("set","test",TestRedis{"haha",11})
	if err2 != nil{
		fmt.Println("set key test err:",err2)
	}
	tr := TestRedis{"test",20}
	bs,_ := json.Marshal(tr)
	c.Do("set","test1",string(bs))
	ge,_ := redis.String(c.Do("get","test1"))
	get := TestRedis{}
	json.Unmarshal([]byte(ge),&get)
	println(age)
	println(string(name))
	println(get.Age)
}
