// Code generated from Hare.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hare

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HareListener is a complete listener for a parse tree produced by HareParser.
type HareListener interface {
	antlr.ParseTreeListener

	// EnterLnot is called when entering the lnot production.
	EnterLnot(c *LnotContext)

	// EnterNequal is called when entering the nequal production.
	EnterNequal(c *NequalContext)

	// EnterModulo is called when entering the modulo production.
	EnterModulo(c *ModuloContext)

	// EnterBand is called when entering the band production.
	EnterBand(c *BandContext)

	// EnterLand is called when entering the land production.
	EnterLand(c *LandContext)

	// EnterLandeq is called when entering the landeq production.
	EnterLandeq(c *LandeqContext)

	// EnterBandeq is called when entering the bandeq production.
	EnterBandeq(c *BandeqContext)

	// EnterLparen is called when entering the lparen production.
	EnterLparen(c *LparenContext)

	// EnterRparen is called when entering the rparen production.
	EnterRparen(c *RparenContext)

	// EnterTimes is called when entering the times production.
	EnterTimes(c *TimesContext)

	// EnterTimeseq is called when entering the timeseq production.
	EnterTimeseq(c *TimeseqContext)

	// EnterPlus is called when entering the plus production.
	EnterPlus(c *PlusContext)

	// EnterPluseq is called when entering the pluseq production.
	EnterPluseq(c *PluseqContext)

	// EnterComma is called when entering the comma production.
	EnterComma(c *CommaContext)

	// EnterMinus is called when entering the minus production.
	EnterMinus(c *MinusContext)

	// EnterMinuseq is called when entering the minuseq production.
	EnterMinuseq(c *MinuseqContext)

	// EnterDot is called when entering the dot production.
	EnterDot(c *DotContext)

	// EnterDoubleDot is called when entering the doubleDot production.
	EnterDoubleDot(c *DoubleDotContext)

	// EnterEllipsis is called when entering the ellipsis production.
	EnterEllipsis(c *EllipsisContext)

	// EnterDiv is called when entering the div production.
	EnterDiv(c *DivContext)

	// EnterDiveq is called when entering the diveq production.
	EnterDiveq(c *DiveqContext)

	// EnterColon is called when entering the colon production.
	EnterColon(c *ColonContext)

	// EnterDoubleColon is called when entering the doubleColon production.
	EnterDoubleColon(c *DoubleColonContext)

	// EnterSemicolon is called when entering the semicolon production.
	EnterSemicolon(c *SemicolonContext)

	// EnterLess is called when entering the less production.
	EnterLess(c *LessContext)

	// EnterLshift is called when entering the lshift production.
	EnterLshift(c *LshiftContext)

	// EnterLshifteq is called when entering the lshifteq production.
	EnterLshifteq(c *LshifteqContext)

	// EnterLesseq is called when entering the lesseq production.
	EnterLesseq(c *LesseqContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterLequal is called when entering the lequal production.
	EnterLequal(c *LequalContext)

	// EnterArrow is called when entering the arrow production.
	EnterArrow(c *ArrowContext)

	// EnterGt is called when entering the gt production.
	EnterGt(c *GtContext)

	// EnterGteq is called when entering the gteq production.
	EnterGteq(c *GteqContext)

	// EnterRshift is called when entering the rshift production.
	EnterRshift(c *RshiftContext)

	// EnterRshifteq is called when entering the rshifteq production.
	EnterRshifteq(c *RshifteqContext)

	// EnterQuestion is called when entering the question production.
	EnterQuestion(c *QuestionContext)

	// EnterLbracket is called when entering the lbracket production.
	EnterLbracket(c *LbracketContext)

	// EnterRbracket is called when entering the rbracket production.
	EnterRbracket(c *RbracketContext)

	// EnterBxor is called when entering the bxor production.
	EnterBxor(c *BxorContext)

	// EnterBxoreq is called when entering the bxoreq production.
	EnterBxoreq(c *BxoreqContext)

	// EnterLxor is called when entering the lxor production.
	EnterLxor(c *LxorContext)

	// EnterLxoreq is called when entering the lxoreq production.
	EnterLxoreq(c *LxoreqContext)

	// EnterLbrace is called when entering the lbrace production.
	EnterLbrace(c *LbraceContext)

	// EnterRbrace is called when entering the rbrace production.
	EnterRbrace(c *RbraceContext)

	// EnterBor is called when entering the bor production.
	EnterBor(c *BorContext)

	// EnterBoreq is called when entering the boreq production.
	EnterBoreq(c *BoreqContext)

	// EnterLor is called when entering the lor production.
	EnterLor(c *LorContext)

	// EnterLoreq is called when entering the loreq production.
	EnterLoreq(c *LoreqContext)

	// EnterBnot is called when entering the bnot production.
	EnterBnot(c *BnotContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// EnterAttributes is called when entering the attributes production.
	EnterAttributes(c *AttributesContext)

	// EnterInvalidAttribute is called when entering the invalidAttribute production.
	EnterInvalidAttribute(c *InvalidAttributeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterStructUnionType is called when entering the structUnionType production.
	EnterStructUnionType(c *StructUnionTypeContext)

	// EnterStructUnionFields is called when entering the structUnionFields production.
	EnterStructUnionFields(c *StructUnionFieldsContext)

	// EnterStructUnionField is called when entering the structUnionField production.
	EnterStructUnionField(c *StructUnionFieldContext)

	// EnterTupleType is called when entering the tupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterTupleTypes is called when entering the tupleTypes production.
	EnterTupleTypes(c *TupleTypesContext)

	// EnterTaggedUnionType is called when entering the taggedUnionType production.
	EnterTaggedUnionType(c *TaggedUnionTypeContext)

	// EnterTaggedTypes is called when entering the taggedTypes production.
	EnterTaggedTypes(c *TaggedTypesContext)

	// EnterSliceArrayType is called when entering the sliceArrayType production.
	EnterSliceArrayType(c *SliceArrayTypeContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterPrototype is called when entering the prototype production.
	EnterPrototype(c *PrototypeContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterAliasType is called when entering the aliasType production.
	EnterAliasType(c *AliasTypeContext)

	// EnterUnwrappedAlias is called when entering the unwrappedAlias production.
	EnterUnwrappedAlias(c *UnwrappedAliasContext)

	// EnterIntegerType is called when entering the integerType production.
	EnterIntegerType(c *IntegerTypeContext)

	// EnterFloatingType is called when entering the floatingType production.
	EnterFloatingType(c *FloatingTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterStorageClass is called when entering the storageClass production.
	EnterStorageClass(c *StorageClassContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNondigit is called when entering the nondigit production.
	EnterNondigit(c *NondigitContext)

	// EnterDecimalDigit is called when entering the decimalDigit production.
	EnterDecimalDigit(c *DecimalDigitContext)

	// EnterAlnum is called when entering the alnum production.
	EnterAlnum(c *AlnumContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFloatingLiteral is called when entering the floatingLiteral production.
	EnterFloatingLiteral(c *FloatingLiteralContext)

	// EnterFloatingSuffix is called when entering the floatingSuffix production.
	EnterFloatingSuffix(c *FloatingSuffixContext)

	// EnterDecimalDigitsWithoutSeparators is called when entering the decimalDigitsWithoutSeparators production.
	EnterDecimalDigitsWithoutSeparators(c *DecimalDigitsWithoutSeparatorsContext)

	// EnterDecimalDigits is called when entering the decimalDigits production.
	EnterDecimalDigits(c *DecimalDigitsContext)

	// EnterNonzeroDecimalDigits is called when entering the nonzeroDecimalDigits production.
	EnterNonzeroDecimalDigits(c *NonzeroDecimalDigitsContext)

	// EnterNonzeroDecimalDigit is called when entering the nonzeroDecimalDigit production.
	EnterNonzeroDecimalDigit(c *NonzeroDecimalDigitContext)

	// EnterHexDigits is called when entering the hexDigits production.
	EnterHexDigits(c *HexDigitsContext)

	// EnterHexDigit is called when entering the hexDigit production.
	EnterHexDigit(c *HexDigitContext)

	// EnterDecimalExponent is called when entering the decimalExponent production.
	EnterDecimalExponent(c *DecimalExponentContext)

	// EnterBinaryExponent is called when entering the binaryExponent production.
	EnterBinaryExponent(c *BinaryExponentContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterDecimalExponentChar is called when entering the decimalExponentChar production.
	EnterDecimalExponentChar(c *DecimalExponentCharContext)

	// EnterBinaryExponentChar is called when entering the binaryExponentChar production.
	EnterBinaryExponentChar(c *BinaryExponentCharContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterIntegerSuffix is called when entering the integerSuffix production.
	EnterIntegerSuffix(c *IntegerSuffixContext)

	// EnterBinaryDigit is called when entering the binaryDigit production.
	EnterBinaryDigit(c *BinaryDigitContext)

	// EnterOctalDigit is called when entering the octalDigit production.
	EnterOctalDigit(c *OctalDigitContext)

	// EnterBinaryDigits is called when entering the binaryDigits production.
	EnterBinaryDigits(c *BinaryDigitsContext)

	// EnterOctalDigits is called when entering the octalDigits production.
	EnterOctalDigits(c *OctalDigitsContext)

	// EnterPositiveDecimalExponent is called when entering the positiveDecimalExponent production.
	EnterPositiveDecimalExponent(c *PositiveDecimalExponentContext)

	// EnterRuneLiteral is called when entering the runeLiteral production.
	EnterRuneLiteral(c *RuneLiteralContext)

	// EnterRune is called when entering the rune production.
	EnterRune(c *RuneContext)

	// EnterEscapeSequence is called when entering the escapeSequence production.
	EnterEscapeSequence(c *EscapeSequenceContext)

	// EnterFourbyte is called when entering the fourbyte production.
	EnterFourbyte(c *FourbyteContext)

	// EnterEightbyte is called when entering the eightbyte production.
	EnterEightbyte(c *EightbyteContext)

	// EnterNamedEscape is called when entering the namedEscape production.
	EnterNamedEscape(c *NamedEscapeContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterStringSection is called when entering the stringSection production.
	EnterStringSection(c *StringSectionContext)

	// EnterStringChars is called when entering the stringChars production.
	EnterStringChars(c *StringCharsContext)

	// EnterStringChar is called when entering the stringChar production.
	EnterStringChar(c *StringCharContext)

	// EnterRawstringChars is called when entering the rawstringChars production.
	EnterRawstringChars(c *RawstringCharsContext)

	// EnterRawstringChar is called when entering the rawstringChar production.
	EnterRawstringChar(c *RawstringCharContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterArrayMembers is called when entering the arrayMembers production.
	EnterArrayMembers(c *ArrayMembersContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterStructInitializer is called when entering the structInitializer production.
	EnterStructInitializer(c *StructInitializerContext)

	// EnterFieldValues is called when entering the fieldValues production.
	EnterFieldValues(c *FieldValuesContext)

	// EnterFieldValue is called when entering the fieldValue production.
	EnterFieldValue(c *FieldValueContext)

	// EnterTupleLiteral is called when entering the tupleLiteral production.
	EnterTupleLiteral(c *TupleLiteralContext)

	// EnterTupleItems is called when entering the tupleItems production.
	EnterTupleItems(c *TupleItemsContext)

	// EnterPlainExpression is called when entering the plainExpression production.
	EnterPlainExpression(c *PlainExpressionContext)

	// EnterNestedExpression is called when entering the nestedExpression production.
	EnterNestedExpression(c *NestedExpressionContext)

	// EnterAllocationExpression is called when entering the allocationExpression production.
	EnterAllocationExpression(c *AllocationExpressionContext)

	// EnterFreeExpression is called when entering the freeExpression production.
	EnterFreeExpression(c *FreeExpressionContext)

	// EnterAssertionExpression is called when entering the assertionExpression production.
	EnterAssertionExpression(c *AssertionExpressionContext)

	// EnterStaticAssertionExpression is called when entering the staticAssertionExpression production.
	EnterStaticAssertionExpression(c *StaticAssertionExpressionContext)

	// EnterCallExpression is called when entering the callExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterMeasurementExpression is called when entering the measurementExpression production.
	EnterMeasurementExpression(c *MeasurementExpressionContext)

	// EnterAlignExpression is called when entering the alignExpression production.
	EnterAlignExpression(c *AlignExpressionContext)

	// EnterSizeExpression is called when entering the sizeExpression production.
	EnterSizeExpression(c *SizeExpressionContext)

	// EnterLengthExpression is called when entering the lengthExpression production.
	EnterLengthExpression(c *LengthExpressionContext)

	// EnterOffsetExpression is called when entering the offsetExpression production.
	EnterOffsetExpression(c *OffsetExpressionContext)

	// EnterOffsetOperand is called when entering the offsetOperand production.
	EnterOffsetOperand(c *OffsetOperandContext)

	// EnterFieldAccessExpression is called when entering the fieldAccessExpression production.
	EnterFieldAccessExpression(c *FieldAccessExpressionContext)

	// EnterIndexingExpression is called when entering the indexingExpression production.
	EnterIndexingExpression(c *IndexingExpressionContext)

	// EnterSlicingExpression is called when entering the slicingExpression production.
	EnterSlicingExpression(c *SlicingExpressionContext)

	// EnterSliceMutationExpression is called when entering the sliceMutationExpression production.
	EnterSliceMutationExpression(c *SliceMutationExpressionContext)

	// EnterAppendExpression is called when entering the appendExpression production.
	EnterAppendExpression(c *AppendExpressionContext)

	// EnterInsertExpression is called when entering the insertExpression production.
	EnterInsertExpression(c *InsertExpressionContext)

	// EnterInsertOperand is called when entering the insertOperand production.
	EnterInsertOperand(c *InsertOperandContext)

	// EnterDeleteExpression is called when entering the deleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterDeleteOperand is called when entering the deleteOperand production.
	EnterDeleteOperand(c *DeleteOperandContext)

	// EnterErrorCheckingExpression is called when entering the errorCheckingExpression production.
	EnterErrorCheckingExpression(c *ErrorCheckingExpressionContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterObjectSelector is called when entering the objectSelector production.
	EnterObjectSelector(c *ObjectSelectorContext)

	// EnterVariadicExpression is called when entering the variadicExpression production.
	EnterVariadicExpression(c *VariadicExpressionContext)

	// EnterBuiltinExpression is called when entering the builtinExpression production.
	EnterBuiltinExpression(c *BuiltinExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterNullableType is called when entering the nullableType production.
	EnterNullableType(c *NullableTypeContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalXorExpression is called when entering the logicalXorExpression production.
	EnterLogicalXorExpression(c *LogicalXorExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterIfExpression is called when entering the ifExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterConditionalBranch is called when entering the conditionalBranch production.
	EnterConditionalBranch(c *ConditionalBranchContext)

	// EnterForLoop is called when entering the forLoop production.
	EnterForLoop(c *ForLoopContext)

	// EnterForPredicate is called when entering the forPredicate production.
	EnterForPredicate(c *ForPredicateContext)

	// EnterIterableBinding is called when entering the iterableBinding production.
	EnterIterableBinding(c *IterableBindingContext)

	// EnterIterableBindingLeft is called when entering the iterableBindingLeft production.
	EnterIterableBindingLeft(c *IterableBindingLeftContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterSwitchExpression is called when entering the switchExpression production.
	EnterSwitchExpression(c *SwitchExpressionContext)

	// EnterSwitchCases is called when entering the switchCases production.
	EnterSwitchCases(c *SwitchCasesContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterCaseOptions is called when entering the caseOptions production.
	EnterCaseOptions(c *CaseOptionsContext)

	// EnterMatchExpression is called when entering the matchExpression production.
	EnterMatchExpression(c *MatchExpressionContext)

	// EnterMatchCases is called when entering the matchCases production.
	EnterMatchCases(c *MatchCasesContext)

	// EnterMatchCase is called when entering the matchCase production.
	EnterMatchCase(c *MatchCaseContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssignmentTarget is called when entering the assignmentTarget production.
	EnterAssignmentTarget(c *AssignmentTargetContext)

	// EnterIndirectAssignmentTarget is called when entering the indirectAssignmentTarget production.
	EnterIndirectAssignmentTarget(c *IndirectAssignmentTargetContext)

	// EnterSlicingAssignmentTarget is called when entering the slicingAssignmentTarget production.
	EnterSlicingAssignmentTarget(c *SlicingAssignmentTargetContext)

	// EnterAssignmentOp is called when entering the assignmentOp production.
	EnterAssignmentOp(c *AssignmentOpContext)

	// EnterBindingList is called when entering the bindingList production.
	EnterBindingList(c *BindingListContext)

	// EnterBindings is called when entering the bindings production.
	EnterBindings(c *BindingsContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterBindingName is called when entering the bindingName production.
	EnterBindingName(c *BindingNameContext)

	// EnterTupleBindingNames is called when entering the tupleBindingNames production.
	EnterTupleBindingNames(c *TupleBindingNamesContext)

	// EnterTupleBindingName is called when entering the tupleBindingName production.
	EnterTupleBindingName(c *TupleBindingNameContext)

	// EnterDeferExpression is called when entering the deferExpression production.
	EnterDeferExpression(c *DeferExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterCompoundExpression is called when entering the compoundExpression production.
	EnterCompoundExpression(c *CompoundExpressionContext)

	// EnterControlExpression is called when entering the controlExpression production.
	EnterControlExpression(c *ControlExpressionContext)

	// EnterYieldExpression is called when entering the yieldExpression production.
	EnterYieldExpression(c *YieldExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterGlobalDeclaration is called when entering the globalDeclaration production.
	EnterGlobalDeclaration(c *GlobalDeclarationContext)

	// EnterGlobalBindings is called when entering the globalBindings production.
	EnterGlobalBindings(c *GlobalBindingsContext)

	// EnterGlobalBinding is called when entering the globalBinding production.
	EnterGlobalBinding(c *GlobalBindingContext)

	// EnterDeclAttr is called when entering the declAttr production.
	EnterDeclAttr(c *DeclAttrContext)

	// EnterConstantDeclaration is called when entering the constantDeclaration production.
	EnterConstantDeclaration(c *ConstantDeclarationContext)

	// EnterConstantBindings is called when entering the constantBindings production.
	EnterConstantBindings(c *ConstantBindingsContext)

	// EnterConstantBinding is called when entering the constantBinding production.
	EnterConstantBinding(c *ConstantBindingContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterTypeBindings is called when entering the typeBindings production.
	EnterTypeBindings(c *TypeBindingsContext)

	// EnterTypeBinding is called when entering the typeBinding production.
	EnterTypeBinding(c *TypeBindingContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterEnumValues is called when entering the enumValues production.
	EnterEnumValues(c *EnumValuesContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterEnumStorage is called when entering the enumStorage production.
	EnterEnumStorage(c *EnumStorageContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFndecAttr is called when entering the fndecAttr production.
	EnterFndecAttr(c *FndecAttrContext)

	// EnterSubUnit is called when entering the subUnit production.
	EnterSubUnit(c *SubUnitContext)

	// EnterImports is called when entering the imports production.
	EnterImports(c *ImportsContext)

	// EnterUseDirective is called when entering the useDirective production.
	EnterUseDirective(c *UseDirectiveContext)

	// EnterMemberList is called when entering the memberList production.
	EnterMemberList(c *MemberListContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// ExitLnot is called when exiting the lnot production.
	ExitLnot(c *LnotContext)

	// ExitNequal is called when exiting the nequal production.
	ExitNequal(c *NequalContext)

	// ExitModulo is called when exiting the modulo production.
	ExitModulo(c *ModuloContext)

	// ExitBand is called when exiting the band production.
	ExitBand(c *BandContext)

	// ExitLand is called when exiting the land production.
	ExitLand(c *LandContext)

	// ExitLandeq is called when exiting the landeq production.
	ExitLandeq(c *LandeqContext)

	// ExitBandeq is called when exiting the bandeq production.
	ExitBandeq(c *BandeqContext)

	// ExitLparen is called when exiting the lparen production.
	ExitLparen(c *LparenContext)

	// ExitRparen is called when exiting the rparen production.
	ExitRparen(c *RparenContext)

	// ExitTimes is called when exiting the times production.
	ExitTimes(c *TimesContext)

	// ExitTimeseq is called when exiting the timeseq production.
	ExitTimeseq(c *TimeseqContext)

	// ExitPlus is called when exiting the plus production.
	ExitPlus(c *PlusContext)

	// ExitPluseq is called when exiting the pluseq production.
	ExitPluseq(c *PluseqContext)

	// ExitComma is called when exiting the comma production.
	ExitComma(c *CommaContext)

	// ExitMinus is called when exiting the minus production.
	ExitMinus(c *MinusContext)

	// ExitMinuseq is called when exiting the minuseq production.
	ExitMinuseq(c *MinuseqContext)

	// ExitDot is called when exiting the dot production.
	ExitDot(c *DotContext)

	// ExitDoubleDot is called when exiting the doubleDot production.
	ExitDoubleDot(c *DoubleDotContext)

	// ExitEllipsis is called when exiting the ellipsis production.
	ExitEllipsis(c *EllipsisContext)

	// ExitDiv is called when exiting the div production.
	ExitDiv(c *DivContext)

	// ExitDiveq is called when exiting the diveq production.
	ExitDiveq(c *DiveqContext)

	// ExitColon is called when exiting the colon production.
	ExitColon(c *ColonContext)

	// ExitDoubleColon is called when exiting the doubleColon production.
	ExitDoubleColon(c *DoubleColonContext)

	// ExitSemicolon is called when exiting the semicolon production.
	ExitSemicolon(c *SemicolonContext)

	// ExitLess is called when exiting the less production.
	ExitLess(c *LessContext)

	// ExitLshift is called when exiting the lshift production.
	ExitLshift(c *LshiftContext)

	// ExitLshifteq is called when exiting the lshifteq production.
	ExitLshifteq(c *LshifteqContext)

	// ExitLesseq is called when exiting the lesseq production.
	ExitLesseq(c *LesseqContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitLequal is called when exiting the lequal production.
	ExitLequal(c *LequalContext)

	// ExitArrow is called when exiting the arrow production.
	ExitArrow(c *ArrowContext)

	// ExitGt is called when exiting the gt production.
	ExitGt(c *GtContext)

	// ExitGteq is called when exiting the gteq production.
	ExitGteq(c *GteqContext)

	// ExitRshift is called when exiting the rshift production.
	ExitRshift(c *RshiftContext)

	// ExitRshifteq is called when exiting the rshifteq production.
	ExitRshifteq(c *RshifteqContext)

	// ExitQuestion is called when exiting the question production.
	ExitQuestion(c *QuestionContext)

	// ExitLbracket is called when exiting the lbracket production.
	ExitLbracket(c *LbracketContext)

	// ExitRbracket is called when exiting the rbracket production.
	ExitRbracket(c *RbracketContext)

	// ExitBxor is called when exiting the bxor production.
	ExitBxor(c *BxorContext)

	// ExitBxoreq is called when exiting the bxoreq production.
	ExitBxoreq(c *BxoreqContext)

	// ExitLxor is called when exiting the lxor production.
	ExitLxor(c *LxorContext)

	// ExitLxoreq is called when exiting the lxoreq production.
	ExitLxoreq(c *LxoreqContext)

	// ExitLbrace is called when exiting the lbrace production.
	ExitLbrace(c *LbraceContext)

	// ExitRbrace is called when exiting the rbrace production.
	ExitRbrace(c *RbraceContext)

	// ExitBor is called when exiting the bor production.
	ExitBor(c *BorContext)

	// ExitBoreq is called when exiting the boreq production.
	ExitBoreq(c *BoreqContext)

	// ExitLor is called when exiting the lor production.
	ExitLor(c *LorContext)

	// ExitLoreq is called when exiting the loreq production.
	ExitLoreq(c *LoreqContext)

	// ExitBnot is called when exiting the bnot production.
	ExitBnot(c *BnotContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)

	// ExitAttributes is called when exiting the attributes production.
	ExitAttributes(c *AttributesContext)

	// ExitInvalidAttribute is called when exiting the invalidAttribute production.
	ExitInvalidAttribute(c *InvalidAttributeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitStructUnionType is called when exiting the structUnionType production.
	ExitStructUnionType(c *StructUnionTypeContext)

	// ExitStructUnionFields is called when exiting the structUnionFields production.
	ExitStructUnionFields(c *StructUnionFieldsContext)

	// ExitStructUnionField is called when exiting the structUnionField production.
	ExitStructUnionField(c *StructUnionFieldContext)

	// ExitTupleType is called when exiting the tupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitTupleTypes is called when exiting the tupleTypes production.
	ExitTupleTypes(c *TupleTypesContext)

	// ExitTaggedUnionType is called when exiting the taggedUnionType production.
	ExitTaggedUnionType(c *TaggedUnionTypeContext)

	// ExitTaggedTypes is called when exiting the taggedTypes production.
	ExitTaggedTypes(c *TaggedTypesContext)

	// ExitSliceArrayType is called when exiting the sliceArrayType production.
	ExitSliceArrayType(c *SliceArrayTypeContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitPrototype is called when exiting the prototype production.
	ExitPrototype(c *PrototypeContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitAliasType is called when exiting the aliasType production.
	ExitAliasType(c *AliasTypeContext)

	// ExitUnwrappedAlias is called when exiting the unwrappedAlias production.
	ExitUnwrappedAlias(c *UnwrappedAliasContext)

	// ExitIntegerType is called when exiting the integerType production.
	ExitIntegerType(c *IntegerTypeContext)

	// ExitFloatingType is called when exiting the floatingType production.
	ExitFloatingType(c *FloatingTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitStorageClass is called when exiting the storageClass production.
	ExitStorageClass(c *StorageClassContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNondigit is called when exiting the nondigit production.
	ExitNondigit(c *NondigitContext)

	// ExitDecimalDigit is called when exiting the decimalDigit production.
	ExitDecimalDigit(c *DecimalDigitContext)

	// ExitAlnum is called when exiting the alnum production.
	ExitAlnum(c *AlnumContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFloatingLiteral is called when exiting the floatingLiteral production.
	ExitFloatingLiteral(c *FloatingLiteralContext)

	// ExitFloatingSuffix is called when exiting the floatingSuffix production.
	ExitFloatingSuffix(c *FloatingSuffixContext)

	// ExitDecimalDigitsWithoutSeparators is called when exiting the decimalDigitsWithoutSeparators production.
	ExitDecimalDigitsWithoutSeparators(c *DecimalDigitsWithoutSeparatorsContext)

	// ExitDecimalDigits is called when exiting the decimalDigits production.
	ExitDecimalDigits(c *DecimalDigitsContext)

	// ExitNonzeroDecimalDigits is called when exiting the nonzeroDecimalDigits production.
	ExitNonzeroDecimalDigits(c *NonzeroDecimalDigitsContext)

	// ExitNonzeroDecimalDigit is called when exiting the nonzeroDecimalDigit production.
	ExitNonzeroDecimalDigit(c *NonzeroDecimalDigitContext)

	// ExitHexDigits is called when exiting the hexDigits production.
	ExitHexDigits(c *HexDigitsContext)

	// ExitHexDigit is called when exiting the hexDigit production.
	ExitHexDigit(c *HexDigitContext)

	// ExitDecimalExponent is called when exiting the decimalExponent production.
	ExitDecimalExponent(c *DecimalExponentContext)

	// ExitBinaryExponent is called when exiting the binaryExponent production.
	ExitBinaryExponent(c *BinaryExponentContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitDecimalExponentChar is called when exiting the decimalExponentChar production.
	ExitDecimalExponentChar(c *DecimalExponentCharContext)

	// ExitBinaryExponentChar is called when exiting the binaryExponentChar production.
	ExitBinaryExponentChar(c *BinaryExponentCharContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitIntegerSuffix is called when exiting the integerSuffix production.
	ExitIntegerSuffix(c *IntegerSuffixContext)

	// ExitBinaryDigit is called when exiting the binaryDigit production.
	ExitBinaryDigit(c *BinaryDigitContext)

	// ExitOctalDigit is called when exiting the octalDigit production.
	ExitOctalDigit(c *OctalDigitContext)

	// ExitBinaryDigits is called when exiting the binaryDigits production.
	ExitBinaryDigits(c *BinaryDigitsContext)

	// ExitOctalDigits is called when exiting the octalDigits production.
	ExitOctalDigits(c *OctalDigitsContext)

	// ExitPositiveDecimalExponent is called when exiting the positiveDecimalExponent production.
	ExitPositiveDecimalExponent(c *PositiveDecimalExponentContext)

	// ExitRuneLiteral is called when exiting the runeLiteral production.
	ExitRuneLiteral(c *RuneLiteralContext)

	// ExitRune is called when exiting the rune production.
	ExitRune(c *RuneContext)

	// ExitEscapeSequence is called when exiting the escapeSequence production.
	ExitEscapeSequence(c *EscapeSequenceContext)

	// ExitFourbyte is called when exiting the fourbyte production.
	ExitFourbyte(c *FourbyteContext)

	// ExitEightbyte is called when exiting the eightbyte production.
	ExitEightbyte(c *EightbyteContext)

	// ExitNamedEscape is called when exiting the namedEscape production.
	ExitNamedEscape(c *NamedEscapeContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitStringSection is called when exiting the stringSection production.
	ExitStringSection(c *StringSectionContext)

	// ExitStringChars is called when exiting the stringChars production.
	ExitStringChars(c *StringCharsContext)

	// ExitStringChar is called when exiting the stringChar production.
	ExitStringChar(c *StringCharContext)

	// ExitRawstringChars is called when exiting the rawstringChars production.
	ExitRawstringChars(c *RawstringCharsContext)

	// ExitRawstringChar is called when exiting the rawstringChar production.
	ExitRawstringChar(c *RawstringCharContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitArrayMembers is called when exiting the arrayMembers production.
	ExitArrayMembers(c *ArrayMembersContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitStructInitializer is called when exiting the structInitializer production.
	ExitStructInitializer(c *StructInitializerContext)

	// ExitFieldValues is called when exiting the fieldValues production.
	ExitFieldValues(c *FieldValuesContext)

	// ExitFieldValue is called when exiting the fieldValue production.
	ExitFieldValue(c *FieldValueContext)

	// ExitTupleLiteral is called when exiting the tupleLiteral production.
	ExitTupleLiteral(c *TupleLiteralContext)

	// ExitTupleItems is called when exiting the tupleItems production.
	ExitTupleItems(c *TupleItemsContext)

	// ExitPlainExpression is called when exiting the plainExpression production.
	ExitPlainExpression(c *PlainExpressionContext)

	// ExitNestedExpression is called when exiting the nestedExpression production.
	ExitNestedExpression(c *NestedExpressionContext)

	// ExitAllocationExpression is called when exiting the allocationExpression production.
	ExitAllocationExpression(c *AllocationExpressionContext)

	// ExitFreeExpression is called when exiting the freeExpression production.
	ExitFreeExpression(c *FreeExpressionContext)

	// ExitAssertionExpression is called when exiting the assertionExpression production.
	ExitAssertionExpression(c *AssertionExpressionContext)

	// ExitStaticAssertionExpression is called when exiting the staticAssertionExpression production.
	ExitStaticAssertionExpression(c *StaticAssertionExpressionContext)

	// ExitCallExpression is called when exiting the callExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitMeasurementExpression is called when exiting the measurementExpression production.
	ExitMeasurementExpression(c *MeasurementExpressionContext)

	// ExitAlignExpression is called when exiting the alignExpression production.
	ExitAlignExpression(c *AlignExpressionContext)

	// ExitSizeExpression is called when exiting the sizeExpression production.
	ExitSizeExpression(c *SizeExpressionContext)

	// ExitLengthExpression is called when exiting the lengthExpression production.
	ExitLengthExpression(c *LengthExpressionContext)

	// ExitOffsetExpression is called when exiting the offsetExpression production.
	ExitOffsetExpression(c *OffsetExpressionContext)

	// ExitOffsetOperand is called when exiting the offsetOperand production.
	ExitOffsetOperand(c *OffsetOperandContext)

	// ExitFieldAccessExpression is called when exiting the fieldAccessExpression production.
	ExitFieldAccessExpression(c *FieldAccessExpressionContext)

	// ExitIndexingExpression is called when exiting the indexingExpression production.
	ExitIndexingExpression(c *IndexingExpressionContext)

	// ExitSlicingExpression is called when exiting the slicingExpression production.
	ExitSlicingExpression(c *SlicingExpressionContext)

	// ExitSliceMutationExpression is called when exiting the sliceMutationExpression production.
	ExitSliceMutationExpression(c *SliceMutationExpressionContext)

	// ExitAppendExpression is called when exiting the appendExpression production.
	ExitAppendExpression(c *AppendExpressionContext)

	// ExitInsertExpression is called when exiting the insertExpression production.
	ExitInsertExpression(c *InsertExpressionContext)

	// ExitInsertOperand is called when exiting the insertOperand production.
	ExitInsertOperand(c *InsertOperandContext)

	// ExitDeleteExpression is called when exiting the deleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitDeleteOperand is called when exiting the deleteOperand production.
	ExitDeleteOperand(c *DeleteOperandContext)

	// ExitErrorCheckingExpression is called when exiting the errorCheckingExpression production.
	ExitErrorCheckingExpression(c *ErrorCheckingExpressionContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitObjectSelector is called when exiting the objectSelector production.
	ExitObjectSelector(c *ObjectSelectorContext)

	// ExitVariadicExpression is called when exiting the variadicExpression production.
	ExitVariadicExpression(c *VariadicExpressionContext)

	// ExitBuiltinExpression is called when exiting the builtinExpression production.
	ExitBuiltinExpression(c *BuiltinExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitNullableType is called when exiting the nullableType production.
	ExitNullableType(c *NullableTypeContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalXorExpression is called when exiting the logicalXorExpression production.
	ExitLogicalXorExpression(c *LogicalXorExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitIfExpression is called when exiting the ifExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitConditionalBranch is called when exiting the conditionalBranch production.
	ExitConditionalBranch(c *ConditionalBranchContext)

	// ExitForLoop is called when exiting the forLoop production.
	ExitForLoop(c *ForLoopContext)

	// ExitForPredicate is called when exiting the forPredicate production.
	ExitForPredicate(c *ForPredicateContext)

	// ExitIterableBinding is called when exiting the iterableBinding production.
	ExitIterableBinding(c *IterableBindingContext)

	// ExitIterableBindingLeft is called when exiting the iterableBindingLeft production.
	ExitIterableBindingLeft(c *IterableBindingLeftContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitSwitchExpression is called when exiting the switchExpression production.
	ExitSwitchExpression(c *SwitchExpressionContext)

	// ExitSwitchCases is called when exiting the switchCases production.
	ExitSwitchCases(c *SwitchCasesContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitCaseOptions is called when exiting the caseOptions production.
	ExitCaseOptions(c *CaseOptionsContext)

	// ExitMatchExpression is called when exiting the matchExpression production.
	ExitMatchExpression(c *MatchExpressionContext)

	// ExitMatchCases is called when exiting the matchCases production.
	ExitMatchCases(c *MatchCasesContext)

	// ExitMatchCase is called when exiting the matchCase production.
	ExitMatchCase(c *MatchCaseContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssignmentTarget is called when exiting the assignmentTarget production.
	ExitAssignmentTarget(c *AssignmentTargetContext)

	// ExitIndirectAssignmentTarget is called when exiting the indirectAssignmentTarget production.
	ExitIndirectAssignmentTarget(c *IndirectAssignmentTargetContext)

	// ExitSlicingAssignmentTarget is called when exiting the slicingAssignmentTarget production.
	ExitSlicingAssignmentTarget(c *SlicingAssignmentTargetContext)

	// ExitAssignmentOp is called when exiting the assignmentOp production.
	ExitAssignmentOp(c *AssignmentOpContext)

	// ExitBindingList is called when exiting the bindingList production.
	ExitBindingList(c *BindingListContext)

	// ExitBindings is called when exiting the bindings production.
	ExitBindings(c *BindingsContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitBindingName is called when exiting the bindingName production.
	ExitBindingName(c *BindingNameContext)

	// ExitTupleBindingNames is called when exiting the tupleBindingNames production.
	ExitTupleBindingNames(c *TupleBindingNamesContext)

	// ExitTupleBindingName is called when exiting the tupleBindingName production.
	ExitTupleBindingName(c *TupleBindingNameContext)

	// ExitDeferExpression is called when exiting the deferExpression production.
	ExitDeferExpression(c *DeferExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitCompoundExpression is called when exiting the compoundExpression production.
	ExitCompoundExpression(c *CompoundExpressionContext)

	// ExitControlExpression is called when exiting the controlExpression production.
	ExitControlExpression(c *ControlExpressionContext)

	// ExitYieldExpression is called when exiting the yieldExpression production.
	ExitYieldExpression(c *YieldExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitGlobalDeclaration is called when exiting the globalDeclaration production.
	ExitGlobalDeclaration(c *GlobalDeclarationContext)

	// ExitGlobalBindings is called when exiting the globalBindings production.
	ExitGlobalBindings(c *GlobalBindingsContext)

	// ExitGlobalBinding is called when exiting the globalBinding production.
	ExitGlobalBinding(c *GlobalBindingContext)

	// ExitDeclAttr is called when exiting the declAttr production.
	ExitDeclAttr(c *DeclAttrContext)

	// ExitConstantDeclaration is called when exiting the constantDeclaration production.
	ExitConstantDeclaration(c *ConstantDeclarationContext)

	// ExitConstantBindings is called when exiting the constantBindings production.
	ExitConstantBindings(c *ConstantBindingsContext)

	// ExitConstantBinding is called when exiting the constantBinding production.
	ExitConstantBinding(c *ConstantBindingContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitTypeBindings is called when exiting the typeBindings production.
	ExitTypeBindings(c *TypeBindingsContext)

	// ExitTypeBinding is called when exiting the typeBinding production.
	ExitTypeBinding(c *TypeBindingContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitEnumValues is called when exiting the enumValues production.
	ExitEnumValues(c *EnumValuesContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitEnumStorage is called when exiting the enumStorage production.
	ExitEnumStorage(c *EnumStorageContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFndecAttr is called when exiting the fndecAttr production.
	ExitFndecAttr(c *FndecAttrContext)

	// ExitSubUnit is called when exiting the subUnit production.
	ExitSubUnit(c *SubUnitContext)

	// ExitImports is called when exiting the imports production.
	ExitImports(c *ImportsContext)

	// ExitUseDirective is called when exiting the useDirective production.
	ExitUseDirective(c *UseDirectiveContext)

	// ExitMemberList is called when exiting the memberList production.
	ExitMemberList(c *MemberListContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)
}
