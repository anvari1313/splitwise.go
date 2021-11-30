package splitwise

import (
	"context"
)

// Comments contains methods to access and modify comments on expenses
type Comments interface {
	// Comments is responsible to get expense comments
	Comments(ctx context.Context) ([]Comment, error)

	// CreateComment creates a comment for an expense
	CreateComment(ctx context.Context, dto *CreateCommentDTO) (*Comment, error)
}

type Comment struct {
}

type CreateCommentDTO struct {
}
