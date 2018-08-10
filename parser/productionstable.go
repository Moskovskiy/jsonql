// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"fmt"
	"strconv"
        "strings"

        "github.com/teslamotors/jsonql/ast"
        "github.com/teslamotors/jsonql/token"
)

func String(a Attrib) string {
     return string(a.(*token.Token).Lit) 
}

func SingleUnquote(tok string) (string, error) {
     val := tok[1:len(tok)-1] 
     // do something slow and hacky so we can re-use strconv.Unquote
     val = strings.Replace(val, "\\'", "'", -1)
     val = strings.Replace(val, "\"", "\\\"", -1)
     val = fmt.Sprintf("\"%s\"", val)
     unq, err := strconv.Unquote(val)
     fmt.Printf("unq: %#v => %#v (err: %v)\n", val, unq, err)
     return unq, err
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Expression	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : OrExpr	<<  >>`,
		Id:         "Expression",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OrExpr : AndExpr	<<  >>`,
		Id:         "OrExpr",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OrExpr : OrExpr "||" AndExpr	<< ast.Or(X[0], X[2]) >>`,
		Id:         "OrExpr",
		NTType:     2,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Or(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `AndExpr : CompareExpr	<<  >>`,
		Id:         "AndExpr",
		NTType:     3,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AndExpr : AndExpr "&&" CompareExpr	<< ast.And(X[0], X[2]) >>`,
		Id:         "AndExpr",
		NTType:     3,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.And(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : ValueExpr	<<  >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "=" ValueExpr	<< ast.Eq(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Eq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "==" ValueExpr	<< ast.Eq(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Eq(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "!=" ValueExpr	<< ast.NE(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NE(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "<=" ValueExpr	<< ast.LE(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LE(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr ">=" ValueExpr	<< ast.GE(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GE(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr ">" ValueExpr	<< ast.GT(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "<" ValueExpr	<< ast.LT(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LT(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "is" ExistentialWord	<< ast.Is(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Is(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "isnot" ExistentialWord	<< ast.IsNot(X[0], X[2]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.IsNot(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : CompareExpr "is" "not" ExistentialWord	<< ast.IsNot(X[0], X[3]) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      16,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.IsNot(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `ExistentialWord : "null"	<<  >>`,
		Id:         "ExistentialWord",
		NTType:     5,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ExistentialWord : "defined"	<<  >>`,
		Id:         "ExistentialWord",
		NTType:     5,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ValueExpr : RegexpExpr	<<  >>`,
		Id:         "ValueExpr",
		NTType:     6,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ValueExpr : AddExpr	<<  >>`,
		Id:         "ValueExpr",
		NTType:     6,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RegexpArgument : Identifier	<<  >>`,
		Id:         "RegexpArgument",
		NTType:     7,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RegexpArgument : StringLiteral	<<  >>`,
		Id:         "RegexpArgument",
		NTType:     7,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RegexpExpr : RegexpArgument "~=" StringLiteral	<< ast.RegexpMatch(X[0], X[2]) >>`,
		Id:         "RegexpExpr",
		NTType:     8,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.RegexpMatch(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `RegexpExpr : RegexpArgument "!~=" StringLiteral	<< ast.RegexpNegMatch(X[0], X[2]) >>`,
		Id:         "RegexpExpr",
		NTType:     8,
		Index:      24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.RegexpNegMatch(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `AddExpr : MulExpr	<<  >>`,
		Id:         "AddExpr",
		NTType:     9,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddExpr : AddExpr "+" MulExpr	<< ast.Add(X[0], X[2]) >>`,
		Id:         "AddExpr",
		NTType:     9,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Add(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `AddExpr : AddExpr "-" MulExpr	<< ast.Sub(X[0], X[2]) >>`,
		Id:         "AddExpr",
		NTType:     9,
		Index:      27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Sub(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `MulExpr : ExpExpr	<<  >>`,
		Id:         "MulExpr",
		NTType:     10,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulExpr : MulExpr "*" ExpExpr	<< ast.Mul(X[0], X[2]) >>`,
		Id:         "MulExpr",
		NTType:     10,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Mul(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `MulExpr : MulExpr "/" ExpExpr	<< ast.Div(X[0], X[2]) >>`,
		Id:         "MulExpr",
		NTType:     10,
		Index:      30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Div(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `MulExpr : MulExpr "%!"(MISSING) ExpExpr	<< ast.Mod(X[0], X[2]) >>`,
		Id:         "MulExpr",
		NTType:     10,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Mod(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `ExpExpr : UnaryExpr	<<  >>`,
		Id:         "ExpExpr",
		NTType:     11,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ExpExpr : ExpExpr "^" UnaryExpr	<< ast.Exp(X[0], X[2]) >>`,
		Id:         "ExpExpr",
		NTType:     11,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Exp(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `UnaryExpr : ParenExpr	<<  >>`,
		Id:         "UnaryExpr",
		NTType:     12,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryExpr : "-" ParenExpr	<< ast.Negative(X[1]) >>`,
		Id:         "UnaryExpr",
		NTType:     12,
		Index:      35,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Negative(X[1])
		},
	},
	ProdTabEntry{
		String: `UnaryExpr : "!" UnaryExpr	<< ast.Not(X[1]) >>`,
		Id:         "UnaryExpr",
		NTType:     12,
		Index:      36,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Not(X[1])
		},
	},
	ProdTabEntry{
		String: `ParenExpr : Term	<<  >>`,
		Id:         "ParenExpr",
		NTType:     13,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ParenExpr : "(" Expression ")"	<< X[1], nil >>`,
		Id:         "ParenExpr",
		NTType:     13,
		Index:      38,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Term : Literal	<<  >>`,
		Id:         "Term",
		NTType:     14,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Identifier	<<  >>`,
		Id:         "Term",
		NTType:     14,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NullLiteral : "null"	<< nil, nil >>`,
		Id:         "NullLiteral",
		NTType:     15,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `BooleanLiteral : "true"	<< true, nil >>`,
		Id:         "BooleanLiteral",
		NTType:     16,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `BooleanLiteral : "false"	<< false, nil >>`,
		Id:         "BooleanLiteral",
		NTType:     16,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `NumericLiteral : intLit	<< strconv.ParseInt(String(X[0]), 0, 64) >>`,
		Id:         "NumericLiteral",
		NTType:     17,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return strconv.ParseInt(String(X[0]), 0, 64)
		},
	},
	ProdTabEntry{
		String: `NumericLiteral : floatLit	<< strconv.ParseFloat(String(X[0]), 64) >>`,
		Id:         "NumericLiteral",
		NTType:     17,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return strconv.ParseFloat(String(X[0]), 64)
		},
	},
	ProdTabEntry{
		String: `StringLiteral : doubleStringLit	<< strconv.Unquote(String(X[0])) >>`,
		Id:         "StringLiteral",
		NTType:     18,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return strconv.Unquote(String(X[0]))
		},
	},
	ProdTabEntry{
		String: `StringLiteral : singleStringLit	<< SingleUnquote(String(X[0])) >>`,
		Id:         "StringLiteral",
		NTType:     18,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return SingleUnquote(String(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : NullLiteral	<< ast.Literal(X[0]) >>`,
		Id:         "Literal",
		NTType:     19,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Literal(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : BooleanLiteral	<< ast.Literal(X[0]) >>`,
		Id:         "Literal",
		NTType:     19,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Literal(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : NumericLiteral	<< ast.Literal(X[0]) >>`,
		Id:         "Literal",
		NTType:     19,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Literal(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : StringLiteral	<< ast.Literal(X[0]) >>`,
		Id:         "Literal",
		NTType:     19,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Literal(X[0])
		},
	},
	ProdTabEntry{
		String: `ObjectKey : symbol	<< String(X[0]), nil >>`,
		Id:         "ObjectKey",
		NTType:     20,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return String(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Identifier : ObjectKey	<< ast.ObjectKey(X[0]) >>`,
		Id:         "Identifier",
		NTType:     21,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ObjectKey(X[0])
		},
	},
	ProdTabEntry{
		String: `Identifier : Identifier "." ObjectKey	<< ast.SelectKey(X[0].(ast.Expr), X[2]) >>`,
		Id:         "Identifier",
		NTType:     21,
		Index:      54,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SelectKey(X[0].(ast.Expr), X[2])
		},
	},
	ProdTabEntry{
		String: `Identifier : Identifier "[" ValueExpr "]"	<< ast.Index(X[0].(ast.Expr), X[2].(ast.Expr)) >>`,
		Id:         "Identifier",
		NTType:     21,
		Index:      55,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Index(X[0].(ast.Expr), X[2].(ast.Expr))
		},
	},
	ProdTabEntry{
		String: `Identifier : Identifier ".[" ValueExpr "]"	<< ast.Index(X[0].(ast.Expr), X[2].(ast.Expr)) >>`,
		Id:         "Identifier",
		NTType:     21,
		Index:      56,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Index(X[0].(ast.Expr), X[2].(ast.Expr))
		},
	},
}
