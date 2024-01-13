package utils

import (
	"fmt"
	"strings"
)

type Input struct {
	Command string
	Payload string
	Args    []string
}

func NewInput(command string, payload string) Input {
	p := Input{
		Command: command,
		Payload: payload,
	}
	p.Args = p.parseQuotes()
	return p
}

// parseQuotes handles quotes in a telegram command payload. This method needs refactoring as its dogshit, but it works
// for now. Returned is string slice.
func (p *Input) parseQuotes() []string {
	results := []string{}
	parts := []string{}
	payload := strings.TrimSpace(p.Payload)
	if len(p.Payload) == 0 {
		return []string{}
	}
	if !strings.Contains(payload, " ") {
		parts = []string{payload}
	} else {
		parts = strings.Split(payload, " ")
	}

	grabQuote := func(partInd int) (string, int) {
		append := func(existing, new string) string {
			if existing == "" {
				return new
			} else {
				return fmt.Sprintf("%s %s", existing, new)
			}
		}
		builder := ""
		for offset := 0; offset+partInd < len(parts); offset++ {
			if parts[partInd+offset][len(parts[partInd+offset])-1] == '"' {
				builder = append(builder, strings.Split(parts[partInd+offset], `"`)[0])
				return builder, partInd + offset
			} else {
				if offset == 0 {
					builder = append(builder, strings.Split(parts[partInd+offset], `"`)[1])
				} else {
					builder = append(builder, parts[partInd+offset])
				}
			}
		}
		return "", -1
	}
	for i := 0; i < len(parts); i++ {
		if parts[i][0] == '"' && parts[i][len(parts[i])-1] == '"' {
			results = append(results, strings.Split(parts[i], `"`)[1])
			continue
		}
		if parts[i][0] == '"' {
			quote, endingInd := grabQuote(i)
			if endingInd == -1 {
				// no finishing quote
				results = append(results, parts[i])
			} else {
				results = append(results, quote)
				i = endingInd
			}
		} else {
			results = append(results, parts[i])
		}
	}
	return results
}
