package main

import (
	"fmt"
)

func main() {
	target := 500
	for n := 2.0; ; n++ {
		tri := int((n + 1.0) / 2.0 * n)
		numFactors := 2
		for i := 2; i*i < tri; i++ {
			if tri%i == 0 {
				numFactors += 2
			}
		}
		if numFactors > target {
			fmt.Printf("%d: %d has %d factors\n", int(n), tri, numFactors)
			break
		}
	}
}
