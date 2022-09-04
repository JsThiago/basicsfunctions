package basicsfunctions

import (
	"testing"

)


func TestMax(t *testing.T){

	max := Max([]int{1,2,3,4,5,6,10}) 
	if(max != 10){
		t.Errorf("Max was incorrect, got: %d but want: %d",max,10)
	}

}
