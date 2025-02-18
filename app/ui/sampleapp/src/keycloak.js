import Keycloak from "keycloak-js";

const keycloak = new Keycloak({
  url: "https://sso.uds.dev",
  realm: "uds",
  clientId: "react-auth",
});

// Disable iframe checking to prevent MutationObserver-related errors
const keycloakInitOptions = {
  checkLoginIframe: false,
  onLoad: "check-sso",
  pkceMethod: "S256",  // ✅ Secure proof key for auth
  onLoad: "check-sso",
};

export { keycloak, keycloakInitOptions };
export default keycloak;