// generated by Textmapper; DO NOT EDIT

package simple

import (
	"fmt"
)

var afterListStates = map[int]bool{
	6: true,
	7: true,
}

const atBState = 1

var tmNonterminals = [...]string{
	"Bar_list",
	"Foo_list",
	"Xyz_list",
	"input",
	"Foo",
	"Bar",
	"Xyz",
}

func symbolName(sym int32) string {
	if sym == noToken {
		return "<no-token>"
	}
	if sym < int32(NumTokens) {
		return Token(sym).String()
	}
	if i := int(sym) - int(NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
	-1, -1, 12, 11, 13, -3, -9, -15, 3, 1, 5, 7, 6, 0, 2, 4, -1, -2,
}

var tmLalr = []int32{
	3, -1, 0, 10, -1, -2, 4, -1, 0, 9, -1, -2, 5, -1, 0, 8, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 4, 8, 14, 20, 22, 24, 26, 28, 32, 36, 42,
}

var tmFromTo = []int8{
	16, 17, 0, 1, 0, 2, 5, 2, 0, 3, 1, 11, 6, 3, 0, 4, 1, 4, 7, 4, 0, 5, 0, 6, 0,
	7, 0, 16, 0, 8, 6, 14, 0, 9, 5, 13, 0, 10, 1, 12, 7, 15,
}

var tmRuleLen = []int8{
	2, 1, 2, 1, 2, 1, 2, 2, 1, 1, 1, 1, 1, 1,
}

var tmRuleSymbol = []int32{
	6, 6, 7, 7, 8, 8, 9, 9, 9, 9, 9, 10, 11, 12,
}

var tmRuleType = [...]NodeType{
	0, // Bar_list : Bar_list Bar
	0, // Bar_list : Bar
	0, // Foo_list : Foo_list Foo
	0, // Foo_list : Foo
	0, // Xyz_list : Xyz_list Xyz
	0, // Xyz_list : Xyz
	0, // input : 'simple' Xyz
	0, // input : 'simple' .atB 'b'
	0, // input : Xyz_list .afterList
	0, // input : Foo_list .afterList
	0, // input : Bar_list
	0, // Foo : 'b'
	0, // Bar : 'a'
	0, // Xyz : 'c'
}

// set(follow 'simple') = CHAR_B, CHAR_C
var afterSimple = []int32{
	4, 5,
}
