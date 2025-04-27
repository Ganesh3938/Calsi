package main

import "errors"

func Division(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	} else {
		return a / b, nil
	}
}
