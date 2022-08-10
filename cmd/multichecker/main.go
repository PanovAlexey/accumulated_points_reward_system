package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/asmdecl"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/cgocall"
	"golang.org/x/tools/go/analysis/passes/composite"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/framepointer"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/sigchanyzer"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/testinggoroutine"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unsafeptr"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"
)

func main() {
	analyzerList := []*analysis.Analyzer{}

	// Add standard packet analyzers (golang.org/x/tools/go/analysis/passes)
	analyzerList = append(analyzerList,
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		errorsas.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		stdmethods.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		testinggoroutine.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)

	analyzerList = append(analyzerList,
		// staticcheck.io package SA class analyzers
		staticcheck.Analyzers["SA1000"],
		staticcheck.Analyzers["SA1001"],
		staticcheck.Analyzers["SA1002"],
		staticcheck.Analyzers["SA1003"],
		staticcheck.Analyzers["SA1004"],
		staticcheck.Analyzers["SA1005"],
		staticcheck.Analyzers["SA1006"],
		staticcheck.Analyzers["SA1007"],
		staticcheck.Analyzers["SA1008"],
		staticcheck.Analyzers["SA1010"],
		staticcheck.Analyzers["SA1011"],
		staticcheck.Analyzers["SA1012"],
		staticcheck.Analyzers["SA1013"],
		staticcheck.Analyzers["SA1014"],
		staticcheck.Analyzers["SA1015"],
		staticcheck.Analyzers["SA1016"],
		staticcheck.Analyzers["SA1017"],
		staticcheck.Analyzers["SA1018"],
		staticcheck.Analyzers["SA1019"],
		staticcheck.Analyzers["SA1020"],
		staticcheck.Analyzers["SA1021"],
		staticcheck.Analyzers["SA1023"],
		staticcheck.Analyzers["SA1024"],
		staticcheck.Analyzers["SA1025"],
		staticcheck.Analyzers["SA1026"],
		staticcheck.Analyzers["SA1027"],
		staticcheck.Analyzers["SA1028"],
		staticcheck.Analyzers["SA1029"],

		staticcheck.Analyzers["SA2000"],
		staticcheck.Analyzers["SA2001"],
		staticcheck.Analyzers["SA2002"],
		staticcheck.Analyzers["SA2003"],

		staticcheck.Analyzers["SA3000"],
		staticcheck.Analyzers["SA3001"],

		staticcheck.Analyzers["SA4000"],
		staticcheck.Analyzers["SA4001"],
		staticcheck.Analyzers["SA4003"],
		staticcheck.Analyzers["SA4004"],
		staticcheck.Analyzers["SA4006"],
		staticcheck.Analyzers["SA4008"],
		staticcheck.Analyzers["SA4009"],
		staticcheck.Analyzers["SA4010"],
		staticcheck.Analyzers["SA4011"],
		staticcheck.Analyzers["SA4012"],
		staticcheck.Analyzers["SA4013"],
		staticcheck.Analyzers["SA4014"],
		staticcheck.Analyzers["SA4015"],
		staticcheck.Analyzers["SA4016"],
		staticcheck.Analyzers["SA4017"],
		staticcheck.Analyzers["SA4018"],
		staticcheck.Analyzers["SA4019"],
		staticcheck.Analyzers["SA4020"],
		staticcheck.Analyzers["SA4021"],

		staticcheck.Analyzers["SA5000"],
		staticcheck.Analyzers["SA5001"],
		staticcheck.Analyzers["SA5002"],
		staticcheck.Analyzers["SA5003"],
		staticcheck.Analyzers["SA5004"],
		staticcheck.Analyzers["SA5005"],
		staticcheck.Analyzers["SA5007"],
		staticcheck.Analyzers["SA5008"],
		staticcheck.Analyzers["SA5009"],
		staticcheck.Analyzers["SA5010"],
		staticcheck.Analyzers["SA5011"],

		staticcheck.Analyzers["SA6000"],
		staticcheck.Analyzers["SA6001"],
		staticcheck.Analyzers["SA6002"],
		staticcheck.Analyzers["SA6003"],
		staticcheck.Analyzers["SA6005"],

		staticcheck.Analyzers["SA9001"],
		staticcheck.Analyzers["SA9002"],
		staticcheck.Analyzers["SA9003"],
		staticcheck.Analyzers["SA9004"],
		staticcheck.Analyzers["SA9005"],

		// staticcheck.io package other analyzers
		staticcheck.Analyzers["SA4022"],

		simple.Analyzers["S1009"],

		stylecheck.Analyzers["ST1013"],
		stylecheck.Analyzers["ST1012"],
	)

	multichecker.Main(analyzerList...)
}
