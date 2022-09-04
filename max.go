package basicsfunctions
func Max[V int | int64 | float32 | float64](m []V) V{
	if(len(m) <= 0){
		return -1
	}
	max := m[0]
	for _, num:= range m{
		if(num>max){
			max = num
		}
	}
	return max

}


