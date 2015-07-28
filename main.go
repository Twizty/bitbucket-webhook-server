package main

import (
  "fmt"
  "net/http"
  "os"
  "encoding/json"

  "github.com/Twizty/bitbucket-webhook-server/structs"
)

func WebhooksHandler(res http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var push structs.Response
  if err := decoder.Decode(&push); err != nil {
    res.WriteHeader(500)
    fmt.Println(err)
    return
  }
  b, err := json.MarshalIndent(push, "", "  ")
  if err != nil {
    fmt.Println("Error in marshalling")
  }
  fmt.Printf("%s\n", b)
  res.WriteHeader(204)
  res.Write([]byte(""))
}
 
func main() {
  port := os.Args[1]
  http.HandleFunc("/webhooks", WebhooksHandler)
  http.ListenAndServe(":" + port, nil)
}
