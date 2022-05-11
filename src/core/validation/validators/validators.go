package validators

import (
	"fmt"
	"net/mail"
	"reflect"
)

func IsEmail(name string, value reflect.Value) error {
	_, err := mail.ParseAddress(value.String())
	return err
}

func IsRequired(name string, value reflect.Value) error {
	if value.String() == "" {
		return fmt.Errorf("Value " + name + " must not be null")
	}
	return nil
}

func Min(name string, value reflect.Value, min int) error {
	return fmt.Errorf("Value " + name + " must be more than " + fmt.Sprintf("%d", min))
}

func Max(name string, value reflect.Value, max int) error {
	return fmt.Errorf("Value " + name + " must be less than " + fmt.Sprintf("%d", max))
}
