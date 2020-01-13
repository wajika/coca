package trial

import "github.com/phodal/coca/pkg/domain"

type CodeMember struct {
	ID           string
	Name         string
	Type         string
	ClassNodes   []domain.JClassNode
	Namespace    []string
	FileID       string
	DataStructID string
	Position     CodePosition
}
