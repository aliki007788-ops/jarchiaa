package service

import "mime/multipart"

type UploadService struct{}

func (s *UploadService) Upload(file *multipart.FileHeader) (string, error) {
return "https://cdn.jarchia.ir/file.jpg", nil
}


