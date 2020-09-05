package a

import "errors"

func f() (int, error) {
	return 0, nil
}

var (
	_, err1 = f() // OK
)

func main() {
	_, err := f() // OK
	if err != nil {
		panic(err)
	}

	if _, err := f(); err != nil { // OK
		panic(err)
	}

	if _, err := f(); err == nil { // OK
		panic(err)
	}

	_, err = f() // want "error received but not handled"

	err = func() error {
		return nil
	}() // want "error received but not handled"

	if a, err := f(); a != 0 { // want "error received but not handled"
		panic(err)
	}

	if a, err := f(); a != 0 && err != nil { // OK
		panic(err)
	}

	_, err = f() // want "error received but not handled"

	b := errors.New("foo") // want "error received but not handled"

	print(b)
}
