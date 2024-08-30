package parser

import (
	"fmt"
)

var _ LRReducer = &Reducer{}

type Reducer struct {
}

func (Reducer) ToGrammar(
	defs []Definition,
	additionalSections []*AdditionalSection) (
	*Grammar,
	error) {

	return NewGrammar(defs, additionalSections), nil
}

func (Reducer) AddToAdditionalSections(
	sections []*AdditionalSection,
	section *AdditionalSection) (
	[]*AdditionalSection,
	error) {

	return append(sections, section), nil
}

func (Reducer) NilToAdditionalSections() ([]*AdditionalSection, error) {
	return []*AdditionalSection{}, nil
}

func (Reducer) ToAdditionalSection(
	marker LRGenericSymbol,
	name *Token,
	content *Token) (
	*AdditionalSection,
	error) {

	return NewAdditionalSection(name, content), nil
}

func (Reducer) AddToDefs(
	defs []Definition,
	def Definition) (
	[]Definition,
	error) {

	return append(defs, def), nil
}

func (Reducer) AddExplicitToDefs(
	defs []Definition,
	def Definition,
	terminator LRGenericSymbol) (
	[]Definition, error) {

	return append(defs, def), nil
}

func (Reducer) DefToDefs(def Definition) ([]Definition, error) {
	return []Definition{def}, nil
}

func (Reducer) ExplicitDefToDefs(
	def Definition,
	terminator LRGenericSymbol) (
	[]Definition,
	error) {

	return []Definition{def}, nil
}

func (Reducer) TermDeclToDef(
	rword LRGenericSymbol,
	lt LRGenericSymbol,
	value *Token,
	gt LRGenericSymbol,
	terms []*Token) (
	Definition,
	error) {

	return NewTermDeclaration(rword, value, terms), nil
}

func (Reducer) UntypedTermDeclToDef(
	rword LRGenericSymbol,
	terms []*Token) (
	Definition,
	error) {

	return NewTermDeclaration(rword, nil, terms), nil
}

func (Reducer) StartDeclToDef(
	startKw LRGenericSymbol,
	ruleNames []*Token) (
	Definition,
	error) {

	return NewStartDeclaration(startKw, ruleNames), nil
}

func (Reducer) RuleToDef(rule *Rule) (Definition, error) {
	return rule, nil
}

func (Reducer) TokenToRword(tokenKw LRGenericSymbol) (LRGenericSymbol, error) {
	return tokenKw, nil
}

func (Reducer) TypeToRword(typeKw LRGenericSymbol) (LRGenericSymbol, error) {
	return typeKw, nil
}

func (Reducer) AddToNonemptyIdentList(
	identList []*Token,
	ident *Token) (
	[]*Token,
	error) {

	return append(identList, ident), nil
}

func (Reducer) IdentToNonemptyIdentList(ident *Token) ([]*Token, error) {
	return []*Token{ident}, nil
}

func (Reducer) AddIdToNonemptyIdOrCharList(
	list []*Token,
	id *Token) (
	[]*Token,
	error) {

	return append(list, id), nil
}

func (Reducer) AddCharToNonemptyIdOrCharList(
	list []*Token,
	char *Token) (
	[]*Token,
	error) {

	return append(list, char), nil
}

func (Reducer) IdToNonemptyIdOrCharList(id *Token) ([]*Token, error) {
	return []*Token{id}, nil
}

func (Reducer) CharToNonemptyIdOrCharList(char *Token) ([]*Token, error) {
	return []*Token{char}, nil
}

func (Reducer) ListToIdOrCharList(list []*Token) ([]*Token, error) {
	return list, nil
}

func (Reducer) NilToIdOrCharList() ([]*Token, error) {
	return []*Token{}, nil
}

func (Reducer) UnlabeledClauseToRule(
	ruleName *RuleDef,
	clauseBody []*Token) (
	*Rule,
	error) {

	return NewRule(ruleName, []*Clause{NewClause(nil, clauseBody)}), nil
}

func (Reducer) ToRule(
	ruleDef *RuleDef,
	clauses []*Clause) (
	*Rule,
	error) {

	if len(clauses) > 1 {
		for idx, clause := range clauses {
			if clause.Label != nil { // explicitly labelled
				continue
			}

			if len(clause.Body) == 1 && clause.Body[0].Id() == LRIdentifierToken {
				clause.Label = clause.Body[0]
			} else {
				return nil, fmt.Errorf(
					"rule %s (%s) clause %d must be explicitly named",
					ruleDef.Name.Value,
					ruleDef.Loc(),
					idx)
			}
		}
	}

	return NewRule(ruleDef, clauses), nil
}

func (Reducer) AddToClauses(
	clauses []*Clause,
	or LRGenericSymbol,
	clause *Clause) (
	[]*Clause,
	error) {
	return append(clauses, clause), nil
}

func (Reducer) ClauseToClauses(clause *Clause) ([]*Clause, error) {
	return []*Clause{clause}, nil
}

func (Reducer) UnlabeledToClause(
	clauseBody []*Token) (
	*Clause,
	error) {
	return NewClause(nil, clauseBody), nil
}

func (Reducer) LabeledToClause(
	label *Token,
	clauseBody []*Token) (
	*Clause,
	error) {
	return NewClause(label, clauseBody), nil
}
