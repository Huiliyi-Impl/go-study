package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex" doc:"性别"`
}

func findTag(obj interface{}) {
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Printf("info:%s,doc:%s\n", tagInfo, tagDoc)
	}
}
func main() {
	r := resume{
		Name: "小明",
		Sex:  "男",
	}
	fmt.Println(r)
	findTag(&r)
}
