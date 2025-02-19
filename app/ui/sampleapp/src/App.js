import React, { useEffect, useState } from "react";
import logo from "./logo.svg";
import "./App.css";

function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/api/userinfo", { credentials: "include" })  // Ensure cookies are sent
      .then((res) => res.json())
      .then((data) => {
        setUser(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Failed to fetch user info:", err);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <p>Loading user info...</p>;
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>{user ? `Hello, ${user.preferred_username}` : "Not logged in"}</p>

        {user && (
          <div style={{ textAlign: "left", padding: "10px", background: "#282c34", borderRadius: "8px" }}>
            <h3>User Info:</h3>
            <pre style={{ whiteSpace: "pre-wrap", wordBreak: "break-word" }}>
              {JSON.stringify(user, null, 2)}
            </pre>
          </div>
        )}
      </header>
    </div>
  );
}

export default App;