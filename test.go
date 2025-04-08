package yargs
import(
	"fmt"
)
type Type1 struct {
	One string
	Two string
}
type TypeA struct {
	A string
	B string
	C bool
	D int32
	F float32
	Item Type1
}
	
func test() {
	a := TypeA{A:"Aaaa", B:"Bbbb", C:true, D:12, F:37.2, Item:Type1{One:"Won", Two:"Too"}}
	Unmarshal([]string{"--A=Qqqq", "--C=false", "--D=21", "--F=24.2", "--Item.One=Tree"}, &a)
	fmt.Printf("%v", a)
}