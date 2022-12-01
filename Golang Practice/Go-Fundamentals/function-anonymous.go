package main

import "fmt"

type Member func(string) bool

func registerMember(name string, member Member) {
	if member(name) {
		fmt.Println("Hallo", name)
	} else {
		fmt.Println("Register Now!!")
	}
}

func main() {
	memberlist := func(name string) bool {
		return name == "Member"
	}

	registerMember("Member", memberlist)
	registerMember("Jamil", memberlist)

}
