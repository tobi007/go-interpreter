package lexer

import (
	"go-interpreter/token"
	"testing"
)
type NextTokenTestCases struct {
	expectedType	token.TokenType
	expectedLiteral	string
}

func nextTokenTest(l *Lexer, testsCases []NextTokenTestCases, t *testing.T)  {
	for i, tt := range testsCases {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	t.Run("Parse First", func(t *testing.T) {
		input :=
			`
			let five = 5;
			let ten = 10;

			let add = fn(x, y) {
				x + y;
			};
			
			let result = add(five, ten);
		`

		tests := []NextTokenTestCases{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
		}

		nextTokenTest(New(input), tests, t)
	})

	t.Run("Parse Second", func(t *testing.T) {
		input :=
			`
			!-/*5;
			5 < 10 > 5;
		`
		tests := []NextTokenTestCases{
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.GT, ">"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		nextTokenTest(New(input), tests, t)
	})

	t.Run("Parse Third", func(t *testing.T) {
		input :=
		`
			if (5 < 10) {
				return true;
			} else {
				return false;
			}
		`
		tests := []NextTokenTestCases{
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.TRUE, "true"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.ELSE, "else"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.FALSE, "false"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}

		nextTokenTest(New(input), tests, t)
	})

	t.Run("Parse Fourth", func(t *testing.T) {
		input :=
			`
				10 == 10;
				10 != 9;
			`
		tests := []NextTokenTestCases{
			{token.INT, "10"},
			{token.EQ, "=="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.INT, "10"},
			{token.NOT_EQ, "!="},
			{token.INT, "9"},
			{token.SEMICOLON, ";"},
		}

		nextTokenTest(New(input), tests, t)
	})
}
