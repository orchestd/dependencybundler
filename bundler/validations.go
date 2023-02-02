package bundler

import (
	constructor "github.com/orchestd/dependencybundler/constructors/validations"
	validations2 "github.com/orchestd/dependencybundler/interfaces/validations"
	"github.com/orchestd/validations"
	"github.com/orchestd/validations/bvalidator"
	"github.com/orchestd/validations/bvalidator/customValidations"
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
