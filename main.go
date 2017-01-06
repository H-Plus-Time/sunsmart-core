package main

import (
    "fmt"
    "net/http"
    "encoding/xml"
    "io/ioutil"
    "reflect"
    "sync"
    // "time"
    // "github.com/robfig/cron"
)

const (
     serverKey = "YOUR-KEY"
     MIN_THRESHOLD=0
     MAX_THRESHOLD=16
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

func publish_worker(time_of_day int) {
  resp, err := http.Get("http://www.arpansa.gov.au/uvindex/realtime/xml/uvvalues.xml")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()

  var stations Arpansa_Query;
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err.Error())
  }

	xml.Unmarshal(body, &stations)
  // do stuff with stations
  for _,element := range stations.LocationList {
    // element is the element from someSlice for where we are
    fmt.Println(element)
    for i := 0; i <= 16; i++ {
        fmt.Println(i)
        if element.Index >= float64(i) {
          publishUVIndex(time_of_day, i, element)
        }
    }
    // get firebase topic subscriptions
  }
}

func publishUVIndex(time_of_day int, threshold int, data Location) {
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

  messages := make(chan int)
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    publish_worker(0)
    messages <- 1
  }()

  go func() {
      for i := range messages {
          fmt.Println(i)
      }
  }()
  wg.Wait()
}
