package pkg_utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

func GetUploadedFileDirectory() (string, error) {
	var rootDirectory string
	if path, pathErr := os.Getwd(); pathErr != nil {
		return "", pkg_types.NewServerError(
			"error in getting root path directory",
			"UserHandler.UploadProfile",
			pathErr.Error(),
		)
	} else {
		rootDirectory = path
	}
	return filepath.Join(rootDirectory, "/uploads"), nil
}

func GenerateFilePath(file multipart.File, fileHandler *multipart.FileHeader) (string, string, error) {
	var baseFilename string

	extname := filepath.Ext(fileHandler.Filename)

	rootDir, rootDirectoryErr := GetUploadedFileDirectory()

	if rootDirectoryErr != nil {
		return "", "", rootDirectoryErr
	}

	if generatedCode, generatedCodeErr := GenerateCode(50); generatedCodeErr != nil {
		return "", "", pkg_types.NewServerError(
			"error in generating random code",
			"UserHandle.UploadProfile.GenerateCode",
			generatedCodeErr.Error(),
		)
	} else {
		baseFilename = generatedCode
	}
	filename := baseFilename + extname
	return filepath.Join(rootDir, filename), filename, nil
}

func FileUploader(file multipart.File, fileHandler *multipart.FileHeader) (string, error) {
	baseFilePath, filename, generateFilenameErr := GenerateFilePath(file, fileHandler)
	if generateFilenameErr != nil {
		return "", generateFilenameErr
	}
	uploadedFile, uploadedFileErr := os.Create(baseFilePath)
	if uploadedFileErr != nil {
		return "", pkg_types.NewServerError(
			"error in storing file in filesystem",
			"UserHandle.UploadProfile.OsCreate",
			uploadedFileErr.Error(),
		)
	}
	defer uploadedFile.Close()
	if _, copyErr := io.Copy(uploadedFile, file); copyErr != nil {
		return "", pkg_types.NewServerError(
			"error in copy file in filesystem",
			"UserHandle.UploadProfile.Copy",
			copyErr.Error(),
		)
	}
	return filename, nil
}
