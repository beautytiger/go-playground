package main

import (
	"fmt"
	//错，不可以调用internal包
	//https://stackoverflow.com/questions/59342373/use-of-internal-package-not-allowed
	"helm.sh/helm/v3/internal/version"
)

func main() {
	fmt.Println(version.GetVersion())
}
