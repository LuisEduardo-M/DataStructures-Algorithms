package Google_Video

// Video: https://www.youtube.com/watch?v=XKu_SEDAykw

/* First Solution - Brute force
func hasPairWithSum(arr []int, sum int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				return true
			}
		}
	}
	return false
}
*/

// |=======================================================|
// |====================Better Solution====================|
// |=======================================================|

// Sets in Go: https://www.sohamkamani.com/golang/sets/

type set map[int]struct{}

// Adds an integer to the set
func (s set) add(n int) {
	s[n] = struct{}{}
}

// Returns a boolean value describing if the integer exists in the set
func (s set) has(n int) bool {
	_, ok := s[n]
	return ok
}

func hasPairWithSum(arr []int, sum int) bool {
	mySet := set{}

	for i := 0; i < len(arr); i++ {
		if mySet.has(arr[i]) {
			return true
		}
		mySet.add(sum - arr[i])
	}
	return false
}

// func main() {
// 	s := []int{6, 4, 3, 2, 1, 8}

// 	r := hasPairWithSum(s, 7) // true
// 	r := hasPairWithSum(s, 22) // false
// 	fmt.Println(r)
// }
