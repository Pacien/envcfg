package envcfg

import (
	"os"
	"testing"
)

type taggedStruct struct {
	Field string `env:"CUSTOM_POTATOE"`
}

func TestTaggedField(t *testing.T) {
	const ENV_KEY = "CUSTOM_POTATOE"
	const ENV_VAL = "Remember: testing is the future!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := taggedStruct{}

	ReadInto(&s)

	if s.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.Field)
	}
}

type taggedSubStruct struct {
	SubStruct struct {
		Field string
	} `env:"POTATOE"`
}

func TestTaggedSubStruct(t *testing.T) {
	const ENV_KEY = "POTATOE_FIELD"
	const ENV_VAL = "Remember: testing is the future!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := taggedSubStruct{}

	ReadInto(&s)

	if s.SubStruct.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.SubStruct.Field)
	}
}

type absTaggedField struct {
	SubStruct struct {
		Field string `env:"POTATOE" absenv:"true"`
	}
}

func TestAbsTaggedField(t *testing.T) {
	const ENV_KEY = "POTATOE"
	const ENV_VAL = "Remember: testing is the future!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := absTaggedField{}

	ReadInto(&s)

	if s.SubStruct.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.SubStruct.Field)
	}
}
