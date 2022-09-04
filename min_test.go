package basicsfunctions

import (
	"testing"

)


func TestMin(t *testing.T){

	min := Min([]int{1,2,3,4,5,6,10}) 
	if(min != 1){
		t.Errorf("Max was incorrect, got: %d but want: %d",min,10)
	}

}
