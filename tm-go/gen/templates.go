package gen

type file struct {
	name     string
	template string
}

var lexerFiles = []file{
	{"token/token.go", tokenTpl},
	{"lexer_tables.go", lexerTablesTpl},
	{"lexer.go", lexerTpl},
}

var parserFiles = []file{
	{"parser.go", parserTpl},
	{"parser_tables.go", parserTablesTpl},
}

var listenerFile = file{"listener.go", parserListenerTpl}

const sharedDefs = `
{{define "header" -}}
// generated by Textmapper; DO NOT EDIT

{{end}}

`

const tokenTpl = `
{{- template "header" . -}}
package token

// Token is an enum of all terminal symbols of the {{.Name}} language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
{{- range .Tokens}}
	{{.ID}}{{if .Comment}}  // {{.Comment}}{{end}}
{{- end}}

	NumTokens
)

var tokenStr = [...]string{
{{- range .Tokens}}
	{{if .Comment}}{{str_literal .Comment}}{{else}}{{str_literal .ID}}{{end}},
{{- end}}
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return "fmt".Sprintf("token(%d)", tok)
}
`

const lexerTablesTpl = `
{{- define "tokenType"}}"{{.Options.Package}}/token".Token{{end -}}
{{- template "header" . -}}
package {{.Name}}

const tmNumClasses = {{.Lexer.Tables.NumSymbols}}

{{$runeType := bits .Lexer.Tables.NumSymbols -}}
{{if gt .Lexer.Tables.LastMapEntry.Start 2048 -}}
type mapRange struct {
	lo         rune
	hi         rune
	defaultVal uint{{$runeType}}
	val        []uint{{$runeType}}
}

func mapRune(c rune) int {
	lo := 0
	hi := len(tmRuneRanges)
	for lo < hi {
		m := lo + (hi-lo)/2
		r := tmRuneRanges[m]
		if c < r.lo {
			hi = m
		} else if c >= r.hi {
			lo = m + 1
		} else {
			i := int(c - r.lo)
			if i < len(r.val) {
				return int(r.val[i])
			}
			return int(r.defaultVal)
		}
	}
	return 1
}

// Latin-1 characters.
var tmRuneClass = []uint{{$runeType}}{
{{- int_array (.Lexer.Tables.SymbolArr 256) "\t" 79 -}}
}

const tmRuneClassLen = 256
const tmFirstRule = {{.Lexer.Tables.ActionStart}}

var tmRuneRanges = []mapRange{
{{range .Lexer.Tables.CompressedMap 256}}	{ {{- .Lo}}, {{.Hi}}, {{.DefaultVal}}, {{if .Vals}}[]uint{{$runeType}}{
{{- int_array .Vals "\t\t" 78}}	}{{else}}nil{{end -}} },
{{end -}}
}

{{else -}}
{{ $runeArr := .Lexer.Tables.SymbolArr 0 -}}
var tmRuneClass = []uint{{$runeType}}{
{{- int_array $runeArr "\t" 79 -}}
}

const tmRuneClassLen = {{len $runeArr}}
const tmFirstRule = {{.Lexer.Tables.ActionStart}}

{{end -}}
var tmStateMap = []int{
{{- int_array .Lexer.Tables.StateMap "\t" 79 -}}
}

{{if .Lexer.RuleToken -}}
var tmToken = []{{template "tokenType" .}}{
{{- int_array .Lexer.RuleToken "\t" 79 -}}
}

{{end -}}
var tmLexerAction = []int{{bits_per_element .Lexer.Tables.Dfa}}{
{{- int_array .Lexer.Tables.Dfa "\t" 79 -}}
}

{{- if .Lexer.Tables.Backtrack}}

var tmBacktracking = []int{
{{- range .Lexer.Tables.Backtrack}}
	{{.Action}}, {{.NextState}},{{if .Details}} // {{.Details}}{{end}}
{{- end}}
}
{{- end}}
`

const lexerTpl = `
{{- template "header" . -}}
package {{.Name}}

{{- if gt (len .Lexer.StartConditions) 1}}

// Lexer states.
const (
{{- range $index, $el := .Lexer.StartConditions}}
	State{{title .}} = {{$index}}
{{- end}}
)
{{- end}}
{{block "onBeforeLexer" .}}{{end}}
{{template "lexerType" .}}
{{template "lexerInit" .}}
{{template "lexerNext" .}}
{{template "lexerPos" .}}
{{- if .Options.TokenLine}}
{{template "lexerLine" .}}
{{- end}}
{{- if .Options.TokenColumn}}
{{template "lexerColumn" .}}
{{- end}}
{{template "lexerText" .}}
{{template "lexerValue" .}}
{{template "lexerRewind" .}}
{{- block "onAfterLexer" .}}{{end}}

{{- define "tokenType"}}"{{.Options.Package}}/token".Token{{end -}}
{{- define "tokenPkg"}}"{{.Options.Package}}/token".{{end -}}

{{- define "lexerType" -}}
// Lexer uses a generated DFA to scan through a utf-8 encoded input string. If
// the string starts with a BOM character, it gets skipped.
type Lexer struct {
	source string

	ch          rune // current character, -1 means EOI
	offset      int  // character offset
	tokenOffset int  // last token offset
{{- if .Options.TokenLine}}
	line        int  // current line number (1-based)
	tokenLine   int  // last token line
{{- end}}
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
	lineOffset  int  // current line offset
{{- end}}
{{- if .Options.TokenColumn}}
	tokenColumn int  // last token column (in bytes)
{{- end}}
	scanOffset  int  // scanning offset
	value       interface{}

	State int // lexer state, modifiable
{{ block "stateVars" .}}{{end -}}
}
{{end -}}

{{- define "lexerInit" -}}
var bomSeq = "\xef\xbb\xbf"

// Init prepares the lexer l to tokenize source by performing the full reset
// of the internal state.
func (l *Lexer) Init(source string) {
	l.source = source

	l.ch = 0
	l.offset = 0
	l.tokenOffset = 0
{{- if .Options.TokenLine}}
	l.line = 1
	l.tokenLine = 1
{{- end}}
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
	l.lineOffset = 0
{{- end}}
{{- if .Options.TokenColumn}}
	l.tokenColumn = 1
{{- end}}
	l.State = 0
{{ block "initStateVars" .}}{{end}}
	if "strings".HasPrefix(source, bomSeq) {
		l.offset += len(bomSeq)
	}

	l.rewind(l.offset)
}
{{end -}}

{{- define "lexerNext" -}}
// Next finds and returns the next token in l.source. The source end is
// indicated by Token.EOI.
//
// The token text can be retrieved later by calling the Text() method.
func (l *Lexer) Next() {{template "tokenType" .}} {
{{ block "onBeforeNext" .}}{{end -}}
{{ $spaceRules := .SpaceActions -}}
{{ if or $spaceRules .Lexer.RuleToken -}}
restart:
{{ end -}}
{{ if .Options.TokenLine}}	l.tokenLine = l.line
{{ end -}}
{{ if .Options.TokenColumn}}	l.tokenColumn = l.offset-l.lineOffset+1
{{ end -}}
	l.tokenOffset = l.offset

	state := tmStateMap[l.State]
{{- if .Lexer.ClassActions}}
	hash := uint32(0)
{{- end}}
{{- if .Lexer.Tables.Backtrack}}
	backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} := -1
	var backupOffset int
{{- if .Lexer.ClassActions}}
	backupHash := hash
{{- end}}
{{- end}}
	for state >= 0 {
		var ch int
		if uint(l.ch) < tmRuneClassLen {
			ch = int(tmRuneClass[l.ch])
		} else if l.ch < 0 {
			state = int(tmLexerAction[state*tmNumClasses])
{{- if .Lexer.Tables.Backtrack}}
			if state > tmFirstRule && state < 0 {
				state = (-1 - state) * 2
				backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} = tmBacktracking[state]
				backupOffset = l.offset
{{- if .Lexer.ClassActions}}
				backupHash = hash
{{- end}}
				state = tmBacktracking[state+1]
			}
{{- end}}
			continue
		} else {
{{- if gt .Lexer.Tables.LastMapEntry.Start 2048}}
			ch = mapRune(l.ch)
{{- else}}
			ch = 1
{{- end}}
		}
		state = int(tmLexerAction[state*tmNumClasses+ch])
		if state > tmFirstRule {
{{- if .Lexer.Tables.Backtrack}}
			if state < 0 {
				state = (-1 - state) * 2
				backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} = tmBacktracking[state]
				backupOffset = l.offset
{{- if .Lexer.ClassActions}}
				backupHash = hash
{{- end}}
				state = tmBacktracking[state+1]
			}
{{- end}}
{{- if .Lexer.ClassActions}}
			hash = hash*uint32(31) + uint32(l.ch)
{{end}}
{{- if .Options.TokenLine}}
			if l.ch == '\n' {
				l.line++
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
				l.lineOffset = l.offset
{{- end}}
			}
{{end}}
			// Scan the next character.
			// Note: the following code is inlined to avoid performance implications.
			l.offset = l.scanOffset
			if l.offset < len(l.source) {
				r, w := rune(l.source[l.offset]), 1
				if r >= 0x80 {
					// not ASCII
					r, w = "unicode/utf8".DecodeRuneInString(l.source[l.offset:])
				}
				l.scanOffset += w
				l.ch = r
			} else {
				l.ch = -1 // EOI
			}
		}
	}
{{if .Lexer.RuleToken}}
	rule := tmFirstRule - state
{{- else}}
	tok := {{template "tokenType" .}}(tmFirstRule - state)
{{- end}}
{{- if .Lexer.Tables.Backtrack}}
recovered:
{{- end}}
{{- if .Lexer.ClassActions}}
	switch {{if .Lexer.RuleToken}}rule{{else}}tok{{end}} {
{{- range .Lexer.ClassActions}}
{{- if $.Lexer.RuleToken}}
	case {{sum .Action 2}}:
{{- else}}
	case {{template "tokenPkg" $}}{{(index $.Syms .Action).ID}}:
{{- end}}
{{- with string_switch .Custom }}
		hh := hash & {{.Mask}}
		switch hh {
{{- range .Cases}}
		case {{.Value}}:
{{- range .Subcases}}
			if hash == {{hex .Hash}} && {{quote .Str}} == l.source[l.tokenOffset:l.offset] {
{{- if $.Lexer.RuleToken}}
				rule = {{sum .Action 2}}
{{- else}}
				tok = {{template "tokenPkg" $}}{{(index $.Syms .Action).ID}}
{{- end}}
				break
			}
{{- end}}
{{- end}}
		}
{{- end}}
{{- end}}
	}
{{- end}}
{{- if .Lexer.RuleToken}}

	tok := tmToken[rule]
	var space bool
{{- if .Lexer.Actions}}
	switch rule {
	case 0:
{{- template "handleInvalidToken" .}}
{{- range .Lexer.Actions}}
	case {{sum .Action 2}}:{{if .Comments}} // {{join .Comments ", "}}{{end}}
{{- if .Space }}
		space = true
{{- end}}
{{- if .Code }}
{{lexer_action .Code}}
{{- end}}
{{- end}}
	}
{{- else}}
	if rule == 0 {
{{- template "handleInvalidToken" .}}
	}
{{- end}}
	if space {
		goto restart
	}
{{- else}}
	switch tok {
	case {{template "tokenPkg" .}}{{(index $.Syms .Lexer.InvalidToken).ID}}:
{{- template "handleInvalidToken" .}}
{{- if $spaceRules}}
	case {{range $i, $val := $spaceRules}}{{if gt $i 0}}, {{end}}{{$val}}{{end}}:
		goto restart
{{- end}}
	}
{{- end}}
{{ block "onAfterNext" .}}{{end -}}
{{/**/}}	return tok
}
{{end -}}

{{- define "lexerPos" -}}
// Pos returns the start and end positions of the last token returned by Next().
func (l *Lexer) Pos() (start, end int) {
	start = l.tokenOffset
	end = l.offset
	return
}
{{end -}}

{{- define "lexerLine" -}}
// Line returns the line number of the last token returned by Next() (1-based).
func (l *Lexer) Line() int {
	return l.tokenLine
}
{{end -}}

{{- define "lexerColumn" -}}
// Column returns the column of the last token returned by Next() (in bytes, 1-based).
func (l *Lexer) Column() int {
	return l.tokenColumn
}
{{end -}}

{{- define "lexerText" -}}
// Text returns the substring of the input corresponding to the last token.
func (l *Lexer) Text() string {
	return l.source[l.tokenOffset:l.offset]
}
{{end -}}

{{- define "lexerValue" -}}
// Value returns the value associated with the last returned token.
func (l *Lexer) Value() interface{} {
	return l.value
}
{{end -}}

{{- define "lexerRewind" -}}
// rewind can be used in lexer actions to accept a portion of a scanned token, or to include
// more text into it.
func (l *Lexer) rewind(offset int) {
{{- if .Options.TokenLine}}
	if offset < l.offset {
		l.line -= "strings".Count(l.source[offset:l.offset], "\n")
	} else {
		if offset > len(l.source) {
			offset = len(l.source)
		}
		l.line += "strings".Count(l.source[l.offset:offset], "\n")
	}
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
	l.lineOffset = 1 + "strings".LastIndexByte(l.source[:offset], '\n')
{{- end}}
{{end}}
	// Scan the next character.
	l.scanOffset = offset
	l.offset = offset
	if l.offset < len(l.source) {
		r, w := rune(l.source[l.offset]), 1
		if r >= 0x80 {
			// not ASCII
			r, w = "unicode/utf8".DecodeRuneInString(l.source[l.offset:])
		}
		l.scanOffset += w
		l.ch = r
	} else {
		l.ch = -1 // EOI
	}
}
{{end -}}

{{- define "handleInvalidToken" -}}
{{if .Lexer.Tables.Backtrack}}
		if backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} >= 0 {
{{- if .Lexer.RuleToken}}
			rule = backupRule
{{- else}}
			tok = {{template "tokenType" .}}(backupToken)
{{- end}}
{{- if .Lexer.ClassActions}}
			hash = backupHash
{{- end}}
			l.rewind(backupOffset)
		} else if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
{{- if .Lexer.RuleToken}}
		if rule != 0 {
{{- else}}
		if tok != {{template "tokenPkg" .}}{{(index $.Syms .Lexer.InvalidToken).ID}} {
{{- end}}
			goto recovered
		}
{{- else}}
		if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
{{- end -}}
{{end -}}
`

const bisonTpl = `%{
%}
{{range .Parser.Inputs}}
%start {{(index $.Parser.Nonterms .Nonterm).Name}}{{if .NoEoi}} // no-eoi{{end}}
{{- end}}
{{range .Parser.Prec}}
%{{.Associativity}}{{range .Terminals}} {{(index $.Syms .).ID}}{{end}}
{{- end}}
{{- range slice .TokensWithoutPrec 1}}
%token {{.ID}}
{{- end}}

%%
{{- range .Parser.Nonterms}}

{{ if eq .Value.Kind 11 -}}
// lookahead: {{ range $i, $it := .Value.Sub }}{{if gt $i 0}} & {{end}}{{$it}}{{end}}
{{ end -}}
{{.Name}} :
{{- if eq .Value.Kind 2 }}
{{- range $i, $rule := .Value.Sub}}
{{ if eq $i 0}}  {{else}}| {{end}}{{$.ExprString $rule}}
{{- end}}
{{- else if eq .Value.Kind 11 }}
  %empty
{{- else }}
  {{$.ExprString .Value}}
{{- end }}
;
{{- end}}

%%

`

const parserTpl = `
{{- define "flushPending"}}
{{- if .ReportTokens true }}
	if len(p.pending) > 0 {
		for _, tok := range p.pending {
			p.reportIgnoredToken(tok)
		}
		p.pending = p.pending[:0]
	}
{{- end}}
{{- end}}

{{- define "tokenType"}}"{{.Options.Package}}/token".Token{{end -}}
{{- define "tokenPkg"}}"{{.Options.Package}}/token".{{end -}}

{{- define "reportConsumedNext" -}}
{{- if .ReportTokens false }}
			switch {{template "tokenType" .}}(p.next.symbol) {
{{- range .Parser.MappedTokens}}
{{- $sym := index $.Syms .Token}}
{{- if not (or $sym.Space (eq $sym.Name "invalid_token")) }}
	case {{template "tokenPkg" $}}{{$sym.ID}}:
		p.listener({{.Name}}, {{if $.Parser.UsesFlags}}0, {{end}}p.next.offset, p.next.endoffset)
{{- end}}
{{- end}}
			}
{{- end}}
{{- end}}


{{- template "header" . -}}
package {{.Name}}

{{- if .Parser.IsRecovering }}
{{- block "errorHandler" .}}

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// StopOnFirstError is an error handler that forces the parser to stop on and return the first
// error.
func StopOnFirstError(_ SyntaxError) bool { return false }
{{- end }}
{{- end }}

{{$stateType := bits_per_element .Parser.Tables.FromTo -}}
{{ if .Options.IsEnabled "Parser" -}}
// Parser is a table-driven LALR parser for {{.Name}}.
type Parser struct {
{{- if .Parser.IsRecovering }}
	eh ErrorHandler
{{- end}}
{{- if .Parser.Types }}
	listener Listener
{{- end}}

	next symbol
{{- if .Parser.IsRecovering }}
	endState  int{{$stateType}}
{{- end}}

{{- if .ReportTokens true }}

	// Tokens to be reported with the next shift. Only non-empty when next.symbol != noToken.
	pending []symbol
{{- end }}
{{ block "parserVars" .}}{{end -}}
}

{{ end -}}
{{ block "syntaxError" . -}}
type SyntaxError struct {
{{- if .Options.TokenLine }}
	Line      int
{{- end }}
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
{{- if .Options.TokenLine }}
	return "fmt".Sprintf("syntax error at line %v", e.Line)
{{- else}}
	return "syntax error"
{{- end }}
}

{{ end -}}

{{ if .Options.IsEnabled "symbol" -}}
type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

{{ end -}}
{{ if .Options.IsEnabled "stackEntry" -}}
type stackEntry struct {
	sym   symbol
	state int{{$stateType}}
{{- if .Parser.HasAssocValues }}
	value interface{}
{{- end}}
}

{{ end -}}
{{ if .Options.IsEnabled "ParserInit" -}}
func (p *Parser) Init({{if .Parser.IsRecovering }}eh ErrorHandler{{end}}{{if .Parser.Types }}{{if .Parser.IsRecovering }}, {{end}}l Listener{{end}}) {
{{- if .Parser.IsRecovering }}
	p.eh = eh
{{- end}}
{{- if .Parser.Types }}
	p.listener = l
{{- end}}
{{- if .ReportTokens true }}
	if cap(p.pending) < startTokenBufferSize {
		p.pending = make([]symbol, 0, startTokenBufferSize)
	}
{{- end}}
{{ block "initParserVars" .}}{{end -}}
}

{{ end -}}
const (
	startStackSize = 256
{{- if .ReportTokens true }}
	startTokenBufferSize = 16
{{- end}}
	noToken        = int32({{template "tokenPkg" .}}UNAVAILABLE)
	eoiToken       = int32({{template "tokenPkg" .}}EOI)
	debugSyntax    = {{ .Options.DebugParser }}
)

{{ range $index, $inp := .Parser.Inputs -}}
{{ if $inp.Synthetic }}{{continue}}{{end -}}
{{ $nt := index $.Parser.Nonterms $inp.Nonterm -}}
func (p *Parser) Parse{{if $.Parser.HasMultipleUserInputs}}{{$.NontermID $inp.Nonterm}}{{end}}({{if $.Options.Cancellable}}ctx "context".Context, {{end}}lexer *Lexer) {{if eq $nt.Type ""}}error{{else}}({{$nt.Type}}, error){{end}} {
{{- if $.Parser.HasInputAssocValues}}
	{{if ne $nt.Type ""}}v{{else}}_{{end}}, err := p.parse({{if $.Options.Cancellable}}ctx, {{end}}{{$index}}, {{index $.Parser.Tables.FinalStates $index}}, lexer)
{{- if ne $nt.Type ""}}
	val, _ := v.({{$nt.Type}})
	return val, err
{{- else}}
	return err
{{- end}}
{{- else}}
	return p.parse({{if $.Options.Cancellable}}ctx, {{end}}{{$index}}, {{index $.Parser.Tables.FinalStates $index}}, lexer)
{{- end}}
}

{{end -}}
{{ block "session" . -}}
{{ if and .NeedsSession (.Options.IsEnabled "session") -}}
type session struct {
{{- if $.Options.Cancellable}}
	shiftCounter int32
{{- end }}
{{- if .Options.RecursiveLookaheads }}
	cache map[uint64]bool
{{- end }}
}

{{- end}}
{{- end}}
{{ block "parseFunc" . -}}
{{ $stateType := bits_per_element .Parser.Tables.FromTo -}}
{{ if .Options.IsEnabled "parse" -}}
func (p *Parser) parse({{if $.Options.Cancellable}}ctx "context".Context, {{end}}start, end int{{$stateType}}, lexer *Lexer) {{if .Parser.HasInputAssocValues}}(interface{}, error){{else}}error{{end}} { 
{{- if .ReportTokens true }}
	p.pending = p.pending[:0]
{{- end}}
{{- if .NeedsSession}}
	var s session
{{- if .Options.RecursiveLookaheads }}
	s.cache = make(map[uint64]bool)
{{- end}}
{{- else if .Options.Cancellable }}
	var shiftCounter int
{{- end}}
	state := start
{{- if .Parser.IsRecovering }}
	var lastErr SyntaxError
	recovering := 0
{{- end}}

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
{{- if .Parser.IsRecovering }}
	p.endState = end
{{- end}}
	p.fetchNext(lexer, stack)

	for state != end {
		action := tmAction[state]
{{- if .Parser.Tables.Lalr}}
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			action = lalr(action, p.next.symbol)
		}
{{- end}}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			rhs := stack[len(stack)-ln:]
			stack = stack[:len(stack)-ln]
			if ln == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				entry.sym.offset, entry.sym.endoffset = p.next.offset, p.next.offset
			} else {
				entry.sym.offset = rhs[0].sym.offset
				entry.sym.endoffset = rhs[ln-1].sym.endoffset
			}
			if err := p.applyRule({{if .Options.Cancellable}}ctx, {{end}}rule, &entry, rhs, lexer{{if .NeedsSession}}, &s{{end}}); err != nil {
				return {{if .Parser.HasInputAssocValues}}nil, {{end}}err
			}
			if debugSyntax {
				"fmt".Printf("reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
{{- if .Options.Cancellable }}
			if {{if .NeedsSession}}s.{{end}}shiftCounter++; {{if .NeedsSession}}s.{{end}}shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return {{if .Parser.HasInputAssocValues}}nil, {{end}}ctx.Err()
				default:
				}
			}
{{end}}
			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			state = gotoState(state, p.next.symbol)
			if state >= 0 {
				stack = append(stack, stackEntry{
					sym:   p.next,
					state: state,
{{- if .Parser.HasAssocValues }}
					value: lexer.Value(),
{{- end}}
				})
				if debugSyntax {
					"fmt".Printf("shift: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
				}
{{- block "onAfterShift" .}}{{end}}
{{- template "flushPending" .}}
				if p.next.symbol != eoiToken {
{{- template "reportConsumedNext" .}}
					p.next.symbol = noToken
				}
{{- if .Parser.IsRecovering }}
				if recovering > 0 {
					recovering--
				}
{{- end}}
			}
		}

		if action == -2 || state == -1 {
{{- if .Parser.IsRecovering }}
			if recovering == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				lastErr = SyntaxError{
{{- if .Options.TokenLine}}
					Line:      lexer.Line(),
{{- end}}
					Offset:    p.next.offset,
					Endoffset: p.next.endoffset,
				}
				if !p.eh(lastErr) {
{{- template "flushPending" .}}
					return {{if .Parser.HasInputAssocValues}}nil, {{end}}lastErr
				}
			}

			if stack = p.recoverFromError(lexer, stack); stack == nil {
{{- template "flushPending" .}}
				return {{if .Parser.HasInputAssocValues}}nil, {{end}}lastErr
			}
			state = stack[len(stack)-1].state
			recovering = 4
{{- else}}
			break
{{- end}}
		}
	}

{{- if not .Parser.IsRecovering }}

	if state != end {
		if p.next.symbol == noToken {
			p.fetchNext(lexer, stack)
		}
		err := SyntaxError{
{{- if .Options.TokenLine}}
			Line:      lexer.Line(),
{{- end}}
			Offset:    p.next.offset,
			Endoffset: p.next.endoffset,
		}
		return {{if .Parser.HasInputAssocValues}}nil, {{end}}err
	}
{{- end}}

	return {{if .Parser.HasInputAssocValues}}stack[len(stack)-2].value, {{end}}nil
}

{{ end -}}
{{ end -}}

{{ if .Parser.IsRecovering -}}
const errSymbol = {{ .Parser.ErrorSymbol }}

{{ block "willShift" . -}}
{{$stateType := bits_per_element .Parser.Tables.FromTo -}}
// willShift checks if "symbol" is going to be shifted in the given state.
// This function does not support empty productions and returns false if they occur before "symbol".
func (p *Parser) willShift(stackPos int, state int{{$stateType}}, symbol int32, stack []stackEntry) bool {
	if state == -1 {
		return false
	}

	for state != p.endState {
		action := tmAction[state]
{{- if .Parser.Tables.Lalr}}
		if action < -2 {
			action = lalr(action, symbol)
		}
{{- end}}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])
			if ln == 0 {
				// we do not support empty productions
				return false
			}
			stackPos -= ln - 1
			state = gotoState(stack[stackPos-1].state, tmRuleSymbol[rule])
		} else {
			return action == -1 && gotoState(state, symbol) >= 0
		}
	}
	return symbol == eoiToken
}
{{ end }}
{{ block "skipBrokenCode" . -}}
func (p *Parser) skipBrokenCode(lexer *Lexer, stack []stackEntry, canRecover func (symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		if debugSyntax {
			"fmt".Printf("skipped while recovering: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
		}
{{- template "flushPending" .}}
{{- template "reportConsumedNext" .}}
		e = p.next.endoffset
		p.fetchNext(lexer, stack)
	}
	return e
}
{{ end }}
{{ block "recoverFromError" . -}}
{{ if .Options.IsEnabled "recoverFromError" -}}
func (p *Parser) recoverFromError(lexer *Lexer, stack []stackEntry) []stackEntry {
	var recoverSyms [1 + {{template "tokenPkg" .}}NumTokens/8]uint8
	var recoverPos []int

	if debugSyntax {
		"fmt".Printf("broke at %v\n", symbolName(p.next.symbol))
	}
	for size := len(stack); size > 0; size-- {
		if gotoState(stack[size-1].state, errSymbol) == -1 {
			continue
		}
		recoverPos = append(recoverPos, size)
{{- range .Parser.Tables.Markers}}
{{- if eq (lower .Name) "recoveryscope" }}
{{- if eq (len .States) 1}}
		if {{.Name}}State == stack[size-1].state {
			break
		}
{{- else}}
		if {{.Name}}States[int(stack[size-1].state)] {
			break
		}
{{- end}}
{{- end}}
{{- end}}
	}
	if len(recoverPos) == 0 {
		return nil
	}

	for _, v := range afterErr {
		recoverSyms[v/8] |= 1 << uint32(v%8)
	}
	canRecover := func (symbol int32) bool {
		return recoverSyms[symbol/8]&(1<<uint32(symbol%8)) != 0
	}
	if p.next.symbol == noToken {
		p.fetchNext(lexer, stack)
	}
	// By default, insert 'error' in front of the next token.
	s := p.next.offset
	e := s
{{- if .ReportsInvalidToken}}
	for _, tok := range p.pending {
		// Try to cover all nearby invalid tokens.
		if {{template "tokenType" .}}(tok.symbol) == {{template "tokenPkg" .}}{{(index .Syms .Lexer.InvalidToken).ID}} {
			if s > tok.offset {
				s = tok.offset
			}
			e = tok.endoffset
		}
	}
{{- end}}
	for {
		if endoffset := p.skipBrokenCode(lexer, stack, canRecover); endoffset > e {
			e = endoffset
		}

		var matchingPos int
		if debugSyntax {
			"fmt".Printf("trying to recover on %v\n", symbolName(p.next.symbol))
		}
		for _, pos := range recoverPos {
			if p.willShift(pos, gotoState(stack[pos-1].state, errSymbol), p.next.symbol, stack) {
				matchingPos = pos
				break
			}
		}
		if matchingPos == 0 {
			if p.next.symbol == eoiToken {
				return nil
			}
			recoverSyms[p.next.symbol/8] &^= 1 << uint32(p.next.symbol%8)
			continue
		}

		if matchingPos < len(stack) {
			if s == e {
				// Avoid producing syntax problems covering trailing whitespace.
				e = stack[len(stack)-1].sym.endoffset
			}
			s = stack[matchingPos].sym.offset
{{- if .ReportTokens true }}
		} else if s == e && len(p.pending) > 0 {
			// This means pending tokens don't contain InvalidTokens.
			for _, tok := range p.pending {
				p.reportIgnoredToken(tok)
			}
			p.pending = p.pending[:0]
{{- end}}
		}
{{- if .ReportsInvalidToken}}
		if s != e {
			// Consume trailing invalid tokens.
			for _, tok := range p.pending {
				if {{template "tokenType" .}}(tok.symbol) == {{template "tokenPkg" .}}{{(index .Syms .Lexer.InvalidToken).ID}} && tok.endoffset > e {
					e = tok.endoffset
				}
			}
			var consumed int
			for ; consumed < len(p.pending); consumed++ {
				tok := p.pending[consumed]
				if tok.offset >= e {
					break
				}
				p.reportIgnoredToken(tok)
			}
			newSize := len(p.pending) - consumed
			copy(p.pending[:newSize], p.pending[consumed:])
			p.pending = p.pending[:newSize]
		}
{{- end}}
		if debugSyntax {
			for i := len(stack)-1; i >= matchingPos; i-- {
				"fmt".Printf("dropped from stack: %v\n", symbolName(stack[i].sym.symbol))
			}
			"fmt".Println("recovered")
		}
		stack = append(stack[:matchingPos], stackEntry{
			sym:   symbol{errSymbol, s, e},
			state: gotoState(stack[matchingPos-1].state, errSymbol),
		})
		return stack
	}
}

{{ end -}}
{{ end -}}
{{ end -}}

{{ block "lalr" . -}}
{{ if .Parser.Tables.Lalr -}}
func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

{{end -}}
{{end -}}

{{ block "gotoState" . -}}
{{$stateType := bits_per_element .Parser.Tables.FromTo -}}
func gotoState(state int{{$stateType}}, symbol int32) int{{$stateType}} {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

{{ end -}}

{{ block "fetchNext" . -}}
{{ if .Options.IsEnabled "fetchNext" -}}
func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry) {
restart:
	tok := lexer.Next()
	switch tok {
{{- if .ReportTokens true }}
	case {{range $ind, $tok := .ReportTokens true}}{{if ne $ind 0}}, {{end}}{{template "tokenPkg" $}}{{.ID}}{{end}}:
		s, e := lexer.Pos()
		tok := symbol{int32(tok), s, e}
		p.pending = append(p.pending, tok)
		goto restart
{{- end}}
{{- if not .ReportsInvalidToken}}
	case {{template "tokenPkg" .}}{{(index .Syms .Lexer.InvalidToken).ID}}:
		goto restart
{{- end}}
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = lexer.Pos()
}

{{ end -}}
{{ end -}}

{{ block "lookahead" . -}}
{{ if and .Parser.Tables.Lookaheads (.Options.IsEnabled "lookaheadNext") -}}
func lookaheadNext(lexer *Lexer) int32 {
restart:
	tok := lexer.Next()
	switch tok {
{{- if .ReportTokens true }}
	case {{range $ind, $tok := .ReportTokens true}}{{if ne $ind 0}}, {{end}}{{template "tokenPkg" $}}{{.ID}}{{end}}:
		goto restart
{{- end}}
{{- if not .ReportsInvalidToken}}
	case {{template "tokenPkg" .}}{{(index .Syms .Lexer.InvalidToken).ID}}:
		goto restart
{{- end}}
	}
	return int32(tok)
}

{{ end -}}
{{ end -}}

{{ block "lookaheadRule" . -}}
{{ if and .Parser.Tables.Lookaheads .Options.RecursiveLookaheads -}}
func lookaheadRule({{if $.Options.Cancellable}}ctx "context".Context, {{end}}lexer *Lexer, next, rule int32, s *session) (sym int32{{if $.Options.Cancellable}}, err error{{end}}) {
	switch rule {
{{- range $index, $rule := .Parser.Tables.Lookaheads }}
	case {{sum $index (len $.Parser.Rules)}}:
{{- if $.Options.Cancellable}}
		var ok bool
{{- end}}
		{{ range $rule.Cases }}
		{{- $sym := index $.Syms (sum $.NumTokens (index $.Parser.Inputs .Predicate.Input).Nonterm) -}}
		if {{if $.Options.Cancellable}}ok, err = {{else}}{{if .Predicate.Negated}}!{{end}}{{end -}}
		   lookahead({{if $.Options.Cancellable}}ctx, {{end}}lexer, next, {{.Predicate.Input}}, {{index $.Parser.Tables.FinalStates .Predicate.Input}}{{if $.NeedsSession}}, s{{end}})
		{{- if $.Options.Cancellable}}; {{if .Predicate.Negated}}!{{end}}ok{{end}} {
			sym = {{.Target}} /* {{(index $.Syms .Target).Name}} */
		} else {{end}}{
			sym = {{.DefaultTarget}} /* {{(index $.Syms .DefaultTarget).Name}} */
		}
		return
{{- end}}
	}
	return 0{{if $.Options.Cancellable}}, nil{{end}}
}

{{ end -}}
{{ end -}}

{{ block "lookaheadMethods" . -}}
{{ range $ind, $inp := .Parser.Inputs -}}
{{ if and .Synthetic .NoEoi -}}
{{ $sym := index $.Syms (sum $.NumTokens .Nonterm) -}}
func At{{$sym.Name}}({{if $.Options.Cancellable}}ctx "context".Context, {{end}}lexer *Lexer, next int32{{if $.NeedsSession}}, s *session{{end}}) {{if $.Options.Cancellable}}(bool, error){{else}}bool{{end}} {
	return lookahead({{if $.Options.Cancellable}}ctx, {{end}}lexer, next, {{$ind}}, {{index $.Parser.Tables.FinalStates $ind}}{{if $.NeedsSession}}, s{{end}});
}

{{ end -}}
{{ end -}}
{{ end -}}

{{- define "callLookaheadNext"}}{{/*(memoization)*/}}lookaheadNext(&lexer){{end -}}

{{ block "lookaheadFunc" . -}}
{{ if .Parser.Tables.Lookaheads -}}
{{$stateType := bits_per_element .Parser.Tables.FromTo -}}
func lookahead({{if $.Options.Cancellable}}ctx "context".Context, {{end}}l *Lexer, next int32, start, end int{{$stateType}}{{if $.NeedsSession}}, s *session{{end}}) {{if $.Options.Cancellable}}(bool, error){{else}}bool{{end}} {
{{ block "setupLookaheadLexer" . -}}
{{/**/}}	var lexer Lexer = *l
{{end -}}
{{ if .Options.RecursiveLookaheads }}
	// Use memoization for recursive lookaheads.
	if next == noToken {
		next = {{template "callLookaheadNext" true}}
	}
	key := uint64(l.tokenOffset) + uint64(end)<<40
	if ret, ok := s.cache[key]; ok {
		return ret{{if $.Options.Cancellable}}, nil{{end}}
	}
{{end}}
	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
{{- if .Parser.Tables.Lalr}}
		if action < -2 {
			// Lookahead is needed.
			if next == noToken {
				next = {{template "callLookaheadNext" false}}
			}
			action = lalr(action, next)
		}
{{- end}}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
{{- if .Options.RecursiveLookaheads }}
			sym{{if $.Options.Cancellable}}, err{{end}} := lookaheadRule({{if $.Options.Cancellable}}ctx, {{end}}&lexer, next, rule, s)
{{- if $.Options.Cancellable}}
			if err != nil {
				return false, err
			}
{{- end}}
			if sym != 0 {
				entry.sym.symbol = sym
			}
{{- end}}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
{{- if .Options.Cancellable }}
			if {{if .NeedsSession}}s.{{end}}shiftCounter++; {{if .NeedsSession}}s.{{end}}shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return false, ctx.Err()
				default:
				}
			}
{{end}}
			// Shift.
			if next == noToken {
				next = {{template "callLookaheadNext" false}}
			}
			state = gotoState(state, next)
			stack = append(stack, stackEntry{
				sym:   symbol{symbol: next},
				state: state,
			})
			if state != -1 && next != eoiToken {
				next = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

{{ if .Options.RecursiveLookaheads -}}
	s.cache[key] = state == end
{{ end -}}
	return state == end{{if $.Options.Cancellable}}, nil{{end}}
}

{{ end -}}
{{ end -}}

{{ block "applyRule" . -}}
func (p *Parser) applyRule({{if $.Options.Cancellable}}ctx "context".Context, {{end}}rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer{{if .NeedsSession}}, s *session{{end}}) (err error) {
{{- if or .Parser.HasActions .Parser.Tables.Lookaheads }}
	switch rule {
{{- range $index, $rule := .Parser.Rules}}
{{- $fixWS := and $.Options.FixWhitespace ($.HasTrailingNulls $rule) }}
{{- if or (ne $rule.Action 0) $fixWS }}
{{- $act := index $.Parser.Actions $rule.Action }}
{{- if or (ne $act.Code "") $act.Report $fixWS}}
	case {{$index}}: // {{$.RuleString $rule}}
{{- if $fixWS }}
		fixTrailingWS(lhs, rhs)
{{- end}}
{{- range $act.Report}}
{{- $val := index $.Parser.Types.RangeTypes .Type }}
{{- if $.Parser.UsesFlags}}
		p.listener({{$val.Name}}, {{if .Flags}}{{join .Flags "|"}}{{else}}0{{end}}, rhs[{{.Start}}].sym.offset, rhs[{{minus1 .End}}].sym.endoffset)
{{- else}}
		p.listener({{$val.Name}}, rhs[{{.Start}}].sym.offset, rhs[{{minus1 .End}}].sym.endoffset)
{{- end}}
{{- end}}
{{- if $act.Code }}
{{go_parser_action $act.Code $act.Vars $act.Origin}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}

{{- range $index, $rule := .Parser.Tables.Lookaheads }}
	case {{sum $index (len $.Parser.Rules)}}:
{{- if $.Options.Cancellable}}
		var ok bool
{{- end}}
		{{ range $rule.Cases }}
		{{- $sym := index $.Syms (sum $.NumTokens (index $.Parser.Inputs .Predicate.Input).Nonterm) -}}
		if {{if $.Options.Cancellable}}ok, err = {{else}}{{if .Predicate.Negated}}!{{end}}{{end -}}
			At{{$sym.Name}}({{if $.Options.Cancellable}}ctx, {{end}}lexer, p.next.symbol{{if $.NeedsSession}}, s{{end}})
		{{- if $.Options.Cancellable}}; {{if .Predicate.Negated}}!{{end}}ok{{end}} {
			lhs.sym.symbol = {{.Target}} /* {{(index $.Syms .Target).Name}} */
		} else {{end}}{
			lhs.sym.symbol = {{.DefaultTarget}} /* {{(index $.Syms .DefaultTarget).Name}} */
		}
		return
{{- end}}
	}
{{- end}}
{{- if .Parser.Types }}
	if nt := tmRuleType[rule]; nt != 0 {
{{- if .Parser.UsesFlags}}
		p.listener({{ref "NodeType"}}(nt&0xffff), {{ref "NodeFlags"}}(nt>>16), lhs.sym.offset, lhs.sym.endoffset)
{{- else}}
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
{{- end}}
	}
{{- end}}
	return
}

{{ end -}}

{{ if .Options.FixWhitespace -}}
{{ block "fixTrailingWS" . -}}
func fixTrailingWS(lhs *stackEntry, rhs []stackEntry) {
	last := len(rhs)-1
	if last < 0 {
		return
	}
	for last >= 0 && rhs[last].sym.offset == rhs[last].sym.endoffset {
		last--
	}
	if last >= 0 {
		lhs.sym.endoffset = rhs[last].sym.endoffset
	} else {
		lhs.sym.endoffset = lhs.sym.offset
	}
}

{{ end -}}
{{ end -}}

{{ if .ReportTokens true -}}
{{ block "reportIgnoredToken" . -}}
func (p *Parser) reportIgnoredToken(tok symbol) {
	var t {{ref "NodeType"}}
	switch {{template "tokenType" .}}(tok.symbol) {
{{- range .Parser.MappedTokens}}
{{- $sym := index $.Syms .Token}}
{{- if or $sym.Space (eq $sym.Name "invalid_token") }}
	case {{template "tokenPkg" $}}{{$sym.ID}}:
		t = {{.Name}}
{{- end}}
{{- end}}
	default:
		return
	}
	if debugSyntax {
		"fmt".Printf("ignored: %v as %v\n", {{template "tokenType" .}}(tok.symbol), t)
	}
	p.listener(t, {{if .Parser.UsesFlags}}0, {{end}}tok.offset, tok.endoffset)
}
{{ end -}}
{{ end -}}
`

const parserTablesTpl = `
{{- define "tokenType"}}"{{.Options.Package}}/token".Token{{end -}}
{{- define "tokenPkg"}}"{{.Options.Package}}/token".{{end -}}

{{- template "header" . -}}
package {{.Name}}

{{- range .Parser.Tables.Markers}}
{{- if and (ne .Name "lr0") (ne .Name "greedy")}}
{{if eq (len .States) 1}}
const {{.Name}}State = {{index .States 0}}
{{- else}}
var {{.Name}}States = map[int]bool{
{{- range .States}}
	{{.}}: true,
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}

var tmNonterminals = [...]string{
{{- range .Parser.Nonterms}}
	"{{.Name}}",
{{- end}}
}

func symbolName(sym int32) string {
	if sym == noToken {
		return "<no-token>"
	}
	if sym < int32({{template "tokenPkg" .}}NumTokens) {
		return {{template "tokenType" .}}(sym).String()
	}
	if i := int(sym) - int({{template "tokenPkg" .}}NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return "fmt".Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
{{- int_array .Parser.Tables.Action "\t" 79 -}}
}
{{- if .Parser.Tables.Lalr}}

var tmLalr = []int32{
{{- int_array .Parser.Tables.Lalr "\t" 79 -}}
}
{{- end}}

var tmGoto = []int32{
{{- int_array .Parser.Tables.Goto "\t" 79 -}}
}

{{$stateType := bits_per_element .Parser.Tables.FromTo -}}
var tmFromTo = []int{{$stateType}}{
{{- int_array .Parser.Tables.FromTo "\t" 79 -}}
}

var tmRuleLen = []int{{bits_per_element .Parser.Tables.RuleLen}}{
{{- int_array .Parser.Tables.RuleLen "\t" 79 -}}
}

var tmRuleSymbol = []int32{
{{- int_array .Parser.Tables.RuleSymbol "\t" 79 -}}
}

{{- if .Parser.UsesFlags}}

var tmRuleType = [...]uint32{
{{- range .Parser.Rules}}
{{- if ne .Type -1 }}
{{- $val := index $.Parser.Types.RangeTypes .Type }}
	{{if ne $val.Name $.Options.FileNode}}uint32({{$val.Name}}){{if .Flags}} + uint32({{join .Flags " | "}})<<16{{end}}{{else}}0{{end}}, // {{$.RuleString .}}
{{- else }}
	0, // {{$.RuleString .}}
{{- end}}
{{- end}}
}
{{- else }}

var tmRuleType = [...]{{ref "NodeType"}}{
{{- range .Parser.Rules}}
{{- if ne .Type -1 }}
{{- $val := index $.Parser.Types.RangeTypes .Type }}
	{{if ne $val.Name $.Options.FileNode}}{{$val.Name}}{{else}}0{{end}}, // {{$.RuleString .}}
{{- else }}
	0, // {{$.RuleString .}}
{{- end}}
{{- end}}
}
{{- end }}

{{- range .Sets}}

// {{.Expr}} = {{.ValueString $}}
var {{.Name}} = []int32{
{{- if gt (len .Terminals) 0 -}}
{{- int_array .Terminals "\t" 79 -}}
{{- end -}}
}
{{- end}}
`

const parserListenerTpl = `
{{- template "header" . -}}
package {{.Name}}
{{if .Parser.UsesFlags}}
type NodeType uint16

type NodeFlags uint16

type Listener func(t NodeType, flags NodeFlags, offset, endoffset int)
{{- else}}
type NodeType int

type Listener func(t NodeType, offset, endoffset int)
{{- end}}

const (
	NoType NodeType = iota
{{- range .Parser.Types.RangeTypes }}
	{{.Name}}    {{- if gt (len .Fields) 0}}  // {{.Descriptor}}{{end}}
{{- end}}
{{- range .Options.ExtraTypes }}
	{{.}}
{{- end}}
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
{{- range .Parser.Types.RangeTypes }}
	"{{.Name}}",
{{- end}}
{{- range .Options.ExtraTypes }}
	"{{.}}",
{{- end}}
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return "fmt".Sprintf("node(%d)", t)
}

{{- range .Parser.Types.Categories }}
var {{.Name}} = []NodeType{
{{- range .Types }}
    {{.}},
{{- end}}
}
{{end}}
`
