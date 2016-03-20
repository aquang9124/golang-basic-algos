package main
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("The sum of all multiples of 3 or 5 below 1000:", threeOrFive(1000))
	fmt.Println("evenFibonacci function output:", evenFibonacci(12))
	fmt.Println("The factorial of 5 is:", factorial(5))
	fmt.Println("The reverse of the string 'dlrow olleh' is:", reverseStr("dlrow olleh"))
	fmt.Println("The total sum of all integers up to and including 10 is:", sumNums(10))
	fmt.Println("Output of longestWord:", longestWord("Hello there my friend"))
	fmt.Println("Output of vowelCount:", vowelCount("aeioua"))
	fmt.Println("'racecar' is a palindrome:", palindrome("racecar"))
}

// This function will find and return the sum of all those
// integers which are multiples of 3 or 5 below x.
func threeOrFive(x int) int {
	// sum is initialized to 0 using short declaration
	sum := 0

	for i := 0; i < x; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	// At the end return sum
	return sum
}

/* -------------------- evenFibonacci --------------------------
   Currently this function will generate a fibonacci sequence
   that contains 'x' extra numbers past the original 0, 1 and
   it will then sum up all of the numbers that are even.
   It will break out of the initial for loop if at any point
   the values in the fibonacci sequence become greater than
   4 million 
------------------------------------------------------------- */
func evenFibonacci(x int) int {
	// A slice called k is created with a length of 2.
	k := make([]int, 2)
	// Since k is initialized as an array of 0s, we manually set k[1] = 1
	k[1] = 1
	for i, j, total := 0, 2, 1; i < x; i, j, total = i+1, j+1, k[j - 1] + k[j - 2] {
		if total >= 4000000 {
			break
		}
		k = append(k, (k[j - 1] + k[j - 2]))
	}
	// sum is initialized to 0 using short declaration
	sum := 0
	for i, j := range k {
		// if an element of index i in k is even then it gets added to the sum
		if k[i] % 2 == 0 {
			sum += j
		}
	}
	// At the very end we return sum
	return sum
}

// This is a simple factorial function
func factorial(n int) int {
	// Here the variable fact of type int is initialized to 'n'
	// I could have declared it using short declaration 'fact := n'
	// or even 'var fact = n' as well.
	var fact int = n

	if n == 0 {
		return 1
	}
	// If n > 0, then this for loop will run and calculate the result
	for i := n - 1; i >= 1; i = i - 1 {
		fact = fact * i
	}
	// Finally, we return fact
	return fact
}

// This function takes a string as input and returns a reverse
// of that string.
func reverseStr(str string) string {
	reversedStr := ""
	s := strings.Split(str, "")
	for i := 0; i < len(s); i++ {
		reversedStr = s[i] + reversedStr
	}
	return reversedStr
}

// This function takes an integer, 'num', as input and
// returns the sum of all integers between zero and num,
// up to and including num.
func sumNums(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	return total
}

// This function takes in a string and returns the
// longest word in that string. This is assuming
// that the string contains only letters and spaces.
func longestWord(str string) string {
	var longWord, currWord string
	splitStr := strings.Split(str, " ")

	for i := 0; i < len(splitStr); i++ {
		currWord = splitStr[i]		
		if len(currWord) > len(longWord) || longWord == "" {
			longWord = currWord
		} 
	}
	return longWord
}

// This function will take in a string and return
// the number of vowels in that string.
func vowelCount(str string) int {
	arr := [5]string{"a", "e", "i", "o", "u"}
	count := 0
	splitStr := strings.Split(str, "")
	for i := 0; i < len(splitStr); i++ {
		for _, val := range arr {
			if splitStr[i] == val {
				count += 1
			}
		}
	}
	return count
}

// This function will take a string and return true if it is 
// a palindrome; false if it isn't. For example: "racecar" is a
// palindrome.
func palindrome(str string) bool {
	splitStr := strings.Split(str, "")
	for i, j := 0, len(splitStr) - 1; i < j; i, j = i + 1, j - 1 {
		if splitStr[i] != splitStr[j] {
			return false
		}
	}
	return true 
}