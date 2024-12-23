package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"nganterin-go/mapper"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func (s *compServices) FileUpload(file []byte, data dto.FilesInputDTO) (*dto.FilesOutputDTO, error) {
	modelData := mapper.MapFilesInputToModel(data)

	uniqueName := helpers.GenerateUniqueFileName() + "." + data.Extension

	publicURL, metadata, err := s.SaveFileToDrive(file, uniqueName, data.MimeType)
	if err != nil {
		return nil, err
	}

	modelData.PublicURL = *publicURL
	modelData.Meta = *metadata

	repoData, err := s.repo.FileUpload(modelData)
	if err != nil {
		return nil, err
	}

	modelData.ID = repoData.ID

	result := mapper.MapFilesModelToOutput(*repoData)

	return &result, nil
}

func (s *compServices) SaveFileToDrive(file []byte, name, mimeType string) (*string, *string, error) {
	APPLICATION_CREDENTIALS := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	APPLICATION_FOLDER_ID := os.Getenv("APPLICATION_FOLDER_ID")

	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithCredentialsJSON([]byte(APPLICATION_CREDENTIALS)))
	if err != nil {
		return nil, nil, errors.New("failed to create drive service: " + err.Error())
	}

	fileMetadata := &drive.File{
		Name:    name,
		Parents: []string{APPLICATION_FOLDER_ID},
		MimeType: mimeType,

	}

	fileReader := bytes.NewReader(file)
	uploadedFile, err := driveService.Files.Create(fileMetadata).
		Media(fileReader).
		Fields("id, name, mimeType, size, createdTime").
		Do()
	if err != nil {
		return nil, nil, errors.New("failed to upload file to drive: " + err.Error())
	}

	_, err = driveService.Permissions.Create(uploadedFile.Id, &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}).Do()
	if err != nil {
		return nil, nil, errors.New("failed to set reader permissions: " + err.Error())
	}

	publicLink := fmt.Sprintf("https://lh3.googleusercontent.com/d/%s", uploadedFile.Id)

	metadata := map[string]interface{}{
		"id":          uploadedFile.Id,
		"name":        uploadedFile.Name,
		"mimeType":    uploadedFile.MimeType,
		"size":        uploadedFile.Size,
		"createdTime": uploadedFile.CreatedTime,
		"publicLink":  publicLink,
	}

	stringifiedMetadata, err := json.Marshal(metadata)
	if err != nil {
		return nil, nil, errors.New("failed to encode file metadata: " + err.Error())
	}

	return &publicLink, helpers.StringPointer(stringifiedMetadata), nil
}
