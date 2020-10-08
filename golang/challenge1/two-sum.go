package challenge

func twoSum(nums []int, target int) []int {
	// index is responsible for saving the serial number of map[integer] integer
	index := make(map[int]int, len(nums))

	// Get the serial number of b through the for loop
	for  i , b  : =  range  nums {
		// Get the serial number of a = target-b by querying the map
		if  j , ok  : =  index [ target - b ]; ok {
			// ok is true
			// Explain that there is nums[j] == a before i
			return []int{j, i}
			// Note that the order is j, i
		}

		// Store the values ​​of b and i in the map
		index[b] = i
	}

	return nil
}
