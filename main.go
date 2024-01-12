package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Yank struct {
	Command string `json:"command"`
}

func main() {
	cmd := os.Args[1]
	cmdId := os.Args[2]
	cmdCommand := os.Args[3]

	switch cmd {
	case "new":
		yanks := getYanks()

		yanks[cmdId] = Yank{Command: cmdCommand}

		fmt.Println(yanks)

		yankBytes, err := json.MarshalIndent(yanks, "", "  ")
		if err != nil {
			panic(err)
		}

		err = writeToFile("yanks.json", yankBytes)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Yank added successfully.")
		break
	}
}

func getYanks() map[string]Yank {
	_, err := os.Stat("yanks.json")

	if err == nil {
		existingYanks, err := os.ReadFile("yanks.json")

		// Check if the file is empty
		if len(existingYanks) == 0 {
			return make(map[string]Yank)
		}

		if err != nil {
			panic(err)
		}

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
		return make(map[string]Yank)
	} else {
		fmt.Println(err)
	}
	return nil
}

func writeToFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

