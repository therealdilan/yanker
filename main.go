package main 

import (
  "fmt"
	"encoding/json"
  "os"
)

type newYank struct {
  Command string
  Id string
}

func main() {
  var n = newYank{}

  fmt.Scanln(&n.Command)
  fmt.Scanln(&n.Id)

  fmt.Println(n)

  content, err := json.Marshal(n)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(string(content))

  err = os.WriteFile("yank.json",content,0644)
  if err != nil {
    fmt.Println(err)
    return
  }
} 
