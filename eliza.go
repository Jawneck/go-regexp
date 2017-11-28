package main

import(
	"fmt"
	"regexp"
	"time"
	"math/rand"
)

func ElizaResponse(input string) string{
	if matched, _ := regexp.MatchString(`(?i).*\bFrench\b.*`, input); matched {
		return "France is a beautiful country"
	}
	if matched, _ := regexp.MatchString(`(?i).*\bweekend\b.*`, input); matched {
		return "I love the weekend!"
	}
	if matched, _ := regexp.MatchString(`(?i).*I was.*`, input); matched {
		return "Thats very nice"
	}
	if matched, _ := regexp.MatchString(`(?i).*father was.*`, input); matched {
		return "A teacher! Tell me more about him?"
	}
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

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()

}