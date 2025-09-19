// Calc.g4
grammar Hare;

// operators :)
lnot : '!';
nequal : '!=';
modulo : '%';
band : '&';
land : '&&';
landeq : '&&=';
bandeq : '&=';
lparen : '(';
rparen : ')';
times : '*';
timeseq : '*';
plus : '+';
pluseq : '+=';
comma : ',';
minus : '-';
minuseq : '-=';
dot : '.';
double_dot : '..';
ellipsis : '...';
div : '/';
diveq : '/=';
colon : ':';
double_colon : '::';
semicolon : ';';
less : '<';
lshift : '<<';
lshifteq : '<<=';
lesseq : '<=';
equal : '=';
lequal : '==';
arrow : '=>';
gt : '>';
gteq : '>=';
rshift : '>>';
rshifteq : '>>=';
question : '?';
lbracket : '[';
rbracket : ']';
bxor : '^';
bxoreq : '^=';
lxor : '^^';
lxoreq : '^^=';
lbrace : '{';
rbrace : '}';
bor : '|';
boreq : '|=';
lor : '||';
loreq : '||=';
bnot : '~';

keywords
    : 'abort'
    | 'align' 
    | 'alloc' 
    | 'append' 
    | 'as' 
    | 'assert' 
    | 'bool' 
    | 'break' 
    | 'case' 
    | 'const' 
    | 'continue' 
    | 'def' 
    | 'defer' 
    | 'delete' 
    | 'done' 
    | 'else' 
    | 'enum' 
    | 'export' 
    | 'f32' 
    | 'f64' 
    | 'false' 
    | 'fn' 
    | 'for' 
    | 'free' 
    | 'i16' 
    | 'i32' 
    | 'i64' 
    | 'i8' 
    | 'if' 
    | 'insert' 
    | 'int' 
    | 'is' 
    | 'len' 
    | 'let' 
    | 'match' 
    | 'never' 
    | 'nomem' 
    | 'null' 
    | 'nullable' 
    | 'offset' 
    | 'opaque' 
    | 'return' 
    | 'rune' 
    | 'size' 
    | 'static' 
    | 'str' 
    | 'struct' 
    | 'switch' 
    | 'true' 
    | 'type' 
    | 'u16' 
    | 'u32' 
    | 'u64' 
    | 'u8' 
    | 'uint' 
    | 'uintptr' 
    | 'union' 
    | 'use' 
    | 'vaarg' 
    | 'vaend' 
    | 'valist' 
    | 'vastart' 
    | 'void' 
    | 'yield' 
    | '\_'
    ;


attributes
    : '@fini'
    | '@init'
    | '@packed'
    | '@symbol'
    | '@test'
    | '@threadlocal'
    ;


// todo: falta name: invalid_attribute : '@ name'


// types

type: 'const'? '!'? storage-class

primitive-type
    : integer-type 
    | floating-type 
    | 'bool' 
    | 'done' 
    | 'never' 
    | 'nomem' 
    | 'opaque' 
    | 'rune' 
    | 'str' 
    | 'valist' 
    | 'void'
    ;

pointer-type: '*', type | 'nullable', '*', type

struct-union-type
    : 'struct' '@packed'? '{' struct-union-fields '}'
    | union '{ struct-union-fields '}'
    ;

struct-union-fields
    : struct-union-field ','?
    | struct-union-field ',' struct-union-fields
    ;

struct-union-field
    : '_' ':' type
    | name ':' type
    | struct-union-type
    | identifier
    ;

tuple-type
    : '(' tupyle-types ')'
    ;

tuple-types
    : type ',' type ','?
    | type ',' tuple-types
    ;

tagged-union-type
    : '( tagged-types ')'
    ;

tagged-types
    : type '|' type '|'?
    | type '|' tagged-types
    ;

slice-array-type
    : '[' ']' type
    | '[' expression ']' type
    | '[' '*' ']' type
    | '[' '_' ']' type
    ;

function-type
    : 'fn' prototype
    ;

prototype
    : '(' parameter-list ')' type
    ;

parameter-list
    : parameters ','?
    | parameters '...'
    | parameters ',' '...'
    | '...'
    ;

parameters
    : parameter
    | parameters ',' parameter
    ;

parameter
    : name ':' type default-value?
    | type default-value?
    ;

default-value
    : '=' expression
    ;

alias-type
    : identifier
    ;

unwrapped-alias
    : '...' identifier
    ;

integer-type
    : 'i8' 
    | 'i16' 
    | 'i32' 
    | 'i64' 
    | 'u8' 
    | 'u16' 
    | 'u32' 
    | 'u64' 
    | 'int' 
    | 'uint' 
    | 'size' 
    | 'uintptr'
    ;

floating-type : 'f32' | 'f64';

primitive_type
    : integer_type
    | float_type
    | 'bool' 
    | 'done' 
    | 'never' 
    | 'nomem' 
    | 'opaque' 
    | 'rune' 
    | 'str' 
    | 'valist' 
    | 'void'
    ;

storage-class
    : primitive-type
    | pointer-type
    | struct-union-type
    | tuple-type
    | tagged-union-type
    | slice-array-type
    | function-type
    | alias-type
    | unwrapped-type
    ;

// type declarations

enum_storage : integer_type | 'rune';
enum_values : enum_value ','? | enum_value ',' enum_values;
enum_value :  name | name '=' expression;


// identifier


name: nondigit | name alnum;

nondigit
    : 'a'
    | 'b'
    | 'c' 
    | 'd' 
    | 'e' 
    | 'f' 
    | 'g' 
    | 'h' 
    | 'i' 
    | 'j' 
    | 'k' 
    | 'l' 
    | 'm' 
    | 'n' 
    | 'o' 
    | 'p' 
    | 'q' 
    | 'r' 
    | 's' 
    | 't' 
    | 'u' 
    | 'v' 
    | 'w' 
    | 'x' 
    | 'y' 
    | 'z' 
    | 'a' 
    | 'b' 
    | 'c' 
    | 'd' 
    | 'e' 
    | 'f' 
    | 'g' 
    | 'h' 
    | 'i' 
    | 'j' 
    | 'k' 
    | 'l' 
    | 'm' 
    | 'n' 
    | 'o' 
    | 'p' 
    | 'q' 
    | 'r' 
    | 's' 
    | 't' 
    | 'u' 
    | 'v' 
    | 'w' 
    | 'x' 
    | 'y' 
    | 'z' 
    | '_';


decimal_digit
    : '0'
    | '1'
    | '2'
    | '3'
    | '4'
    | '5'
    | '6'
    | '7'
    | '8'
    | '9'
    ;

alnum: decimal_digit | nondigit;

identifier : name | name '::' identifier;




// expression

literal
    : integer_literal
    | floating_literal
    | rune_literal
    | string_literal 
    | array_literal 
    | struct_literal 
    | tuple_literal 
    | 'true' 
    | 'false' 
    | 'nomem' 
    | 'null' 
    | 'void' 
    | 'done'
    ;


floating-literal
    : nonzero-decimal-digits '.' decimal-digits decimal-exponent? floating-suffix?
    | nonzero-decimal-digits decimal-exponent? floating-suffix
    | '0x' hex-digits '.' hex-digits binary-exponent floating-suffix?
    | '0x' hex-digits binary-exponent floating-suffix?
    ;

floating-suffix
    : 'f32'
    | 'f64'
    ;

decimal-digits-without-separators
    : decimal-digit decimal-digits-without-separators?
    ;

decimal-digits
    : decimal-digit decimaldigits?
    | decimal_digit '_' decimaldigits
    ;

nonzero-decimal-digits
    : '0'
    | nonzero-decimal-digit decimal-digits?
    | nonzero-decimal-digit '_' decimal_digits
    ;

nonzero-decimal-digit
    : '1'
    | '2'
    | '3'
    | '4'
    | '5'
    | '6'
    | '7'
    | '8'
    | '9'
    ;

hex-digits
    : hex-digit hex-digits?
    : hex-digit '_' hex-digits
    ;

hex-digit
    : '0'
    | '1'
    | '2'
    | '3'
    | '4'
    | '5'
    | '6'
    | '7'
    | '8'
    | '9'
    | '0'
    | 'A'
    | 'a'
    | 'B'
    | 'b'
    | 'C'
    | 'c'
    | 'D'
    | 'd'
    | 'E'
    | 'e'
    | 'F'
    | 'f'
    ;

decimal-exponent
    : decimal-exponent-char sign? decimal-digits-without-separators
    ;

binary-exponent
    : binary-exponent-char sign? decimal-digits-without-separators
    ;

sign
    : '+'
    | '-'
    ;

decimal-exponent-char
    : 'e'
    | 'E'
    ;

binary-exponent-char
    : 'p'
    | 'P'
    ;



integer-literal
: '0x' hex-digits integer-suffix?
| '0o' octal-digits integer-suffix?
| '0b' binary-digits integer-suffix?
| nonzero-decimal-digits positive-decimal-exponent? integer-suffix?
;

integer_suffix
    : 'i'
    | 'u'
    | 'z'
    | 'i8'
    | 'i16'
    | 'i32'
    | 'i64'
    | 'u8'
    | 'u16'
    | 'u32'
    | 'u64'
    ;

binary_digit: '0' | '1';
octal_digit: '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7';

binary_digits
    : binary_digit binary_digits?
    | binary_digit '_' binary_digits
    ;

octal_digits
    : octal_digit octal_digits?
    | octal_digit '_' octal_digits
    ;

positive-decimal-exponent
    : decimal-exponent-char '+'? decimal-digits-without-separators
    ;




// Other

whitespace
    : [ \t]+ -> channel(hidden)
    ;

newline
    : ('\r' '\n'? | '\n') -> channel(hidden)
    ;