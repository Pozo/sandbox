package main

import (
	"fmt"
	"reflect"
	"strings"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"os/exec"
)
type BuildSteps struct {
	Before_install []string
	Install []string
	Before_script []string
	Script []string
	After_success []string
	After_failure []string
	After_script []string
	Before_deploy []string
	Deploy []string
	After_deploy []string

}

func main() {
	filename, _ := filepath.Abs("./build.yml")
    yamlFile, err := ioutil.ReadFile(filename)
	
	fmt.Printf("Value: %#v\n", string(yamlFile))
	
	if err != nil {
	    panic(err)
	}
	var (
		buildSteps BuildSteps
		cmdOut []byte
	)
	
	err = yaml.Unmarshal(yamlFile, &buildSteps)
	if err != nil {
	  panic(err)
	}
	fmt.Printf("Value: %s\n", buildSteps)
	
	v := reflect.ValueOf(buildSteps)
	values := make([]reflect.Value, v.NumField())
	
	for i := 0; i < v.NumField(); i++ {
		length := v.Field(i).Len()
        values[i] = v.Field(i).Slice(0,length)
		
		for j := 0; j < values[i].Len(); j++ {
			command := values[i].Index(j).String()
			fmt.Printf("command :%v\n", command)
			
			args := strings.Split(command," ");

		    if cmdOut,err = exec.Command(args[0],).Output(); err != nil { 
		        fmt.Println("Error: ", err)
		    }
			fmt.Printf(string(cmdOut))
		}
    }
}
