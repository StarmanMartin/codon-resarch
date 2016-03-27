package resarch

import (
	"github.com/starmanmartin/codon-resarch/ctools"
)

func (c *CodonGraph) IsSelfComplementary() {
	c.PropertyOne, c.PropertyTwo = c.hasProperties()
	hasComp := make([]bool, len(c.List))
	c.StrongNotSelfComplementary = true
    isSelfComplementary := true
	for idx, val := range c.List {
		if !hasComp[idx] {
			comp := ctools.MakeComplementar(val)
			cIdx, hasNot := indexOf(c.List, comp)
			if hasNot {
                isSelfComplementary = false
                if !c.StrongNotSelfComplementary {
                   return
                }
			} else {
				c.StrongNotSelfComplementary = false
				hasComp[cIdx] = true
			}
		}
	}

	c.SelfComplementary = !c.StrongNotSelfComplementary && isSelfComplementary
}

func (c *CodonGraph) hasProperties() (bool, bool) {
	hasComp := make([]bool, len(c.DinucleotideNodes))
	isPTwo := true
	for tIdx := 0; tIdx < len(baseList); tIdx += 2 {
		oneIndex, _ := indexOf(c.Nucleotide, baseList[tIdx])
		twoIndex, _ := indexOf(c.Nucleotide, baseList[tIdx+1])

		if len(c.TetranucleotideNodes[oneIndex*2]) != len(c.TetranucleotideNodes[twoIndex*2+1]) {
			isPTwo = false
			break
		}

		if len(c.TetranucleotideNodes[oneIndex*2+1]) != len(c.TetranucleotideNodes[twoIndex*2]) {
			isPTwo = false
			break
		}
	}

	for idx, val := range c.DinucleotideNodes {
		if !hasComp[idx] {
			comp := ctools.MakeComplementar(val)
			cIdx, hasNot := indexOf(c.DinucleotideNodes, comp)
			if hasNot {
				return false, false
			}

			hasComp[cIdx] = true

			if isPTwo {
				countO, countT := 0, 0

				for tIdx := 0; tIdx < len(c.TetranucleotideNodes); tIdx += 2 {
					var hasNot bool
					_, hasNot = indexOfInt(c.TetranucleotideNodes[tIdx], cIdx)
					if !hasNot {
						countO++
					}
					_, hasNot = indexOfInt(c.TetranucleotideNodes[tIdx+1], idx)
					if !hasNot {
						countO--
					}

					_, hasNot = indexOfInt(c.TetranucleotideNodes[tIdx+1], cIdx)
					if !hasNot {
						countT++
					}
					_, hasNot = indexOfInt(c.TetranucleotideNodes[tIdx], idx)
					if !hasNot {
						countT--
					}
				}

				if countO != 0 || countT != 0 {
					isPTwo = false
				}
			}
		}
	}

	return true, isPTwo
}
