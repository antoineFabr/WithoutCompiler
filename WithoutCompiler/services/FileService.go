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
// DownloadFile télécharge une URL et l'enregistre dans un dossier spécifique
func (s *FileService) DownloadDockerFile(url string, targetFolder string, fileName string) string {
    
   

    // 2. Construire le chemin complet (ex: C:/Users/Downloads/docker.zip)
    fullPath := filepath.Join(targetFolder, fileName)

    // 3. Créer le fichier vide sur le disque
    out, err := os.Create(fullPath)
    if err != nil {
        return "Erreur création fichier: " + err.Error()
    }
    defer out.Close() // On fermera le fichier à la fin

    // 4. Lancer la requête HTTP (le "curl")
    resp, err := http.Get(url)
    if err != nil {
        return "Erreur téléchargement: " + err.Error()
    }
    defer resp.Body.Close() // On fermera la connexion réseau à la fin

    // Vérifier que le serveur a répondu 200 OK
    if resp.StatusCode != http.StatusOK {
        return fmt.Sprintf("Erreur serveur: %s", resp.Status)
    }

    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return "Erreur lors de l'écriture: " + err.Error()
    }

    return "Succès : Fichier téléchargé dans " + fullPath
}