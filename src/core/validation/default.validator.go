package validation

import (
	"api/src/core/validation/validators"
	"reflect"
	"strconv"
	"strings"
)

type PropArray = []FieldValidation

type ValidationAttributes struct {
	Required bool
	Type     string
	Min      int
	Max      int
}

type FieldValidation struct {
	Name  string
	Attrs ValidationAttributes
	Value interface{}
}

func DefaultValidator[T comparable](body T) (bool, []string) {
	return validate[T](prepare(body))
}

func splitProperties(tag string, name string, value interface{}) FieldValidation {
	rawAttrs := strings.Split(tag, " ")

	for i := 0; i < len(rawAttrs); i++ {
		splitted := strings.Split(rawAttrs[i], ":")
		if splitted[0] == "validate" {
			fv := FieldValidation{
				Name:  name,
				Value: value,
				Attrs: validationAttributesFactory(strings.Split(splitted[1], ",")),
			}
			return fv
		}
	}
	return FieldValidation{}
}

func validationAttributesFactory(tag []string) ValidationAttributes {
	result := ValidationAttributes{}

	for _, field := range tag {
		clean := strings.Split(strings.Trim(field, "\""), "=")
		field = clean[0]

		switch field {
		case "required":
			result.Required = true
			break
		case "type":
			result.Type = clean[1]
			break
		case "min":
			min, err := strconv.Atoi(clean[1])
			if err != nil {
				return result
			}
			result.Min = int(min)
		case "max":
			max, err := strconv.Atoi(clean[1])
			if err != nil {
				return result
			}
			result.Max = int(max)
		}
	}
	return result
}

func prepare[T comparable](body T) PropArray {
	numField := reflect.TypeOf(body).NumField()
	var validationItems PropArray

	for i := 0; i < numField; i++ {
		field := reflect.TypeOf(body).Field(i)
		tag := string(field.Tag)
		name := string(field.Name)
		value := reflect.Indirect(reflect.ValueOf(body)).FieldByName(name)
		validationItems = append(validationItems, splitProperties(tag, name, value))
	}

	return validationItems
}

func validate[T comparable](items PropArray) (bool, []string) {
	var errs []string

	for i := 0; i < len(items); i++ {
		attrs := items[i].Attrs

		if attrs.Required {
			if err := validators.IsRequired(items[i].Name, items[i].Value.(reflect.Value)); err != nil {
				errs = append(errs, err.Error())
			}
		}

		if attrs.Type != "" {
			if attrs.Type == "email" {
				if err := validators.IsEmail(items[i].Name, items[i].Value.(reflect.Value)); err != nil {
					errs = append(errs, err.Error())
				}
			}
		}

		if attrs.Min > 0 {
			if err := validators.Min(items[i].Name, items[i].Value.(reflect.Value), items[i].Attrs.Min); err != nil {
				errs = append(errs, err.Error())
			}
		}

		if attrs.Max > 0 {
			if err := validators.Max(items[i].Name, items[i].Value.(reflect.Value), items[i].Attrs.Max); err != nil {
				errs = append(errs, err.Error())
			}
		}
	}

	if len(errs) > 0 {
		return false, errs
	}
	return true, nil
}
