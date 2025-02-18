// generated by Textmapper; DO NOT EDIT

package js

import (
	"context"
	"fmt"
)

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// StopOnFirstError is an error handler that forces the parser to stop on and return the first
// error.
func StopOnFirstError(_ SyntaxError) bool { return false }

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

func (p *Parser) ParseModule(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 8, 8673, lexer)
}

func (p *Parser) ParseTypeSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 9, 8674, lexer)
}

func (p *Parser) ParseExpressionSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 10, 8675, lexer)
}

func lookaheadRule(ctx context.Context, lexer *Lexer, next, rule int32, s *session) (sym int32, err error) {
	switch rule {
	case 4979:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 0, 8662, s); ok {
			sym = 776 /* lookahead_StartOfArrowFunction */
		} else {
			sym = 177 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 4980:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 3, 8665, s); ok {
			sym = 868 /* lookahead_StartOfTypeImport */
		} else {
			sym = 869 /* lookahead_notStartOfTypeImport */
		}
		return
	case 4981:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 1, 8663, s); ok {
			sym = 358 /* lookahead_StartOfParametrizedCall */
		} else {
			sym = 330 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 4982:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 4, 8666, s); ok {
			sym = 932 /* lookahead_StartOfIs */
		} else {
			sym = 934 /* lookahead_notStartOfIs */
		}
		return
	case 4983:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 6, 8668, s); ok {
			sym = 968 /* lookahead_StartOfMappedType */
		} else {
			sym = 958 /* lookahead_notStartOfMappedType */
		}
		return
	case 4984:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 5, 8667, s); ok {
			sym = 999 /* lookahead_StartOfFunctionType */
		} else {
			sym = 951 /* lookahead_notStartOfFunctionType */
		}
		return
	case 4985:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 7, 8669, s); ok {
			sym = 972 /* lookahead_StartOfTupleElementName */
		} else {
			sym = 973 /* lookahead_notStartOfTupleElementName */
		}
		return
	case 4986:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 2, 8664, s); ok {
			sym = 837 /* lookahead_StartOfExtendsTypeRef */
		} else {
			sym = 838 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	}
	return 0, nil
}

func AtStartOfArrowFunction(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 0, 8662, s)
}

func AtStartOfParametrizedCall(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 1, 8663, s)
}

func AtStartOfExtendsTypeRef(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 2, 8664, s)
}

func AtStartOfTypeImport(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 3, 8665, s)
}

func AtStartOfIs(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 4, 8666, s)
}

func AtStartOfFunctionType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 5, 8667, s)
}

func AtStartOfMappedType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 6, 8668, s)
}

func AtStartOfTupleElementName(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 7, 8669, s)
}

func lookahead(ctx context.Context, l *Lexer, next int32, start, end int16, s *session) (bool, error) {
	var lexer Lexer
	lexer.source = l.source
	lexer.ch = l.ch
	lexer.offset = l.offset
	lexer.tokenOffset = l.tokenOffset
	lexer.line = l.line
	lexer.tokenLine = l.tokenLine
	lexer.scanOffset = l.scanOffset
	lexer.State = l.State
	lexer.Dialect = l.Dialect
	lexer.token = l.token
	// Note: Stack is intentionally omitted.

	// Use memoization for recursive lookaheads.
	if next == noToken {
		next = lookaheadNext(&lexer, end, nil /*empty stack*/)
	}
	key := uint64(l.tokenOffset) + uint64(end)<<40
	if ret, ok := s.cache[key]; ok {
		return ret, nil
	}

	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if next == noToken {
				next = lookaheadNext(&lexer, end, stack)
			}
			action = lalr(action, next)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
			sym, err := lookaheadRule(ctx, &lexer, next, rule, s)
			if err != nil {
				return false, err
			}
			if sym != 0 {
				entry.sym.symbol = sym
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if s.shiftCounter++; s.shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return false, ctx.Err()
				default:
				}
			}

			// Shift.
			if next == noToken {
				next = lookaheadNext(&lexer, end, stack)
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

	s.cache[key] = state == end
	return state == end, nil
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

func gotoState(state int16, symbol int32) int16 {
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

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer, s *session) (err error) {
	switch rule {
	case 1253: // Elision : ','
		p.listener(NoElement, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1254: // Elision : Elision ','
		p.listener(NoElement, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 1322: // LiteralPropertyName : PrivateIdentifier
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1326: // LiteralPropertyName_WithoutNew : PrivateIdentifier
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1432: // MemberExpression_Await_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1497: // MemberExpression_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1562: // MemberExpression_Yield_Await_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1609: // MemberExpression_Yield_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 2497: // BinaryExpression : BinaryExpression .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2524: // BinaryExpression_Await : BinaryExpression_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2576: // BinaryExpression_Await_NoLet : BinaryExpression_Await_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2603: // BinaryExpression_Await_NoObjLiteral : BinaryExpression_Await_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2631: // BinaryExpression_In : BinaryExpression_In .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2659: // BinaryExpression_In_Await : BinaryExpression_In_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2713: // BinaryExpression_In_Await_NoObjLiteral : BinaryExpression_In_Await_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2741: // BinaryExpression_In_NoFuncClass : BinaryExpression_In_NoFuncClass .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2795: // BinaryExpression_In_NoObjLiteral : BinaryExpression_In_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2823: // BinaryExpression_In_Yield : BinaryExpression_In_Yield .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2851: // BinaryExpression_In_Yield_Await : BinaryExpression_In_Yield_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2955: // BinaryExpression_NoLet : BinaryExpression_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2982: // BinaryExpression_NoObjLiteral : BinaryExpression_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3009: // BinaryExpression_Yield : BinaryExpression_Yield .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3036: // BinaryExpression_Yield_Await : BinaryExpression_Yield_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3088: // BinaryExpression_Yield_Await_NoLet : BinaryExpression_Yield_Await_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3140: // BinaryExpression_Yield_NoLet : BinaryExpression_Yield_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3660: // ElementElision : ','
		p.listener(NoElement, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3661: // ElementElision : Elision ','
		p.listener(NoElement, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3737: // IterationStatement : 'for' '(' 'var' VariableDeclarationList ';' .forSC ForCondition ';' .forSC ForFinalExpression ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3741: // IterationStatement : 'for' '(' 'var' ForBinding 'in' Expression_In ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3744: // IterationStatement : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In ')' Statement
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3745: // IterationStatement : 'for' '(' 'var' ForBinding 'of' AssignmentExpression_In ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3752: // IterationStatement_Await : 'for' '(' 'var' VariableDeclarationList_Await ';' .forSC ForCondition_Await ';' .forSC ForFinalExpression_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3756: // IterationStatement_Await : 'for' '(' 'var' ForBinding_Await 'in' Expression_In_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3758: // IterationStatement_Await : 'for' 'await' '(' LeftHandSideExpression_Await_NoAsync_NoLet 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3760: // IterationStatement_Await : 'for' 'await' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(ReferenceIdent, rhs[3].sym.offset, rhs[3].sym.endoffset)
		p.listener(IdentExpr, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3761: // IterationStatement_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3762: // IterationStatement_Await : 'for' 'await' '(' 'var' ForBinding_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(Var, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3763: // IterationStatement_Await : 'for' '(' 'var' ForBinding_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3764: // IterationStatement_Await : 'for' 'await' '(' ForDeclaration_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3771: // IterationStatement_Yield : 'for' '(' 'var' VariableDeclarationList_Yield ';' .forSC ForCondition_Yield ';' .forSC ForFinalExpression_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3775: // IterationStatement_Yield : 'for' '(' 'var' ForBinding_Yield 'in' Expression_In_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3778: // IterationStatement_Yield : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3779: // IterationStatement_Yield : 'for' '(' 'var' ForBinding_Yield 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3786: // IterationStatement_Yield_Await : 'for' '(' 'var' VariableDeclarationList_Yield_Await ';' .forSC ForCondition_Yield_Await ';' .forSC ForFinalExpression_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3790: // IterationStatement_Yield_Await : 'for' '(' 'var' ForBinding_Yield_Await 'in' Expression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3792: // IterationStatement_Yield_Await : 'for' 'await' '(' LeftHandSideExpression_Yield_Await_NoAsync_NoLet 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3794: // IterationStatement_Yield_Await : 'for' 'await' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(ReferenceIdent, rhs[3].sym.offset, rhs[3].sym.endoffset)
		p.listener(IdentExpr, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3795: // IterationStatement_Yield_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3796: // IterationStatement_Yield_Await : 'for' 'await' '(' 'var' ForBinding_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(Var, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3797: // IterationStatement_Yield_Await : 'for' '(' 'var' ForBinding_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3798: // IterationStatement_Yield_Await : 'for' 'await' '(' ForDeclaration_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4233: // ImportDeclaration : 'import' lookahead_StartOfTypeImport 'type' ImportClause FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4234: // ImportDeclaration : 'import' lookahead_StartOfTypeImport 'type' ImportClause FromClause ';'
		p.listener(TsTypeOnly, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4250: // ImportRequireDeclaration : 'export' 'import' lookahead_notStartOfTypeImport BindingIdentifier '=' 'require' '(' StringLiteral ')' ';'
		p.listener(TsExport, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4270: // ExportDeclaration : 'export' 'type' '*' 'as' ImportedBinding FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4271: // ExportDeclaration : 'export' 'type' '*' 'as' ImportedBinding FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4272: // ExportDeclaration : 'export' 'type' '*' FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4273: // ExportDeclaration : 'export' 'type' '*' FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4278: // ExportDeclaration : 'export' 'type' ExportClause FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4279: // ExportDeclaration : 'export' 'type' ExportClause FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4282: // ExportDeclaration : 'export' 'type' ExportClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4304: // DecoratorMemberExpression : DecoratorMemberExpression '.' IdentifierName
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4405: // TypePredicate : 'asserts' lookahead_StartOfIs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4407: // TypePredicate_NoQuest : 'asserts' lookahead_StartOfIs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4409: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs 'this' 'is' Type
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4410: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs 'this'
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4411: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs 'is' Type
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4412: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4442: // TypeOperator : 'infer' IdentifierName
		p.listener(ReferenceIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4447: // TypeOperator_NoQuest : 'infer' IdentifierName
		p.listener(ReferenceIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4595: // TupleElementType : '...' lookahead_StartOfTupleElementName IdentifierName '?' ':' Type
		p.listener(RestType, rhs[5].sym.offset, rhs[5].sym.endoffset)
	case 4596: // TupleElementType : '...' lookahead_StartOfTupleElementName IdentifierName ':' Type
		p.listener(RestType, rhs[4].sym.offset, rhs[4].sym.endoffset)
	case 4628: // ConstructorType : 'abstract' 'new' TypeParameters ParameterList '=>' Type
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4629: // ConstructorType : 'abstract' 'new' ParameterList '=>' Type
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4632: // ConstructorType_NoQuest : 'abstract' 'new' TypeParameters ParameterList '=>' Type_NoQuest
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4633: // ConstructorType_NoQuest : 'abstract' 'new' ParameterList '=>' Type_NoQuest
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4836: // IndexSignature : Modifiers '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4837: // IndexSignature : '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4838: // IndexSignature_WithDeclare : Modifiers_WithDeclare '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4839: // IndexSignature_WithDeclare : '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4853: // EnumDeclaration : 'const' 'enum' BindingIdentifier EnumBody
		p.listener(TsConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4879: // AmbientVariableDeclaration : 'var' AmbientBindingList ';'
		p.listener(Var, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4880: // AmbientVariableDeclaration : 'let' AmbientBindingList ';'
		p.listener(LetOrConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4881: // AmbientVariableDeclaration : 'const' AmbientBindingList ';'
		p.listener(LetOrConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4895: // AmbientEnumDeclaration : 'const' 'enum' BindingIdentifier EnumBody
		p.listener(TsConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4979:
		var ok bool
		if ok, err = AtStartOfArrowFunction(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 776 /* lookahead_StartOfArrowFunction */
		} else {
			lhs.sym.symbol = 177 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 4980:
		var ok bool
		if ok, err = AtStartOfTypeImport(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 868 /* lookahead_StartOfTypeImport */
		} else {
			lhs.sym.symbol = 869 /* lookahead_notStartOfTypeImport */
		}
		return
	case 4981:
		var ok bool
		if ok, err = AtStartOfParametrizedCall(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 358 /* lookahead_StartOfParametrizedCall */
		} else {
			lhs.sym.symbol = 330 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 4982:
		var ok bool
		if ok, err = AtStartOfIs(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 932 /* lookahead_StartOfIs */
		} else {
			lhs.sym.symbol = 934 /* lookahead_notStartOfIs */
		}
		return
	case 4983:
		var ok bool
		if ok, err = AtStartOfMappedType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 968 /* lookahead_StartOfMappedType */
		} else {
			lhs.sym.symbol = 958 /* lookahead_notStartOfMappedType */
		}
		return
	case 4984:
		var ok bool
		if ok, err = AtStartOfFunctionType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 999 /* lookahead_StartOfFunctionType */
		} else {
			lhs.sym.symbol = 951 /* lookahead_notStartOfFunctionType */
		}
		return
	case 4985:
		var ok bool
		if ok, err = AtStartOfTupleElementName(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 972 /* lookahead_StartOfTupleElementName */
		} else {
			lhs.sym.symbol = 973 /* lookahead_notStartOfTupleElementName */
		}
		return
	case 4986:
		var ok bool
		if ok, err = AtStartOfExtendsTypeRef(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 837 /* lookahead_StartOfExtendsTypeRef */
		} else {
			lhs.sym.symbol = 838 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

const errSymbol = 2

func (p *Parser) skipBrokenCode(lexer *Lexer, stack []stackEntry, canRecover func(symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		if debugSyntax {
			fmt.Printf("skipped while recovering: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
		}
		if len(p.pending) > 0 {
			for _, tok := range p.pending {
				p.reportIgnoredToken(tok)
			}
			p.pending = p.pending[:0]
		}
		switch Token(p.next.symbol) {
		case NOSUBSTITUTIONTEMPLATE:
			p.listener(NoSubstitutionTemplate, p.next.offset, p.next.endoffset)
		case TEMPLATEHEAD:
			p.listener(TemplateHead, p.next.offset, p.next.endoffset)
		case TEMPLATEMIDDLE:
			p.listener(TemplateMiddle, p.next.offset, p.next.endoffset)
		case TEMPLATETAIL:
			p.listener(TemplateTail, p.next.offset, p.next.endoffset)
		}
		e = p.next.endoffset
		p.fetchNext(lexer, stack)
	}
	return e
}

// willShift checks if "symbol" is going to be shifted in the given state.
// This function does not support empty productions and returns false if they occur before "symbol".
func (p *Parser) willShift(stackPos int, state int16, symbol int32, stack []stackEntry) bool {
	if state == -1 {
		return false
	}

	for state != p.endState {
		action := tmAction[state]
		if action < -2 {
			action = lalr(action, symbol)
		}

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
	p.listener(t, tok.offset, tok.endoffset)
}
