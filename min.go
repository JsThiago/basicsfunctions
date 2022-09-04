package basicsfunctions
func Min[V int | int64 | float32 | float64](m []V) V{
	if(len(m) <= 0){
		return -1
	}
	min := m[0]
	for _, num:= range m{
		if(num<min){
			min = num
		}
	}
	return min

}
