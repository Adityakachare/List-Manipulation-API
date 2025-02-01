package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var intList []int

// Function to manipulate the list
func manipulateList(input int) {
	log.Printf("Received input: %d", input)

	// Ignore 0 input
	if input == 0 {
		log.Println("Received 0. No changes made.")
		return
	}

	// If list is empty, simply append the input
	if len(intList) == 0 {
		intList = append(intList, input)
		log.Printf("List was empty. Added %d, updated list: %v", input, intList)
		return
	}

	// If input has the same sign as the first element, append
	if sameSign(intList[0], input) {
		intList = append(intList, input)
		log.Printf("Same sign detected. Appended %d, updated list: %v", input, intList)
		return
	}

	// Opposite sign: Reduce FIFO
	remaining := abs(input)
	log.Printf("Opposite sign detected. Attempting to update %d from list: %v", remaining, intList)

	newList := []int{}
	for _, val := range intList {
		if remaining == 0 {
			newList = append(newList, val)
			continue
		}

		absVal := abs(val)

		if remaining >= absVal {
			log.Printf("Removing %d completely from list", val)
			remaining -= absVal
		} else {
			log.Printf("Reducing %d from %d", remaining, val)
			newList = append(newList, val-sign(val)*remaining) 
			remaining = 0
		}
	}

	// If there's remaining input after full cancellation, add it back
	if remaining > 0 {
		newList = append(newList, input)
		log.Printf("Remaining %d added back to list", input)
	}

	intList = newList
	log.Printf("Updated list after operation: %v", intList)
}

// Helper function to check if two numbers have the same sign
func sameSign(a, b int) bool {
	return (a >= 0 && b >= 0) || (a < 0 && b < 0)
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Helper function to get sign of a number
func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

// Main function to set up Gin server
func main() {
	route := gin.Default()

	// API to accept number as input
	route.GET("/manipulate", func(c *gin.Context) {
		numStr := c.DefaultQuery("number", "0")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Println("Invalid input received:", numStr)
			c.JSON(400, gin.H{"error": "Invalid input, must be an integer"})
			return
		}

		log.Printf("Processing input: %d", num)
		manipulateList(num)

		c.JSON(200, gin.H{"updatedList": intList})
	})

	log.Println("Server started on port 8080")
	route.Run(":8080")
}
