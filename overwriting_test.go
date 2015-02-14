package envcfg

import (
	"os"
	"testing"
)

func TestKeeping(t *testing.T) {
	const ORIG_VAL = "Remember: testing is the future!"

	os.Clearenv()

	s := struct{ Field string }{ORIG_VAL}

	ReadInto(&s)

	if s.Field != ORIG_VAL {
		t.Errorf("expected '%s', got '%s'", ORIG_VAL, s.Field)
	}
}

func TestOverwriting(t *testing.T) {
	const ENV_KEY = "FIELD"
	const ENV_VAL = "Remember: testing is the future!"
	const ORIG_VAL = "Testing is pointless!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := struct{ Field string }{ORIG_VAL}

	ReadInto(&s)

	if s.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.Field)
	}
}

type superStruct struct {
	SubStruct nestedFields
}

type nestedFields struct {
	KeepMe      string
	OverwriteMe string
}

func TestMultiOverwriting(t *testing.T) {
	const ENV_KEY = "SUBSTRUCT_OVERWRITEME"
	const ENV_VAL = "Remember: testing is the future!"
	const ORIG_VAL = "Testing is pointless!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := superStruct{nestedFields{ORIG_VAL, ORIG_VAL}}

	ReadInto(&s)

	if s.SubStruct.KeepMe != ORIG_VAL {
		t.Errorf("expected '%s', got '%s'", ORIG_VAL, s.SubStruct.KeepMe)
	}

	if s.SubStruct.OverwriteMe != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.SubStruct.OverwriteMe)
	}
}
