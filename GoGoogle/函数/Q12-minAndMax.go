func max(l []int)(max int){
	max=l[0]
	for _,v:=range l{
		if v>max{
			max=v
		}
	}
	return 
}

func min(l []int)(min int){
	min=l[0]
	for _v:=range l{
		if v<min{
			min=v
		}
	}
	return
}

