package services

import (
	"context"
	"fmt"
	"os/exec"
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

	url := "https://raw.githubusercontent.com/antoineFabr/WithoutCompiler/refs/heads/main/dockerFile/"+language

	

	res := p.fileService.DownloadDockerConfig(url,path,"dockerfile")

	commands := [][]string{
		{"docker", "compose", "run", "--rm", "app", "npm", "init", "-y"},
        {"docker", "compose", "up", "-d"},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = path
		output, err := cmd.CombinedOutput()
		if err != nil {
            return fmt.Sprintf("erreur sur la commande %s: %s\nSortie: %s", args, err, string(output))
        }
	}
	return fmt.Sprintf(res)
}

 