package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Participant struct {
	ID                string `json:"id"`
	StudentCode       string `json:"student_code"`
	StudentName       string `json:"student_name"`
	StudentAddress    string `json:"student_address"`
	StudentOccupation string `json:"student_occupation"`
	JoiningReason     string `json:"joining_reason"`
}

type ParticipantList struct {
	Participants []Participant `json:"participants"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <student_code>")
		return
	}

	fileData, err := ioutil.ReadFile("participants.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var participants ParticipantList
	err = json.Unmarshal(fileData, &participants)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	studentCode := os.Args[1]
	found := false

	for _, participant := range participants.Participants {
		if participant.StudentCode == studentCode {
			found = true
			fmt.Printf("Name: %s\n", participant.StudentName)
			fmt.Printf("Address: %s\n", participant.StudentAddress)
			fmt.Printf("Occupation: %s\n", participant.StudentOccupation)
			fmt.Printf("Joining Reason: %s\n", participant.JoiningReason)
			break
		}
	}

	if !found {
		fmt.Println("Student not found.")
	}
}
