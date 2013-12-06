//
// Main tes
// @author Yu Dongxu(iamyudongxu@gmail.com)
// @date 2013-12-06
//
package main

import "fmt"
import "commonlib/array"
import "commonlib/timeutil"

func main(){
	arr := array.New()
    arr.Push("aaa")
	arr.Append('a',2,true)
	arr.Push("c")

	fmt.Println(arr.Size())
	fmt.Println(arr.Array())
	top,_ := arr.Pop()
	fmt.Println(top)
	fmt.Println(arr.Size())

	t,err := timeutil.ParseCn("2015-02-05 16:30:35")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
}
