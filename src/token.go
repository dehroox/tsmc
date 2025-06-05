package main

type TokenClass int
type TokenId int

// token ids
// sourced from https://github.com/guidobouman/typescript/blob/master/src/compiler/tokens.ts
const (
	Any TokenId = iota
	Bool
	Break
	Case
	Catch
	Class
	Const
	Continue
	Debugger
	Default
	Delete
	Do
	Else
	Enum
	Export
	Extends
	Declare
	False
	Finally
	For
	Function
	Constructor
	Get
	If
	Implements
	Import
	In
	InstanceOf
	Interface
	Let
	Module
	New
	Number
	Null
	Package
	Private
	Protected
	Public
	Return
	Set
	Static
	String
	Super
	Switch
	This
	Throw
	True
	Try
	TypeOf
	Var
	Void
	With
	While
	Yield
	// Punctuation
	Semicolon
	OpenParen
	CloseParen
	OpenBracket
	CloseBracket
	OpenBrace
	CloseBrace
	Comma
	Equals
	PlusEquals
	MinusEquals
	AsteriskEquals
	SlashEquals
	PercentEquals
	AmpersandEquals
	CaretEquals
	BarEquals
	LessThanLessThanEquals
	GreaterThanGreaterThanEquals
	GreaterThanGreaterThanGreaterThanEquals
	Question
	Colon
	BarBar
	AmpersandAmpersand
	Bar
	Caret
	And
	EqualsEquals
	ExclamationEquals
	EqualsEqualsEquals
	ExclamationEqualsEquals
	LessThan
	LessThanEquals
	GreaterThan
	GreaterThanEquals
	LessThanLessThan
	GreaterThanGreaterThan
	GreaterThanGreaterThanGreaterThan
	Plus
	Minus
	Asterisk
	Slash
	Percent
	Tilde
	Exclamation
	PlusPlus
	MinusMinus
	Dot
	DotDotDot
	Error
	EndOfFile
	EqualsGreaterThan
	Identifier
	StringLiteral
	RegularExpressionLiteral
	NumberLiteral
	Whitespace
	Comment
	Lim
	LimFixed   = EqualsGreaterThan
	LimKeyword = Yield
)

// token classes
const (
	Punctuation TokenClass = iota
	Keyword
	Operator
	Comment_
	Whitespace_
	Identifier_
	NumberLiteral_
	StringLiteral_
	RegExpLiteral
)

var idName = map[TokenId]string{
	Any:                                     "Any",
	Bool:                                    "Bool",
	Break:                                   "Break",
	Case:                                    "Case",
	Catch:                                   "Catch",
	Class:                                   "Class",
	Const:                                   "Const",
	Continue:                                "Continue",
	Debugger:                                "Debugger",
	Default:                                 "Default",
	Delete:                                  "Delete",
	Do:                                      "Do",
	Else:                                    "Else",
	Enum:                                    "Enum",
	Export:                                  "Export",
	Extends:                                 "Extends",
	Declare:                                 "Declare",
	False:                                   "False",
	Finally:                                 "Finally",
	For:                                     "For",
	Function:                                "Function",
	Constructor:                             "Constructor",
	Get:                                     "Get",
	If:                                      "If",
	Implements:                              "Implements",
	Import:                                  "Import",
	In:                                      "In",
	InstanceOf:                              "InstanceOf",
	Interface:                               "Interface",
	Let:                                     "Let",
	Module:                                  "Module",
	New:                                     "New",
	Number:                                  "Number",
	Null:                                    "Null",
	Package:                                 "Package",
	Private:                                 "Private",
	Protected:                               "Protected",
	Public:                                  "Public",
	Return:                                  "Return",
	Set:                                     "Set",
	Static:                                  "Static",
	String:                                  "String",
	Super:                                   "Super",
	Switch:                                  "Switch",
	This:                                    "This",
	Throw:                                   "Throw",
	True:                                    "True",
	Try:                                     "Try",
	TypeOf:                                  "TypeOf",
	Var:                                     "Var",
	Void:                                    "Void",
	With:                                    "With",
	While:                                   "While",
	Yield:                                   "Yield",
	Semicolon:                               "Semicolon",
	OpenParen:                               "OpenParen",
	CloseParen:                              "CloseParen",
	OpenBracket:                             "OpenBracket",
	CloseBracket:                            "CloseBracket",
	OpenBrace:                               "OpenBrace",
	CloseBrace:                              "CloseBrace",
	Comma:                                   "Comma",
	Equals:                                  "Equals",
	PlusEquals:                              "PlusEquals",
	MinusEquals:                             "MinusEquals",
	AsteriskEquals:                          "AsteriskEquals",
	SlashEquals:                             "SlashEquals",
	PercentEquals:                           "PercentEquals",
	AmpersandEquals:                         "AmpersandEquals",
	CaretEquals:                             "CaretEquals",
	BarEquals:                               "BarEquals",
	LessThanLessThanEquals:                  "LessThanLessThanEquals",
	GreaterThanGreaterThanEquals:            "GreaterThanGreaterThanEquals",
	GreaterThanGreaterThanGreaterThanEquals: "GreaterThanGreaterThanGreaterThanEquals",
	Question:                                "Question",
	Colon:                                   "Colon",
	BarBar:                                  "BarBar",
	AmpersandAmpersand:                      "AmpersandAmpersand",
	Bar:                                     "Bar",
	Caret:                                   "Caret",
	And:                                     "And",
	EqualsEquals:                            "EqualsEquals",
	ExclamationEquals:                       "ExclamationEquals",
	EqualsEqualsEquals:                      "EqualsEqualsEquals",
	ExclamationEqualsEquals:                 "ExclamationEqualsEquals",
	LessThan:                                "LessThan",
	LessThanEquals:                          "LessThanEquals",
	GreaterThan:                             "GreaterThan",
	GreaterThanEquals:                       "GreaterThanEquals",
	LessThanLessThan:                        "LessThanLessThan",
	GreaterThanGreaterThan:                  "GreaterThanGreaterThan",
	GreaterThanGreaterThanGreaterThan:       "GreaterThanGreaterThanGreaterThan",
	Plus:                                    "Plus",
	Minus:                                   "Minus",
	Asterisk:                                "Asterisk",
	Slash:                                   "Slash",
	Percent:                                 "Percent",
	Tilde:                                   "Tilde",
	Exclamation:                             "Exclamation",
	PlusPlus:                                "PlusPlus",
	MinusMinus:                              "MinusMinus",
	Dot:                                     "Dot",
	DotDotDot:                               "DotDotDot",
	Error:                                   "Error",
	EndOfFile:                               "EndOfFile",
	EqualsGreaterThan:                       "EqualsGreaterThan",
	Identifier:                              "Identifier",
	StringLiteral:                           "StringLiteral",
	RegularExpressionLiteral:                "RegularExpressionLiteral",
	NumberLiteral:                           "NumberLiteral",
	Whitespace:                              "Whitespace",
	Comment:                                 "Comment",
	Lim:                                     "Lim",
}

var tokenName = map[TokenClass]string{
	Punctuation:    "Punctuation",
	Keyword:        "Keyword",
	Operator:       "Operator",
	Comment_:       "Comment",
	Whitespace_:    "Whitespace",
	Identifier_:    "Identifier",
	NumberLiteral_: "NumberLiteral",
	StringLiteral_: "StringLiteral",
	RegExpLiteral:  "RegExpLiteral",
}

type Token struct {
	id      TokenId
	class   TokenClass
	lexeme  string
	literal string
}

func (t Token) ToString() string {
	return tokenName[t.class] + " " + t.lexeme + " " + t.literal + " " + idName[t.id]
}
