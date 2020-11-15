package main

import (
	"fmt"
	"strconv"
	"strings"
)
//####################################################################################
// token
type Type string

type Token struct {
        Type    Type
        Literal string
}

const (
        ILLEGAL = "ILLEGAL"
        EOF = "EOF"
        COMMENT = "COMMENT"
        IDENT = "IDENT"
        INT = "INT"
        STRING = "STRING"
        BIND = ":="
        ASSIGN = "="
        PLUS = "+"
        MINUS = "-"
        MULTIPLY = "*"
        DIVIDE = "/"
        MODULO = "%"
        LeftShift = "<<"
        RightShift = ">>"
        BitwiseAND = "&"
        BitwiseOR = "|"
        BitwiseXOR = "^"
        BitwiseNOT = "~"
        NOT = "!"
        AND = "&&"
        OR = "||"
        LT = "<"
        LTE = "<="
        GT = ">"
        GTE = ">="
        EQ = "=="
        NEQ = "!="
        COMMA = ","
        SEMICOLON = ";"
        COLON = ":"
        DOT = "."
        LPAREN = "("
        RPAREN = ")"
        LBRACE = "{"
        RBRACE = "}"
        LBRACKET = "["
        RBRACKET = "]"
        FUNCTION = "FUNCTION"
        TRUE = "TRUE"
        FALSE = "FALSE"
        NULL = "NULL"
        IF = "IF"
        ELSE = "ELSE"
        RETURN = "RETURN"
        WHILE = "WHILE"
        IMPORT = "IMPORT"
)

var keywords = map[string]Type{
        "fn":     FUNCTION,
        "true":   TRUE,
        "false":  FALSE,
        "null":   NULL,
        "if":     IF,
        "else":   ELSE,
        "return": RETURN,
        "while":  WHILE,
        "import": IMPORT,
}


func LookupIdent(ident string) Type {
        if token, ok := keywords[ident]; ok {
                return token
        }
        return IDENT
}


//####################################################################################
//Lexer

func isDigit(ch byte) bool {
        return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
        return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

type Lexer struct {
        input        string
        position     int  // current position in input (points to current char)
        readPosition int  // current reading position in input (after current char)
        ch           byte // current char under examination
        prevCh       byte // previous char read
}

func newToken(tokenType Type, ch byte) Token {
        return Token{Type: tokenType, Literal: string(ch)}
}

// New returns a new Lexer
func New(input string) *Lexer {
        l := &Lexer{input: input}
        l.readChar()
        return l
}

func (l *Lexer) readChar() {
        l.prevCh = l.ch
        if l.readPosition >= len(l.input) {
                l.ch = 0
        } else {
                l.ch = l.input[l.readPosition]
        }

        l.position = l.readPosition
        l.readPosition++
}

func (l *Lexer) peekChar() byte {
        if l.readPosition >= len(l.input) {
                return 0
        } else {
                return l.input[l.readPosition]
        }
}

func (l *Lexer) skipWhitespace() {
        for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
                l.readChar()
        }
}

func (l *Lexer) readNumber() string {
        position := l.position
        for isDigit(l.ch) {
                l.readChar()
        }
        return l.input[position:l.position]
}
func (l *Lexer) readIdentifier() string {
        position := l.position
        for isLetter(l.ch) {
                l.readChar()
        }
        return l.input[position:l.position]
}


func (l *Lexer) readLine() string {
        position := l.position + 1
        for {
                l.readChar()
                if l.ch == '\r' || l.ch == '\n' || l.ch == 0 {
                        break
                }
        }
        return l.input[position:l.position]
}



func (l *Lexer) NextToken() Token {
        var tok Token
        l.skipWhitespace()
        switch l.ch {
        case '#':
                tok.Type = COMMENT
                tok.Literal = l.readLine()
        case '=':
                if l.peekChar() == '=' {
                        ch := l.ch
                        l.readChar()
                        literal := string(ch) + string(l.ch)
                        tok = Token{Type: EQ, Literal: literal}
                } else {
                        tok = newToken(ASSIGN, l.ch)
                }
        case '(':
                tok = newToken(LPAREN, l.ch)
        case ')':
                tok = newToken(RPAREN, l.ch)
        case ';':
                tok = newToken(SEMICOLON, l.ch)
        case ':':
                if l.peekChar() == '=' {
                        ch := l.ch
                        l.readChar()
                        literal := string(ch) + string(l.ch)
                        tok = Token{Type: BIND, Literal: literal}
                } else {
                        tok = newToken(COLON, l.ch)
                }
        case '/':
                if l.peekChar() == '/' {
                        l.readChar() // skip over the '/'
                        tok.Type = COMMENT
                        tok.Literal = l.readLine()
                } else {
                        tok = newToken(DIVIDE, l.ch)
                }
        case '*':
                tok = newToken(MULTIPLY, l.ch)
        case '%':
                tok = newToken(MODULO, l.ch)
        case '+':
                tok = newToken(PLUS, l.ch)
        case '-':
                tok = newToken(MINUS, l.ch)
        case 0:
                tok.Literal = ""
                tok.Type = EOF
        default:
                if isLetter(l.ch) {
                        tok.Literal = l.readIdentifier()
                        tok.Type = LookupIdent(tok.Literal)
                        return tok
                } else if isDigit(l.ch) {
                        tok.Type = INT
                        tok.Literal = l.readNumber()
                        return tok
                } else {
                        tok = newToken(ILLEGAL, l.ch)
                }

        }

        l.readChar()
        return tok
}

//############################################################
// AST

// Comment a comment
type Comment struct {
        Token Token // the token.COMMENT token
        Value string
}

// Statement defines the interface for all statement nodes.
type Statement interface {
        Node
        statementNode()
}



type Identifier struct {
        Token Token // the token.IDENT token
        Value string
}

// Node defines an interface for all nodes in the AST.
type Node interface {
        TokenLiteral() string
        String() string
}


// Expression defines the interface for all expression nodes.
type Expression interface {
        Node
        expressionNode()
}

// AssignmentExpression represents an assignment expression of the form:
// x = 1 or xs[1] = 2
type AssignmentExpression struct {
        Token Token // The = token
        Left  Expression
        Value Expression
}


type (
        prefixParseFn func() Expression
        infixParseFn  func(Expression) Expression
)

// Program is the root node. All programs consist of a slice of Statement(s)
type Program struct {
        Statements []Statement
}

type IntegerLiteral struct {
        Token Token
        Value int64
}


type CallExpression struct {
        Token     Token // The '(' token
        Function  Expression  // Identifier or FunctionLiteral
        Arguments []Expression
}



//#############################################################
// parser
type Parser struct {
        l      *Lexer
        errors []string

        curToken  Token
        prevToken Token
        peekToken Token

        prefixParseFns map[Type]prefixParseFn
        infixParseFns  map[Type]infixParseFn
}

func New_Parser (l *Lexer) *Parser {
        p := &Parser{
                l:      l,
                errors: []string{},
        }

        return p
}

func (p *Parser) curTokenIs(t Type) bool {
        return p.curToken.Type == t
}

func (p *Parser) nextToken() {
	p.prevToken = p.curToken
        p.curToken = p.peekToken
        p.peekToken = p.l.NextToken()
}


func (p *Parser) ParseProgram(id int,EOL Type) *Program {
        program := &Program{}
        program.Statements = []Statement{}
	p.nextToken()
	p.nextToken()
LOOP:   for !p.curTokenIs(EOF) {

		//fmt.Printf("%s1) %d %#v \n",strings.Repeat(" ", id),id,p.curToken) 
		if p.curToken.Type==COMMENT {
//			u := Comment{Token: p.curToken, Value: p.curToken.Literal}
//			fmt.Printf("2)%#v\n",u)			
		}  else if p.curToken.Type==INT { 	
			
        		lit := &IntegerLiteral{Token: p.curToken}

		        value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
		        if err != nil {
		                msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		                p.errors = append(p.errors, msg)
		                return nil
		        }	
        		lit.Value = value
			fmt.Printf("%s3-1)%T %#v \n",strings.Repeat(" ", id),lit,lit)
			if p.peekToken.Type==PLUS || p.peekToken.Type==MINUS || p.peekToken.Type==MULTIPLY || p.peekToken.Type==DIVIDE || p.peekToken.Type==MODULO {
				fmt.Printf("%s3-1) ОПЕРАТОР %T %#v \n",strings.Repeat(" ", id),p.peekToken,p.peekToken)			
			}
		} else if p.curToken.Type==IDENT { 			
			if p.peekToken.Type==ASSIGN {
				u := Identifier{Token: p.curToken, Value: p.curToken.Literal}
				fmt.Printf("%s3-2)%T %#v \n",strings.Repeat(" ", id),u,u)
				// парсинг математики
				pm := p.ParseProgram(id+1,SEMICOLON)
				fmt.Printf("ПЕРЕМЕННАЯ внесения %s3-2)%T %#v \n",strings.Repeat(" ", id),pm,pm);
			} else if p.peekToken.Type==LPAREN {
				u := CallExpression{Token: p.curToken}
				fmt.Printf("%s3-3)%T %#v \n",strings.Repeat(" ", id),u,u)
				pm := p.ParseProgram(id+1,RPAREN)	
				fmt.Printf("%s3-3)%T %#v \n",strings.Repeat(" ", id),pm,pm);
				// парсинг аргументов
			} else if p.peekToken.Type==LPAREN || p.peekToken.Type==RPAREN || p.peekToken.Type==PLUS || p.peekToken.Type==MINUS || p.peekToken.Type==MULTIPLY || p.peekToken.Type==DIVIDE || p.peekToken.Type==MODULO  {
				u := AssignmentExpression{Token: p.curToken }
				fmt.Printf(" ПЕРЕМЕНАЯ %s3-4-0)%T %#v \n",strings.Repeat(" ", id),u,u)	
				if p.peekToken.Type==PLUS || p.peekToken.Type==MINUS || p.peekToken.Type==MULTIPLY || p.peekToken.Type==DIVIDE || p.peekToken.Type==MODULO {
					fmt.Printf("%s3-4-1) ОПЕРАЦИЯ %#v \n",strings.Repeat(" ", id),p.peekToken.Type);
				}
				p.nextToken()
				//fmt.Printf("%s3-4-2)%T %#v \n",strings.Repeat(" ", id),p.curToken,p.curToken)				
				//fmt.Printf("%s3-4-3)%T %#v \n",strings.Repeat(" ", id),p.peekToken ,p.peekToken )
				if p.curToken.Type==RPAREN{
					return program
				} else 	if p.peekToken.Type==LPAREN {
					pm := p.ParseProgram(id+1,RPAREN)
					fmt.Printf("%s3-5)%T %#v \n",strings.Repeat(" ", id),pm,pm);
				}		
			}
		}  else	if p.curToken.Type==EOL {
			break LOOP;
		}

                p.nextToken()
        }
//	p.nextToken()
        return program
}



func main() {
	code := "# ewedqasd fsad asd \n a=10;b=3; v=444*b+(4*2)/3; \n print(a); print(b);\n print(a+v);"
	
	
	l :=	New(code);// fmt.Printf("%T %#v\n",l,l)
	p := New_Parser(l);
	v := p.ParseProgram(0,EOF)

	fmt.Printf("%T %#v \n",v,v) 
//	fmt.Printf("%T %#v",b,b) 

}
