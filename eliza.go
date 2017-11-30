package main

import(
	"fmt"
	"regexp"
	"time"
	"math/rand"
	"strings"
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
	re := regexp.MustCompile(`(?i).*\bI am\b|\bI'm\b|\bIm\b.*`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1")
	}
	if matched, _:= regexp.MatchString(`(?i).*\bHello\b|\bHi\b|\bHey\b.*`, input);matched{
		return "Hello there, how are you"
	}
	if matched, _:= regexp.MatchString(`(?i).*\bHelp\b.*`,input);matched{
		return "If we continue talking, maybe I can help you?"
	}
	if matched, _:= regexp.MatchString(`(?i).*\bGoodbye|Bye\b.*`,input);matched{
		return "Farewell friend, until next time"
	}
	//This randomly returns one of the following strings if nothing is matched.
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	return answers[rand.Intn(len(answers))]
}

//Adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
func reflectPronouns(input string) string {
	//Splitting the input on word.
	re := regexp.MustCompile(`\b`)
	tokens := re.Split(input, -1)
	
	// List of the reflections.
	reflectionMap := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	//Looping through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflectionMap := range reflectionMap {
			if matched, _ := regexp.MatchString(reflectionMap[0], token); matched {
				tokens[i] = reflectionMap[1]
				break
			}
		}
	}
	
	//Putting the tokens back together.
	return strings.Join(tokens, ``)
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

	fmt.Println("Im happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()

	fmt.Println("I'm not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()

	fmt.Println("Hello Eliza!")
	fmt.Println(ElizaResponse("Hello Eliza!"))
	fmt.Println()

	fmt.Println("Goodbye for now, until we speak again")
	fmt.Println(ElizaResponse("Goodbye for now, until we speak again"))
	fmt.Println()

	fmt.Println("Can you help me please")
	fmt.Println(ElizaResponse("Can you help me please"))
	fmt.Println()
	

	fmt.Println(reflectPronouns("You are my friend."))
	
}