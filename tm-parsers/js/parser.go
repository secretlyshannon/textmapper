// generated by Textmapper; DO NOT EDIT

package js

func (p *Parser) Parse(lexer *Lexer) bool {
	return p.parse(0, 2686, lexer)
}

func (p *Parser) applyRule(rule int32, node *node, rhs []node) {
	nt := ruleNodeType[rule]
	if nt == 0 {
		return
	}
	p.listener(nt, node.sym.offset, node.sym.endoffset)
}

const errSymbol = 113
