package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notfoundError struct {
	Message string
}

func (n *notfoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id is required"}
} 

if id != "daffa" {
	return &notfoundError{Message: "data not found"}
}

return nil
}

func main() {
	 err := SaveData("agus", nil)

	 if err != nil {
			 if validationErr, ok := err.(*validationError); ok {
					 fmt.Println("validation error:", validationErr.Error())
			 } else if notfoundError, ok := err.(*notfoundError); ok {
				fmt.Println("not found error:", notfoundError.Error())
			 } else {
					 fmt.Println("unknown error:", err.Error())
			 }
			 
	 } else {
			//  sukses
			fmt.Println("sukses")

	 }
	 
	 
}