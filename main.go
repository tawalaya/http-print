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

var rawString string
func cliSetup() {
	viper.SetDefault("port", 80)
	viper.SetDefault("status", 200)

	flag.Int("port", 80, "set the port to listen to")
	flag.Int("status", 200, "set the status to response with")
	flag.String("response", "", "file to send on each response")
	flag.String("raw","","instead of file you can also specify a raw string as a response")
	flag.String("type", "application/json", "Content Type of each response")
	flag.Bool("verbose", true, "if set the server will print all body and header information")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func main() {
	cliSetup()

	setupServer()
}

type PrintServer struct {
	response []byte
}

func setupServer() {

	printer := &PrintServer{}
	if viper.GetString("response") != "" {
		data, err := ioutil.ReadFile(viper.GetString("response"))
		if err != nil {
			fmt.Printf("failed to read response file - using none : %+v\n", err)
		} else {
			printer.response = data
		}
	} else if response:=viper.GetString("raw") ; response != ""{
		printer.response = []byte(response)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", viper.GetInt("port")),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      printer,
	}

	server.ListenAndServe()
}

func (p *PrintServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("[%s] %s\n", r.Method, r.RequestURI)

	if viper.GetBool("verbose") {

		fmt.Printf("Header:%+v\n", r.Header)
		fmt.Printf("Agent:%+v\n", r.UserAgent())
		fmt.Printf("Query:%+v\n", r.URL.Query())

		fmt.Println(string(data))
	}

	w.Header().Add("Content-Type", viper.GetString("type"))
	if p.response != nil {
		w.Write(p.response)
	} else {
		w.WriteHeader(viper.GetInt("status"))
	}

}
