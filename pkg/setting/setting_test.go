package setting

import(
	"fmt"
	"testing"
)

func TestS(t *testing.T){
	fmt.Println(LoadConfig("database","PASSWORD"))
}