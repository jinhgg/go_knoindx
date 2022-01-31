package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "knoindx_backend/conf"
)

func main() {
	fmt.Println(viper.GetString("version"))
}
