package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure"
	"path/filepath"
)

var allApis []domain.RestApi

type JavaApiApp struct {
}

func (j *JavaApiApp) AnalysisPath(codeDir string, parsedDeps []domain.JClassNode, identifiersMap map[string]domain.JIdentifier, diMap map[string]string) []domain.RestApi {
	files := infrastructure.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := infrastructure.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaApiListener(identifiersMap, diMap)
		listener.appendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		allApis = listener.getClassApis()
	}

	return *&allApis
}
