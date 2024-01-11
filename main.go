package main 

import (
  "fmt"
	"encoding/json"
  "os"
)

type Yank struct {
  Command string `json:"command"`
}

func main() {
  cmd := os.Args[1]
  cmdId := os.Args[2]
  cmdCommand := os.Args[3]

  yanks := make(map[string]Yank)

  switch cmd {
    case "new":
      yanks[cmdId] = Yank{Command: cmdCommand}      
      yankBytes, err := json.Marshal(yanks)
      if err != nil {
        panic(err)
      }

      file, err := os.OpenFile("yanks.json", os.O_WRONLY|os.O_TRUNC, 0644)
        if err != nil {
          fmt.Println("Error opening file:", err)
          return
      }
      defer file.Close()

      encoder := json.NewEncoder(file)
      err = encoder.Encode(yankBytes)
      if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
      }
      break;
  }
}

func getYanks() map[string]Yank {
  _, err := os.Stat("yanks.json")

  if err == nil {
    existingYanks, err := os.ReadFile("yanks.json")
    yanks := make(map[string]Yank)

    fmt.Println(existingYanks)

    err = json.Unmarshal(existingYanks, &yanks)

    if err != nil {
      panic(err)
    }

    fmt.Println(yanks)

    return yanks 
  } else if os.IsNotExist(err) {
      os.Create("yanks.json")
      return nil
  } else {
    fmt.Println(err)
  }
  return nil
}
