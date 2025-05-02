package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	hasil := sumOfMultiples(n)
	fmt.Print(hasil)
}
func sumOfMultiples(n int)int{
	var i, jum int
	for i = 1; i <= n; i++{
		if i%3 == 0{
			jum+= i
		}else if i%5 == 0{
			jum+= i
		}else if i%7 == 0{
			jum+= i
		}
	}
	return jum
}