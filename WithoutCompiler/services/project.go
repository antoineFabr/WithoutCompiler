package services

import (
	"context"
	"fmt"
)

type ProjectService struct {
	ctx context.Context
	fileService *FileService

}

func NewProjectService(fs *FileService) *ProjectService {
	return &ProjectService{
		fileService: fs,

	}
}

func (p *ProjectService) Startup(ctx context.Context) {
	p.ctx = ctx
}

func (p *ProjectService) CreatProject(path string, language string) string {

	url := "https://raw.githubusercontent.com/antoineFabr/WithoutCompiler/refs/heads/main/dockerFile/"+language+"/dockerfile"

	res := p.fileService.DownloadDockerFile(url,path,"dockerfile")
	return fmt.Sprintf(res)
}

 