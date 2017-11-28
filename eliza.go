package main

import(
	"fmt"
	"regexp"
	"time"
	"math/rand"
)

func ElizaResponse(input string) string{
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?"
	}

	//This randomly returns one of the following strings if nothing is matched.
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	return answers[rand.Intn(len(answers))]
}

func main() {
	//Using the current time to seed.
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

}