package eddnc

import "fmt"

func ExampleScmInfos() {
	fmt.Println(ScmInfos[0].Ref)
	fmt.Println(ScmInfos[0].Topic)
	fmt.Println(ScmInfos[0].Version)
	// Output:
	// https://eddn.edcd.io/schemas/approachsettlement/1
	// https://eddn.edcd.io/schemas/approachsettlement/
	// 1
}
