package services

import (
	"context"
	"fmt"
	"os/exec"
)

type ContainerService struct {
	ctx context.Context
	fileService *FileService

}

func NewContainerService(fs *FileService) *ContainerService {
	return &ContainerService{
		fileService: fs,
	}
}

func (c *ContainerService) Startup(ctx context.Context) {
	c.ctx = ctx
}

func (c *ContainerService) MountContainer(path string)string {
	cmd := exec.Command("docker","compose","up","-d")
	cmd.Dir = path
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Sprintf("Erreur: %s\nDétails: %s", err.Error(), string(output))
	}
	
	return "Container bien créé"

}

 