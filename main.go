package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {

}

func cliSetup() {
	viper.SetDefault("port", 80)
	viper.SetDefault("status", 200)

	flag.Int("port", 80, "set the port to listen to")
	flag.Int("status", 200, "set the status to response with")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func main() {
	cliSetup()

	setupServer()
}

type PrintServer struct{}

func setupServer() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", viper.GetInt("port")),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      &PrintServer{},
	}

	server.ListenAndServe()
}

func (p *PrintServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("[%s] %s", r.Method, r.RequestURI)

	fmt.Print(string(data))

	fmt.Println()

	w.WriteHeader(viper.GetInt("status"))
}
