// Calc.g4
grammar Hare;

// operators :)
Lnot : '!';
Nequal : '!=';
Modulo : '%';
Band : '&';
Land : '&&';
Landeq : '&&=';
Bandeq : '&=';
Lparen : '(';
Rparen : ')';
Times : '*';
Timeseq : '*=';
Plus : '+';
Pluseq : '+=';
Comma : ',';
Minus : '-';
Minuseq : '-=';
Dot : '.';
DoubleDot : '..';
Ellipsis : '...';
Div : '/';
Diveq : '/=';
Colon : ':';
DoubleColon : '::';
Semicolon : ';';
Less : '<';
Lshift : '<<';
Lshifteq : '<<=';
Lesseq : '<=';
Equal : '=';
Lequal : '==';
Arrow : '=>';
Gt : '>';
Gteq : '>=';
Rshift : '>>';
Rshifteq : '>>=';
Question : '?';
Lbracket : '[';
Rbracket : ']';
Bxor : '^';
Bxoreq : '^=';
Lxor : '^^';
Lxoreq : '^^=';
Lbrace : '{';
Rbrace : '}';
Bor : '|';
Boreq : '|=';
Lor : '||';
Loreq : '||=';
Bnot : '~';


attributes
    : '@fini'
    | '@init'
    | '@packed'
    | '@symbol'
    | '@test'
    | '@threadlocal'
    ;


invalidAttribute
    : '@' Name
    ;


type
    : 'const'? '!'? storageClass
    ;

pointerType: Times type | 'nullable' Times type ;

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
    | Name ':' type
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
    | '[' Times ']' type
    | '[' '_' ']' type
    ;

functionType
    : 'fn' prototype
    ;

prototype
    : '(' parameterList? ')' type
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
    : Name ':' type defaultValue?
    | type defaultValue?
    ;

defaultValue
    : Equal expression
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
    | Int
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


// Name: Nondigit | Name Alnum;
Name: [a-zA-Z_][a-zA-Z0-9_]*;

Nondigit
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
    | '_'
;


DecimalDigit
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

Alnum: DecimalDigit | Nondigit;

identifier : Name | Name '::' identifier;




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
    : nonzeroDecimalDigits '.' DecimalDigits decimalExponent? floatingSuffix?
    | nonzeroDecimalDigits decimalExponent? floatingSuffix
    | '0x' hexDigits '.' hexDigits binaryExponent floatingSuffix?
    | '0x' hexDigits binaryExponent floatingSuffix?
    ;

floatingSuffix
    : 'f32'
    | 'f64'
    ;

DecimalDigitsWithoutSeparators
    : DecimalDigit DecimalDigitsWithoutSeparators?
    ;

DecimalDigits
    : DecimalDigit DecimalDigits?
    | DecimalDigit '_' DecimalDigits
    ;

nonzeroDecimalDigits
    : '0'
    | nonzeroDecimalDigit DecimalDigits?
    | nonzeroDecimalDigit '_' DecimalDigits
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
    : DecimalExponentChar sign? DecimalDigitsWithoutSeparators
    ;

binaryExponent
    : binaryExponentChar sign? DecimalDigitsWithoutSeparators
    ;

sign
    : '+'
    | '-'
    ;

DecimalExponentChar
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
    : DecimalExponentChar '+'? DecimalDigitsWithoutSeparators
    ;

runeLiteral
    : '\'' rune '\''
    ;

rune 
    : ~('\\'|'\'')
    | escapeSequence
    ;

escapeSequence
    : NamedEscape
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


NamedEscape
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
    : Name Equal expression
    | Name ':' type Equal expression
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
    : postfixExpression '.' Name
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
    | postfixExpression '.' Name
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
    | Times
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
    | multiplicativeExpression Times castExpression
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
    : ':' Name
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
    : 'case' 'let' Name ':' type '=>' expressionList
    | 'case' nullableType? '=>' expressionList
    ;



assignment
    : assignmentTarget assignmentOp expression
    | slicingAssignmentTarget Equal expression
    ;

assignmentTarget
    : objectSelector
    | indirectAssignmentTarget
    ;


indirectAssignmentTarget
    : Times unaryExpression
    ;

slicingAssignmentTarget
    : slicingExpression
    | '(' slicingAssignmentTarget ')'
    ;


assignmentOp
    : Equal
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
    : bindingName  Equal  expression 
    | bindingName  ':'  type  Equal  expression 
    ;

bindingName
    : Name 
    | '('  tupleBindingNames  ')' 
    ;

tupleBindingNames
    : tupleBindingName  ','  tupleBindingName  ','? 
    | tupleBindingName  ','  tupleBindingNames 
    ;

tupleBindingName
    : Name 
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
    | declAttr?  '@threadlocal'?  identifier  ':'  type  Equal  expression 
    | declAttr?  '@threadlocal'?  identifier  Equal  expression 
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
    : identifier  ':'  type  Equal  expression 
    | identifier  Equal  expression 
    ;
typeDeclaration
    : 'type'  typeBindings 
    ;
typeBindings
    : typeBinding  ','? 
    | typeBinding  ','  typeBindings 
    ;
typeBinding
    : identifier  Equal  type 
    | identifier  Equal  enumType 
    ;
enumType
    : 'enum'  enumStorage?  '{'  enumValues  '}' 
    ;
enumValues
    : enumValue  ','? 
    | enumValue  ','  enumValues 
    ;
enumValue
    : Name 
    | Name  Equal  expression 
    ;
enumStorage
    : integerType 
    | 'rune' 
    ;
functionDeclaration
    : fndecAttr?  'fn'  identifier  prototype 
    | fndecAttr?  'fn'  identifier  prototype  Equal  expression 
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
    | 'use'  Name  Equal  identifier  ';' 
    | 'use'  identifier  '::'  '{'  memberList  '}'  ';' 
    | 'use'  identifier  '::'  Times  ';' 
    ;
memberList
    : Name  ','? 
    | Name  ','  memberList 
    ;

Int
    : 'int'
    ;

// :)
start : subUnit EOF;

// Other

Whitespace
    : [ \t]+ -> channel(HIDDEN)
    ;

Newline
    : ('\r' '\n'? | '\n') -> channel(HIDDEN)
    ;