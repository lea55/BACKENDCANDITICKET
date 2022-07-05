package local

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/lea55/BACKENDCANDITICKET/src/global/core"
)

type AssetFile string

const (
	Image    AssetFile = "IMAGE"
	Document AssetFile = "DOCUMENT"
)

type File struct {
	c *core.AppConfig
}

func NewFile() *File {
	config := core.GetAppConfig()
	return &File{c: config}
}

func (f File) Delete(fileName string, fileType AssetFile) error {
	var initialPath string
	if fileType == Image {
		initialPath = f.c.ImagePath + fileName
	} else {
		initialPath = f.c.DocsPath + fileName
	}

	err := os.RemoveAll(initialPath)
	if err != nil {
		return err
	}

	return nil
}

func (f File) ValidateImageExtension(extension string) error {
	ext := filepath.Ext(extension)

	if ext != ".jpg" && ext != ".jpeg" && ext != ".tiff" && ext != ".png" {
		return errors.New("El archivo no es una imagen válida")
	}

	return nil
}

func (f File) Get(fileName string, fileType AssetFile) string {
	var completePath string

	if fileType == Image {
		completePath = f.c.ImagePath + fileName
	}

	if fileType == Document {
		completePath = f.c.DocsPath + fileName
	}

	_, err := os.Stat(completePath)
	if err != nil {
		return f.c.AssetsImageDefault
	}

	return completePath
}

func (f File) Save(file *multipart.FileHeader, docType AssetFile) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", errors.New("No se ha detectado una imagen valida")
	}
	defer src.Close()

	filename, _ := uuid.NewRandom()
	extensionSplit := strings.Split(file.Filename, ".")
	extensionName := extensionSplit[len(extensionSplit)-1]
	completeName := filename.String() + "." + extensionName

	var imagePath string

	if docType == Image {
		imagePath = f.c.ImagePath + completeName
	} else {
		imagePath = f.c.DocsPath + completeName
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return "", errors.New("Error en la generación de imagen")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", errors.New("Error al guardar imagen en servidor")
	}

	imageLink := f.GenerateDefaultDocumentLink(completeName, docType)

	return imageLink, nil
}

func (f File) GenerateDefaultDocumentLink(documentName string, docType AssetFile) string {
	var initialPath string
	if docType == Image {
		initialPath = "files/image/"
	} else {
		initialPath = "files/document/"
	}

	return f.c.ServerUrl + f.c.BaseUrl + initialPath + documentName
}
