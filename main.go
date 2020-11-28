package main

import (
	"fmt"
	"github.com/headend/share-module/model"
)

func main()  {
	var test model.WorkerUpdateSignal
	test.FilePath = "ss"
	fmt.Printf("%#v" ,test)
	print("Shared module\n")
}
