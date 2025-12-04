package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileService struct{
	ctx context.Context
}

func NewFileService() *FileService {
    return &FileService{}
}
func (c *FileService) Startup(ctx context.Context) {
	c.ctx = ctx
}

func (s *FileService) DownloadDockerConfig(url string, targetFolder string, fileName string) string {
    
   

    urlDockerFile := url +"/dockerfile"
    urlDockerCompose := url + "/docker-compose.yml"
    fullPathDockerFile := filepath.Join(targetFolder, "dockerfile")
    fullPathDockerCompose := filepath.Join(targetFolder, "docker-compose.yml")

    outDockerFile, err := os.Create(fullPathDockerCompose)
    if err != nil {
        return "Erreur création fichier: " + err.Error()
    }
    defer outDockerFile.Close()

    outDockerCompose, err := os.Create(fullPathDockerFile)
    if err != nil {
        return "Erreur création fichier: " + err.Error()
    }
    defer outDockerCompose.Close() 

    // 4. Lancer la requête HTTP (le "curl")
    respDockerFile, errDockerFile := http.Get(urlDockerFile)
    if errDockerFile != nil {
        return "Erreur téléchargement: " + errDockerFile.Error()
    }
    defer respDockerFile.Body.Close()

    respDockerCompose, errDockerCompose := http.Get(urlDockerCompose)
    if errDockerCompose != nil {
        return "Erreur téléchargement: " + errDockerCompose.Error()
    }
    defer respDockerCompose.Body.Close()


    // Vérifier que le serveur a répondu 200 OK
    if respDockerFile.StatusCode != http.StatusOK {
        return fmt.Sprintf("Erreur serveur: %s", respDockerFile.Status)
    }

    if respDockerCompose.StatusCode != http.StatusOK {
        return fmt.Sprintf("Erreur serveur: %s", respDockerCompose.Status)
    }

    _, err = io.Copy(outDockerFile, respDockerFile.Body)
    if err != nil {
        return "Erreur lors de l'écriture: " + err.Error()
    }

    _, err = io.Copy(outDockerCompose, respDockerCompose.Body)
    if err != nil {
        return "Erreur lors de l'écriture: " + err.Error()
    }

    return "Succès"
}