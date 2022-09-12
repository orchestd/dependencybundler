package bundler

import (
	constructor "bitbucket.org/HeilaSystems/dependencybundler/constructors/validations"
	validations2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/validations"
	"bitbucket.org/HeilaSystems/validations"
	"bitbucket.org/HeilaSystems/validations/bvalidator"
	"bitbucket.org/HeilaSystems/validations/bvalidator/customValidations"
	"go.uber.org/fx"
)

const CustomValidationsGroup = "customValidations"

func ValidationsFxOption() fx.Option {
	israeliPhoneAll := validations2.NewValidatorCustomValidation("israeliPhoneAll", customValidations.ValidateIsraeliPhoneAll)
	israeliPhoneMobile := validations2.NewValidatorCustomValidation("israeliPhoneMobile", customValidations.ValidateIsraeliPhoneMobile)
	longDate := validations2.NewValidatorCustomValidation("longDate", customValidations.ValidateLongDate)
	regexp := validations2.NewValidatorCustomValidation("regexp", customValidations.ValidateRegexp)

	return fx.Options(
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: israeliPhoneAll}),
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: israeliPhoneMobile}),
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: longDate}),
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: regexp}),
		fx.Provide(func() validations.Builder {
			builder := bvalidator.Builder()
			return builder
		}),
		fx.Provide(constructor.DefaultValidations),
	)
}
