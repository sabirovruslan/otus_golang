package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	StrindLenAndIn struct {
		Value string `validate:"len:5|in:admin,stuff"`
	}

	StringRegexp struct {
		Value string `validate:"regexp:\\d+"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	Empty struct{}
)

type validationSuite struct {
	suite.Suite
}

func TestRunValidationSuite(t *testing.T) {
	suite.Run(t, new(validationSuite))
}

func (s *validationSuite) TestValidateTypeError() {
	s.EqualError(Validate(1), "expected a struct, but received int")
	s.EqualError(Validate("test"), "expected a struct, but received string")
}

func (s *validationSuite) TestValidate() {
	s.Equal(Validate(Empty{}), nil)

	errs := Validate(App{"01"})
	s.NotEqual(errs, nil)

	errs = Validate(App{"valid"})
	s.Equal(errs, nil)

	errs = Validate(StrindLenAndIn{"admin"})
	s.Equal(errs, nil)

	errs = Validate(StrindLenAndIn{"valid"})
	s.EqualError(errs, "Value: value: valid is not included in set\n")

	errs = Validate(StrindLenAndIn{"why"})
	s.EqualError(errs, "Value: size less than 5\nValue: value: why is not included in set\n")

	errs = Validate(StringRegexp{"11"})
	s.Equal(errs, nil)
	errs = Validate(StringRegexp{"0"})
	s.Equal(errs, nil)
	errs = Validate(StringRegexp{"test"})
	s.EqualError(errs, "Value: not match: test\n")

}

func (s *validationSuite) TestValidatinErrors() {
	vErrors := make(ValidationErrors, 0)
	err := ValidationError{"test", errors.New("Error test")}
	vErrors = append(vErrors, err)
	s.EqualError(vErrors, "test: Error test\n")
}

//func TestValidate(t *testing.T) {
//	tests := []struct {
//		in          interface{}
//		expectedErr error
//	}{
//		{
//			// Place your code here.
//		},
//		// ...
//		// Place your code here.
//	}
//
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
//			tt := tt
//			t.Parallel()
//
//			// Place your code here.
//			_ = tt
//		})
//	}
//}
