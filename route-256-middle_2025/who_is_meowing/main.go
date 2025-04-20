package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		scores := make(map[string]int)
		var action string
		suspects := make(map[string]bool)

		for j := 0; j < n; j++ {
			var statement string
			line, _ := reader.ReadString('\n')
			statement = strings.TrimSpace(line)
			parts := strings.SplitN(statement, ": ", 2)
			speaker := parts[0]
			suspects[speaker] = true
			opinion := parts[1]

			if strings.HasPrefix(opinion, "I am ") {
				targetActionParts := strings.SplitN(opinion, "I am ", 2)
				isNot := false
				actWithPunctuation := targetActionParts[1]
				if strings.HasPrefix(actWithPunctuation, "not ") {
					isNot = true
					actWithPunctuation = strings.TrimPrefix(actWithPunctuation, "not ")
				}
				act := strings.TrimSuffix(actWithPunctuation, "!") // Remove potential exclamation mark
				act = strings.TrimSuffix(act, ".")                 // Remove potential period (though unlikely here)
				if action == "" {
					action = act
				}
				if !isNot {
					scores[speaker] += 2
				} else {
					scores[speaker] -= 1
				}
			} else {
				opinionParts := strings.SplitN(opinion, " is ", 2)
				target := opinionParts[0]
				targetActionWithPunctuation := opinionParts[1]
				isNot := false
				if strings.HasPrefix(targetActionWithPunctuation, "not ") {
					isNot = true
					targetActionWithPunctuation = strings.TrimPrefix(targetActionWithPunctuation, "not ")
				}
				targetAction := strings.TrimSuffix(targetActionWithPunctuation, "!") // Remove potential exclamation mark
				targetAction = strings.TrimSuffix(targetAction, ".")                 // Remove potential period
				suspects[target] = true
				if action == "" {
					action = targetAction
				}
				if !isNot {
					scores[target] += 1
				} else {
					scores[target] -= 1
				}
			}
		}

		maxScore := -1000000000
		winners := []string{}
		for name, score := range scores {
			if score > maxScore {
				maxScore = score
				winners = []string{name}
			} else if score == maxScore {
				winners = append(winners, name)
			}
		}
		sort.Strings(winners)
		for _, winner := range winners {
			fmt.Printf("%s is %s.\n", winner, action)
		}
	}
}
