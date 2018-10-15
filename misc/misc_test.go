package misc_test

import (
	"reflect"
	"testing"

	"commitr/misc"
)

var (
	toFilter []string = []string{"test1", "", "test3", "", "", "test6", ""}
	filtered []string = []string{"test1", "test3", "test6"}
)

func TestFilter(t *testing.T) {
	if !reflect.DeepEqual(misc.Filter(toFilter, misc.RemoveEmpty), filtered) {
		t.Error("Filtered array doesn't match expected output from Filter.")
	}
}

func TestRemoveEmpty(t *testing.T) {
	if misc.RemoveEmpty("") != false {
		t.Error("RemoveEmpty doesn't return false on empty strings")
	}
}
