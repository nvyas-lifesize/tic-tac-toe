package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Structure to store server response
type responseBody struct {
	Grid     [3][3]int `json:"grid"`
	Result   string    `json:"result"`
	Computer int       `json:"computer"`
}

// Structure to store server request
type requestBody struct {
	Grid [3][3]int `json:"grid"`
	Move int       `json:"move"`
}

// Converts the player marker int to marker character X or O
func getMarker(i int) string {
	if i == 1 {
		return "X"
	} else if i == 0 {
		return "O"
	}
	return "."
}

// Print the 3X3 grid
func PrintGrid(grid [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(getMarker(grid[i][j]))
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}

// Checking valid move
func IsValid(grid [3][3]int, move int) bool {
	move--
	return grid[move/3][move%3] == -1
}

func main() {
	const port = "9999"
	const url = "http://localhost"
	resp, err1 := http.Get(url + ":" + port + "/start")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	defer resp.Body.Close()
	var response responseBody
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	var data requestBody
	data.Grid = response.Grid

	PrintGrid(response.Grid)
	for true {
		fmt.Print("Enter your choice from 1 to 9: ")
		var input int
		fmt.Scanln(&input)
		if input < 1 || input > 9 {
			fmt.Println("Enter the number from 1 to 9 range only!!")
			continue
		}
		if !IsValid(data.Grid, input) {
			fmt.Println("Position already filled, choose the position which is available!!")
			continue
		}
		fmt.Println("Client move to:", input)
		data.Move = input
		requestByte, _ := json.Marshal(data)

		Moveresp, error := http.Post(url+":"+port+"/move", "application/json", bytes.NewReader(requestByte))

		if error != nil {
			print("Error:", error)
		}
		defer Moveresp.Body.Close()
		body, err2 = ioutil.ReadAll(Moveresp.Body)
		var resp responseBody
		if err := json.Unmarshal(body, &resp); err != nil {
			panic(err)
		}
		fmt.Println("Computer move to:", resp.Computer)
		PrintGrid(resp.Grid)
		if resp.Result != "" {
			fmt.Println(resp.Result)
			break
		}
		data.Grid = resp.Grid
	}
}
