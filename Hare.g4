// Calc.g4
grammar Hare;

WHITESPACE: [ \r\n\t]+ -> skip;

// OPERATORS :)
LNOT : "!";
NEQUAL : "!=";
MODULO : "%";
BAND : "&";
LAND : "&&";
LANDEQ : "&&=";
BANDEQ : "&=";
LPAREN : "(";
RPAREN : ")";
TIMES : "*";
TIMESEQ : "*";
PLUS : "+";
PLUSEQ : "+=";
COMMA : ",";
MINUS : "-";
MINUSEQ : "-=";
DOT : ".";
DOUBLE_DOT : "..";
ELLIPSIS : "...";
DIV : "/";
DIVEQ : "/=";
COLON : ":";
DOUBLE_COLON : "::";
SEMICOLON : ";";
LESS : "<";
LSHIFT : "<<";
LSHIFTEQ : "<<=";
LESSEQ : "<=";
EQUAL : "=";
LEQUAL : "==";
ARROW : "=>";
GT : ">";
GTEQ : ">=";
RSHIFT : ">>";
RSHIFTEQ : ">>=";
QUESTION : "?";
LBRACKET : "[";
RBRACKET : "]";
BXOR : "^";
BXOREQ : "^=";
LXOR : "^^";
LXOREQ : "^^=";
LBRACE : "{";
RBRACE : "}";
BOR : "|";
BOREQ : "|=";
LOR : "||";
LOREQ : "||=";
BNOT : "~";

OCTAL-

KEYWORDS : 'abort' | 
    'align' | 
    'alloc' | 
    'append' | 
    'as' | 
    'assert' | 
    'bool' | 
    'break' | 
    'case' | 
    'const' | 
    'continue' | 
    'def' | 
    'defer' | 
    'delete' | 
    'done' | 
    'else' | 
    'enum' | 
    'export' | 
    'f32' | 
    'f64' | 
    'false' | 
    'fn' | 
    'for' | 
    'free' | 
    'i16' | 
    'i32' | 
    'i64' | 
    'i8' | 
    'if' | 
    'insert' | 
    'int' | 
    'is' | 
    'len' | 
    'let' | 
    'match' | 
    'never' | 
    'nomem' | 
    'null' | 
    'nullable' | 
    'offset' | 
    'opaque' | 
    'return' | 
    'rune' | 
    'size' | 
    'static' | 
    'str' | 
    'struct' | 
    'switch' | 
    'true' | 
    'type' | 
    'u16' | 
    'u32' | 
    'u64' | 
    'u8' | 
    'uint' | 
    'uintptr' | 
    'union' | 
    'use' | 
    'vaarg' | 
    'vaend' | 
    'valist' | 
    'vastart' | 
    'void' | 
    'yield' | 
    '\_';


ATTRIBUTES : '@fini' | 
    '@init' | 
    '@packed' | 
    '@symbol' | 
    '@test' | 
    '@threadlocal';


// TODO: FALTA NAME: INVALID_ATTRIBUTE : '@ name'


// Types

INTEGER_TYPE : 'i8' | 
    'i16' | 
    'i32' | 
    'i64' | 
    'u8' | 
    'u16' | 
    'u32' | 
    'u64' | 
    'int' | 
    'uint' | 
    'size' | 
    'uintptr';

FLOAT_TYPE : 'f32' | 'f64';

PRIMITIVE_TYPE : INTEGER_TYPE |
    FLOAT_TYPE |
    'bool' | 
    'done' | 
    'never' | 
    'nomem' | 
    'opaque' | 
    'rune' | 
    'str' | 
    'valist' | 
    'void';

// Type declarations

ENUM_STORAGE : INTEGER_TYPE | 'rune';
ENUM_VALUES : ENUM_VALUE, {","} | ENUM_VALUE, ",", ENUM_VALUES;
ENUM_VALUE :  NAME | NAME, '=', EXPRESSION;

def{enum-values} \\
	{enum-value} \optional{{,}} \\
	{enum-value} {,} {enum-values} \\

def{enum-value} \\
	{name} \\
	{name} {=} {expression} \\


// Identifier


NAME: NONDIGIT | NAME, ALNUM;

NONDIGIT : 'a' | 
    'b' | 
    'c' | 
    'd' | 
    'e' | 
    'f' | 
    'g' | 
    'h' | 
    'i' | 
    'j' | 
    'k' | 
    'l' | 
    'm' | 
    'n' | 
    'o' | 
    'p' | 
    'q' | 
    'r' | 
    's' | 
    't' | 
    'u' | 
    'v' | 
    'w' | 
    'x' | 
    'y' | 
    'z' | 
    'A' | 
    'B' | 
    'C' | 
    'D' | 
    'E' | 
    'F' | 
    'G' | 
    'H' | 
    'I' | 
    'J' | 
    'K' | 
    'L' | 
    'M' | 
    'N' | 
    'O' | 
    'P' | 
    'Q' | 
    'R' | 
    'S' | 
    'T' | 
    'U' | 
    'V' | 
    'W' | 
    'X' | 
    'Y' | 
    'Z' | 
    '\_';


DECIMAL_DIGIT: '0' | 
    '1' | 
    '2' | 
    '3' | 
    '4' | 
    '5' | 
    '6' | 
    '7' | 
    '8' | 
    '9';

ALNUM: DECIMAL_DIGIT | NONDIGIT;

IDENTIFIER : NAME | NAME, {::}, IDENTIFIER;




// Expression

LITERAL: INTEGER_LITERAL | 
	FLOATING_LITERAL | 
	RUNE_LITERAL | 
	STRING_LITERAL | 
	ARRAY_LITERAL | 
	STRUCT_LITERAL | 
	TUPLE_LITERAL | 
	"true" | 
	"false" | 
	"nomem" | 
	"null" | 
	"void" | 
	"done";


INTEGER_SUFFIX: 'i' | 
    'u' | 
    'z' | 
    'i8' | 
    'i16' | 
    'i32' | 
    'i64' | 
    'u8' | 
    'u16' | 
    'u32' | 
    'u64';

BINARY_DIGIT: '0' | '1';
OCTAL_DIGIT: '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7';

BINARY_DIGITS: BINARY_DIGIT, {BINARY_DIGITS} | BINARY_DIGIT, '_', BINARY_DIGITS;
OCTAL_DIGITS: OCTAL_DIGIT, {OCTAL_DIGITS} | OCTAL_DIGIT, '_', OCTAL_DIGITS;