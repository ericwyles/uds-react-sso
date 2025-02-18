import React from "react";
import { useKeycloak } from "@react-keycloak/web";
import logo from "./logo.svg";
import "./App.css";

function App() {
  const { keycloak, initialized } = useKeycloak();

  // Ensure authentication is initialized before rendering
  if (!initialized) {
    return <p>Loading authentication...</p>;
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Hello UDS</p>

        {keycloak.authenticated ? (
          <>
            <p>Welcome, {keycloak.tokenParsed?.preferred_username}!</p>
            <button onClick={() => keycloak.logout()}>Logout</button>
          </>
        ) : (
          <>
            <button onClick={() => keycloak.login()}>Login</button>
            <button onClick={() => keycloak.login({ idpHint: "saml" })}>
              Login with Google
            </button>
          </>
        )}
      </header>
    </div>
  );
}

export default App;