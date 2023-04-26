package main

import "fmt"

type errCustom struct {
	err  error
	path string
}

func (ec errCustom) Error() string {
	return fmt.Sprintf("%s: %s", ec.path, ec.err)
}

func XYZ(a int) error { // returning *errFoo would not return nil even though errCustom is a erorr type as it implements Error() method. because  *errCustom is a typed , interface is not nill anymore.
	return nil
}

func main() {
	var err error = XYZ(1)

	if err != nil {
		fmt.Println("OPS")
	} else {
		fmt.Println("OK")
	}
}
