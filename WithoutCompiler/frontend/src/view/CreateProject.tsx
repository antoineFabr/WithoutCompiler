import { useState } from "react";
import { OpenFolderDialog } from "../../wailsjs/go/main/App";

export default function CreateProject() {
  const [path, setPath] = useState("");
  const [language, setLanguage] = useState("");
  const languages = ["JavaScript", "Go", "Rust", "C"];
  const handleSelectFolder = () => {
    OpenFolderDialog().then((selectedPath) => {
      if (selectedPath) {
        setPath(selectedPath);
      }
    });
  };

  const handleCreateProject = () => {};
  return (
    <>
      <h1>Create a project</h1>
      <label>Dossier du projet :</label>

      <div style={{ display: "flex", gap: "10px", alignItems: "center" }}>
        {/* Champ en lecture seule pour afficher le chemin */}
        <input
          type="text"
          value={path}
          readOnly
          placeholder="Aucun dossier sélectionné"
          style={{ width: "300px" }}
        />

        {/* Bouton qui déclenche la fenêtre native */}
        <button onClick={handleSelectFolder}>Parcourir...</button>

        <label>Dans quel langage voulez vous coder ?</label>
        <select
          name="language"
          value={language}
          onChange={(e) => setLanguage(e.target.value)}
        >
          {languages.map((x) => (
            <option key={x} value={x}>
              {x}
            </option>
          ))}
        </select>
        {language && path && <button>créer le projet</button>}
      </div>
    </>
  );
}
