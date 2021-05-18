package permutation

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {

	a:= []interface{}{}
	permute([]interface{}{"a", "b", "c"}, []interface{}{}, &a)
	b:= []interface{}{}
	permute([]interface{}{}, []interface{}{}, &b)

	fmt.Println(a)
	fmt.Println(b)


}
