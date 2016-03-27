package ctools

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
    "time"
	"math/rand"
)

var (
	ruleReg        *regexp.Regexp
	CodonReg       *regexp.Regexp
	BaseList       [4]string
	ComplementList map[rune]string
)

func init() {
	CodonReg = regexp.MustCompile(`([AUGC]{3})`)
	ruleReg = regexp.MustCompile(`([AUCG])\<-\>([AUCG])`)

	BaseList = [...]string{"A", "U", "C", "G"}
	ComplementList = map[rune]string{
		[]rune(BaseList[0])[0]: BaseList[1],
		[]rune(BaseList[1])[0]: BaseList[0],
		[]rune(BaseList[3])[0]: BaseList[2],
		[]rune(BaseList[2])[0]: BaseList[3],
	}
}

func preparePermutateRule(rules string) ([]string, error) {
	ruleRegList := ruleReg.FindAllStringSubmatch(rules, -1)
	if len(ruleRegList) > 2 || len(ruleRegList) <= 0 {
		return nil, errors.New("Rule not correct")
	}

	ruleList := make([]string, len(ruleRegList)*2)
	for idx, sRule := range ruleRegList {
		ruleList[idx*2] = sRule[1]
		ruleList[idx*2+1] = sRule[2]
	}

	return ruleList, nil
}

func ShiftLeft(list []string) string {
	tempStringList := strings.Join(list, "") + string(list[0][0])
	return CodonReg.ReplaceAllString(tempStringList[1:], "$1 ")
}

func ShiftCodonLeft(list []string) string {
	for i, val := range list {
		list[i] = val[1:] + val[:1]
	}

	return strings.Join(list, " ")
}

func FillComlements(list []string) string {
	isCompleemnt := make([]bool, len(list))
	for i := 0; i < len(isCompleemnt); i++ {
		if !isCompleemnt[i] {
			val := list[i]
			comp := MakeComplementar(val)
			index, has := findCodonInList(list[i:], comp)
			if has {
				index += i
				isCompleemnt[index] = true
			} else {
				list = append(list, comp)
			}
		}
	}

	return strings.Join(list, " ")
}

func RemoveComlements(list []string) string {
	newList := make([]string, 0, len(list)/2)
	for i := 0; i < len(list); i++ {
        val := list[i]
        comp := MakeComplementar(val)
        _, has := findCodonInList(newList, comp)
        if !has {
            newList = append(newList, val)
        }
	}

	return strings.Join(newList, " ")
}

func Shuffle(list []string) string {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
    for idx := range list {
        newInsdex := rand.Intn(len(list))
        list[idx], list[newInsdex] = list[newInsdex], list[idx]
    }

	return strings.Join(list, " ")
}

func findCodonInList(list []string, codon string) (int, bool) {
	for idx, val := range list {
		if val == codon {
			return idx, true
		}
	}

	return -1, false
}

// PermutateCodons permutates a list of codons by a rule.
// Rule sample: A<->U;G<->C
func PermutateCodons(codonList []string, rules string) ([]string, error) {
	ruleList, err := preparePermutateRule(rules)
	if err != nil {
		return nil, err
	}

	var lastBase string
	for idx, rule := range ruleList {
		letter := "L"
		if idx%2 == 0 {
			for cIdex, codon := range codonList {
				codonList[cIdex] = strings.Replace(codon, rule, letter, -1)
			}

			lastBase = rule
		} else {
			for cIdex, codon := range codonList {
				codon = strings.Replace(codon, rule, lastBase, -1)
				codonList[cIdex] = strings.Replace(codon, letter, rule, -1)
			}

		}
	}

	return codonList, nil
}

// MakeComplementar
func MakeComplementar(codon string) (comCodon string) {
	runes := []rune(codon)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	for i := 0; i < len(runes); i++ {
		comCodon = fmt.Sprintf("%s%s", comCodon, ComplementList[runes[i]])
	}

	return
}
