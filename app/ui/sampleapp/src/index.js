import React from "react";
import ReactDOM from "react-dom/client";
import { ReactKeycloakProvider } from "@react-keycloak/web";
import keycloak, { keycloakInitOptions } from "./keycloak";
import App from "./App";

const root = ReactDOM.createRoot(document.getElementById("root"));

root.render(
  <ReactKeycloakProvider authClient={keycloak} initOptions={keycloakInitOptions}>
    <App />
  </ReactKeycloakProvider>
);