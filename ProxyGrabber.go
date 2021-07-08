package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	var Choice int
	color.Cyan(" Programmed By : https://instagram.com/6o9s \n\n")
	fmt.Print(" 1 For Http/s, 2 For Soucks4, 3 For Soucks5 : ")
	fmt.Scanln(&Choice)
	i := 1
	if Choice == 1 {
		Url := "https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all"
		for {
			req, err := http.Get(Url)
			if err != nil {
				panic(err)
			}
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic(err)
			}
			file, err := os.OpenFile("Proxies.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
			}
			file.WriteString(string(body))
			fmt.Print(" Proxies Downloaded : ", i, " Times\n\n")
			i++
			time.Sleep(time.Second * 60)
		}
	}
	if Choice == 2 {
		Url := "https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4&timeout=10000&country=all"
		for {
			req, err := http.Get(Url)
			if err != nil {
				panic(err)
			}
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic(err)
			}
			file, err := os.OpenFile("Proxies.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
			}
			file.WriteString(string(body))
			fmt.Print(" Proxies Downloaded : ", i, " Times\n\n")
			i++
			time.Sleep(time.Second * 60)
		}
	}
	if Choice == 3 {
		for {
			Url := "https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5&timeout=10000&country=all"
			req, err := http.Get(Url)
			if err != nil {
				panic(err)
			}
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic(err)
			}
			file, err := os.OpenFile("Proxies.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
			}
			file.WriteString(string(body))
			fmt.Print(" Proxies Downloaded : ", i, " Times\n\n")
			i++
			time.Sleep(time.Second * 60)
		}
	} else {
		color.Red("\n Wrong Choice")
	}
}
