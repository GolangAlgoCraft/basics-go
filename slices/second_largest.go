package slices
//returns the second largest element in a slice of ints.

func SecondLargest(slice []int) int {
	if len(slice) < 2 {
		return 0
	}
	firstLargest := slice[0]
	SecondLargest := slice[0]

	for _,value :=range slice{
		if firstLargest < value{
			SecondLargest = firstLargest
			firstLargest = value
		}else if SecondLargest < value && value != firstLargest{
			SecondLargest = value
		}
	}
	return SecondLargest
}