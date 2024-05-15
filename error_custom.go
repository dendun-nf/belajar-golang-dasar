package main

import "fmt"

func main() {
	err := saveData("")

	if err != nil {
		resolveError(err)
	} else {
		fmt.Println("Sukses")
	}
}

func resolveError(err error) {
	switch errMessage := err.(type) {
	case *validationError:
		fmt.Println("validation error", errMessage)
	case *notFoundError:
		fmt.Println("not found error", errMessage)
	default:
		fmt.Println("unknown error has been occurred")
	}
}

func saveData(id string) error {
	if id == "" {
		return &validationError{Message: "id is required"}
	}

	if id != "deni" {
		return &notFoundError{Message: "data not fonud"}
	}

	return nil
}

type validationError struct {
	Message string
}
type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}
func (n *notFoundError) Error() string {
	return n.Message
}
