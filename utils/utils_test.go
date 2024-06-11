package utils

import (
	"reflect"
	"testing"
)

func TestValToRef(t *testing.T) {
	test := ValToRef[string]("test")

	if *test != "test" {
		t.Error("Unexpected type returned from ValToRef")
	}

	if reflect.ValueOf(test).Kind() != reflect.Ptr {
		t.Error("Unexpected type returned from ValToRef")
	}
}

func TestValToRefSlice(t *testing.T) {
	expectedLen := 2
	value := []string{"test", "test2"}

	newVal := ValToRefSlice[string](value)

	if len(newVal) != expectedLen {
		t.Error("Unexpected type returned from RefToVal")
	}

}

func TestRefToVal(t *testing.T) {
	value := "test"
	test := &value

	newVal := RefToVal[string](test)

	if newVal != "test" {
		t.Error("Unexpected type returned from RefToVal")
	}

	if reflect.ValueOf(test).Kind() != reflect.Ptr {
		t.Error("Unexpected type returned from RefToVal")
	}

	if reflect.ValueOf(newVal).Kind() != reflect.String {
		t.Error("Unexpected type returned from RefToVal")
	}
}

func TestRefToValSlice(t *testing.T) {
	expectedLen := 0
	value := []*string{}

	newVal := RefToValSlice[string](value)

	if len(newVal) != expectedLen {
		t.Error("Unexpected type returned from RefToVal")
	}
}
