package validations

import (
	"bitbucket.org/HeilaSystems/validations"
	iValidations "bitbucket.org/HeilaSystems/dependencybundler/interfaces/validations"
	"go.uber.org/fx"
)

type validationsDeps struct {
	fx.In
	CustomValidations []validations.CustomValidation `group:"customValidations"`
	ValidationsBuilder validations.Builder
}
func DefaultValidations(deps validationsDeps) iValidations.Validations{
	if len(deps.CustomValidations)>0{
		deps.ValidationsBuilder = deps.ValidationsBuilder.AddCustomValidations(deps.CustomValidations...)
	}
	validator,err := deps.ValidationsBuilder.Build()
	if err != nil {
		panic(err)
	}
	return validator
}