package permutation

import "fmt"

func Ints(i []int) [][]int {

	return [][]int{}
}

func Strings(i []string) [][]string {

	return [][]string{}
}

func permute(input, output []interface{}, results *[]interface{}){
	if len(input) == 0 {
		fmt.Printf("%p \n", results)
		fmt.Printf("%v \n", output)

		*results = append(*results, output)
		return
	}

	for i := 0; i < len(input); i++ {
		permute(removeElement(i, input), append(append([]interface{}(nil), output...), input[i]), results)
	}

}

func removeElement(i int, list []interface{}) []interface{} {
	rest := append([]interface{}(nil), list[:i]...)

	rest = append(rest, list[i+1:]...)

	return rest
}
