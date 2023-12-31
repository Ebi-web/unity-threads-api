package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"unity-threads/graph/model"
)

// CreateThread is the resolver for the createThread field.
func (r *mutationResolver) CreateThread(ctx context.Context, input model.NewThread) (*model.Thread, error) {
	panic(fmt.Errorf("not implemented: CreateThread - createThread"))
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Thread is the resolver for the thread field.
func (r *postResolver) Thread(ctx context.Context, obj *model.Post) (*model.Thread, error) {
	panic(fmt.Errorf("not implemented: Thread - thread"))
}

// Threads is the resolver for the threads field.
func (r *queryResolver) Threads(ctx context.Context) ([]*model.Thread, error) {
	panic(fmt.Errorf("not implemented: Threads - threads"))
}

// User is the resolver for the user field.
func (r *threadResolver) User(ctx context.Context, obj *model.Thread) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Posts is the resolver for the posts field.
func (r *threadResolver) Posts(ctx context.Context, obj *model.Thread) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Thread returns ThreadResolver implementation.
func (r *Resolver) Thread() ThreadResolver { return &threadResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type threadResolver struct{ *Resolver }
