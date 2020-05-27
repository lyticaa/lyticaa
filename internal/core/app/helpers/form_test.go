package helpers

import (
	"testing"

	"github.com/rs/zerolog"
)

type FormTestStruct struct {
	Data string `validate:"required,min=5"`
}

func TestValidateInput(t *testing.T) {
	var log zerolog.Logger

	ok, err := ValidateInput(FormTestStruct{Data: ""}, &log)
	if ok {
		t.Error()
	}

	if err == nil {
		t.Error()
	}

	ok, err = ValidateInput(FormTestStruct{Data: "Some Data"}, &log)
	if !ok {
		t.Error()
	}

	if err != nil {
		t.Error()
	}
}
