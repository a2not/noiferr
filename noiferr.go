package noiferr

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "noiferr is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "noiferr",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, f := range funcs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				switch instr := instr.(type) {
				case *ssa.Call:
					if !isHandledWithCond(instr) {
						pass.Reportf(instr.Pos(), "error received but not handled")
					}
				}
			}
		}
	}

	return nil, nil
}

func isError(typ types.Type) bool {
	return types.Implements(typ, errType) || types.Implements(types.NewPointer(typ), errType)
}

func isHandledWithCond(callInstr *ssa.Call) bool {
	switch typ := callInstr.Type().(type) {
	case *types.Tuple:
		for i := 0; i < typ.Len(); i++ {
			if isError(typ.At(i).Type()) {
				ref := (*callInstr.Referrers())[i]
				extract, ok := ref.(*ssa.Extract)
				if ok && !errorVarHandled(extract.Referrers()) {
					return false
				}
			}
		}
	default:
		if isError(typ) {
			return errorVarHandled(callInstr.Referrers())
		}
	}
	return true
}

// func errorVarHandeld(binop *ssa.BinOp) bool {
func errorVarHandled(instrs *[]ssa.Instruction) bool {
	for _, instr := range *instrs {
		switch instr := instr.(type) {
		case *ssa.BinOp:
			if instr.Op != token.NEQ && instr.Op != token.EQL {
				continue
			}
			for _, ref := range *instr.Referrers() {
				_, ok := ref.(*ssa.If)
				if ok {
					return true
				}
			}
		}
	}
	return false
}
