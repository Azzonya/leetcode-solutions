package main

func main() {
	reverseString([]byte{'H', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = tmp
	}
}
