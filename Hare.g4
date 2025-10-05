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
doubleDot : '..';
ellipsis : '...';
div : '/';
diveq : '/=';
colon : ':';
doubleColon : '::';
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
    | '_'
    ;


attributes
    : '@fini'
    | '@init'
    | '@packed'
    | '@symbol'
    | '@test'
    | '@threadlocal'
    ;


invalidAttribute
    : '@' name
    ;



type
    : 'const'? '!'? storageClass
    ;

pointerType: '*' type | 'nullable' '*' type ;

structUnionType
    : 'struct' '@packed'? '{' structUnionFields '}'
    | 'union' '{' structUnionFields '}'
    ;

structUnionFields
    : structUnionField ','?
    | structUnionField ',' structUnionFields
    ;

structUnionField
    : '_' ':' type
    | name ':' type
    | structUnionType
    | identifier
    ;

tupleType
    : '(' tupleTypes ')'
    ;

tupleTypes
    : type ',' type ','?
    | type ',' tupleTypes
    ;

taggedUnionType
    : '(' taggedTypes ')'
    ;

taggedTypes
    : type '|' type '|'?
    | type '|' taggedTypes
    ;

sliceArrayType
    : '[' ']' type
    | '[' expression ']' type
    | '[' '*' ']' type
    | '[' '_' ']' type
    ;

functionType
    : 'fn' prototype
    ;

prototype
    : '(' parameterList ')' type
    ;

parameterList
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
    : name ':' type defaultValue?
    | type defaultValue?
    ;

defaultValue
    : '=' expression
    ;

aliasType
    : identifier
    ;

unwrappedAlias
    : '...' identifier
    ;

integerType
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

floatingType : 'f32' | 'f64';

primitiveType
    : integerType
    | floatingType
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

storageClass
    : primitiveType
    | pointerType
    | structUnionType
    | tupleType
    | taggedUnionType
    | sliceArrayType
    | functionType
    | aliasType
    | unwrappedAlias
    ;

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
    | '_'
;


decimalDigit
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

alnum: decimalDigit | nondigit;

identifier : name | name '::' identifier;




// expression

literal
    : integerLiteral
    | floatingLiteral
    | runeLiteral
    | stringLiteral 
    | arrayLiteral 
    | structLiteral 
    | tupleLiteral 
    | 'true' 
    | 'false' 
    | 'nomem' 
    | 'null' 
    | 'void' 
    | 'done'
    ;


floatingLiteral
    : nonzeroDecimalDigits '.' decimalDigits decimalExponent? floatingSuffix?
    | nonzeroDecimalDigits decimalExponent? floatingSuffix
    | '0x' hexDigits '.' hexDigits binaryExponent floatingSuffix?
    | '0x' hexDigits binaryExponent floatingSuffix?
    ;

floatingSuffix
    : 'f32'
    | 'f64'
    ;

decimalDigitsWithoutSeparators
    : decimalDigit decimalDigitsWithoutSeparators?
    ;

decimalDigits
    : decimalDigit decimalDigits?
    | decimalDigit '_' decimalDigits
    ;

nonzeroDecimalDigits
    : '0'
    | nonzeroDecimalDigit decimalDigits?
    | nonzeroDecimalDigit '_' decimalDigits
    ;

nonzeroDecimalDigit
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

hexDigits
    : hexDigit hexDigits?
    | hexDigit '_' hexDigits
    ;

hexDigit
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

decimalExponent
    : decimalExponentChar sign? decimalDigitsWithoutSeparators
    ;

binaryExponent
    : binaryExponentChar sign? decimalDigitsWithoutSeparators
    ;

sign
    : '+'
    | '-'
    ;

decimalExponentChar
    : 'e'
    | 'E'
    ;

binaryExponentChar
    : 'p'
    | 'P'
    ;



integerLiteral
: '0x' hexDigits integerSuffix?
| '0o' octalDigits integerSuffix?
| '0b' binaryDigits integerSuffix?
| nonzeroDecimalDigits positiveDecimalExponent? integerSuffix?
;

integerSuffix
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

binaryDigit: '0' | '1';
octalDigit: '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7';

binaryDigits
    : binaryDigit binaryDigits?
    | binaryDigit '_' binaryDigits
    ;

octalDigits
    : octalDigit octalDigits?
    | octalDigit '_' octalDigits
    ;

positiveDecimalExponent
    : decimalExponentChar '+'? decimalDigitsWithoutSeparators
    ;

runeLiteral
    : '\'' rune '\''
    ;

rune 
    : ~('\\'|'\'')
    | escapeSequence
    ;

escapeSequence
    : namedEscape
    | '\\x' hexDigit hexDigit
    | '\\u' fourbyte fourbyte
    | '\\U' eightbyte
    ;

fourbyte
    : hexDigit hexDigit hexDigit hexDigit
    ;

eightbyte
    : fourbyte fourbyte
    ;


namedEscape
    : '\\0'
    | '\\a'
    | '\\b'
    | '\\f'
    | '\\n'
    | '\\r'
    | '\\t'
    | '\\v'
    | '\\'
    | '\\\''
    | '\\"'
    ;



stringLiteral
    : stringSection stringLiteral?
    ;

stringSection
    : '"' stringChars? '"'
    | '`' rawstringChars? '`'
    ;

stringChars
    : stringChar stringChars?
    ;


stringChar
    : ~('"'|'\\')
    | escapeSequence
    ;

rawstringChars
    : rawstringChar rawstringChars?
    ;

rawstringChar
    : ~('`')
    ;



arrayLiteral
    : '[' arrayMembers ']'
    ;


arrayMembers
    : expression ','?
    | expression '...'
    | expression ',' arrayMembers
    ;



structLiteral
    : 'struct' '{' fieldValues ','? '}'
    | identifier '{' structInitializer '}'
    ;

structInitializer
    : fieldValues ','?
    | fieldValues ','? '...'
    | '...'
    ;


fieldValues
    : fieldValue
    | fieldValues ',' fieldValue
    ;

fieldValue
    : name '=' expression
    | name ':' type '=' expression
    | structLiteral
    ;



tupleLiteral
    : '(' tupleItems ')'
    ;

tupleItems
    : expression ',' expression ','?
    | expression ',' tupleItems
    ;


plainExpression
    : identifier
    | literal
    ;

nestedExpression
    : plainExpression
    | '(' expression ')'
    ;

allocationExpression
    : 'alloc' '(' expression ')'
    | 'alloc' '(' expression '...' ')'
    | 'alloc' '(' expression ',' expression ')'
    ;

freeExpression
    : 'free' '(' expression ')'
    ;


assertionExpression
    : 'assert' '('  expression ')'
    | 'assert' '(' expression ',' expression ')'
    | 'abort' '(' expression? ')'
    ;

staticAssertionExpression
    : 'static' assertionExpression
    ;


callExpression
    : postfixExpression '(' argumentList? ')'
    ;

argumentList
    : expression ','?
    | expression '...'
    | expression ',' argumentList
    ;


measurementExpression
    : alignExpression
    | sizeExpression
    | lengthExpression
    | offsetExpression
    ;

alignExpression
    : 'align' '(' type ')'
    ;

sizeExpression
    : 'size' '(' type ')'
    ;

lengthExpression
    : 'len' '(' expression ')'
    ;

offsetExpression
    : 'offset' '(' offsetOperand ')'
    ;

offsetOperand
    : fieldAccessExpression
    | '(' offsetOperand ')'
    ;

fieldAccessExpression
    : postfixExpression '.' name
    | postfixExpression '.' integerLiteral
    ;

indexingExpression
    : postfixExpression '[' expression ']'
    ; 

slicingExpression
    : postfixExpression '[' expression? '..' expression? ']'
    ;


sliceMutationExpression
    : appendExpression
    | insertExpression
    | deleteExpression
    ;

appendExpression
    : 'static'? 'append' '(' objectSelector ',' expression ')'
    | 'static'? 'append' '(' objectSelector ',' expression '...' ')'
    | 'static'? 'append' '(' objectSelector ',' expression ',' expression ')'
    ;



insertExpression
    : 'static'? 'insert' '(' insertOperand ',' expression ')'
    | 'static'? 'insert' '(' insertOperand ',' expression '...' ')'
    | 'static'? 'insert' '(' insertOperand ',' expression ',' expression ')'
    ;

insertOperand
    : indexingExpression
    | '(' insertOperand ')'
    ;

deleteExpression
: 'static'? 'delete' '(' deleteOperand ')'
;

deleteOperand
    : indexingExpression
    | slicingExpression
    | '(' deleteOperand ')'
    ;


errorCheckingExpression
    : postfixExpression '?'
    | postfixExpression '!'
    ;

postfixExpression
    : nestedExpression
    | postfixExpression '(' argumentList? ')'
    | postfixExpression '.' name
    | postfixExpression '.' integerLiteral
    | postfixExpression '[' expression ']'
    | postfixExpression '[' expression? '..' expression? ']'
    | postfixExpression '?'
    | postfixExpression '!'
    | builtinExpression
    ;

objectSelector
    : identifier
    | indexingExpression
    | fieldAccessExpression
    | '(' objectSelector ')'
    ;


variadicExpression
    : 'vastart' '(' ')'
    | 'vaarg' '(' objectSelector ',' type ')'
    | 'vaend' '(' objectSelector ')'
    ;

builtinExpression
    : allocationExpression
    | assertionExpression
    | measurementExpression
    | sliceMutationExpression
    | staticAssertionExpression
    | variadicExpression
    ;

unaryExpression
    : postfixExpression
    | compoundExpression
    | matchExpression
    | switchExpression
    | unaryOperator unaryExpression
    ;

unaryOperator
    : '-'
    | '~'
    | '!'
    | '*'
    | '&'
    ;


castExpression
    : unaryExpression
    | castExpression ':' type
    | castExpression 'as' nullableType
    | castExpression 'is' nullableType
    ;

nullableType
    : type
    | 'null'
    ;


multiplicativeExpression
    : castExpression
    | multiplicativeExpression '*' castExpression
    | multiplicativeExpression '/' castExpression
    | multiplicativeExpression '%' castExpression
    ;


additiveExpression
    : multiplicativeExpression
    | additiveExpression '+' multiplicativeExpression
    | additiveExpression '-' multiplicativeExpression
    ;


shiftExpression
    : additiveExpression
    | shiftExpression '<<' additiveExpression
    | shiftExpression '>>' additiveExpression
    ;

andExpression
    : shiftExpression
    | andExpression '&' shiftExpression
    ;

exclusiveOrExpression
    : andExpression
    | exclusiveOrExpression '^' andExpression
    ;

inclusiveOrExpression
    : exclusiveOrExpression
    | inclusiveOrExpression '^' exclusiveOrExpression
    ;

comparisonExpression
    : inclusiveOrExpression
    | comparisonExpression '<' inclusiveOrExpression
    | comparisonExpression '>' inclusiveOrExpression
    | comparisonExpression '<=' inclusiveOrExpression
    | comparisonExpression '>=' inclusiveOrExpression
    ;

equalityExpression
    : comparisonExpression
    | equalityExpression '==' comparisonExpression
    | equalityExpression '!=' comparisonExpression
    ;


logicalAndExpression
    : equalityExpression
    | logicalAndExpression '&&' equalityExpression
    ;

logicalXorExpression
    : logicalAndExpression
    | logicalXorExpression '^^' logicalAndExpression
    ;

logicalOrExpression
    : logicalXorExpression
    | logicalOrExpression '||' logicalXorExpression
    ;

ifExpression
    : 'if' conditionalBranch
    | 'if' conditionalBranch 'else' expression
    ;

conditionalBranch
    : '(' expression ')' expression
    ;

forLoop
    : 'for' label? '(' forPredicate ')' expression
    ;

forPredicate
    : iterableBinding
    | expression
    | bindingList ';' expression
    | expression ';' expression
    | bindingList ';' expression ';' expression
    ;

iterableBinding
    : iterableBindingLeft '..' expression
    | iterableBindingLeft '&' '..' expression
    | iterableBindingLeft '=>' expression
    ;

iterableBindingLeft
    : 'const' bindingName
    | 'const' bindingName ':' type
    | 'let' bindingName 
    | 'let' bindingName ':' type
    ;

label
    : ':' name
    ;



switchExpression
    : 'switch' label? '(' expression ')' '{' switchCases '}'
    ;

switchCases
    : switchCase switchCases?
    ;

switchCase
    : 'case' caseOptions? '=>' expressionList
    ;

caseOptions
    : expression ','
    | expression ',' caseOptions
    ;


matchExpression
    : 'match' label? '(' expression ')' '{' matchCases '}'
    ;


matchCases
    : matchCase matchCases?
    ;

matchCase
    : 'case' 'let' name ':' type '=>' expressionList
    | 'case' nullableType? '=>' expressionList
    ;



assignment
    : assignmentTarget assignmentOp expression
    | slicingAssignmentTarget '=' expression
    ;

assignmentTarget
    : objectSelector
    | indirectAssignmentTarget
    ;


indirectAssignmentTarget
    : '*' unaryExpression
    ;

slicingAssignmentTarget
    : slicingExpression
    | '(' slicingAssignmentTarget ')'
    ;


assignmentOp
    : '='
    | '+='
    | '-='
    | '*='
    | '/='
    | '%='
    | '<<='
    | '>>='
    | '&='
    | '|='
    | '^'
    | '&&='
    | '||='
    | '^^'
    ;


bindingList
    : 'static'?  'let'  bindings 
    | 'static'?  'const'  bindings 
    | 'def'  bindings 
    ;

bindings
    : binding  ','? 
    | binding  ','  bindings 
    ;

binding
    : bindingName  '='  expression 
    | bindingName  ':'  type  '='  expression 
    ;

bindingName
    : name 
    | '('  tupleBindingNames  ')' 
    ;

tupleBindingNames
    : tupleBindingName  ','  tupleBindingName  ','? 
    | tupleBindingName  ','  tupleBindingNames 
    ;

tupleBindingName
    : name 
    | '_' 
    ;

deferExpression
    : 'defer'  expression 
    ;

expressionList
    : expression  ';'  expressionList? 
    | bindingList  ';'  expressionList? 
    | deferExpression  ';'  expressionList? 
    ;

compoundExpression
    : label? '{' expressionList '}' 
    ;

controlExpression
    : 'break'  label? 
    | 'continue'  label? 
    | 'return'  expression? 
    | yieldExpression 
    ;

yieldExpression
    : 'yield' 
    | 'yield'  expression 
    | 'yield'  label 
    | 'yield'  label  ','  expression 
    ;

expression
    : assignment 
    | logicalOrExpression 
    | ifExpression 
    | forLoop 
    | controlExpression 
    ;


// UNIT


declarations
    : 'export'?  declaration  ';'  declarations? 
    | staticAssertionExpression  ';'  declarations? 
    ;
declaration
    : globalDeclaration 
    | constantDeclaration 
    | typeDeclaration 
    | functionDeclaration 
    ;
globalDeclaration
    : 'let'  globalBindings 
    | 'const'  globalBindings 
    ;
globalBindings
    : globalBinding  ','? 
    | globalBinding  ','  globalBindings 
    ;
globalBinding
    : declAttr?  '@threadlocal'?  identifier  ':'  type 
    | declAttr?  '@threadlocal'?  identifier  ':'  type  '='  expression 
    | declAttr?  '@threadlocal'?  identifier  '='  expression 
    ;
declAttr
    : '@symbol'  '('  stringLiteral  ')' 
    ;
constantDeclaration
    : 'def'  constantBindings 
    ;
constantBindings
    : constantBinding  ','? 
    | constantBinding  ','  constantBindings 
    ;
constantBinding
    : identifier  ':'  type  '='  expression 
    | identifier  '='  expression 
    ;
typeDeclaration
    : 'type'  typeBindings 
    ;
typeBindings
    : typeBinding  ','? 
    | typeBinding  ','  typeBindings 
    ;
typeBinding
    : identifier  '='  type 
    | identifier  '='  enumType 
    ;
enumType
    : 'enum'  enumStorage?  '{'  enumValues  '}' 
    ;
enumValues
    : enumValue  ','? 
    | enumValue  ','  enumValues 
    ;
enumValue
    : name 
    | name  '='  expression 
    ;
enumStorage
    : integerType 
    | 'rune' 
    ;
functionDeclaration
    : fndecAttr?  'fn'  identifier  prototype 
    | fndecAttr?  'fn'  identifier  prototype  '='  expression 
    ;
fndecAttr
    : '@fini' 
    | '@init' 
    | '@test' 
    | declAttr 
    ;
subUnit
    : imports?  declarations? 
    ;
imports
    : useDirective  imports? 
    ;
useDirective
    : 'use'  identifier  ';' 
    | 'use'  name  '='  identifier  ';' 
    | 'use'  identifier  '::'  '{'  memberList  '}'  ';' 
    | 'use'  identifier  '::'  '*'  ';' 
    ;
memberList
    : name  ','? 
    | name  ','  memberList 
    ;


// :)
start : subUnit EOF;

// Other

WHITESPACE
    : [ \t]+ -> skip
    ;

NEWLINE
    : ('\r' '\n'? | '\n') -> skip
    ;
