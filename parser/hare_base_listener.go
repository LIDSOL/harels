// Code generated from Hare.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hare

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHareListener is a complete listener for a parse tree produced by HareParser.
type BaseHareListener struct{}

var _ HareListener = &BaseHareListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHareListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHareListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHareListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHareListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAttributes is called when production attributes is entered.
func (s *BaseHareListener) EnterAttributes(ctx *AttributesContext) {}

// ExitAttributes is called when production attributes is exited.
func (s *BaseHareListener) ExitAttributes(ctx *AttributesContext) {}

// EnterInvalidAttribute is called when production invalidAttribute is entered.
func (s *BaseHareListener) EnterInvalidAttribute(ctx *InvalidAttributeContext) {}

// ExitInvalidAttribute is called when production invalidAttribute is exited.
func (s *BaseHareListener) ExitInvalidAttribute(ctx *InvalidAttributeContext) {}

// EnterType is called when production type is entered.
func (s *BaseHareListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseHareListener) ExitType(ctx *TypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseHareListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseHareListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterStructUnionType is called when production structUnionType is entered.
func (s *BaseHareListener) EnterStructUnionType(ctx *StructUnionTypeContext) {}

// ExitStructUnionType is called when production structUnionType is exited.
func (s *BaseHareListener) ExitStructUnionType(ctx *StructUnionTypeContext) {}

// EnterStructUnionFields is called when production structUnionFields is entered.
func (s *BaseHareListener) EnterStructUnionFields(ctx *StructUnionFieldsContext) {}

// ExitStructUnionFields is called when production structUnionFields is exited.
func (s *BaseHareListener) ExitStructUnionFields(ctx *StructUnionFieldsContext) {}

// EnterStructUnionField is called when production structUnionField is entered.
func (s *BaseHareListener) EnterStructUnionField(ctx *StructUnionFieldContext) {}

// ExitStructUnionField is called when production structUnionField is exited.
func (s *BaseHareListener) ExitStructUnionField(ctx *StructUnionFieldContext) {}

// EnterTupleType is called when production tupleType is entered.
func (s *BaseHareListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production tupleType is exited.
func (s *BaseHareListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterTupleTypes is called when production tupleTypes is entered.
func (s *BaseHareListener) EnterTupleTypes(ctx *TupleTypesContext) {}

// ExitTupleTypes is called when production tupleTypes is exited.
func (s *BaseHareListener) ExitTupleTypes(ctx *TupleTypesContext) {}

// EnterTaggedUnionType is called when production taggedUnionType is entered.
func (s *BaseHareListener) EnterTaggedUnionType(ctx *TaggedUnionTypeContext) {}

// ExitTaggedUnionType is called when production taggedUnionType is exited.
func (s *BaseHareListener) ExitTaggedUnionType(ctx *TaggedUnionTypeContext) {}

// EnterTaggedTypes is called when production taggedTypes is entered.
func (s *BaseHareListener) EnterTaggedTypes(ctx *TaggedTypesContext) {}

// ExitTaggedTypes is called when production taggedTypes is exited.
func (s *BaseHareListener) ExitTaggedTypes(ctx *TaggedTypesContext) {}

// EnterSliceArrayType is called when production sliceArrayType is entered.
func (s *BaseHareListener) EnterSliceArrayType(ctx *SliceArrayTypeContext) {}

// ExitSliceArrayType is called when production sliceArrayType is exited.
func (s *BaseHareListener) ExitSliceArrayType(ctx *SliceArrayTypeContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseHareListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseHareListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterPrototype is called when production prototype is entered.
func (s *BaseHareListener) EnterPrototype(ctx *PrototypeContext) {}

// ExitPrototype is called when production prototype is exited.
func (s *BaseHareListener) ExitPrototype(ctx *PrototypeContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseHareListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseHareListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseHareListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseHareListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseHareListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseHareListener) ExitParameter(ctx *ParameterContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseHareListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseHareListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterAliasType is called when production aliasType is entered.
func (s *BaseHareListener) EnterAliasType(ctx *AliasTypeContext) {}

// ExitAliasType is called when production aliasType is exited.
func (s *BaseHareListener) ExitAliasType(ctx *AliasTypeContext) {}

// EnterUnwrappedAlias is called when production unwrappedAlias is entered.
func (s *BaseHareListener) EnterUnwrappedAlias(ctx *UnwrappedAliasContext) {}

// ExitUnwrappedAlias is called when production unwrappedAlias is exited.
func (s *BaseHareListener) ExitUnwrappedAlias(ctx *UnwrappedAliasContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseHareListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseHareListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterFloatingType is called when production floatingType is entered.
func (s *BaseHareListener) EnterFloatingType(ctx *FloatingTypeContext) {}

// ExitFloatingType is called when production floatingType is exited.
func (s *BaseHareListener) ExitFloatingType(ctx *FloatingTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseHareListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseHareListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterStorageClass is called when production storageClass is entered.
func (s *BaseHareListener) EnterStorageClass(ctx *StorageClassContext) {}

// ExitStorageClass is called when production storageClass is exited.
func (s *BaseHareListener) ExitStorageClass(ctx *StorageClassContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseHareListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseHareListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseHareListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseHareListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFloatingLiteral is called when production floatingLiteral is entered.
func (s *BaseHareListener) EnterFloatingLiteral(ctx *FloatingLiteralContext) {}

// ExitFloatingLiteral is called when production floatingLiteral is exited.
func (s *BaseHareListener) ExitFloatingLiteral(ctx *FloatingLiteralContext) {}

// EnterFloatingSuffix is called when production floatingSuffix is entered.
func (s *BaseHareListener) EnterFloatingSuffix(ctx *FloatingSuffixContext) {}

// ExitFloatingSuffix is called when production floatingSuffix is exited.
func (s *BaseHareListener) ExitFloatingSuffix(ctx *FloatingSuffixContext) {}

// EnterNonzeroDecimalDigits is called when production nonzeroDecimalDigits is entered.
func (s *BaseHareListener) EnterNonzeroDecimalDigits(ctx *NonzeroDecimalDigitsContext) {}

// ExitNonzeroDecimalDigits is called when production nonzeroDecimalDigits is exited.
func (s *BaseHareListener) ExitNonzeroDecimalDigits(ctx *NonzeroDecimalDigitsContext) {}

// EnterNonzeroDecimalDigit is called when production nonzeroDecimalDigit is entered.
func (s *BaseHareListener) EnterNonzeroDecimalDigit(ctx *NonzeroDecimalDigitContext) {}

// ExitNonzeroDecimalDigit is called when production nonzeroDecimalDigit is exited.
func (s *BaseHareListener) ExitNonzeroDecimalDigit(ctx *NonzeroDecimalDigitContext) {}

// EnterHexDigits is called when production hexDigits is entered.
func (s *BaseHareListener) EnterHexDigits(ctx *HexDigitsContext) {}

// ExitHexDigits is called when production hexDigits is exited.
func (s *BaseHareListener) ExitHexDigits(ctx *HexDigitsContext) {}

// EnterHexDigit is called when production hexDigit is entered.
func (s *BaseHareListener) EnterHexDigit(ctx *HexDigitContext) {}

// ExitHexDigit is called when production hexDigit is exited.
func (s *BaseHareListener) ExitHexDigit(ctx *HexDigitContext) {}

// EnterDecimalExponent is called when production decimalExponent is entered.
func (s *BaseHareListener) EnterDecimalExponent(ctx *DecimalExponentContext) {}

// ExitDecimalExponent is called when production decimalExponent is exited.
func (s *BaseHareListener) ExitDecimalExponent(ctx *DecimalExponentContext) {}

// EnterBinaryExponent is called when production binaryExponent is entered.
func (s *BaseHareListener) EnterBinaryExponent(ctx *BinaryExponentContext) {}

// ExitBinaryExponent is called when production binaryExponent is exited.
func (s *BaseHareListener) ExitBinaryExponent(ctx *BinaryExponentContext) {}

// EnterSign is called when production sign is entered.
func (s *BaseHareListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BaseHareListener) ExitSign(ctx *SignContext) {}

// EnterBinaryExponentChar is called when production binaryExponentChar is entered.
func (s *BaseHareListener) EnterBinaryExponentChar(ctx *BinaryExponentCharContext) {}

// ExitBinaryExponentChar is called when production binaryExponentChar is exited.
func (s *BaseHareListener) ExitBinaryExponentChar(ctx *BinaryExponentCharContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseHareListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseHareListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterIntegerSuffix is called when production integerSuffix is entered.
func (s *BaseHareListener) EnterIntegerSuffix(ctx *IntegerSuffixContext) {}

// ExitIntegerSuffix is called when production integerSuffix is exited.
func (s *BaseHareListener) ExitIntegerSuffix(ctx *IntegerSuffixContext) {}

// EnterBinaryDigit is called when production binaryDigit is entered.
func (s *BaseHareListener) EnterBinaryDigit(ctx *BinaryDigitContext) {}

// ExitBinaryDigit is called when production binaryDigit is exited.
func (s *BaseHareListener) ExitBinaryDigit(ctx *BinaryDigitContext) {}

// EnterOctalDigit is called when production octalDigit is entered.
func (s *BaseHareListener) EnterOctalDigit(ctx *OctalDigitContext) {}

// ExitOctalDigit is called when production octalDigit is exited.
func (s *BaseHareListener) ExitOctalDigit(ctx *OctalDigitContext) {}

// EnterBinaryDigits is called when production binaryDigits is entered.
func (s *BaseHareListener) EnterBinaryDigits(ctx *BinaryDigitsContext) {}

// ExitBinaryDigits is called when production binaryDigits is exited.
func (s *BaseHareListener) ExitBinaryDigits(ctx *BinaryDigitsContext) {}

// EnterOctalDigits is called when production octalDigits is entered.
func (s *BaseHareListener) EnterOctalDigits(ctx *OctalDigitsContext) {}

// ExitOctalDigits is called when production octalDigits is exited.
func (s *BaseHareListener) ExitOctalDigits(ctx *OctalDigitsContext) {}

// EnterPositiveDecimalExponent is called when production positiveDecimalExponent is entered.
func (s *BaseHareListener) EnterPositiveDecimalExponent(ctx *PositiveDecimalExponentContext) {}

// ExitPositiveDecimalExponent is called when production positiveDecimalExponent is exited.
func (s *BaseHareListener) ExitPositiveDecimalExponent(ctx *PositiveDecimalExponentContext) {}

// EnterRuneLiteral is called when production runeLiteral is entered.
func (s *BaseHareListener) EnterRuneLiteral(ctx *RuneLiteralContext) {}

// ExitRuneLiteral is called when production runeLiteral is exited.
func (s *BaseHareListener) ExitRuneLiteral(ctx *RuneLiteralContext) {}

// EnterRune is called when production rune is entered.
func (s *BaseHareListener) EnterRune(ctx *RuneContext) {}

// ExitRune is called when production rune is exited.
func (s *BaseHareListener) ExitRune(ctx *RuneContext) {}

// EnterEscapeSequence is called when production escapeSequence is entered.
func (s *BaseHareListener) EnterEscapeSequence(ctx *EscapeSequenceContext) {}

// ExitEscapeSequence is called when production escapeSequence is exited.
func (s *BaseHareListener) ExitEscapeSequence(ctx *EscapeSequenceContext) {}

// EnterFourbyte is called when production fourbyte is entered.
func (s *BaseHareListener) EnterFourbyte(ctx *FourbyteContext) {}

// ExitFourbyte is called when production fourbyte is exited.
func (s *BaseHareListener) ExitFourbyte(ctx *FourbyteContext) {}

// EnterEightbyte is called when production eightbyte is entered.
func (s *BaseHareListener) EnterEightbyte(ctx *EightbyteContext) {}

// ExitEightbyte is called when production eightbyte is exited.
func (s *BaseHareListener) ExitEightbyte(ctx *EightbyteContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseHareListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseHareListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterStringSection is called when production stringSection is entered.
func (s *BaseHareListener) EnterStringSection(ctx *StringSectionContext) {}

// ExitStringSection is called when production stringSection is exited.
func (s *BaseHareListener) ExitStringSection(ctx *StringSectionContext) {}

// EnterStringChars is called when production stringChars is entered.
func (s *BaseHareListener) EnterStringChars(ctx *StringCharsContext) {}

// ExitStringChars is called when production stringChars is exited.
func (s *BaseHareListener) ExitStringChars(ctx *StringCharsContext) {}

// EnterStringChar is called when production stringChar is entered.
func (s *BaseHareListener) EnterStringChar(ctx *StringCharContext) {}

// ExitStringChar is called when production stringChar is exited.
func (s *BaseHareListener) ExitStringChar(ctx *StringCharContext) {}

// EnterRawstringChars is called when production rawstringChars is entered.
func (s *BaseHareListener) EnterRawstringChars(ctx *RawstringCharsContext) {}

// ExitRawstringChars is called when production rawstringChars is exited.
func (s *BaseHareListener) ExitRawstringChars(ctx *RawstringCharsContext) {}

// EnterRawstringChar is called when production rawstringChar is entered.
func (s *BaseHareListener) EnterRawstringChar(ctx *RawstringCharContext) {}

// ExitRawstringChar is called when production rawstringChar is exited.
func (s *BaseHareListener) ExitRawstringChar(ctx *RawstringCharContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseHareListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseHareListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterArrayMembers is called when production arrayMembers is entered.
func (s *BaseHareListener) EnterArrayMembers(ctx *ArrayMembersContext) {}

// ExitArrayMembers is called when production arrayMembers is exited.
func (s *BaseHareListener) ExitArrayMembers(ctx *ArrayMembersContext) {}

// EnterStructLiteral is called when production structLiteral is entered.
func (s *BaseHareListener) EnterStructLiteral(ctx *StructLiteralContext) {}

// ExitStructLiteral is called when production structLiteral is exited.
func (s *BaseHareListener) ExitStructLiteral(ctx *StructLiteralContext) {}

// EnterStructInitializer is called when production structInitializer is entered.
func (s *BaseHareListener) EnterStructInitializer(ctx *StructInitializerContext) {}

// ExitStructInitializer is called when production structInitializer is exited.
func (s *BaseHareListener) ExitStructInitializer(ctx *StructInitializerContext) {}

// EnterFieldValues is called when production fieldValues is entered.
func (s *BaseHareListener) EnterFieldValues(ctx *FieldValuesContext) {}

// ExitFieldValues is called when production fieldValues is exited.
func (s *BaseHareListener) ExitFieldValues(ctx *FieldValuesContext) {}

// EnterFieldValue is called when production fieldValue is entered.
func (s *BaseHareListener) EnterFieldValue(ctx *FieldValueContext) {}

// ExitFieldValue is called when production fieldValue is exited.
func (s *BaseHareListener) ExitFieldValue(ctx *FieldValueContext) {}

// EnterTupleLiteral is called when production tupleLiteral is entered.
func (s *BaseHareListener) EnterTupleLiteral(ctx *TupleLiteralContext) {}

// ExitTupleLiteral is called when production tupleLiteral is exited.
func (s *BaseHareListener) ExitTupleLiteral(ctx *TupleLiteralContext) {}

// EnterTupleItems is called when production tupleItems is entered.
func (s *BaseHareListener) EnterTupleItems(ctx *TupleItemsContext) {}

// ExitTupleItems is called when production tupleItems is exited.
func (s *BaseHareListener) ExitTupleItems(ctx *TupleItemsContext) {}

// EnterPlainExpression is called when production plainExpression is entered.
func (s *BaseHareListener) EnterPlainExpression(ctx *PlainExpressionContext) {}

// ExitPlainExpression is called when production plainExpression is exited.
func (s *BaseHareListener) ExitPlainExpression(ctx *PlainExpressionContext) {}

// EnterNestedExpression is called when production nestedExpression is entered.
func (s *BaseHareListener) EnterNestedExpression(ctx *NestedExpressionContext) {}

// ExitNestedExpression is called when production nestedExpression is exited.
func (s *BaseHareListener) ExitNestedExpression(ctx *NestedExpressionContext) {}

// EnterAllocationExpression is called when production allocationExpression is entered.
func (s *BaseHareListener) EnterAllocationExpression(ctx *AllocationExpressionContext) {}

// ExitAllocationExpression is called when production allocationExpression is exited.
func (s *BaseHareListener) ExitAllocationExpression(ctx *AllocationExpressionContext) {}

// EnterFreeExpression is called when production freeExpression is entered.
func (s *BaseHareListener) EnterFreeExpression(ctx *FreeExpressionContext) {}

// ExitFreeExpression is called when production freeExpression is exited.
func (s *BaseHareListener) ExitFreeExpression(ctx *FreeExpressionContext) {}

// EnterAssertionExpression is called when production assertionExpression is entered.
func (s *BaseHareListener) EnterAssertionExpression(ctx *AssertionExpressionContext) {}

// ExitAssertionExpression is called when production assertionExpression is exited.
func (s *BaseHareListener) ExitAssertionExpression(ctx *AssertionExpressionContext) {}

// EnterStaticAssertionExpression is called when production staticAssertionExpression is entered.
func (s *BaseHareListener) EnterStaticAssertionExpression(ctx *StaticAssertionExpressionContext) {}

// ExitStaticAssertionExpression is called when production staticAssertionExpression is exited.
func (s *BaseHareListener) ExitStaticAssertionExpression(ctx *StaticAssertionExpressionContext) {}

// EnterCallExpression is called when production callExpression is entered.
func (s *BaseHareListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production callExpression is exited.
func (s *BaseHareListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseHareListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseHareListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterMeasurementExpression is called when production measurementExpression is entered.
func (s *BaseHareListener) EnterMeasurementExpression(ctx *MeasurementExpressionContext) {}

// ExitMeasurementExpression is called when production measurementExpression is exited.
func (s *BaseHareListener) ExitMeasurementExpression(ctx *MeasurementExpressionContext) {}

// EnterAlignExpression is called when production alignExpression is entered.
func (s *BaseHareListener) EnterAlignExpression(ctx *AlignExpressionContext) {}

// ExitAlignExpression is called when production alignExpression is exited.
func (s *BaseHareListener) ExitAlignExpression(ctx *AlignExpressionContext) {}

// EnterSizeExpression is called when production sizeExpression is entered.
func (s *BaseHareListener) EnterSizeExpression(ctx *SizeExpressionContext) {}

// ExitSizeExpression is called when production sizeExpression is exited.
func (s *BaseHareListener) ExitSizeExpression(ctx *SizeExpressionContext) {}

// EnterLengthExpression is called when production lengthExpression is entered.
func (s *BaseHareListener) EnterLengthExpression(ctx *LengthExpressionContext) {}

// ExitLengthExpression is called when production lengthExpression is exited.
func (s *BaseHareListener) ExitLengthExpression(ctx *LengthExpressionContext) {}

// EnterOffsetExpression is called when production offsetExpression is entered.
func (s *BaseHareListener) EnterOffsetExpression(ctx *OffsetExpressionContext) {}

// ExitOffsetExpression is called when production offsetExpression is exited.
func (s *BaseHareListener) ExitOffsetExpression(ctx *OffsetExpressionContext) {}

// EnterOffsetOperand is called when production offsetOperand is entered.
func (s *BaseHareListener) EnterOffsetOperand(ctx *OffsetOperandContext) {}

// ExitOffsetOperand is called when production offsetOperand is exited.
func (s *BaseHareListener) ExitOffsetOperand(ctx *OffsetOperandContext) {}

// EnterFieldAccessExpression is called when production fieldAccessExpression is entered.
func (s *BaseHareListener) EnterFieldAccessExpression(ctx *FieldAccessExpressionContext) {}

// ExitFieldAccessExpression is called when production fieldAccessExpression is exited.
func (s *BaseHareListener) ExitFieldAccessExpression(ctx *FieldAccessExpressionContext) {}

// EnterIndexingExpression is called when production indexingExpression is entered.
func (s *BaseHareListener) EnterIndexingExpression(ctx *IndexingExpressionContext) {}

// ExitIndexingExpression is called when production indexingExpression is exited.
func (s *BaseHareListener) ExitIndexingExpression(ctx *IndexingExpressionContext) {}

// EnterSlicingExpression is called when production slicingExpression is entered.
func (s *BaseHareListener) EnterSlicingExpression(ctx *SlicingExpressionContext) {}

// ExitSlicingExpression is called when production slicingExpression is exited.
func (s *BaseHareListener) ExitSlicingExpression(ctx *SlicingExpressionContext) {}

// EnterSliceMutationExpression is called when production sliceMutationExpression is entered.
func (s *BaseHareListener) EnterSliceMutationExpression(ctx *SliceMutationExpressionContext) {}

// ExitSliceMutationExpression is called when production sliceMutationExpression is exited.
func (s *BaseHareListener) ExitSliceMutationExpression(ctx *SliceMutationExpressionContext) {}

// EnterAppendExpression is called when production appendExpression is entered.
func (s *BaseHareListener) EnterAppendExpression(ctx *AppendExpressionContext) {}

// ExitAppendExpression is called when production appendExpression is exited.
func (s *BaseHareListener) ExitAppendExpression(ctx *AppendExpressionContext) {}

// EnterInsertExpression is called when production insertExpression is entered.
func (s *BaseHareListener) EnterInsertExpression(ctx *InsertExpressionContext) {}

// ExitInsertExpression is called when production insertExpression is exited.
func (s *BaseHareListener) ExitInsertExpression(ctx *InsertExpressionContext) {}

// EnterInsertOperand is called when production insertOperand is entered.
func (s *BaseHareListener) EnterInsertOperand(ctx *InsertOperandContext) {}

// ExitInsertOperand is called when production insertOperand is exited.
func (s *BaseHareListener) ExitInsertOperand(ctx *InsertOperandContext) {}

// EnterDeleteExpression is called when production deleteExpression is entered.
func (s *BaseHareListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production deleteExpression is exited.
func (s *BaseHareListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterDeleteOperand is called when production deleteOperand is entered.
func (s *BaseHareListener) EnterDeleteOperand(ctx *DeleteOperandContext) {}

// ExitDeleteOperand is called when production deleteOperand is exited.
func (s *BaseHareListener) ExitDeleteOperand(ctx *DeleteOperandContext) {}

// EnterErrorCheckingExpression is called when production errorCheckingExpression is entered.
func (s *BaseHareListener) EnterErrorCheckingExpression(ctx *ErrorCheckingExpressionContext) {}

// ExitErrorCheckingExpression is called when production errorCheckingExpression is exited.
func (s *BaseHareListener) ExitErrorCheckingExpression(ctx *ErrorCheckingExpressionContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseHareListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseHareListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterObjectSelector is called when production objectSelector is entered.
func (s *BaseHareListener) EnterObjectSelector(ctx *ObjectSelectorContext) {}

// ExitObjectSelector is called when production objectSelector is exited.
func (s *BaseHareListener) ExitObjectSelector(ctx *ObjectSelectorContext) {}

// EnterVariadicExpression is called when production variadicExpression is entered.
func (s *BaseHareListener) EnterVariadicExpression(ctx *VariadicExpressionContext) {}

// ExitVariadicExpression is called when production variadicExpression is exited.
func (s *BaseHareListener) ExitVariadicExpression(ctx *VariadicExpressionContext) {}

// EnterBuiltinExpression is called when production builtinExpression is entered.
func (s *BaseHareListener) EnterBuiltinExpression(ctx *BuiltinExpressionContext) {}

// ExitBuiltinExpression is called when production builtinExpression is exited.
func (s *BaseHareListener) ExitBuiltinExpression(ctx *BuiltinExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseHareListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseHareListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseHareListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseHareListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseHareListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseHareListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterNullableType is called when production nullableType is entered.
func (s *BaseHareListener) EnterNullableType(ctx *NullableTypeContext) {}

// ExitNullableType is called when production nullableType is exited.
func (s *BaseHareListener) ExitNullableType(ctx *NullableTypeContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseHareListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseHareListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseHareListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseHareListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseHareListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseHareListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseHareListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseHareListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseHareListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseHareListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseHareListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseHareListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseHareListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseHareListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseHareListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseHareListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseHareListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseHareListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalXorExpression is called when production logicalXorExpression is entered.
func (s *BaseHareListener) EnterLogicalXorExpression(ctx *LogicalXorExpressionContext) {}

// ExitLogicalXorExpression is called when production logicalXorExpression is exited.
func (s *BaseHareListener) ExitLogicalXorExpression(ctx *LogicalXorExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseHareListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseHareListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterIfExpression is called when production ifExpression is entered.
func (s *BaseHareListener) EnterIfExpression(ctx *IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *BaseHareListener) ExitIfExpression(ctx *IfExpressionContext) {}

// EnterConditionalBranch is called when production conditionalBranch is entered.
func (s *BaseHareListener) EnterConditionalBranch(ctx *ConditionalBranchContext) {}

// ExitConditionalBranch is called when production conditionalBranch is exited.
func (s *BaseHareListener) ExitConditionalBranch(ctx *ConditionalBranchContext) {}

// EnterForLoop is called when production forLoop is entered.
func (s *BaseHareListener) EnterForLoop(ctx *ForLoopContext) {}

// ExitForLoop is called when production forLoop is exited.
func (s *BaseHareListener) ExitForLoop(ctx *ForLoopContext) {}

// EnterForPredicate is called when production forPredicate is entered.
func (s *BaseHareListener) EnterForPredicate(ctx *ForPredicateContext) {}

// ExitForPredicate is called when production forPredicate is exited.
func (s *BaseHareListener) ExitForPredicate(ctx *ForPredicateContext) {}

// EnterIterableBinding is called when production iterableBinding is entered.
func (s *BaseHareListener) EnterIterableBinding(ctx *IterableBindingContext) {}

// ExitIterableBinding is called when production iterableBinding is exited.
func (s *BaseHareListener) ExitIterableBinding(ctx *IterableBindingContext) {}

// EnterIterableBindingLeft is called when production iterableBindingLeft is entered.
func (s *BaseHareListener) EnterIterableBindingLeft(ctx *IterableBindingLeftContext) {}

// ExitIterableBindingLeft is called when production iterableBindingLeft is exited.
func (s *BaseHareListener) ExitIterableBindingLeft(ctx *IterableBindingLeftContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseHareListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseHareListener) ExitLabel(ctx *LabelContext) {}

// EnterSwitchExpression is called when production switchExpression is entered.
func (s *BaseHareListener) EnterSwitchExpression(ctx *SwitchExpressionContext) {}

// ExitSwitchExpression is called when production switchExpression is exited.
func (s *BaseHareListener) ExitSwitchExpression(ctx *SwitchExpressionContext) {}

// EnterSwitchCases is called when production switchCases is entered.
func (s *BaseHareListener) EnterSwitchCases(ctx *SwitchCasesContext) {}

// ExitSwitchCases is called when production switchCases is exited.
func (s *BaseHareListener) ExitSwitchCases(ctx *SwitchCasesContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseHareListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseHareListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterCaseOptions is called when production caseOptions is entered.
func (s *BaseHareListener) EnterCaseOptions(ctx *CaseOptionsContext) {}

// ExitCaseOptions is called when production caseOptions is exited.
func (s *BaseHareListener) ExitCaseOptions(ctx *CaseOptionsContext) {}

// EnterMatchExpression is called when production matchExpression is entered.
func (s *BaseHareListener) EnterMatchExpression(ctx *MatchExpressionContext) {}

// ExitMatchExpression is called when production matchExpression is exited.
func (s *BaseHareListener) ExitMatchExpression(ctx *MatchExpressionContext) {}

// EnterMatchCases is called when production matchCases is entered.
func (s *BaseHareListener) EnterMatchCases(ctx *MatchCasesContext) {}

// ExitMatchCases is called when production matchCases is exited.
func (s *BaseHareListener) ExitMatchCases(ctx *MatchCasesContext) {}

// EnterMatchCase is called when production matchCase is entered.
func (s *BaseHareListener) EnterMatchCase(ctx *MatchCaseContext) {}

// ExitMatchCase is called when production matchCase is exited.
func (s *BaseHareListener) ExitMatchCase(ctx *MatchCaseContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseHareListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseHareListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentTarget is called when production assignmentTarget is entered.
func (s *BaseHareListener) EnterAssignmentTarget(ctx *AssignmentTargetContext) {}

// ExitAssignmentTarget is called when production assignmentTarget is exited.
func (s *BaseHareListener) ExitAssignmentTarget(ctx *AssignmentTargetContext) {}

// EnterIndirectAssignmentTarget is called when production indirectAssignmentTarget is entered.
func (s *BaseHareListener) EnterIndirectAssignmentTarget(ctx *IndirectAssignmentTargetContext) {}

// ExitIndirectAssignmentTarget is called when production indirectAssignmentTarget is exited.
func (s *BaseHareListener) ExitIndirectAssignmentTarget(ctx *IndirectAssignmentTargetContext) {}

// EnterSlicingAssignmentTarget is called when production slicingAssignmentTarget is entered.
func (s *BaseHareListener) EnterSlicingAssignmentTarget(ctx *SlicingAssignmentTargetContext) {}

// ExitSlicingAssignmentTarget is called when production slicingAssignmentTarget is exited.
func (s *BaseHareListener) ExitSlicingAssignmentTarget(ctx *SlicingAssignmentTargetContext) {}

// EnterAssignmentOp is called when production assignmentOp is entered.
func (s *BaseHareListener) EnterAssignmentOp(ctx *AssignmentOpContext) {}

// ExitAssignmentOp is called when production assignmentOp is exited.
func (s *BaseHareListener) ExitAssignmentOp(ctx *AssignmentOpContext) {}

// EnterBindingList is called when production bindingList is entered.
func (s *BaseHareListener) EnterBindingList(ctx *BindingListContext) {}

// ExitBindingList is called when production bindingList is exited.
func (s *BaseHareListener) ExitBindingList(ctx *BindingListContext) {}

// EnterBindings is called when production bindings is entered.
func (s *BaseHareListener) EnterBindings(ctx *BindingsContext) {}

// ExitBindings is called when production bindings is exited.
func (s *BaseHareListener) ExitBindings(ctx *BindingsContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseHareListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseHareListener) ExitBinding(ctx *BindingContext) {}

// EnterBindingName is called when production bindingName is entered.
func (s *BaseHareListener) EnterBindingName(ctx *BindingNameContext) {}

// ExitBindingName is called when production bindingName is exited.
func (s *BaseHareListener) ExitBindingName(ctx *BindingNameContext) {}

// EnterTupleBindingNames is called when production tupleBindingNames is entered.
func (s *BaseHareListener) EnterTupleBindingNames(ctx *TupleBindingNamesContext) {}

// ExitTupleBindingNames is called when production tupleBindingNames is exited.
func (s *BaseHareListener) ExitTupleBindingNames(ctx *TupleBindingNamesContext) {}

// EnterTupleBindingName is called when production tupleBindingName is entered.
func (s *BaseHareListener) EnterTupleBindingName(ctx *TupleBindingNameContext) {}

// ExitTupleBindingName is called when production tupleBindingName is exited.
func (s *BaseHareListener) ExitTupleBindingName(ctx *TupleBindingNameContext) {}

// EnterDeferExpression is called when production deferExpression is entered.
func (s *BaseHareListener) EnterDeferExpression(ctx *DeferExpressionContext) {}

// ExitDeferExpression is called when production deferExpression is exited.
func (s *BaseHareListener) ExitDeferExpression(ctx *DeferExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseHareListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseHareListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterCompoundExpression is called when production compoundExpression is entered.
func (s *BaseHareListener) EnterCompoundExpression(ctx *CompoundExpressionContext) {}

// ExitCompoundExpression is called when production compoundExpression is exited.
func (s *BaseHareListener) ExitCompoundExpression(ctx *CompoundExpressionContext) {}

// EnterControlExpression is called when production controlExpression is entered.
func (s *BaseHareListener) EnterControlExpression(ctx *ControlExpressionContext) {}

// ExitControlExpression is called when production controlExpression is exited.
func (s *BaseHareListener) ExitControlExpression(ctx *ControlExpressionContext) {}

// EnterYieldExpression is called when production yieldExpression is entered.
func (s *BaseHareListener) EnterYieldExpression(ctx *YieldExpressionContext) {}

// ExitYieldExpression is called when production yieldExpression is exited.
func (s *BaseHareListener) ExitYieldExpression(ctx *YieldExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHareListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHareListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseHareListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseHareListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseHareListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseHareListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterGlobalDeclaration is called when production globalDeclaration is entered.
func (s *BaseHareListener) EnterGlobalDeclaration(ctx *GlobalDeclarationContext) {}

// ExitGlobalDeclaration is called when production globalDeclaration is exited.
func (s *BaseHareListener) ExitGlobalDeclaration(ctx *GlobalDeclarationContext) {}

// EnterGlobalBindings is called when production globalBindings is entered.
func (s *BaseHareListener) EnterGlobalBindings(ctx *GlobalBindingsContext) {}

// ExitGlobalBindings is called when production globalBindings is exited.
func (s *BaseHareListener) ExitGlobalBindings(ctx *GlobalBindingsContext) {}

// EnterGlobalBinding is called when production globalBinding is entered.
func (s *BaseHareListener) EnterGlobalBinding(ctx *GlobalBindingContext) {}

// ExitGlobalBinding is called when production globalBinding is exited.
func (s *BaseHareListener) ExitGlobalBinding(ctx *GlobalBindingContext) {}

// EnterDeclAttr is called when production declAttr is entered.
func (s *BaseHareListener) EnterDeclAttr(ctx *DeclAttrContext) {}

// ExitDeclAttr is called when production declAttr is exited.
func (s *BaseHareListener) ExitDeclAttr(ctx *DeclAttrContext) {}

// EnterConstantDeclaration is called when production constantDeclaration is entered.
func (s *BaseHareListener) EnterConstantDeclaration(ctx *ConstantDeclarationContext) {}

// ExitConstantDeclaration is called when production constantDeclaration is exited.
func (s *BaseHareListener) ExitConstantDeclaration(ctx *ConstantDeclarationContext) {}

// EnterConstantBindings is called when production constantBindings is entered.
func (s *BaseHareListener) EnterConstantBindings(ctx *ConstantBindingsContext) {}

// ExitConstantBindings is called when production constantBindings is exited.
func (s *BaseHareListener) ExitConstantBindings(ctx *ConstantBindingsContext) {}

// EnterConstantBinding is called when production constantBinding is entered.
func (s *BaseHareListener) EnterConstantBinding(ctx *ConstantBindingContext) {}

// ExitConstantBinding is called when production constantBinding is exited.
func (s *BaseHareListener) ExitConstantBinding(ctx *ConstantBindingContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseHareListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseHareListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterTypeBindings is called when production typeBindings is entered.
func (s *BaseHareListener) EnterTypeBindings(ctx *TypeBindingsContext) {}

// ExitTypeBindings is called when production typeBindings is exited.
func (s *BaseHareListener) ExitTypeBindings(ctx *TypeBindingsContext) {}

// EnterTypeBinding is called when production typeBinding is entered.
func (s *BaseHareListener) EnterTypeBinding(ctx *TypeBindingContext) {}

// ExitTypeBinding is called when production typeBinding is exited.
func (s *BaseHareListener) ExitTypeBinding(ctx *TypeBindingContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseHareListener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseHareListener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterEnumValues is called when production enumValues is entered.
func (s *BaseHareListener) EnterEnumValues(ctx *EnumValuesContext) {}

// ExitEnumValues is called when production enumValues is exited.
func (s *BaseHareListener) ExitEnumValues(ctx *EnumValuesContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseHareListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseHareListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterEnumStorage is called when production enumStorage is entered.
func (s *BaseHareListener) EnterEnumStorage(ctx *EnumStorageContext) {}

// ExitEnumStorage is called when production enumStorage is exited.
func (s *BaseHareListener) ExitEnumStorage(ctx *EnumStorageContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseHareListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseHareListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFndecAttr is called when production fndecAttr is entered.
func (s *BaseHareListener) EnterFndecAttr(ctx *FndecAttrContext) {}

// ExitFndecAttr is called when production fndecAttr is exited.
func (s *BaseHareListener) ExitFndecAttr(ctx *FndecAttrContext) {}

// EnterSubUnit is called when production subUnit is entered.
func (s *BaseHareListener) EnterSubUnit(ctx *SubUnitContext) {}

// ExitSubUnit is called when production subUnit is exited.
func (s *BaseHareListener) ExitSubUnit(ctx *SubUnitContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseHareListener) EnterImports(ctx *ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseHareListener) ExitImports(ctx *ImportsContext) {}

// EnterUseDirective is called when production useDirective is entered.
func (s *BaseHareListener) EnterUseDirective(ctx *UseDirectiveContext) {}

// ExitUseDirective is called when production useDirective is exited.
func (s *BaseHareListener) ExitUseDirective(ctx *UseDirectiveContext) {}

// EnterMemberList is called when production memberList is entered.
func (s *BaseHareListener) EnterMemberList(ctx *MemberListContext) {}

// ExitMemberList is called when production memberList is exited.
func (s *BaseHareListener) ExitMemberList(ctx *MemberListContext) {}

// EnterStart is called when production start is entered.
func (s *BaseHareListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseHareListener) ExitStart(ctx *StartContext) {}
