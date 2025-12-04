package services

import (
	"context"
)

type ProjectService struct {
	ctx context.Context
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (p *ProjectService) Startup(ctx context.Context) {
	p.ctx = ctx
}

func (p *ProjectService) CreatProject(path string, language string) string {
	
 return "Fonctionne"
}

 