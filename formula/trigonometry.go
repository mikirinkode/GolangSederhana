package formula

import (
	"errors"
	"fmt"
	"math"
)

func GetSudutTrigonometri(x float64, jenis string) (result float64, err error) {
		switch jenis {
		case "sin":
			result = math.Sin(x)
		case "cos":
			result = math.Cos(x)
		case "tan":
			result = math.Tan(x)
	default:
		err = errors.New("tipe tidak didefinisikan")
	}
	return result, err
}

func displayTrigonometryResult(result float64, errorStatus error){
	if errorStatus == nil {
		fmt.Printf("%.2f", result)
		fmt.Println()
	} else {
		fmt.Println(errorStatus.Error())
	}
}