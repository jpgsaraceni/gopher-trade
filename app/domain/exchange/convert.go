package exchange

import "context"

func (UseCase) Convert(_ context.Context, input ConvertInput) ConvertOutput {
	convertedAmount := input.Exchange.Convert(input.FromAmount)

	return ConvertOutput{
		ConvertedAmount: convertedAmount,
	}
}
