// generated by Textmapper; DO NOT EDIT

package tm

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the tm language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
	EOI

	REGEXP
	SCON
	ICON
	WHITESPACE
	COMMENT
	MULTILINE_COMMENT
	REM              // %
	COLONCOLONASSIGN // ::=
	COLONCOLON       // ::
	OR               // |
	OROR             // ||
	ASSIGN           // =
	ASSIGNASSIGN     // ==
	EXCLASSIGN       // !=
	ASSIGNGT         // =>
	SEMICOLON        // ;
	DOT              // .
	COMMA            // ,
	COLON            // :
	LBRACK           // [
	RBRACK           // ]
	LPAREN           // (
	RPAREN           // )
	LBRACETILDE      // {~
	RBRACE           // }
	LT               // <
	GT               // >
	MULT             // *
	PLUS             // +
	PLUSASSIGN       // +=
	QUEST            // ?
	EXCL             // !
	TILDE            // ~
	AND              // &
	ANDAND           // &&
	DOLLAR           // $
	ATSIGN           // @
	ERROR
	INVALID_TOKEN
	ID
	AS         // as
	FALSE      // false
	IMPORT     // import
	SEPARATOR  // separator
	SET        // set
	TRUE       // true
	ASSERT     // assert
	BRACKETS   // brackets
	CLASS      // class
	EMPTY      // empty
	EXPLICIT   // explicit
	FLAG       // flag
	GENERATE   // generate
	GLOBAL     // global
	INLINE     // inline
	INPUT      // input
	INTERFACE  // interface
	LALR       // lalr
	LANGUAGE   // language
	LAYOUT     // layout
	LEFT       // left
	LEXER      // lexer
	LOOKAHEAD  // lookahead
	NOMINUSEOI // no-eoi
	NONASSOC   // nonassoc
	NONEMPTY   // nonempty
	PARAM      // param
	PARSER     // parser
	PREC       // prec
	RETURNS    // returns
	RIGHT      // right
	SHIFT      // shift
	SOFT       // soft
	SPACE      // space
	VOID       // void
	CODE       // {
	LBRACE     // {

	NumTokens
)

var tokenStr = [...]string{
	"EOI",

	"REGEXP",
	"SCON",
	"ICON",
	"WHITESPACE",
	"COMMENT",
	"MULTILINE_COMMENT",
	"%",
	"::=",
	"::",
	"|",
	"||",
	"=",
	"==",
	"!=",
	"=>",
	";",
	".",
	",",
	":",
	"[",
	"]",
	"(",
	")",
	"{~",
	"}",
	"<",
	">",
	"*",
	"+",
	"+=",
	"?",
	"!",
	"~",
	"&",
	"&&",
	"$",
	"@",
	"ERROR",
	"INVALID_TOKEN",
	"ID",
	"as",
	"false",
	"import",
	"separator",
	"set",
	"true",
	"assert",
	"brackets",
	"class",
	"empty",
	"explicit",
	"flag",
	"generate",
	"global",
	"inline",
	"input",
	"interface",
	"lalr",
	"language",
	"layout",
	"left",
	"lexer",
	"lookahead",
	"no-eoi",
	"nonassoc",
	"nonempty",
	"param",
	"parser",
	"prec",
	"returns",
	"right",
	"shift",
	"soft",
	"space",
	"void",
	"{",
	"{",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
