package tempconv
import "fmt"
type Celsius float64
type Fahrenheit float64
const (
	AbsoluteZeroc Celsius = -273.15
	FreezingC Celsius = 0
	Boiling Celsius = 100

)

func (c Celsius) String() string  {
	return fmt.Sprintf("%g~C",c)
	

}

func (f Celsius) String() string  {
	return fmt.Sprintf("%g~C",f)
	

}