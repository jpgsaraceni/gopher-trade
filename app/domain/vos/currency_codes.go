package vos

type CurrencyCode string

func (c CurrencyCode) String() string {
	return string(c)
}

func ListOfCodesToString(codes ...CurrencyCode) []string {
	output := make([]string, 0, len(codes))
	for _, code := range codes {
		output = append(output, code.String())
	}

	return output
}
