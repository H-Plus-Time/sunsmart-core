package main

import (
    "fmt"
    "net/http"
    "encoding/xml"
    "io/ioutil"
    "reflect"
    // "time"
    // "github.com/robfig/cron"
)

const (
     serverKey = "YOUR-KEY"
)

type Location struct {
  Name string `xml:"id,attr"`
  ShortName string `xml:"name"`
  Index float64 `xml:"index"`
  Time string `xml:"time"`
  Date string `xml:"date"`
  Fulldate string `xml:"fulldate"`
  Utcdatetime string `xml:"utcdatetime"`
  Status string `xml:"status"`
}

type Arpansa_Query struct {
  LocationList []Location `xml:"location"`
}
// construct funcs for each interval

func publishUVIndex(scheduled_time string, threshold float64) {
    // data := map[string]string{
    //     "msg": fmt.Sprintf("Max UV: %f", max_uv),
    //     "sum": "Sunsmart UV Warning",
    // }
    //
    // c := fcm.NewFcmClient(serverKey)
    // c.NewFcmMsgTo(topic, data)
    //
    //
    // status, err := c.Send()
    //
    //
    // if err == nil {
    // status.PrintResults()
    // } else {
    //     fmt.Println(err)
    // }
}


func main() {
  resp, err := http.Get("http://www.arpansa.gov.au/uvindex/realtime/xml/uvvalues.xml")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()

  var q Arpansa_Query;
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err.Error())
  }

  fmt.Println(reflect.TypeOf(body))
  fmt.Println(string(body))
	xml.Unmarshal(body, &q)
  fmt.Println(q)
}
