package main
import "fmt"
func main() {
var bil int
fmt.Scan(&bil)
if bil == 0 {
	fmt.Println(0)
	return
	}
for bil > 0 {
	digit := bil % 10
	fmt.Println(digit)
	bil = bil / 10
	}
}