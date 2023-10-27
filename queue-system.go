package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strings"
)

type Patient struct {
  MRNumber string
  Gender   string
}

var patientQueue []Patient
var roundRobinMode bool
var roundRobinGender string // Initialize with Female

func isMRNumberUnique(mrNumber string) bool {
  for _, patient := range patientQueue {
    if patient.MRNumber == mrNumber {
      return false
    }
  }
  return true
}

func printQueue() {
  fmt.Println("Queue:")
  for _, patient := range patientQueue {
    fmt.Printf("%s %s\n", patient.MRNumber, patient.Gender)
  }
}

func processCommand(command string) {
  args := strings.Fields(command)

  if args[0] == "IN" {
    if len(args) != 3 {
      fmt.Println("error: Missing arguments, please input for example IN MR1234 M")
      return
    }

    mrNumber := args[1]
    gender := args[2]

    if gender != "M" && gender != "F" {
      fmt.Println("error: please input gender as M or F")
      return
    }

    mrNumberPattern := regexp.MustCompile(`^MR\d{4}$`)
    if !mrNumberPattern.MatchString(mrNumber) {
      fmt.Println("error: Invalid MRNumber format")
      return
    }

    if !isMRNumberUnique(mrNumber) {
      fmt.Printf("error: Patient with %s already in queue\n", mrNumber)
      return
    }

    patientQueue = append(patientQueue, Patient{MRNumber: mrNumber, Gender: gender})
    fmt.Printf("Added %s %s to the queue.\n", mrNumber, gender)
  } else if args[0] == "OUT" {
    if len(patientQueue) == 0 {
      fmt.Println("Queue is empty.")
      return
    }

    var patient Patient

    if roundRobinMode {
      for _, p := range patientQueue {
        if p.Gender == roundRobinGender {
          patient = p
          break
        }
      }
    } else {
      patient = patientQueue[0]
      patientQueue = patientQueue[1:]
    }

    // Remove patient from the queue
    var newQueue []Patient
    for _, p := range patientQueue {
      if p.MRNumber != patient.MRNumber {
        newQueue = append(newQueue, p)
      }
    }
    patientQueue = newQueue

    if patient.MRNumber == "" {
      roundRobinMode = false
      fmt.Println("No eligible patient found for round-robin. Switching to default.")
      return
    }

    fmt.Printf("send: %s %s > OUT\n", patient.MRNumber, patient.Gender)
  } else if args[0] == "ROUNDROBIN" {
    roundRobinMode = true
    fmt.Println("Switched to round-robin mode.")
  } else if args[0] == "DEFAULT" {
    roundRobinMode = false
    fmt.Println("Switched to default mode.")
  } else if args[0] == "EXIT" {
    os.Exit(0)
  } else {
    fmt.Println("error: Invalid command")
  }
}

func main() {
  fmt.Println("Welcome to the Hospital Patient Queue System!")
  fmt.Println("Commands: IN, OUT, ROUNDROBIN, DEFAULT, EXIT")

	reader := bufio.NewReader(os.Stdin)

	for {
    fmt.Print("> ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    processCommand(input)
    printQueue()
  }
}