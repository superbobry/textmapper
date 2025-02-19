// generated by Textmapper; DO NOT EDIT

package test

import (
	"context"
	"fmt"
)

// Parser is a table-driven LALR parser for test.
type Parser struct {
	listener Listener

	next symbol

	// Tokens to be reported with the next shift. Only non-empty when next.symbol != noToken.
	pending []symbol
}

type SyntaxError struct {
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return "syntax error"
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int8
	value interface{}
}

func (p *Parser) Init(l Listener) {
	p.listener = l
	if cap(p.pending) < startTokenBufferSize {
		p.pending = make([]symbol, 0, startTokenBufferSize)
	}
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(UNAVAILABLE)
	eoiToken             = int32(EOI)
	debugSyntax          = false
)

func (p *Parser) ParseTest(ctx context.Context, lexer *Lexer) error {
	_, err := p.parse(ctx, 0, 105, lexer)
	return err
}

func (p *Parser) ParseDecl1(ctx context.Context, lexer *Lexer) (int, error) {
	v, err := p.parse(ctx, 1, 106, lexer)
	val, _ := v.(int)
	return val, err
}

func (p *Parser) parse(ctx context.Context, start, end int8, lexer *Lexer) (interface{}, error) {
	p.pending = p.pending[:0]
	var shiftCounter int
	state := start

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	p.fetchNext(lexer, stack)

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			action = lalr(action, p.next.symbol)
		}

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
			if err := p.applyRule(ctx, rule, &entry, rhs, lexer); err != nil {
				return nil, err
			}
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if shiftCounter++; shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				default:
				}
			}

			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			state = gotoState(state, p.next.symbol)
			if state >= 0 {
				stack = append(stack, stackEntry{
					sym:   p.next,
					state: state,
					value: lexer.Value(),
				})
				if debugSyntax {
					fmt.Printf("shift: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
				}
				if len(p.pending) > 0 {
					for _, tok := range p.pending {
						p.reportIgnoredToken(tok)
					}
					p.pending = p.pending[:0]
				}
				if p.next.symbol != eoiToken {
					switch Token(p.next.symbol) {
					case IDENTIFIER:
						p.listener(Identifier, 0, p.next.offset, p.next.endoffset)
					}
					p.next.symbol = noToken
				}
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	if state != end {
		if p.next.symbol == noToken {
			p.fetchNext(lexer, stack)
		}
		err := SyntaxError{
			Offset:    p.next.offset,
			Endoffset: p.next.endoffset,
		}
		return nil, err
	}

	return stack[len(stack)-2].value, nil
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int8, symbol int32) int8 {
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

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry) {
restart:
	tok := lexer.Next()
	switch tok {
	case MULTILINECOMMENT, SINGLELINECOMMENT, INVALID_TOKEN:
		s, e := lexer.Pos()
		tok := symbol{int32(tok), s, e}
		p.pending = append(p.pending, tok)
		goto restart
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = lexer.Pos()
}

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer) (err error) {
	switch rule {
	case 5: // Declaration : '{' '-' '-' Declaration_list '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[2].sym.endoffset)
	case 6: // Declaration : '{' '-' '-' '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[2].sym.endoffset)
	case 7: // Declaration : '{' '-' Declaration_list '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 8: // Declaration : '{' '-' '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 12: // Declaration : IntegerConstant '[' ']'
		nn0, _ := rhs[0].value.(int)
		{
			switch nn0 {
			case 7:
				p.listener(Int7, 0, rhs[0].sym.offset, rhs[2].sym.endoffset)
			case 9:
				p.listener(Int9, 0, rhs[0].sym.offset, rhs[2].sym.endoffset)
			}
		}
	case 13: // Declaration : IntegerConstant
		nn0, _ := rhs[0].value.(int)
		{
			switch nn0 {
			case 7:
				p.listener(Int7, 0, rhs[0].sym.offset, rhs[0].sym.endoffset)
			case 9:
				p.listener(Int9, 0, rhs[0].sym.offset, rhs[0].sym.endoffset)
			}
		}
	case 15: // Declaration : 'test' '(' empty1 ')'
		p.listener(Empty1, 0, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 16: // Declaration : 'test' IntegerConstant
		p.listener(Icon, InTest, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 17: // Declaration : 'eval' '(' expr ')' empty1
		fixTrailingWS(lhs, rhs)
	case 20: // Declaration : 'decl2' ':' QualifiedNameopt
		fixTrailingWS(lhs, rhs)
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(NodeType(nt&0xffff), NodeFlags(nt>>16), lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func fixTrailingWS(lhs *stackEntry, rhs []stackEntry) {
	last := len(rhs) - 1
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

func (p *Parser) reportIgnoredToken(tok symbol) {
	var t NodeType
	switch Token(tok.symbol) {
	case MULTILINECOMMENT:
		t = MultiLineComment
	case SINGLELINECOMMENT:
		t = SingleLineComment
	case INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", Token(tok.symbol), t)
	}
	p.listener(t, 0, tok.offset, tok.endoffset)
}
