package cThree

import (
	"strings"
	"fmt"
	"github.com/starmanmartin/codon-resarch/ctools"
)

var Representer = []string{
"AUU ACC AGG CUU GGC AAC ACG GUU GCC AAG AGU UGA CGU ACU UCA AAU GGU CCU GCU AGC",
"AUU ACC CUU GGC AAC GUU GCC AAG CAG GAC GAG AUC CUC AAU GGU GUA CUG GUC UAC GAU",
"AUU AGG CUU CGG AAC CAC ACG GUU CCG AAG CAG CGU CUA AAU AUG CCU CAU CUG GUG UAG",
"UUA CUU CGG AAC CAC GUU CCG AAG CAG GAC GAG CUA CUC AUG UAA CAU CUG GUC GUG UAG",
"UUA CUU GGC AAC CAC GCA GUU GCC AAG GAC GAG AUC UGC CUC GUA UAA GUC UAC GAU GUG",
"AUU ACC AGG CGG AAC UCU ACG GUU CCG AGU UGA CGU ACU UCA AAU GGU CCU GCU AGA AGC",
"AUU ACC AGG GGC AAC UCU ACG GUU GCC AGU UGA CGU ACU UCA AAU GGU CCU GCU AGA AGC",
"AUU ACC AGG CGG AAC UCU ACG GUU CCG AGU CGU ACU AAU AUG GGU CCU CAU GCU AGA AGC",
"GGA AUU ACC GGC AAC UCU ACG GUU GCC AGU AUC CGU ACU AAU GGU UCC GCU AGA GAU AGC",
"AUU AGG CGG AAC UCU CAC ACG GUU CCG CAG AGU CGU ACU AAU AUG CCU CAU CUG AGA GUG",
"GGA AUU GGC AAC UCU CAC GUU GCC GAC AGU AUC ACU AAU UCC GUC GCU AGA GAU GUG AGC",
"AUU UUC GGC AAC CAC GUU GAA GCC CAG GAC GAG AGU AUC ACU CUC AAU CUG GUC GAU GUG",
"GGA AUU UUC GGC AAC CAC GUU GAA GCC GAC AUC AAU UCC GUA GUC UAC GCU GAU GUG AGC",
"AUU ACC AGG CUU CGG AAC ACG GUU CCG AAG AGU CGU ACU AAU AUG GGU CCU CAU GCU AGC",
"AUU ACC AGG CUU CGG AAC ACG GUU CCG AAG CAG AGU CGU ACU AAU AUG GGU CCU CAU CUG",
"AUU ACC CUU CGG AAC ACG GUU CCG AAG CAG GAG AGU CGU ACU CUC AAU AUG GGU CAU CUG",
"AUU ACC AGG CUU GGC AAC GUU GCC AAG GAC AGU AUC ACU AAU GGU CCU GUC GCU GAU AGC",
"AUU AGG CUU CGG AAC CAC ACG GUU CCG AAG AGU CGU ACU AAU AUG CCU CAU GCU GUG AGC",
"AUU AGG CUU GGC AAC CAC GUU GCC AAG GAC AGU AUC ACU AAU CCU GUC GCU GAU GUG AGC",
"AUU CUU CGG AAC CAC GUU CCG AAG CAG GAC GAG AGU ACU CUC AAU AUG CAU CUG GUC GUG",
"AUU AGG CUU CGG AAC CAC ACG GUU CCG AAG CAG AGU CGU ACU AAU AUG CCU CAU CUG GUG",
"AUU CUU CGG AAC CAC GUU CCG AAG CAG GAC GAG CUA CUC AAU AUG CAU CUG GUC GUG UAG",
"AUU CUU GGC AAC CAC GUU GCC AAG GAC GAG AUC CUC AAU GUA GUC UAC GCU GAU GUG AGC",
"GGA AUU ACC GGC AAC UCU GUU GCC GAC AGU AUC ACU AAU GGU UCC GUC GCU AGA GAU AGC",
"GGA AUU GGC AAC UCU CAC GUU GCC GAC AUC AAU UCC GUA GUC UAC GCU AGA GAU GUG AGC",
"AUU UUC GGC AAC CAC GUU GAA GCC CAG GAC GAG AUC CUC AAU GUA CUG GUC UAC GAU GUG",
"GGA UUA UUC GGC AAC CAC GCA GUU GAA GCC GAC UGA UGC UCA UCC GUA UAA GUC UAC GUG",
}

func GetSingleSwitchPath() map[int]string {
	retMap := make(map[int]string)
	idx := 0
	count := 0
	hasCoosen := make([]bool, len(Representer))
	for count < len(Representer) {
		code := strings.Split(Representer[idx], " ")
		for innerIdx, compareText := range Representer {
			if !hasCoosen[innerIdx] && innerIdx != idx {
				compareVal := strings.Split(compareText, " ")
				codeList, ok := diffJustCodonPaar(code, compareVal)
				if ok {
					fmt.Println(idx, codeList)
					hasCoosen[idx] = true
					idx = innerIdx
					retMap[idx] = strings.Join(codeList, " ")
					break
				}
			}
		}
		
		count++
	}
	
	return retMap
}

var permutations = []string {
	"A<->U",
	"G<->C",
	"A<->U;G<->C",
	"A<->C;U<->G",
	"U<->C;A<->G",
}

func diffJustCodonPaar(codeA, codeB []string) ([]string, bool) {
	for _, perm := range permutations {
		tempCode, _ := ctools.PermutateCodons(codeB, perm)
		vl, ok := diffJustOnePerm(codeA, tempCode)
		if ok {
			return vl, ok
		}
	}
	
	return nil, false
}

func diffJustOnePerm(codeA, codeB []string) ([]string, bool) {
	counter := 0
	vals := make([]string, 0 ,2)
	for _, valA := range codeA {
		hit := false
		for _, valB := range codeB {
			if valA == valB {
				counter++
				hit = true
				break
			}
		}
		
		if !hit {
			vals = append(vals, valA)
		}
	}
	
	return vals, counter >= len(codeA) - 2
}