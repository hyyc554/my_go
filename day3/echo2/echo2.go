// echo2 输出命令行参数
package main
import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	// s,sep := "",""
	// for _, arg := range os.Args[1:]{
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)
	// 法2
	// fmt.Println(strings.Join(os.Args[1:]," "))
	// 练习1.1
	// fmt.Println(os.Args[0])
	// 练习1.2
	// s,sep := "",""
	for n, arg := range os.Args[1:]{
		fmt.Println(n,arg)
		
	}
	

}