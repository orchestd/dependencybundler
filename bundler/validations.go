package bundler

import (
	validations2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/validations"
	"bitbucket.org/HeilaSystems/validations"
	"bitbucket.org/HeilaSystems/validations/bvalidator"
	constructor "bitbucket.org/HeilaSystems/dependencybundler/constructors/validations"
	"bitbucket.org/HeilaSystems/validations/bvalidator/customValidations"
	"go.uber.org/fx"
)

const CustomValidationsGroup = "customValidations"

func ValidationsFxOption() fx.Option {
	israeliPhoneAll := validations2.NewValidatorCustomValidation("israeliPhoneAll", customValidations.ValidateIsraeliPhoneAll)
	israeliPhoneMobile := validations2.NewValidatorCustomValidation("israeliPhoneMobile", customValidations.ValidateIsraeliPhoneMobile)
	longDate := validations2.NewValidatorCustomValidation("longDate", customValidations.ValidateLongDate)

	return fx.Options(
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: israeliPhoneAll}),
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: israeliPhoneMobile}),
		fx.Provide(fx.Annotated{Group: CustomValidationsGroup, Target: longDate}),
		fx.Provide(func() validations.Builder {
			builder := bvalidator.Builder()
			return builder
		}),
		fx.Provide(constructor.DefaultValidations),
	)
}
