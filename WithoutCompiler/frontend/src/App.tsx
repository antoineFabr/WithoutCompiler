import { useState } from "react";
import "./App.css";
import { MountContainer } from "../wailsjs/go/services/ContainerService";

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
    <div id="App">
      <h1>Mount container</h1>
      <div id="result" className="result">
        {resultText}
      </div>
      <div id="input" className="input-box">
        <input
          id="name"
          className="input"
          onChange={updatePath}
          autoComplete="off"
          name="input"
          type="text"
        />
        <button className="btn" onClick={mountContainer}>
          Mount
        </button>
      </div>
    </div>
  );
}

export default App;
