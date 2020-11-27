// generated by Textmapper; DO NOT EDIT

package test

import (
	"fmt"
)

var tmNonterminals = [...]string{
	"Declaration_list",
	"Test",
	"Declaration",
	"setof_not_((eoi | '.') | '}')",
	"setof_not_((eoi | '.') | '}')_optlist",
	"empty1",
	"QualifiedName",
	"Decl1",
	"Decl2",
}

func symbolName(sym int32) string {
	if sym < int32(NumTokens) {
		return Token(sym).String()
	}
	if i := int(sym) - int(NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
	-1, -1, -3, 11, -1, -1, 46, -1, -23, 1, 3, 4, -1, 16, 41, 42, -1, 10, -1, -1,
	0, 12, -1, -1, 43, -1, 8, -1, -1, 9, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	14, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 15, 45, -1, 6,
	-1, 7, 44, 5, -1, -1, -2, -2,
}

var tmLalr = []int32{
	14, -1, 0, 13, 5, 13, 6, 13, 7, 13, 8, 13, 9, 13, 10, 13, 11, 13, -1, -2, 5,
	-1, 6, -1, 7, -1, 8, -1, 9, -1, 10, -1, 0, 2, -1, -2,
}

var tmGoto = []int32{
	0, 4, 6, 8, 10, 16, 36, 54, 72, 92, 110, 130, 144, 150, 156, 160, 164, 166,
	168, 170, 176, 178, 180, 182, 184, 186, 188, 196, 198, 214, 216, 218, 220,
	222, 240, 256,
}

var tmFromTo = []int8{
	63, 65, 64, 66, 22, 30, 22, 31, 22, 32, 16, 24, 22, 33, 57, 61, 0, 2, 4, 13,
	7, 2, 8, 2, 18, 2, 19, 2, 22, 34, 27, 2, 28, 2, 59, 2, 0, 3, 7, 3, 8, 3, 18,
	3, 19, 3, 22, 35, 27, 3, 28, 3, 59, 3, 0, 4, 7, 4, 8, 4, 18, 4, 19, 4, 22,
	36, 27, 4, 28, 4, 59, 4, 0, 5, 1, 5, 7, 5, 8, 5, 18, 5, 19, 5, 22, 37, 27, 5,
	28, 5, 59, 5, 0, 6, 7, 6, 8, 6, 18, 6, 19, 6, 22, 38, 27, 6, 28, 6, 59, 6, 0,
	7, 4, 14, 7, 7, 8, 7, 18, 7, 19, 7, 22, 39, 27, 7, 28, 7, 59, 7, 7, 17, 18,
	26, 19, 29, 22, 40, 27, 58, 28, 60, 59, 62, 4, 15, 5, 16, 22, 41, 22, 42, 23,
	55, 25, 56, 2, 12, 22, 43, 12, 21, 22, 44, 25, 57, 22, 45, 22, 46, 7, 18, 18,
	27, 22, 47, 22, 48, 22, 49, 22, 50, 22, 51, 22, 52, 22, 53, 0, 8, 7, 19, 18,
	28, 27, 59, 0, 63, 0, 9, 7, 9, 8, 20, 18, 9, 19, 20, 27, 9, 28, 20, 59, 20,
	22, 54, 14, 22, 15, 23, 16, 25, 0, 10, 1, 64, 7, 10, 8, 10, 18, 10, 19, 10,
	27, 10, 28, 10, 59, 10, 0, 11, 7, 11, 8, 11, 18, 11, 19, 11, 27, 11, 28, 11,
	59, 11,
}

var tmRuleLen = []int8{
	2, 1, 1, 1, 1, 5, 4, 4, 3, 3, 2, 1, 3, 1, 4, 4, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 0, 0, 1, 3, 4, 1,
}

var tmRuleSymbol = []int32{
	26, 26, 27, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 30, 30, 31, 32, 32, 33, 34,
}

var tmRuleType = [...]NodeType{
	0,                              // Declaration_list : Declaration_list Declaration
	0,                              // Declaration_list : Declaration
	Test,                           // Test : Declaration_list
	0,                              // Declaration : Decl1
	0,                              // Declaration : Decl2
	Block,                          // Declaration : '{' '-' '-' Declaration_list '}'
	Block,                          // Declaration : '{' '-' '-' '}'
	Block,                          // Declaration : '{' '-' Declaration_list '}'
	Block,                          // Declaration : '{' '-' '}'
	Block,                          // Declaration : '{' Declaration_list '}'
	Block,                          // Declaration : '{' '}'
	LastInt,                        // Declaration : lastInt
	Int,                            // Declaration : IntegerConstant '[' ']'
	Int,                            // Declaration : IntegerConstant
	TestClause,                     // Declaration : 'test' '{' setof_not_((eoi | '.') | '}')_optlist '}'
	0,                              // Declaration : 'test' '(' empty1 ')'
	TestIntClause | InTest | InFoo, // Declaration : 'test' IntegerConstant
	0,                              // setof_not_((eoi | '.') | '}') : invalid_token
	0,                              // setof_not_((eoi | '.') | '}') : WhiteSpace
	0,                              // setof_not_((eoi | '.') | '}') : SingleLineComment
	0,                              // setof_not_((eoi | '.') | '}') : Identifier
	0,                              // setof_not_((eoi | '.') | '}') : IntegerConstant
	0,                              // setof_not_((eoi | '.') | '}') : lastInt
	0,                              // setof_not_((eoi | '.') | '}') : 'test'
	0,                              // setof_not_((eoi | '.') | '}') : 'decl1'
	0,                              // setof_not_((eoi | '.') | '}') : 'decl2'
	0,                              // setof_not_((eoi | '.') | '}') : '{'
	0,                              // setof_not_((eoi | '.') | '}') : '('
	0,                              // setof_not_((eoi | '.') | '}') : ')'
	0,                              // setof_not_((eoi | '.') | '}') : '['
	0,                              // setof_not_((eoi | '.') | '}') : ']'
	0,                              // setof_not_((eoi | '.') | '}') : ','
	0,                              // setof_not_((eoi | '.') | '}') : ':'
	0,                              // setof_not_((eoi | '.') | '}') : '-'
	0,                              // setof_not_((eoi | '.') | '}') : '->'
	0,                              // setof_not_((eoi | '.') | '}') : SharpAtID
	0,                              // setof_not_((eoi | '.') | '}') : 'Zfoo'
	0,                              // setof_not_((eoi | '.') | '}') : backtrackingToken
	0,                              // setof_not_((eoi | '.') | '}') : error
	0,                              // setof_not_((eoi | '.') | '}') : MultiLineComment
	0,                              // setof_not_((eoi | '.') | '}')_optlist : setof_not_((eoi | '.') | '}')_optlist setof_not_((eoi | '.') | '}')
	0,                              // setof_not_((eoi | '.') | '}')_optlist :
	0,                              // empty1 :
	0,                              // QualifiedName : Identifier
	0,                              // QualifiedName : QualifiedName '.' Identifier
	Decl1,                          // Decl1 : 'decl1' '(' QualifiedName ')'
	Decl2,                          // Decl2 : 'decl2'
}

// set(follow error) =
var afterErr = []int32{}
