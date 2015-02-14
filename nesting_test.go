package envcfg

import (
	"os"
	"testing"
)

type nestedStruct struct {
	SubStruct struct {
		Field string
	}
}

func TestNestedMapping(t *testing.T) {
	const ENV_KEY = "SUBSTRUCT_FIELD"
	const ENV_VAL = "Remember: testing is the future!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := nestedStruct{}

	ReadInto(&s)

	if s.SubStruct.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.SubStruct.Field)
	}
}

type deeplyNestedStruct struct {
	SubStruct struct {
		SubSubStruct struct {
			Field string
		}
	}
}

func TestDeeplyNestedMapping(t *testing.T) {
	const ENV_KEY = "SUBSTRUCT_SUBSUBSTRUCT_FIELD"
	const ENV_VAL = "Remember: testing is the future!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := deeplyNestedStruct{}

	ReadInto(&s)

	if s.SubStruct.SubSubStruct.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.SubStruct.SubSubStruct.Field)
	}
}
