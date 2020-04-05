// An element in a sorted array can be found in O(log n) time via binary search.
// But suppose we rotate an ascending order sorted array at some pivot unknown to you beforehand.
// So for instance, 1 2 3 4 5 might become 3 4 5 1 2.
// Devise a way to find an element in the rotated array in O(log n) time.

package main

func main() {
	arr := [5]int{30, 40, 50, 10, 20}
	key := 10
	index := searchKey(arr, key)
	if index == -1 {
		fmt.Println("key not found in array")
	} else {
		fmt.println("The key is avaliable at index: ", index)
	}
}

//non rotational case not handled
func searchKey(a [5]int, key int) int {
	//find pivot point
	var pivot int
	for i := 0; i < len(a); i++ {
		if a[i] > a[i+1] {
			pivot = i + 1
		}
	}
	//compare key with 1st element
	index := -1
	if a[0] >= key {
		for i = 0; i < pivot; i++ {
			if a[i] == key {
				index = i
			}
		}
	} else {
		for i = pivot; i < len(a); i++ {
			if a[i] == key {
				index = i
			}
		}
	}

	return index
}
