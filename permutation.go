package permutation

func Ints(input []int) [][]int {
	interfaceInput := make([]interface{}, len(input))
	for i := range input {
		interfaceInput[i] = input[i]
	}

	var results [][]interface{}
	permute(interfaceInput, []interface{}{}, &results)

	intResults := make([][]int, len(results))
	for i := range results {
		var ints []int
		for _, v := range results[i] {
			ints = append(ints, v.(int))
		}
		intResults[i] = ints
	}

	return intResults
}

func Strings(input []string) [][]string {
	interfaceInput := make([]interface{}, len(input))
	for i := range input {
		interfaceInput[i] = input[i]
	}

	var results [][]interface{}
	permute(interfaceInput, []interface{}{}, &results)

	stringResults := make([][]string, len(results))
	for i := range results {
		var strings []string
		for _, v := range results[i] {
			strings = append(strings, v.(string))
		}
		stringResults[i] = strings
	}

	return stringResults
}

func permute(input, output []interface{}, results *[][]interface{}) {
	if len(input) == 0 {
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
