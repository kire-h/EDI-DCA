func binarySearch(start int, end int, val int, list []int) int{
	if start > end{
		return -1
	}
	mid := (start+end)/2
	if list[mid] == val{
		return mid
	} else if list[mid]>val{
		return binarySearch(start, mid-1, val, list)
	} else{
		return binarySearch(mid+1, end, val, list)
	}
}

func bin_Search(val int, list []int) int {
    start := 0
    end := len(list) - 1

    for start <= end {
        mid := (start + end) / 2

        if list[mid] == val {
            return mid
        }

        if list[mid] > val {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }

    return -1
}
