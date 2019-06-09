package main

import(
	"github.com/spf13/viper"
	"log"
)

func main(){
	var(
		err error
	)
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/code/src/github.com/shaban/polaris")
	if err = viper.ReadInConfig() ;err!=nil{
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	println("Hello World")
}