// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/inspirer/textmapper/tm-parsers/js"
)

type Selector func(nt js.NodeType) bool

var (
	Any                       = func(t js.NodeType) bool { return true }
	Abstract                  = func(t js.NodeType) bool { return t == js.Abstract }
	AccessibilityModifier     = func(t js.NodeType) bool { return t == js.AccessibilityModifier }
	AdditiveExpr              = func(t js.NodeType) bool { return t == js.AdditiveExpr }
	Arguments                 = func(t js.NodeType) bool { return t == js.Arguments }
	ArrayLiteral              = func(t js.NodeType) bool { return t == js.ArrayLiteral }
	ArrayPattern              = func(t js.NodeType) bool { return t == js.ArrayPattern }
	ArrayType                 = func(t js.NodeType) bool { return t == js.ArrayType }
	ArrowFunc                 = func(t js.NodeType) bool { return t == js.ArrowFunc }
	AssertClause              = func(t js.NodeType) bool { return t == js.AssertClause }
	AssertEntry               = func(t js.NodeType) bool { return t == js.AssertEntry }
	AssertionKey              = func(t js.NodeType) bool { return t == js.AssertionKey }
	AssertsType               = func(t js.NodeType) bool { return t == js.AssertsType }
	AssignmentExpr            = func(t js.NodeType) bool { return t == js.AssignmentExpr }
	AssignmentOperator        = func(t js.NodeType) bool { return t == js.AssignmentOperator }
	AsyncArrowFunc            = func(t js.NodeType) bool { return t == js.AsyncArrowFunc }
	AsyncFunc                 = func(t js.NodeType) bool { return t == js.AsyncFunc }
	AsyncFuncExpr             = func(t js.NodeType) bool { return t == js.AsyncFuncExpr }
	AsyncGeneratorDeclaration = func(t js.NodeType) bool { return t == js.AsyncGeneratorDeclaration }
	AsyncGeneratorExpression  = func(t js.NodeType) bool { return t == js.AsyncGeneratorExpression }
	AsyncGeneratorMethod      = func(t js.NodeType) bool { return t == js.AsyncGeneratorMethod }
	AsyncMethod               = func(t js.NodeType) bool { return t == js.AsyncMethod }
	Await                     = func(t js.NodeType) bool { return t == js.Await }
	AwaitExpr                 = func(t js.NodeType) bool { return t == js.AwaitExpr }
	BindingRestElement        = func(t js.NodeType) bool { return t == js.BindingRestElement }
	BitwiseAND                = func(t js.NodeType) bool { return t == js.BitwiseAND }
	BitwiseOR                 = func(t js.NodeType) bool { return t == js.BitwiseOR }
	BitwiseXOR                = func(t js.NodeType) bool { return t == js.BitwiseXOR }
	Block                     = func(t js.NodeType) bool { return t == js.Block }
	Body                      = func(t js.NodeType) bool { return t == js.Body }
	BreakStmt                 = func(t js.NodeType) bool { return t == js.BreakStmt }
	CallExpr                  = func(t js.NodeType) bool { return t == js.CallExpr }
	CallSignature             = func(t js.NodeType) bool { return t == js.CallSignature }
	Case                      = func(t js.NodeType) bool { return t == js.Case }
	Catch                     = func(t js.NodeType) bool { return t == js.Catch }
	Class                     = func(t js.NodeType) bool { return t == js.Class }
	ClassBody                 = func(t js.NodeType) bool { return t == js.ClassBody }
	ClassExpr                 = func(t js.NodeType) bool { return t == js.ClassExpr }
	CoalesceExpr              = func(t js.NodeType) bool { return t == js.CoalesceExpr }
	CommaExpr                 = func(t js.NodeType) bool { return t == js.CommaExpr }
	ComputedPropertyName      = func(t js.NodeType) bool { return t == js.ComputedPropertyName }
	ConciseBody               = func(t js.NodeType) bool { return t == js.ConciseBody }
	ConditionalExpr           = func(t js.NodeType) bool { return t == js.ConditionalExpr }
	ConstructSignature        = func(t js.NodeType) bool { return t == js.ConstructSignature }
	ConstructorType           = func(t js.NodeType) bool { return t == js.ConstructorType }
	ContinueStmt              = func(t js.NodeType) bool { return t == js.ContinueStmt }
	DebuggerStmt              = func(t js.NodeType) bool { return t == js.DebuggerStmt }
	Declare                   = func(t js.NodeType) bool { return t == js.Declare }
	DecoratorCall             = func(t js.NodeType) bool { return t == js.DecoratorCall }
	DecoratorExpr             = func(t js.NodeType) bool { return t == js.DecoratorExpr }
	Default                   = func(t js.NodeType) bool { return t == js.Default }
	DefaultParameter          = func(t js.NodeType) bool { return t == js.DefaultParameter }
	DoWhileStmt               = func(t js.NodeType) bool { return t == js.DoWhileStmt }
	ElementBinding            = func(t js.NodeType) bool { return t == js.ElementBinding }
	EmptyDecl                 = func(t js.NodeType) bool { return t == js.EmptyDecl }
	EmptyStmt                 = func(t js.NodeType) bool { return t == js.EmptyStmt }
	EqualityExpr              = func(t js.NodeType) bool { return t == js.EqualityExpr }
	ExponentiationExpr        = func(t js.NodeType) bool { return t == js.ExponentiationExpr }
	ExportClause              = func(t js.NodeType) bool { return t == js.ExportClause }
	ExportDecl                = func(t js.NodeType) bool { return t == js.ExportDecl }
	ExportDefault             = func(t js.NodeType) bool { return t == js.ExportDefault }
	ExportSpec                = func(t js.NodeType) bool { return t == js.ExportSpec }
	ExprStmt                  = func(t js.NodeType) bool { return t == js.ExprStmt }
	Extends                   = func(t js.NodeType) bool { return t == js.Extends }
	Finally                   = func(t js.NodeType) bool { return t == js.Finally }
	ForBinding                = func(t js.NodeType) bool { return t == js.ForBinding }
	ForCondition              = func(t js.NodeType) bool { return t == js.ForCondition }
	ForFinalExpr              = func(t js.NodeType) bool { return t == js.ForFinalExpr }
	ForInStmt                 = func(t js.NodeType) bool { return t == js.ForInStmt }
	ForInStmtWithVar          = func(t js.NodeType) bool { return t == js.ForInStmtWithVar }
	ForOfStmt                 = func(t js.NodeType) bool { return t == js.ForOfStmt }
	ForOfStmtWithVar          = func(t js.NodeType) bool { return t == js.ForOfStmtWithVar }
	ForStmt                   = func(t js.NodeType) bool { return t == js.ForStmt }
	ForStmtWithVar            = func(t js.NodeType) bool { return t == js.ForStmtWithVar }
	Func                      = func(t js.NodeType) bool { return t == js.Func }
	FuncExpr                  = func(t js.NodeType) bool { return t == js.FuncExpr }
	FuncType                  = func(t js.NodeType) bool { return t == js.FuncType }
	Generator                 = func(t js.NodeType) bool { return t == js.Generator }
	GeneratorExpr             = func(t js.NodeType) bool { return t == js.GeneratorExpr }
	GeneratorMethod           = func(t js.NodeType) bool { return t == js.GeneratorMethod }
	Getter                    = func(t js.NodeType) bool { return t == js.Getter }
	IdentExpr                 = func(t js.NodeType) bool { return t == js.IdentExpr }
	IfStmt                    = func(t js.NodeType) bool { return t == js.IfStmt }
	ImportDecl                = func(t js.NodeType) bool { return t == js.ImportDecl }
	ImportSpec                = func(t js.NodeType) bool { return t == js.ImportSpec }
	ImportType                = func(t js.NodeType) bool { return t == js.ImportType }
	InExpr                    = func(t js.NodeType) bool { return t == js.InExpr }
	IndexAccess               = func(t js.NodeType) bool { return t == js.IndexAccess }
	IndexSignature            = func(t js.NodeType) bool { return t == js.IndexSignature }
	IndexedAccessType         = func(t js.NodeType) bool { return t == js.IndexedAccessType }
	Initializer               = func(t js.NodeType) bool { return t == js.Initializer }
	InstanceOfExpr            = func(t js.NodeType) bool { return t == js.InstanceOfExpr }
	IntersectionType          = func(t js.NodeType) bool { return t == js.IntersectionType }
	JSXAttributeName          = func(t js.NodeType) bool { return t == js.JSXAttributeName }
	JSXClosingElement         = func(t js.NodeType) bool { return t == js.JSXClosingElement }
	JSXElement                = func(t js.NodeType) bool { return t == js.JSXElement }
	JSXElementName            = func(t js.NodeType) bool { return t == js.JSXElementName }
	JSXExpr                   = func(t js.NodeType) bool { return t == js.JSXExpr }
	JSXLiteral                = func(t js.NodeType) bool { return t == js.JSXLiteral }
	JSXNormalAttribute        = func(t js.NodeType) bool { return t == js.JSXNormalAttribute }
	JSXOpeningElement         = func(t js.NodeType) bool { return t == js.JSXOpeningElement }
	JSXSelfClosingElement     = func(t js.NodeType) bool { return t == js.JSXSelfClosingElement }
	JSXSpreadAttribute        = func(t js.NodeType) bool { return t == js.JSXSpreadAttribute }
	JSXSpreadExpr             = func(t js.NodeType) bool { return t == js.JSXSpreadExpr }
	JSXText                   = func(t js.NodeType) bool { return t == js.JSXText }
	KeyOfType                 = func(t js.NodeType) bool { return t == js.KeyOfType }
	LabelIdent                = func(t js.NodeType) bool { return t == js.LabelIdent }
	LabelledStmt              = func(t js.NodeType) bool { return t == js.LabelledStmt }
	LetOrConst                = func(t js.NodeType) bool { return t == js.LetOrConst }
	LexicalBinding            = func(t js.NodeType) bool { return t == js.LexicalBinding }
	LexicalDecl               = func(t js.NodeType) bool { return t == js.LexicalDecl }
	Literal                   = func(t js.NodeType) bool { return t == js.Literal }
	LiteralPropertyName       = func(t js.NodeType) bool { return t == js.LiteralPropertyName }
	LiteralType               = func(t js.NodeType) bool { return t == js.LiteralType }
	LogicalAND                = func(t js.NodeType) bool { return t == js.LogicalAND }
	LogicalOR                 = func(t js.NodeType) bool { return t == js.LogicalOR }
	MappedType                = func(t js.NodeType) bool { return t == js.MappedType }
	MemberMethod              = func(t js.NodeType) bool { return t == js.MemberMethod }
	MemberVar                 = func(t js.NodeType) bool { return t == js.MemberVar }
	Method                    = func(t js.NodeType) bool { return t == js.Method }
	MethodSignature           = func(t js.NodeType) bool { return t == js.MethodSignature }
	Module                    = func(t js.NodeType) bool { return t == js.Module }
	ModuleSpec                = func(t js.NodeType) bool { return t == js.ModuleSpec }
	MultiplicativeExpr        = func(t js.NodeType) bool { return t == js.MultiplicativeExpr }
	NameIdent                 = func(t js.NodeType) bool { return t == js.NameIdent }
	NameSpaceImport           = func(t js.NodeType) bool { return t == js.NameSpaceImport }
	NamedImports              = func(t js.NodeType) bool { return t == js.NamedImports }
	NamedTupleMember          = func(t js.NodeType) bool { return t == js.NamedTupleMember }
	NewExpr                   = func(t js.NodeType) bool { return t == js.NewExpr }
	NewTarget                 = func(t js.NodeType) bool { return t == js.NewTarget }
	NoElement                 = func(t js.NodeType) bool { return t == js.NoElement }
	NonNullableType           = func(t js.NodeType) bool { return t == js.NonNullableType }
	NullableType              = func(t js.NodeType) bool { return t == js.NullableType }
	ObjectLiteral             = func(t js.NodeType) bool { return t == js.ObjectLiteral }
	ObjectMethod              = func(t js.NodeType) bool { return t == js.ObjectMethod }
	ObjectPattern             = func(t js.NodeType) bool { return t == js.ObjectPattern }
	ObjectType                = func(t js.NodeType) bool { return t == js.ObjectType }
	OptionalCallExpr          = func(t js.NodeType) bool { return t == js.OptionalCallExpr }
	OptionalIndexAccess       = func(t js.NodeType) bool { return t == js.OptionalIndexAccess }
	OptionalPropertyAccess    = func(t js.NodeType) bool { return t == js.OptionalPropertyAccess }
	OptionalTaggedTemplate    = func(t js.NodeType) bool { return t == js.OptionalTaggedTemplate }
	Override                  = func(t js.NodeType) bool { return t == js.Override }
	Parameters                = func(t js.NodeType) bool { return t == js.Parameters }
	Parenthesized             = func(t js.NodeType) bool { return t == js.Parenthesized }
	ParenthesizedType         = func(t js.NodeType) bool { return t == js.ParenthesizedType }
	PostDec                   = func(t js.NodeType) bool { return t == js.PostDec }
	PostInc                   = func(t js.NodeType) bool { return t == js.PostInc }
	PreDec                    = func(t js.NodeType) bool { return t == js.PreDec }
	PreInc                    = func(t js.NodeType) bool { return t == js.PreInc }
	PredefinedType            = func(t js.NodeType) bool { return t == js.PredefinedType }
	Property                  = func(t js.NodeType) bool { return t == js.Property }
	PropertyAccess            = func(t js.NodeType) bool { return t == js.PropertyAccess }
	PropertyBinding           = func(t js.NodeType) bool { return t == js.PropertyBinding }
	PropertySignature         = func(t js.NodeType) bool { return t == js.PropertySignature }
	Readonly                  = func(t js.NodeType) bool { return t == js.Readonly }
	ReadonlyType              = func(t js.NodeType) bool { return t == js.ReadonlyType }
	ReferenceIdent            = func(t js.NodeType) bool { return t == js.ReferenceIdent }
	Regexp                    = func(t js.NodeType) bool { return t == js.Regexp }
	RelationalExpr            = func(t js.NodeType) bool { return t == js.RelationalExpr }
	RestParameter             = func(t js.NodeType) bool { return t == js.RestParameter }
	RestType                  = func(t js.NodeType) bool { return t == js.RestType }
	ReturnStmt                = func(t js.NodeType) bool { return t == js.ReturnStmt }
	Setter                    = func(t js.NodeType) bool { return t == js.Setter }
	ShiftExpr                 = func(t js.NodeType) bool { return t == js.ShiftExpr }
	ShorthandProperty         = func(t js.NodeType) bool { return t == js.ShorthandProperty }
	SingleNameBinding         = func(t js.NodeType) bool { return t == js.SingleNameBinding }
	SpreadElement             = func(t js.NodeType) bool { return t == js.SpreadElement }
	SpreadProperty            = func(t js.NodeType) bool { return t == js.SpreadProperty }
	Static                    = func(t js.NodeType) bool { return t == js.Static }
	StaticBlock               = func(t js.NodeType) bool { return t == js.StaticBlock }
	SuperExpr                 = func(t js.NodeType) bool { return t == js.SuperExpr }
	SwitchStmt                = func(t js.NodeType) bool { return t == js.SwitchStmt }
	SyntaxProblem             = func(t js.NodeType) bool { return t == js.SyntaxProblem }
	TaggedTemplate            = func(t js.NodeType) bool { return t == js.TaggedTemplate }
	TemplateLiteral           = func(t js.NodeType) bool { return t == js.TemplateLiteral }
	This                      = func(t js.NodeType) bool { return t == js.This }
	ThisType                  = func(t js.NodeType) bool { return t == js.ThisType }
	ThrowStmt                 = func(t js.NodeType) bool { return t == js.ThrowStmt }
	TryStmt                   = func(t js.NodeType) bool { return t == js.TryStmt }
	TsAmbientBinding          = func(t js.NodeType) bool { return t == js.TsAmbientBinding }
	TsAmbientClass            = func(t js.NodeType) bool { return t == js.TsAmbientClass }
	TsAmbientEnum             = func(t js.NodeType) bool { return t == js.TsAmbientEnum }
	TsAmbientExportDecl       = func(t js.NodeType) bool { return t == js.TsAmbientExportDecl }
	TsAmbientFunc             = func(t js.NodeType) bool { return t == js.TsAmbientFunc }
	TsAmbientGlobal           = func(t js.NodeType) bool { return t == js.TsAmbientGlobal }
	TsAmbientImportAlias      = func(t js.NodeType) bool { return t == js.TsAmbientImportAlias }
	TsAmbientInterface        = func(t js.NodeType) bool { return t == js.TsAmbientInterface }
	TsAmbientModule           = func(t js.NodeType) bool { return t == js.TsAmbientModule }
	TsAmbientNamespace        = func(t js.NodeType) bool { return t == js.TsAmbientNamespace }
	TsAmbientTypeAlias        = func(t js.NodeType) bool { return t == js.TsAmbientTypeAlias }
	TsAmbientVar              = func(t js.NodeType) bool { return t == js.TsAmbientVar }
	TsAsConstExpr             = func(t js.NodeType) bool { return t == js.TsAsConstExpr }
	TsAsExpr                  = func(t js.NodeType) bool { return t == js.TsAsExpr }
	TsCastExpr                = func(t js.NodeType) bool { return t == js.TsCastExpr }
	TsConditional             = func(t js.NodeType) bool { return t == js.TsConditional }
	TsConst                   = func(t js.NodeType) bool { return t == js.TsConst }
	TsDynamicImport           = func(t js.NodeType) bool { return t == js.TsDynamicImport }
	TsEnum                    = func(t js.NodeType) bool { return t == js.TsEnum }
	TsEnumBody                = func(t js.NodeType) bool { return t == js.TsEnumBody }
	TsEnumMember              = func(t js.NodeType) bool { return t == js.TsEnumMember }
	TsExclToken               = func(t js.NodeType) bool { return t == js.TsExclToken }
	TsExport                  = func(t js.NodeType) bool { return t == js.TsExport }
	TsExportAssignment        = func(t js.NodeType) bool { return t == js.TsExportAssignment }
	TsImplementsClause        = func(t js.NodeType) bool { return t == js.TsImplementsClause }
	TsImportAliasDecl         = func(t js.NodeType) bool { return t == js.TsImportAliasDecl }
	TsImportRequireDecl       = func(t js.NodeType) bool { return t == js.TsImportRequireDecl }
	TsIndexMemberDecl         = func(t js.NodeType) bool { return t == js.TsIndexMemberDecl }
	TsInterface               = func(t js.NodeType) bool { return t == js.TsInterface }
	TsInterfaceExtends        = func(t js.NodeType) bool { return t == js.TsInterfaceExtends }
	TsNamespace               = func(t js.NodeType) bool { return t == js.TsNamespace }
	TsNamespaceBody           = func(t js.NodeType) bool { return t == js.TsNamespaceBody }
	TsNamespaceExportDecl     = func(t js.NodeType) bool { return t == js.TsNamespaceExportDecl }
	TsNonNull                 = func(t js.NodeType) bool { return t == js.TsNonNull }
	TsThisParameter           = func(t js.NodeType) bool { return t == js.TsThisParameter }
	TsTypeOnly                = func(t js.NodeType) bool { return t == js.TsTypeOnly }
	TupleType                 = func(t js.NodeType) bool { return t == js.TupleType }
	TypeAliasDecl             = func(t js.NodeType) bool { return t == js.TypeAliasDecl }
	TypeAnnotation            = func(t js.NodeType) bool { return t == js.TypeAnnotation }
	TypeArguments             = func(t js.NodeType) bool { return t == js.TypeArguments }
	TypeConstraint            = func(t js.NodeType) bool { return t == js.TypeConstraint }
	TypeName                  = func(t js.NodeType) bool { return t == js.TypeName }
	TypeParameter             = func(t js.NodeType) bool { return t == js.TypeParameter }
	TypeParameters            = func(t js.NodeType) bool { return t == js.TypeParameters }
	TypePredicate             = func(t js.NodeType) bool { return t == js.TypePredicate }
	TypeQuery                 = func(t js.NodeType) bool { return t == js.TypeQuery }
	TypeReference             = func(t js.NodeType) bool { return t == js.TypeReference }
	TypeVar                   = func(t js.NodeType) bool { return t == js.TypeVar }
	UnaryExpr                 = func(t js.NodeType) bool { return t == js.UnaryExpr }
	UnionType                 = func(t js.NodeType) bool { return t == js.UnionType }
	UniqueType                = func(t js.NodeType) bool { return t == js.UniqueType }
	Var                       = func(t js.NodeType) bool { return t == js.Var }
	VarDecl                   = func(t js.NodeType) bool { return t == js.VarDecl }
	VarStmt                   = func(t js.NodeType) bool { return t == js.VarStmt }
	WhileStmt                 = func(t js.NodeType) bool { return t == js.WhileStmt }
	WithStmt                  = func(t js.NodeType) bool { return t == js.WithStmt }
	Yield                     = func(t js.NodeType) bool { return t == js.Yield }
	MultiLineComment          = func(t js.NodeType) bool { return t == js.MultiLineComment }
	SingleLineComment         = func(t js.NodeType) bool { return t == js.SingleLineComment }
	InvalidToken              = func(t js.NodeType) bool { return t == js.InvalidToken }
	NoSubstitutionTemplate    = func(t js.NodeType) bool { return t == js.NoSubstitutionTemplate }
	TemplateHead              = func(t js.NodeType) bool { return t == js.TemplateHead }
	TemplateMiddle            = func(t js.NodeType) bool { return t == js.TemplateMiddle }
	TemplateTail              = func(t js.NodeType) bool { return t == js.TemplateTail }
	BindingPattern            = OneOf(js.BindingPattern...)
	CaseClause                = OneOf(js.CaseClause...)
	ClassElement              = OneOf(js.ClassElement...)
	Decl                      = OneOf(js.Decl...)
	Decorator                 = OneOf(js.Decorator...)
	ElementPattern            = OneOf(js.ElementPattern...)
	ExportElement             = OneOf(js.ExportElement...)
	Expr                      = OneOf(js.Expr...)
	IterationStmt             = OneOf(js.IterationStmt...)
	JSXAttribute              = OneOf(js.JSXAttribute...)
	JSXAttributeValue         = OneOf(js.JSXAttributeValue...)
	JSXChild                  = OneOf(js.JSXChild...)
	MethodDefinition          = OneOf(js.MethodDefinition...)
	Modifier                  = OneOf(js.Modifier...)
	ModuleItem                = OneOf(js.ModuleItem...)
	NamedImport               = OneOf(js.NamedImport...)
	Parameter                 = OneOf(js.Parameter...)
	PropertyDefinition        = OneOf(js.PropertyDefinition...)
	PropertyName              = OneOf(js.PropertyName...)
	PropertyPattern           = OneOf(js.PropertyPattern...)
	Stmt                      = OneOf(js.Stmt...)
	StmtListItem              = OneOf(js.StmtListItem...)
	TokenSet                  = OneOf(js.TokenSet...)
	TsAmbientElement          = OneOf(js.TsAmbientElement...)
	TsType                    = OneOf(js.TsType...)
	TupleMember               = OneOf(js.TupleMember...)
	TypeMember                = OneOf(js.TypeMember...)
)

func OneOf(types ...js.NodeType) Selector {
	if len(types) == 0 {
		return func(js.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t js.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
