package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := " http://elastic-chart.inapps.k8s.dsp.local"
	nu, err := url.Parse(u)
	fmt.Println(nu, err)
}
