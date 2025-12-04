package services

import (
	"context"
	"fmt"
	"os/exec"
)

type ContainerService struct {
	ctx context.Context

}

func NewContainerService() *ContainerService {
	return &ContainerService{
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

 