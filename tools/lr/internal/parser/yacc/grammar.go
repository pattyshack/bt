// Code generated by goyacc -o grammar.go -v grammar.output -p Lr grammar.y. DO NOT EDIT.

//line grammar.y:2
package yacc

import __yyfmt__ "fmt"

//line grammar.y:2

import (
	"github.com/pattyshack/gt/tools/lr/internal/parser"
)

//line grammar.y:9
type LrSymType struct {
	yys      int
	Generic_ *parser.LRGenericSymbol

	*parser.Token
	Tokens []*parser.Token

	parser.Definition // interface
	Definitions       []parser.Definition

	*parser.Rule

	Clause  *parser.Clause
	Clauses []*parser.Clause

	*parser.AdditionalSection
	AdditionalSections []*parser.AdditionalSection

	*parser.Grammar
}

const TOKEN = 57346
const TYPE = 57347
const START = 57348
const ARROW = 57349
const RULE_DEF = 57350
const LABEL = 57351
const SECTION_MARKER = 57352
const IDENTIFIER = 57353
const CHARACTER = 57354
const SECTION_CONTENT = 57355

var LrToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TOKEN",
	"TYPE",
	"START",
	"ARROW",
	"RULE_DEF",
	"LABEL",
	"SECTION_MARKER",
	"IDENTIFIER",
	"CHARACTER",
	"SECTION_CONTENT",
	"';'",
	"'<'",
	"'>'",
	"'|'",
}

var LrStatenames = [...]string{}

const LrEofCode = 1
const LrErrCode = 2
const LrInitialStackSize = 16

//line grammar.y:206

func init() {
	LrErrorVerbose = true
}

//line yacctab:1
var LrExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 14,
	11, 21,
	12, 21,
	-2, 10,
	-1, 38,
	11, 21,
	12, 21,
	-2, 9,
}

const LrPrivate = 57344

const LrLast = 43

var LrAct = [...]int{
	22, 23, 15, 32, 35, 14, 16, 17, 27, 37,
	13, 24, 20, 16, 17, 12, 16, 17, 29, 30,
	34, 31, 28, 19, 7, 8, 5, 33, 9, 26,
	24, 1, 3, 10, 36, 11, 38, 25, 21, 6,
	2, 18, 4,
}

var LrPact = [...]int{
	20, -1000, 20, 1, -5, 12, -1000, -1000, -1000, 2,
	19, -6, -1000, 11, -1000, 7, -1000, -1000, 10, -1000,
	7, -14, -1000, -1000, 5, -1000, 9, -1000, -12, -1000,
	-1000, -1000, 21, 7, -4, 5, -1000, -1000, -1000,
}

var LrPgo = [...]int{
	0, 42, 41, 0, 2, 32, 40, 39, 1, 38,
	37, 33, 31,
}

var LrR1 = [...]int{
	0, 12, 11, 11, 10, 6, 6, 6, 6, 5,
	5, 5, 5, 1, 1, 2, 2, 3, 3, 3,
	3, 4, 4, 7, 7, 9, 9, 8,
}

var LrR2 = [...]int{
	0, 2, 2, 0, 3, 2, 3, 1, 2, 5,
	2, 2, 1, 1, 1, 2, 1, 2, 2, 1,
	1, 1, 0, 2, 2, 3, 1, 2,
}

var LrChk = [...]int{
	-1000, -12, -6, -5, -1, 6, -7, 4, 5, 8,
	-11, -5, 14, 15, -3, -4, 11, 12, -2, 11,
	-4, -9, -3, -8, 9, -10, 10, 14, 11, 11,
	12, 11, 17, -4, 11, 16, -8, 13, -3,
}

var LrDef = [...]int{
	0, -2, 3, 7, 0, 0, 12, 13, 14, 22,
	1, 5, 8, 0, -2, 0, 19, 20, 11, 16,
	23, 24, 21, 26, 22, 2, 0, 6, 0, 17,
	18, 15, 0, 27, 0, 0, 25, 4, -2,
}

var LrTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 14,
	15, 3, 16, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 17,
}

var LrTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
}

var LrTok3 = [...]int{
	0,
}

var LrErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	LrDebug        = 0
	LrErrorVerbose = false
)

type LrLexer interface {
	Lex(lval *LrSymType) int
	Error(s string)
}

type LrParser interface {
	Parse(LrLexer) int
	Lookahead() int
}

type LrParserImpl struct {
	lval  LrSymType
	stack [LrInitialStackSize]LrSymType
	char  int
}

func (p *LrParserImpl) Lookahead() int {
	return p.char
}

func LrNewParser() LrParser {
	return &LrParserImpl{}
}

const LrFlag = -1000

func LrTokname(c int) string {
	if c >= 1 && c-1 < len(LrToknames) {
		if LrToknames[c-1] != "" {
			return LrToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func LrStatname(s int) string {
	if s >= 0 && s < len(LrStatenames) {
		if LrStatenames[s] != "" {
			return LrStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func LrErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !LrErrorVerbose {
		return "syntax error"
	}

	for _, e := range LrErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + LrTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := LrPact[state]
	for tok := TOKSTART; tok-1 < len(LrToknames); tok++ {
		if n := base + tok; n >= 0 && n < LrLast && LrChk[LrAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if LrDef[state] == -2 {
		i := 0
		for LrExca[i] != -1 || LrExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; LrExca[i] >= 0; i += 2 {
			tok := LrExca[i]
			if tok < TOKSTART || LrExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if LrExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += LrTokname(tok)
	}
	return res
}

func Lrlex1(lex LrLexer, lval *LrSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = LrTok1[0]
		goto out
	}
	if char < len(LrTok1) {
		token = LrTok1[char]
		goto out
	}
	if char >= LrPrivate {
		if char < LrPrivate+len(LrTok2) {
			token = LrTok2[char-LrPrivate]
			goto out
		}
	}
	for i := 0; i < len(LrTok3); i += 2 {
		token = LrTok3[i+0]
		if token == char {
			token = LrTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = LrTok2[1] /* unknown char */
	}
	if LrDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", LrTokname(token), uint(char))
	}
	return char, token
}

func LrParse(Lrlex LrLexer) int {
	return LrNewParser().Parse(Lrlex)
}

func (Lrrcvr *LrParserImpl) Parse(Lrlex LrLexer) int {
	var Lrn int
	var LrVAL LrSymType
	var LrDollar []LrSymType
	_ = LrDollar // silence set and not used
	LrS := Lrrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Lrstate := 0
	Lrrcvr.char = -1
	Lrtoken := -1 // Lrrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Lrstate = -1
		Lrrcvr.char = -1
		Lrtoken = -1
	}()
	Lrp := -1
	goto Lrstack

ret0:
	return 0

ret1:
	return 1

Lrstack:
	/* put a state and value onto the stack */
	if LrDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", LrTokname(Lrtoken), LrStatname(Lrstate))
	}

	Lrp++
	if Lrp >= len(LrS) {
		nyys := make([]LrSymType, len(LrS)*2)
		copy(nyys, LrS)
		LrS = nyys
	}
	LrS[Lrp] = LrVAL
	LrS[Lrp].yys = Lrstate

Lrnewstate:
	Lrn = LrPact[Lrstate]
	if Lrn <= LrFlag {
		goto Lrdefault /* simple state */
	}
	if Lrrcvr.char < 0 {
		Lrrcvr.char, Lrtoken = Lrlex1(Lrlex, &Lrrcvr.lval)
	}
	Lrn += Lrtoken
	if Lrn < 0 || Lrn >= LrLast {
		goto Lrdefault
	}
	Lrn = LrAct[Lrn]
	if LrChk[Lrn] == Lrtoken { /* valid shift */
		Lrrcvr.char = -1
		Lrtoken = -1
		LrVAL = Lrrcvr.lval
		Lrstate = Lrn
		if Errflag > 0 {
			Errflag--
		}
		goto Lrstack
	}

Lrdefault:
	/* default state action */
	Lrn = LrDef[Lrstate]
	if Lrn == -2 {
		if Lrrcvr.char < 0 {
			Lrrcvr.char, Lrtoken = Lrlex1(Lrlex, &Lrrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if LrExca[xi+0] == -1 && LrExca[xi+1] == Lrstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Lrn = LrExca[xi+0]
			if Lrn < 0 || Lrn == Lrtoken {
				break
			}
		}
		Lrn = LrExca[xi+1]
		if Lrn < 0 {
			goto ret0
		}
	}
	if Lrn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Lrlex.Error(LrErrorMessage(Lrstate, Lrtoken))
			Nerrs++
			if LrDebug >= 1 {
				__yyfmt__.Printf("%s", LrStatname(Lrstate))
				__yyfmt__.Printf(" saw %s\n", LrTokname(Lrtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Lrp >= 0 {
				Lrn = LrPact[LrS[Lrp].yys] + LrErrCode
				if Lrn >= 0 && Lrn < LrLast {
					Lrstate = LrAct[Lrn] /* simulate a shift of "error" */
					if LrChk[Lrstate] == LrErrCode {
						goto Lrstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if LrDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", LrS[Lrp].yys)
				}
				Lrp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if LrDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", LrTokname(Lrtoken))
			}
			if Lrtoken == LrEofCode {
				goto ret1
			}
			Lrrcvr.char = -1
			Lrtoken = -1
			goto Lrnewstate /* try again in the same state */
		}
	}

	/* reduction by production Lrn */
	if LrDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Lrn, LrStatname(Lrstate))
	}

	Lrnt := Lrn
	Lrpt := Lrp
	_ = Lrpt // guard against "declared and not used"

	Lrp -= LrR2[Lrn]
	// Lrp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Lrp+1 >= len(LrS) {
		nyys := make([]LrSymType, len(LrS)*2)
		copy(nyys, LrS)
		LrS = nyys
	}
	LrVAL = LrS[Lrp+1]

	/* consult goto table to find next state */
	Lrn = LrR1[Lrn]
	Lrg := LrPgo[Lrn]
	Lrj := Lrg + LrS[Lrp].yys + 1

	if Lrj >= LrLast {
		Lrstate = LrAct[Lrg]
	} else {
		Lrstate = LrAct[Lrj]
		if LrChk[Lrstate] != -Lrn {
			Lrstate = LrAct[Lrg]
		}
	}
	// dummy call; replaced with literal code
	switch Lrnt {

	case 1:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:74
		{
			Lrlex.(*ParseContext).Grammar, _ = Lrlex.(*ParseContext).ToGrammar(LrDollar[1].Definitions, LrDollar[2].AdditionalSections)
		}
	case 2:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:80
		{
			LrVAL.AdditionalSections, _ = Lrlex.(*ParseContext).AddToAdditionalSections(LrDollar[1].AdditionalSections, LrDollar[2].AdditionalSection)
		}
	case 3:
		LrDollar = LrS[Lrpt-0 : Lrpt+1]
//line grammar.y:84
		{
			LrVAL.AdditionalSections, _ = Lrlex.(*ParseContext).NilToAdditionalSections()
		}
	case 4:
		LrDollar = LrS[Lrpt-3 : Lrpt+1]
//line grammar.y:90
		{
			LrVAL.AdditionalSection, _ = Lrlex.(*ParseContext).ToAdditionalSection(LrDollar[1].Generic_, LrDollar[2].Token, LrDollar[3].Token)
		}
	case 5:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:96
		{
			LrVAL.Definitions, _ = Lrlex.(*ParseContext).AddToDefs(LrDollar[1].Definitions, LrDollar[2].Definition)
		}
	case 6:
		LrDollar = LrS[Lrpt-3 : Lrpt+1]
//line grammar.y:100
		{
			LrVAL.Definitions, _ = Lrlex.(*ParseContext).AddExplicitToDefs(LrDollar[1].Definitions, LrDollar[2].Definition, nil)
		}
	case 7:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:104
		{
			LrVAL.Definitions, _ = Lrlex.(*ParseContext).DefToDefs(LrDollar[1].Definition)
		}
	case 8:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:108
		{
			LrVAL.Definitions, _ = Lrlex.(*ParseContext).ExplicitDefToDefs(LrDollar[1].Definition, nil)
		}
	case 9:
		LrDollar = LrS[Lrpt-5 : Lrpt+1]
//line grammar.y:116
		{
			LrVAL.Definition, _ = Lrlex.(*ParseContext).TermDeclToDef(LrDollar[1].Generic_, nil, LrDollar[3].Token, nil, LrDollar[5].Tokens)
		}
	case 10:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:120
		{
			LrVAL.Definition, _ = Lrlex.(*ParseContext).UntypedTermDeclToDef(LrDollar[1].Generic_, LrDollar[2].Tokens)
		}
	case 11:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:125
		{
			LrVAL.Definition, _ = Lrlex.(*ParseContext).StartDeclToDef(LrDollar[1].Generic_, LrDollar[2].Tokens)
		}
	case 12:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:128
		{
			LrVAL.Definition, _ = Lrlex.(*ParseContext).RuleToDef(LrDollar[1].Rule)
		}
	case 13:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:134
		{
			LrVAL.Generic_, _ = Lrlex.(*ParseContext).TokenToRword(LrDollar[1].Generic_)
		}
	case 14:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:138
		{
			LrVAL.Generic_, _ = Lrlex.(*ParseContext).TypeToRword(LrDollar[1].Generic_)
		}
	case 15:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:144
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).AddToNonemptyIdentList(LrDollar[1].Tokens, LrDollar[2].Token)
		}
	case 16:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:148
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).IdentToNonemptyIdentList(LrDollar[1].Token)
		}
	case 17:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:154
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).AddIdToNonemptyIdOrCharList(LrDollar[1].Tokens, LrDollar[2].Token)
		}
	case 18:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:158
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).AddCharToNonemptyIdOrCharList(LrDollar[1].Tokens, LrDollar[2].Token)
		}
	case 19:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:162
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).IdToNonemptyIdOrCharList(LrDollar[1].Token)
		}
	case 20:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:166
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).CharToNonemptyIdOrCharList(LrDollar[1].Token)
		}
	case 21:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:172
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).ListToIdOrCharList(LrDollar[1].Tokens)
		}
	case 22:
		LrDollar = LrS[Lrpt-0 : Lrpt+1]
//line grammar.y:175
		{
			LrVAL.Tokens, _ = Lrlex.(*ParseContext).NilToIdOrCharList()
		}
	case 23:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:181
		{
			LrVAL.Rule, _ = Lrlex.(*ParseContext).UnlabeledClauseToRule(LrDollar[1].Token, LrDollar[2].Tokens)
		}
	case 24:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:185
		{
			LrVAL.Rule, _ = Lrlex.(*ParseContext).ClausesToRule(LrDollar[1].Token, LrDollar[2].Clauses)
		}
	case 25:
		LrDollar = LrS[Lrpt-3 : Lrpt+1]
//line grammar.y:191
		{
			LrVAL.Clauses, _ = Lrlex.(*ParseContext).AddToLabeledClauses(LrDollar[1].Clauses, nil, LrDollar[3].Clause)
		}
	case 26:
		LrDollar = LrS[Lrpt-1 : Lrpt+1]
//line grammar.y:195
		{
			LrVAL.Clauses, _ = Lrlex.(*ParseContext).ClauseToLabeledClauses(LrDollar[1].Clause)
		}
	case 27:
		LrDollar = LrS[Lrpt-2 : Lrpt+1]
//line grammar.y:201
		{
			LrVAL.Clause, _ = Lrlex.(*ParseContext).ToLabeledClause(LrDollar[1].Token, LrDollar[2].Tokens)
		}
	}
	goto Lrstack /* stack new state and value */
}
