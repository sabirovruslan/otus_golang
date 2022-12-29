package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/suite"
)

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:32"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   string   `validate:"in:admin,stuff"`
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

	SliceString struct {
		Value []string `validate:"len:3"`
	}

	IntMin struct {
		Value int `validate:"min:10"`
	}

	IntMax struct {
		Value int `validate:"max:10"`
	}

	IntIn struct {
		Value int `validate:"in:10,2,40"`
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

func (s *validationSuite) fakeUser() User {
	pAge, _ := faker.RandomInt(18, 50)
	return User{
		ID:     faker.UUIDDigit(),
		Name:   faker.Name(),
		Age:    pAge[0],
		Email:  faker.Email(),
		Role:   "admin",
		Phones: []string{"724891571063"},
		meta:   []byte{},
	}
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

	errs = Validate(IntMin{10})
	s.Equal(errs, nil)
	errs = Validate(IntMin{1})
	s.EqualError(errs, "Value: value less than: 10\n")

	errs = Validate(IntMax{10})
	s.Equal(errs, nil)
	errs = Validate(IntMax{20})
	s.EqualError(errs, "Value: value mush more than: 10\n")

	errs = Validate(IntIn{10})
	s.Equal(errs, nil)
	errs = Validate(IntIn{20})
	s.EqualError(errs, "Value: value '20' is not included in set\n")

	errs = Validate(SliceString{[]string{"ttt", "iii"}})
	s.Equal(errs, nil)
	errs = Validate(SliceString{[]string{"tt", "ii"}})
	s.EqualError(errs, "Value: size less than 3\nValue: size less than 3\n")

	errs = Validate(Response{200, "test"})
	s.Equal(errs, nil)
	errs = Validate(Response{201, "test"})
	s.EqualError(errs, "Code: value '201' is not included in set\n")

	user := s.fakeUser()
	errs = Validate(user)
	s.Equal(errs, nil)
}

func (s *validationSuite) TestValidatinErrors() {
	vErrors := make(ValidationErrors, 0)
	err := ValidationError{"test", errors.New("Error test")}
	vErrors = append(vErrors, err)
	s.EqualError(vErrors, "test: Error test\n")
}
