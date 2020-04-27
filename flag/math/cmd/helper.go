package cmd

import (
	"fmt"
	"errors"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func calc(values []float64, opType OpType) float64 {
	var result float64
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		switch opType {
		case ADD:
			result = result + values[i]
		case MINUS:
			result = result - values[i]
		case MULTIPLY:
			result = result * values[i]
		case DIVIDE:
			fmt.Printf("begin result: %f\n", result)
			fmt.Printf("begin values[i]: %f\n", values[i])
			if values[i] == 0 {
				fmt.Printf("values[i] == 0: %f\n", values[i])
				switch ErrorHandling(dividedByZeroHandling) {
				case ReturnOnDevidedByZero:
					return result
				case PanicOnDividedByZero:
					panic(errors.New("divided by 0"))
				}
			}
			result = result / values[i]
			fmt.Printf("finish result: %f\n", result)
		}
	}
	return result
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args: %v error:%v", cmd.Name(), args, err)
}

func ConvertArgsToFloat64Slice(args []string, errorHandling ErrorHandling) []float64 {
	result := make([]float64, 0, len(args))
	for _, arg := range args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			switch errorHandling {
			case ExitOnParseError:
				fmt.Fprintf(os.Stderr, "invalid number: %s\n", arg)
				os.Exit(1)
			case PanicOnParseError:
				panic(err)
			case ContinueOnParseError:
				continue
			}
		}
		result = append(result, value)
	}
	fmt.Println(result)
	return result
}
