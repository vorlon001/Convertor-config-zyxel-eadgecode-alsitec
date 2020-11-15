package main

import "fmt"
import "strings"

func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func Find(a *[]int, x int) *int {
    for i, n := range *a {
        if x == n {
            return &i
        }
    }
    return nil
}
func Delete(items  *[]int, item  int) *[]int {
	newitems := []int{}
	for _, i := range *items {
		fmt.Println(i)
		if i != item {
			newitems = append(newitems, i)
		}
	}
	fmt.Println(newitems)
	return &newitems
}
func main() {

	A := &[]int{1,4}
	A = Delete(A,1)
	fmt.Println(A)
	A = Delete(A,4093)
	fmt.Println(A)

	fmt.Println(arrayToString(*A, ",")) //1,2,3,4,5,6,7,8,9
}
