package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"graphql-file-info/graph/generated"
	"graphql-file-info/graph/model"
	"graphql-file-info/graph/repository"
	"graphql-file-info/helper"

	"github.com/99designs/gqlgen/graphql"
	"github.com/inhies/go-bytesize"
)

// UploadFile is the resolver for the uploadFile field.
func (r *mutationResolver) UploadFile(ctx context.Context, file graphql.Upload) (*model.UploadFileResult, error) {
	if !helper.IsCorrectSize(uint64(file.Size)) {
		return nil, fmt.Errorf("max file size is %s", helper.MaxFileSize)
	}

	return &model.UploadFileResult{
		File: &model.FileData{
			Name: file.Filename,
			Size: fmt.Sprintf("%s", bytesize.ByteSize(file.Size)),
			Mime: file.ContentType,
		},
	}, nil
}

// UploadFiles is the resolver for the uploadFiles field.
func (r *mutationResolver) UploadFiles(ctx context.Context, files []*graphql.Upload) (*model.UploadFilesResult, error) {
	var result []*model.FileData

	for _, file := range files {
		if !helper.IsCorrectSize(uint64(file.Size)) {
			return nil, fmt.Errorf("max file size is %s", helper.MaxFileSize)
		}

		result = append(result, &model.FileData{
			Name: file.Filename,
			Size: fmt.Sprintf("%s", bytesize.ByteSize(file.Size)),
			Mime: file.ContentType,
		})
	}

	return &model.UploadFilesResult{Files: result}, nil
}

// AddListItem is the resolver for the addListItem field.
func (r *mutationResolver) AddListItem(ctx context.Context, name string) (*model.AddListItemResult, error) {
	item := r.ItemsRepository.Add(name)

	return &model.AddListItemResult{Item: &model.Item{ID: item.ID, Name: item.Name}}, nil
}

// FileInfo is the resolver for the fileInfo field.
func (r *queryResolver) FileInfo(ctx context.Context, file graphql.Upload) (*model.FileInfoResult, error) {
	if !helper.IsCorrectSize(uint64(file.Size)) {
		return nil, fmt.Errorf("max file size is %s", helper.MaxFileSize)
	}

	return &model.FileInfoResult{
		File: &model.FileData{
			Name: file.Filename,
			Size: fmt.Sprintf("%s", bytesize.ByteSize(file.Size)),
			Mime: file.ContentType,
		},
	}, nil
}

// FilesInfo is the resolver for the filesInfo field.
func (r *queryResolver) FilesInfo(ctx context.Context, files []*graphql.Upload) (*model.FilesInfoResult, error) {
	var result []*model.FileData

	for _, file := range files {
		if !helper.IsCorrectSize(uint64(file.Size)) {
			return nil, fmt.Errorf("max file size is %s", helper.MaxFileSize)
		}

		result = append(result, &model.FileData{
			Name: file.Filename,
			Size: fmt.Sprintf("%s", bytesize.ByteSize(file.Size)),
			Mime: file.ContentType,
		})
	}

	return &model.FilesInfoResult{Files: result}, nil
}

// GetListItems is the resolver for the getListItems field.
func (r *queryResolver) GetListItems(ctx context.Context, input *model.ListItemsInput) (*model.ListItemsResult, error) {
	var items []repository.Item
	var out []*model.Item

	if input != nil && len(input.IDAnyOf) > 0 {
		items = r.ItemsRepository.GetByIDs(input.IDAnyOf)
	} else {
		items = r.ItemsRepository.GetAll()
	}

	for _, item := range items {
		out = append(out, &model.Item{ID: item.ID, Name: item.Name})
	}

	return &model.ListItemsResult{
		Items: out,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
