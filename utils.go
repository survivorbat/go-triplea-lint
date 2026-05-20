package linters

import (
	"strings"
	"unicode"
)

// hasLineWithPrefix returns whether one of the lines in the input st arts with the given prefix
func hasLineWithPrefix(str string, sub string) bool {
	replacer := strings.NewReplacer("/", "", "*", "")

	for line := range strings.SplitSeq(str, "\n") {
		// Accomodate for:
		//  Act
		// /// Act
		// ///  Act
		// /** Act
		// //*  Act
		// ///*  Act
		line = replacer.Replace(line)
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, sub) {
			return true
		}
	}

	return false
}

// actCandidates attempts to put together a list of possible candidates to be Act in a test.
// These candidates are grouped by their likelihood of being the act.
//
// The function TestMyService_CreateThing_ActuallyCreatesThing will result in the following
// candidates for the act function:
//
// - MyThingService_CreateThing_ActuallyCreatesThing
//
// - MyThingService
// - CreateThing
//
// - Create
// - Thing
//
// - MyThing
// - ThingService
//
// - My
// - Thing
// - Service
//
// With the third segment being disregarded
func actCandidates(testName string) [][]string {
	name := strings.TrimPrefix(testName, "Test")

	// We consider TestAbc, TestAbc_Def and TestAbc_Def_DoesThing to be the "usual" convention.
	// For that reason we only split 3 segments at maximum
	const usualSegments = 3

	segments := strings.SplitN(name, "_", usualSegments)

	// The name of the test could be a candidate, so we use that first
	candidates := [][]string{{name}}

	switch len(segments) {
	case 3:
		// Consider the second segment the most likely candidate and disregard the final segment.
		// We consider the likelyhood of the third segment containing the name of the act to be extremely low, so
		// we add it in full as a candidate but won't analyse it further
		segments = []string{segments[1], segments[0]}
		candidates = append(candidates, segments)
	case 2:
		// Consider the second segment the most likely candidate, the first one might be the struct
		segments = []string{segments[1], segments[0]}
		candidates = append(candidates, segments)
	default:
		// With only one, we don't rearrange the segments
	}

	// Then we start matching words together for the remaining segments
	for _, segment := range segments {
		words := splitWords(segment)

		if len(words) != 1 {
			candidates = append(candidates, matchWords(words))
		}
	}

	for index, candidateList := range candidates {
		for subIndex, canidate := range candidateList {
			candidates[index][subIndex] = strings.ToLower(canidate)
		}
	}

	return candidates
}

func matchWords(words []string) []string {
	if len(words) < 2 {
		return words
	}

	result := make([]string, 0, len(words))

	for currentIndex, currentWord := range words {
		if currentIndex > len(words) {
			break
		}

		remainingWords := words[currentIndex+1:]

		for index := range remainingWords {
			result = append(result, currentWord+strings.Join(remainingWords[:len(remainingWords)-index], ""))
		}

		result = append(result, currentWord)
	}

	return result
}

// splitWords splits FooBarBaz into [Foo, Bar, Baz]
func splitWords(in string) []string {
	words := make([]string, 0, 1)

	var currentWord strings.Builder

	inRunes := []rune(in)

	for index, character := range inRunes {
		if index != 0 && unicode.IsUpper(character) {
			if unicode.IsLower(inRunes[index-1]) {
				words = append(words, currentWord.String())
				currentWord.Reset()
			} else if len(inRunes) > index+1 && unicode.IsLower(inRunes[index+1]) {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		}

		_, _ = currentWord.WriteRune(character)
	}

	if currentWord.Len() != 0 {
		words = append(words, currentWord.String())
	}

	return words
}
