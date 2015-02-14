package envcfg

import (
	"fmt"
	"os"
	"testing"
)

func TestStringMapping(t *testing.T) {
	const ENV_KEY = "FIELD"
	const ENV_VAL = "Remember: testing is the future!"

	os.Clearenv()
	os.Setenv(ENV_KEY, ENV_VAL)

	s := struct{ Field string }{}

	ReadInto(&s)

	if s.Field != ENV_VAL {
		t.Errorf("expected '%s', got '%s'", ENV_VAL, s.Field)
	}
}

func TestBoolMapping(t *testing.T) {
	const ENV_KEY = "FIELD"
	const ENV_VAL = true

	os.Clearenv()
	os.Setenv(ENV_KEY, fmt.Sprintf("%t", ENV_VAL))

	s := struct{ Field bool }{}

	ReadInto(&s)

	if s.Field != ENV_VAL {
		t.Errorf("expected '%t', got '%t'", ENV_VAL, s.Field)
	}
}

func TestIntMapping(t *testing.T) {
	const ENV_KEY = "FIELD"
	const ENV_VAL = 42

	os.Clearenv()
	os.Setenv(ENV_KEY, fmt.Sprintf("%d", ENV_VAL))

	s := struct{ Field int }{}

	ReadInto(&s)

	if s.Field != ENV_VAL {
		t.Errorf("expected '%d', got '%d'", ENV_VAL, s.Field)
	}
}

func TestFloatMapping(t *testing.T) {
	const ENV_KEY = "FIELD"
	const ENV_VAL = 13.37

	os.Clearenv()
	os.Setenv(ENV_KEY, fmt.Sprintf("%f", ENV_VAL))

	s := struct{ Field float32 }{}

	ReadInto(&s)

	if s.Field != ENV_VAL {
		t.Errorf("expected '%f', got '%f'", ENV_VAL, s.Field)
	}
}
