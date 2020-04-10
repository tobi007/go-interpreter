package token


const (
	ILLEGAL		= 	"ILLEGAL"
	EOF			=	"EOF"
	
	IDENT		=	"IDENT"
	INT			=	"INT"

	ASSIGN		=	"="
	PLUS		=	"+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	COMMA		=	","
	SEMICOLON	=	";"

	LPAREN		=	"("
	RPAREN		=	")"
	LBRACE		=	"{"
	RBRACE		=	"}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	EQ = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

//TokenType a token type
type TokenType string

//Token a token
type Token struct {
	Type		TokenType
	Literal		string
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}