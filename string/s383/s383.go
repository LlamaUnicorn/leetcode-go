package s383

func RansomNote(ransomNote string, magazine string) bool {
	seen := make(map[string]int)
	for _, char := range magazine {
		seen[string(char)] += 1
	}

	for _, char := range ransomNote {
		_, ok := seen[string(char)]
		if !ok {
			return false
		} else {
			seen[string(char)] -= 1
			if seen[string(char)] < 0 {
				return false
			}
		}
	}
	return true
}

//func canConstruct(ransomNote string, magazine string) bool {
//	counts := [26]int{}
//	for i := 0; i < len(magazine); i++ {
//		counts[magazine[i]-'a']++
//	}
//	for i := 0; i < len(ransomNote); i++ {
//		counts[ransomNote[i]-'a']--
//		if counts[ransomNote[i]-'a'] < 0 {
//			return false
//		}
//	}
//	return true
//}
//func main() {
//	result := s383.RansomNote("aa", "ab")
//	fmt.Println(result)
//}
