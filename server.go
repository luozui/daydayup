package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var logdir, uidsdir *string

func usersData() map[string]string {
	users := make(map[string]string)
	file, _ := os.Open(*uidsdir)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		if len(arr) == 2 {
			users[arr[0]] = arr[1]
		}
	}
	return users
}

func logData() string {
	s, _ := os.ReadFile(*logdir)
	return string(s)
}

func index(c *gin.Context) {
	users := usersData()
	var buf string
	for id := range users {
		buf += fmt.Sprintln(id)
	}
	log := logData()
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"users": buf,
		"log":   log,
	})
}

func main() {
	apiaddr := pflag.StringP("apiaddr", "a", ":8080", "Admin api address")
	logdir = pflag.StringP("logdir", "l", "./out.log", "logdir")
	uidsdir = pflag.StringP("uidsdir", "u", "./gzhu_no_clock_in/stu_id.txt", "uidsdir")
	pflag.Parse()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.POST("/", func(c *gin.Context) {
		users := usersData()
		users[c.PostForm("stuid")] = c.PostForm("pwd")
		var buf string
		for id, pass := range users {
			buf += fmt.Sprintln(id, pass)
		}
		ioutil.WriteFile(*uidsdir, []byte(buf), 0644)
		index(c)
	})
	router.GET("/", index)
	router.Run(*apiaddr)
}
