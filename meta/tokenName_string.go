// Code generated by "stringer -type=TokenName -output ./tokenName_string.go"; DO NOT EDIT.

package meta

import "strconv"

const _TokenName_name = "NodeStartNodeEndMagicConstantTokenIncludeTokenIncludeOnceTokenExitTokenIfTokenLnumberTokenDnumberTokenStringTokenStringVarnameTokenNumStringTokenInlineHTMLTokenEncapsedAndWhitespaceTokenConstantEncapsedStringTokenEchoTokenDoTokenWhileTokenEndwhileTokenForInitSemicolonTokenForCondSemicolonTokenForTokenEndforTokenForeachTokenEndforeachTokenDeclareTokenEnddeclareTokenAsTokenSwitchTokenEndswitchTokenCaseTokenDefaultTokenBreakTokenContinueTokenGotoTokenFunctionTokenConstTokenReturnTokenTryTokenCatchTokenFinallyTokenThrowTokenUseTokenInsteadofTokenGlobalTokenVarTokenUnsetTokenIssetTokenEmptyTokenClassTokenTraitTokenInterfaceTokenExtendsTokenImplementsTokenDoubleArrowTokenListTokenArrayTokenCallableTokenStartHeredocTokenCurlyOpenTokenPaamayimNekudotayimTokenNamespaceTokenUseLeadingNsSeparatorTokenNsSeparatorTokenEllipsisTokenEvalTokenRequireTokenRequireOnceTokenLogicalOrTokenLogicalXorTokenLogicalAndTokenInstanceofTokenNewAnchorCloneTokenElseifTokenElseTokenEndifTokenPrintTokenYieldTokenStaticTokenAbstractTokenFinalTokenPrivateTokenProtectedTokenPublicTokenIncTokenDecTokenYieldFromTokenObjectOperatorTokenIntCastTokenDoubleCastTokenStringCastTokenArrayCastTokenObjectCastTokenBoolCastTokenUnsetCastTokenCoalesceTokenSpaceshipTokenPlusEqualTokenMinusEqualTokenMulEqualTokenPowEqualTokenDivEqualTokenConcatEqualTokenModEqualTokenAndEqualTokenOrEqualTokenXorEqualTokenSlEqualTokenSrEqualTokenBooleanOrTokenBooleanAndTokenPowTokenSlTokenSrTokenIsIdenticalTokenIsNotIdenticalTokenIsEqualTokenIsNotEqualTokenIsSmallerOrEqualTokenIsGreaterOrEqualTokenHaltCompilerTokenCaseSeparatorTokenDoubleQuoteTokenBackquoteTokenOpenCurlyBracesTokenCloseCurlyBracesTokenSemiColonTokenColonTokenOpenParenthesisTokenCloseParenthesisTokenOpenSquareBracketCloseSquareBracketQuestionMarkTokenAmpersandTokenMinusTokenPlusTokenExclamationMarkTokenTildeTokenAtTokenCommaTokenVerticalBarTokenEqualTokenCaretTokenAsteriskTokenSlashTokenPercentTokenLessTokenGreaterTokenDotToken"

var _TokenName_index = [...]uint16{0, 9, 16, 34, 46, 62, 71, 78, 90, 102, 113, 131, 145, 160, 186, 213, 222, 229, 239, 252, 273, 294, 302, 313, 325, 340, 352, 367, 374, 385, 399, 408, 420, 430, 443, 452, 465, 475, 486, 494, 504, 516, 526, 534, 548, 559, 567, 577, 587, 597, 607, 617, 631, 643, 658, 674, 683, 693, 706, 723, 737, 761, 775, 801, 817, 830, 839, 851, 867, 881, 896, 911, 926, 935, 945, 956, 965, 975, 985, 995, 1006, 1019, 1029, 1041, 1055, 1066, 1074, 1082, 1096, 1115, 1127, 1142, 1157, 1171, 1186, 1199, 1213, 1226, 1240, 1254, 1269, 1282, 1295, 1308, 1324, 1337, 1350, 1362, 1375, 1387, 1399, 1413, 1428, 1436, 1443, 1450, 1466, 1485, 1497, 1512, 1533, 1554, 1571, 1589, 1605, 1619, 1639, 1660, 1674, 1684, 1704, 1725, 1742, 1760, 1777, 1791, 1801, 1810, 1830, 1840, 1847, 1857, 1873, 1883, 1893, 1906, 1916, 1928, 1937, 1949, 1957}

func (i TokenName) String() string {
	if i < 0 || i >= TokenName(len(_TokenName_index)-1) {
		return "TokenName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenName_name[_TokenName_index[i]:_TokenName_index[i+1]]
}