package main

import "fmt"
import "strings"

func main() {
	fmt.Printf("Hello World!")
	var emails [3]string {"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {

	var superMap map[string]bool = make(map[string]bool)
	for i:=0; i<len(emails); i++{
		var buffer strings.Builder

		index := strings.IndexRune(emails[i],'@')
		firstHalf := emails[i][:index]
		index=strings.IndexRune(firstHalf,'+')
		firstHalf = firstHalf[:index]
		firstHalf = strings.Replace(firstHalf,".","",100)

		secondHalf := emails[i][index+1:]

		buffer.WriteString(firstHalf)
		buffer.WriteRune('@')
		buffer.WriteString(secondHalf)

		superMap[buffer.String()] = true
	}

	return len(superMap)
}


