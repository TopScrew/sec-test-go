package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ttt "io/ioutil"
	sss "net/http"
)

func T1(c *gin.Context) {
	url := c.Query("url")
	res, _ := sss.Get(url)
	defer res.Body.Close()
	body, _ := ttt.ReadAll(res.Body)
	fmt.Print(body)
}

type HttpClient struct {
}

func (*HttpClient) Get(url string) {
	res, _ := sss.Get(url)
	defer res.Body.Close()
	body, _ := ttt.ReadAll(res.Body)
	fmt.Print(body)
}

func T2(c *gin.Context) {
	url := c.Query("url")
	client := HttpClient{}
	client.Get(url)
}