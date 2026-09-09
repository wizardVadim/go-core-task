package main

import "fmt"

func main() {
	obj := NewStringIntMap()
	obj.Add("key", 1)
	obj.Add("sec", 2)
	fmt.Println(obj.Exists("sec"))
	obj.Remove("sec")
	fmt.Println(obj.Get("key"))
	newMap := obj.Copy()
	newMap["key"] = 100
	fmt.Println(obj.source)
	fmt.Println(newMap)
}
