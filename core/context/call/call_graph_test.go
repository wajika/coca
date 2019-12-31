package call_test

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/core/context/call"
	"github.com/phodal/coca/core/domain"
	"path/filepath"
	"testing"
)


func Test_should_generate_correct_files(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []domain.JClassNode
	analyser := call.NewCallGraph()
	codePath := "../../../_fixtures/call/call_api_test.json"
	codePath = filepath.FromSlash(codePath)

	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	dotContent := analyser.Analysis("com.phodal.pholedge.book.BookController.createBook", *&parsedDeps)

	g.Expect(dotContent).To(Equal(`digraph G {
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookFactory.create";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getIsbn";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getName";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookRepository.save";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.Book.getId";
"com.phodal.pholedge.book.BookController.createBook" -> "com.phodal.pholedge.book.BookService.createBook";
}
`))

}
