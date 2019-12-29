package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type loggerData struct {
	Type     string      `json:"type"`
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
	Header   http.Header `json:"header"`
	Time     time.Time   `json:"time"`
	DurMs    int64       `json:"dur_ms"`
	RemoteIp string      `json:"remote_ip"`
	Uri      string      `json:"uri"`
	Err      error       `json:"err"`
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		start := time.Now()
		reqb, err := ioutil.ReadAll(c.Request().Body)
		req := string(reqb)
		if err != nil {
			fmt.Println("logger ioutil readall err", err, "req", req)
			return
		}
		dataJson := &loggerData{
			Type:     "http_req",
			Request:  req,
			Header:   c.Request().Header,
			Time:     start,
			RemoteIp: c.RealIP(),
			Uri:      c.Request().RequestURI,
		}
		reqData, _ := json.Marshal(dataJson)
		fmt.Println(string(reqData))
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqb))
		defer func() {
			dataJson.Type = "http_resp"
			dataJson.Response = GetOutResponse(c)
			dataJson.Time = time.Now()
			dataJson.Header = c.Response().Header()
			dataJson.DurMs = int64(time.Since(start) / time.Millisecond)
			dataJson.Err = err
			respData, _ := json.Marshal(dataJson)
			fmt.Println(string(respData))
		}()
		err = next(c)
		return
	}
}
