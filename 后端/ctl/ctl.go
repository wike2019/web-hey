package ctl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"hey-back/requester"
	"io/ioutil"
	"math"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Ctl struct {
}
func NewCtl() *Ctl{
	return &Ctl{}
}
func(*Ctl)  Name() string{
	return "Ctl"
}
func(*Ctl) getProxy() [][]string {
	data:= make([][]string,0)
	return  data
}

func(this *Ctl)  hey(c *gin.Context) string{

	IpList:=this.getProxy()

	var buf=&bytes.Buffer{}
	sumbmitType:= c.DefaultQuery("type","json")
	bodyAll, _ := ioutil.ReadAll(c.Request.Body)
	var num float64
	var conc float64
	var q float64
	var dur time.Duration
	var t float64
	var url string
	var method string
	var req *http.Request
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyAll))
	if sumbmitType=="file"{
			hey:= c.PostForm("__hey")
			heyData:=make(map[string]interface{})
			json.Unmarshal([]byte(hey),&heyData)
			urlRaw:= c.PostForm("__urlData")
			urlData:=make(map[string]interface{})
			json.Unmarshal([]byte(urlRaw),&urlData)

			num=heyData["n"].(float64)


			conc=heyData["c"].(float64)

			q=heyData["q"].(float64)
			dur, _ = time.ParseDuration(heyData["z"].(string))

			t=heyData["t"].(float64)
			url=  urlData["url"].(string)
			method = strings.ToUpper( urlData["method"].(string))

			header := make(http.Header)

			headerRaw:= c.PostForm("__header")
			headerData:=make(map[string]string)
			json.Unmarshal([]byte(headerRaw),&headerData)
			for k,v:=range  headerData{
				fmt.Println(k,v)
				header.Set(k, v)
			}
			if dur > 0 {
				num = math.MaxInt32
			}
			header.Set("Content-Type", c.Request.Header.Get("Content-Type"))
			req, _ = http.NewRequest(method, url, nil)
			req.Header = header
			req.ContentLength = int64(len(bodyAll))

	}else{
		dataRaw:=make(map[string]interface{})
		json.Unmarshal(bodyAll,&dataRaw)


		hey:=dataRaw["__hey"]
		heyData:=hey.(map[string]interface{})

		urlRaw:= dataRaw["__urlData"]
		urlData:=urlRaw.(map[string]interface{})
		num=heyData["n"].(float64)
		conc=heyData["c"].(float64)
		q=heyData["q"].(float64)

		dur, _ = time.ParseDuration(heyData["z"].(string))
		t=heyData["t"].(float64)
		url=  urlData["url"].(string)
		method = strings.ToUpper( urlData["method"].(string))
		header := make(http.Header)
		headerRaw:= dataRaw["__header"]
		headerData:=headerRaw.(map[string]interface{})

		for k,v:=range  headerData{
			//fmt.Println(k,v)
			header.Set(k, v.(string))
		}
		if dur > 0 {
			num = math.MaxInt32
		}
		header.Set("Content-Type", c.Request.Header.Get("Content-Type"))
		req, _ = http.NewRequest(method, url, nil)
		req.Header = header
		bodyAll,_=json.Marshal(dataRaw["__data"])
		req.ContentLength = int64(len(bodyAll))
	}
	fmt.Println(len(IpList))
	w := &requester.Work{
		Request:            req,
		//N:       20,
		//C:       2,
		RequestBody:        bodyAll,
		N:                  int(num),
		C:                  int(conc),
		QPS:                q,
		Timeout:            int(t),
		Writer: buf,
		IpList:IpList,
		I:len(IpList),
	}

	w.Init()
	if dur > 0 {
		go func() {
			time.Sleep(dur)
			w.Stop()
		}()
	}
	w.Run()
	return  buf.String()
}


func(this *Ctl)  Build(goft *goft.Goft){
	goft.Handle("POST","/hey",this.hey)
	
}

