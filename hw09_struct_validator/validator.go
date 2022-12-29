package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type NotSupportedType error

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var b strings.Builder
	for _, e := range v {
		b.WriteString(fmt.Sprintf("%v: %v\n", e.Field, e.Err))
	}
	return b.String()
}

func Validate(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, but received %T", v)
	}

	vErrors := make(ValidationErrors, 0)

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		vField := rv.Field(i)
		if !vField.CanInterface() {
			continue
		}

		tField := rt.Field(i)
		tValidate, ok := tField.Tag.Lookup("validate")
		if !ok {
			continue
		}

		fv, err := buildFieldValidator(tValidate, vField)
		if err != nil {
			return err
		}

		errs := fv.Run(vField.Interface())
		if len(errs) > 0 {
			for _, e := range errs {
				vErrors = append(vErrors, ValidationError{tField.Name, e})
			}
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}

func buildFieldValidator(tagValue string, rv reflect.Value) (Validator, error) {
	var nErr NotSupportedType
	switch rk := rv.Kind(); rk {
	case reflect.String:
		return buildFieldStringValidatorBy(tagValue)
	case reflect.Int:
		return buildFieldIntValidatorBy(tagValue)
	case reflect.Slice:
		switch vk := rv.Index(0).Kind(); vk {
		case reflect.String:
			return buildFieldStringValidatorBy(tagValue)
		case reflect.Int:
			return buildFieldIntValidatorBy(tagValue)
		default:
			errors.As(fmt.Errorf("not supported type in slice: %s", vk), &nErr)
			return nil, nErr
		}
	default:
		errors.As(fmt.Errorf("not supported type: %s", rk), &nErr)
		return nil, nErr
	}
}

func buildFieldStringValidatorBy(tagValue string) (Validator, error) {
	tValidators := strings.Split(tagValue, "|")
	rValidators := make([]TypeValidator, 0)
	for _, tv := range tValidators {
		sv := strings.Split(tv, ":")
		switch sv[0] {
		case "len":
			size, err := strconv.Atoi(sv[1])
			if err != nil {
				return nil, err
			}
			v := lenStringValidator{size}
			rValidators = append(rValidators, &v)
		case "in":
			options := strings.Split(sv[1], ",")
			v := inStringValidate{options}
			rValidators = append(rValidators, &v)
		case "regexp":
			reg, err := regexp.Compile(sv[1])
			if err != nil {
				return nil, err
			}
			v := regStringValidator{reg}
			rValidators = append(rValidators, &v)
		default:
			return nil, fmt.Errorf("not supported validator: %s", sv[0])
		}
	}
	validator := fieldValidator{rValidators}

	return &validator, nil
}

func buildFieldIntValidatorBy(tagValue string) (Validator, error) {
	tValidators := strings.Split(tagValue, "|")
	rValidators := make([]TypeValidator, 0)
	for _, tv := range tValidators {
		sv := strings.Split(tv, ":")
		switch sv[0] {
		case "min":
			min, err := strconv.Atoi(sv[1])
			if err != nil {
				return nil, err
			}
			v := minIntValidator{min}
			rValidators = append(rValidators, &v)
		case "max":
			max, err := strconv.Atoi(sv[1])
			if err != nil {
				return nil, err
			}
			v := maxIntValidator{max}
			rValidators = append(rValidators, &v)
		case "in":
			options := make([]int, 0)
			for _, o := range strings.Split(sv[1], ",") {
				i, err := strconv.Atoi(o)
				if err != nil {
					return nil, fmt.Errorf("option '%v' is not int", o)
				}
				options = append(options, i)
			}
			v := inIntValidator{options}
			rValidators = append(rValidators, &v)
		default:
			return nil, fmt.Errorf("not supported validator: %s", sv[0])
		}
	}
	validator := fieldValidator{rValidators}

	return &validator, nil
}

type TypeValidator interface {
	validate(value interface{}) error
}

type lenStringValidator struct {
	size int
}

func (v *lenStringValidator) validate(value interface{}) error {
	var val string
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("type assertion error %v", value)
	}
	if len(val) < v.size {
		return fmt.Errorf("size less than %v", v.size)
	}
	return nil
}

type inStringValidate struct {
	options []string
}

func (v *inStringValidate) validate(value interface{}) error {
	for _, o := range v.options {
		if o == value {
			return nil
		}
	}
	return fmt.Errorf("value: %v is not included in set", value)
}

type regStringValidator struct {
	re *regexp.Regexp
}

func (v *regStringValidator) validate(value interface{}) error {
	var val string
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("type assertion error %v", value)
	}
	if res := v.re.MatchString(val); res {
		return nil
	}
	return fmt.Errorf("not match: %v", value)
}

type minIntValidator struct {
	min int
}

func (v *minIntValidator) validate(value interface{}) error {
	var val int
	val, ok := value.(int)
	if !ok {
		return fmt.Errorf("type assertion error %v", value)
	}
	if val < v.min {
		return fmt.Errorf("value less than: %v", v.min)
	}
	return nil
}

type maxIntValidator struct {
	max int
}

func (v *maxIntValidator) validate(value interface{}) error {
	var val int
	val, ok := value.(int)
	if !ok {
		return fmt.Errorf("type assertion error %v", value)
	}
	if val > v.max {
		return fmt.Errorf("value mush more than: %v", v.max)
	}
	return nil
}

type inIntValidator struct {
	options []int
}

func (v *inIntValidator) validate(value interface{}) error {
	for _, o := range v.options {
		if o == value {
			return nil
		}
	}
	return fmt.Errorf("value '%v' is not included in set", value)
}

type Validator interface {
	Run(v interface{}) []error
}

type fieldValidator struct {
	typeValidators []TypeValidator
}

func (sv *fieldValidator) Run(v interface{}) []error {
	errs := make([]error, 0)
	for _, tv := range sv.typeValidators {
		switch v := v.(type) {
		case []string:
			for _, i := range v {
				err := tv.validate(i)
				if err != nil {
					errs = append(errs, err)
				}
			}
		case []int:
			for _, i := range v {
				err := tv.validate(i)
				if err != nil {
					errs = append(errs, err)
				}
			}
		default:
			err := tv.validate(v)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}
