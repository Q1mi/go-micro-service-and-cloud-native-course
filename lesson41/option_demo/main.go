package main

import (
	"fmt"
	"option_demo/basicConfig"
	"option_demo/multiArgsConfig"
	"option_demo/optionConfig"
	"option_demo/advancedOptionConfig"
)

func main() {

	basicConfig := basicConfig.NewServiceConfig("qimi", "西二旗", 10)
	fmt.Printf("basicConfig:%#v\n", basicConfig)
	// basicConfig.c = 100 // cannot modify directly

	multiArgsConfig := multiArgsConfig.NewServiceConfig("qimi", "西二旗", 10)
	fmt.Printf("multiArgsConfig:%#v\n", multiArgsConfig)
	// multiArgsConfig.c = 100 // cannot modify directly

	info := optionConfig.Info{Addr: "127.0.0.1:8080"}
	optionConfig := optionConfig.NewServiceConfig("qimi", "西二旗", optionConfig.WithC(10), optionConfig.WithX(info))
	fmt.Printf("optionConfig:%#v\n", optionConfig)
	// optionConfig.c = 100 // cannot modify directly

	advancedOptionConfig := advancedOptionConfig.NewConfig(10, advancedOptionConfig.WithConfigName("qimi"))
	fmt.Printf("advancedOptionConfig:%#v\n", advancedOptionConfig)
}
