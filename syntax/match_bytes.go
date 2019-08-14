package syntax

import "strings"

func MatchRunes(expr string, prog *Prog) []rune {
	if strings.Contains(expr, "|") {
		return nil
	}
	inst := prog.Inst[prog.Start]
	res := []rune{}
	for inst.Op != InstMatch {
		switch inst.Op {
		default:
			inst = prog.Inst[inst.Out]
		case InstFail:
			panic("unexpected InstFail")
		case InstAlt, InstAltPos, InstAltMatch:
			inst = prog.Inst[inst.Arg]
		case InstRune1:
			res = append(res, inst.Rune...)
			inst = prog.Inst[inst.Out]
		}
	}
	return res
}
