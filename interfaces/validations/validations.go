package validations

import (
	"bitbucket.org/HeilaSystems/validations"
)
import "bitbucket.org/HeilaSystems/validations/bvalidator/customValidations"

type Validations validations.Validations
type CustomValidation validations.CustomValidation

var NewValidatorCustomValidation = customValidations.NewCustomValidation

type Validator validations.Validator

type ValidatorRunner validations.ValidatorRunner

var NewValidatorCont = validations.NewValidatorCont
