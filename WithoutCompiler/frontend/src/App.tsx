import { useState } from "react";
import { HashRouter, Routes, Route, NavLink } from "react-router-dom";
import "./App.css";
import { MountContainer } from "../wailsjs/go/services/ContainerService";
import CreateProject from "./view/CreateProject";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter the path where is the docker-compose file"
  );
  const [path, setPath] = useState("");
  const updatePath = (e: any) => setPath(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function mountContainer() {
    MountContainer(path).then(updateResultText);
  }

  return (
    <>
      <HashRouter basename="/">
        <div
          className="app-container"
          style={{ display: "flex", height: "100vh" }}
        >
          {/* --- MENU LATÉRAL (Sidebar) --- */}
          <nav
            style={{
              width: "200px",
              background: "#333",
              color: "white",
              padding: "20px",
            }}
          >
            <h3>Mon App</h3>
            <ul style={{ listStyle: "none", padding: 0 }}>
              <li>
                <NavLink to="/" style={{ color: "white" }}>
                  Accueil
                </NavLink>
              </li>
              <li>
                <NavLink to="/create" style={{ color: "white" }}>
                  Nouveau Projet
                </NavLink>
              </li>
              <li>
                <NavLink to="/dashboard" style={{ color: "white" }}>
                  Tableau de bord
                </NavLink>
              </li>
            </ul>
          </nav>

          {/* --- ZONE DE CONTENU QUI CHANGE --- */}
          <main style={{ flex: 1, padding: "20px", overflow: "auto" }}>
            <Routes>
              <Route path="/create" element={<CreateProject />} />
            </Routes>
          </main>
        </div>
      </HashRouter>
    </>
  );
}

export default App;
