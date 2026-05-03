package heap

func parentIdxOf(i int) int     { return (i - 1) / 2 }
func leftChildIdxOf(i int) int  { return 2*i + 1 }
func rightChildIdxOf(i int) int { return 2*i + 2 }
