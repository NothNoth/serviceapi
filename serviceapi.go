package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
)

type Config struct {
	HTTPPort        int
	AllowedServices []string
}

var config Config

func startService(w http.ResponseWriter, r *http.Request) {
	service := path.Base(path.Dir(r.URL.Path))

	//Make sure service is really allowed, just in case
	for _, s := range config.AllowedServices {
		if s == service {
			fmt.Println("Start service \"" + service + "\"")
			cmd := exec.Command("/usr/sbin/service", service, "start")
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to start service: " + err.Error())
				http.Error(w, "Internal server error: "+err.Error(), 500)
				return
			}
			return
		}
	}

	fmt.Println("Trying to start un-allowed service " + service)
}

func stopService(w http.ResponseWriter, r *http.Request) {
	service := path.Base(path.Dir(r.URL.Path))

	//Make sure service is really allowed, just in case
	for _, s := range config.AllowedServices {
		if s == service {
			fmt.Println("Stop service \"" + service + "\"")
			cmd := exec.Command("/usr/sbin/service", service, "stop")
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed  to stop service: " + err.Error())
				http.Error(w, "Internal server error: "+err.Error(), 500)
				return
			}

			return
		}
	}

	fmt.Println("Trying to stop un-allowed service " + service)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <config file>\n", os.Args[0])
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read " + os.Args[1])
		return
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Failed to parse " + os.Args[1])
		return
	}

	fmt.Printf("Listening on port: %d\n", config.HTTPPort)
	fmt.Println("Allowed services:")
	for _, s := range config.AllowedServices {
		fmt.Println("  - " + s)
	}

	for _, s := range config.AllowedServices {
		http.HandleFunc(fmt.Sprintf("/%s/start", s), startService)
		http.HandleFunc(fmt.Sprintf("/%s/stop", s), stopService)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", config.HTTPPort), nil)

}
