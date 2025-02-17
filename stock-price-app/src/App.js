import React, { useState, useRef, useEffect } from "react";
import "./App.css";

function App() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [stockSymbol, setStockSymbol] = useState("");
  const [latestPrice, setLatestPrice] = useState("N/A");
  const [priceUpdates, setPriceUpdates] = useState([]);
  const wsRef = useRef(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log("Submitting form...");

    // Create form data for the login API
    const formData = new FormData();
    formData.append("email", email);
    formData.append("password", password);

    try {
      // Login API call
      const response = await fetch("http://localhost:7001/login", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("Login failed.");
      }

      const data = await response.json();
      const token = data.token;
      console.log("Token received:", token);

      // Build the WebSocket URL. Since browsers don’t allow custom headers in native WebSocket,
      // pass the token as a query parameter.
      const wsUrl = `ws://localhost:7001/ws/price?stockID=${encodeURIComponent(
        stockSymbol
      )}&token=${encodeURIComponent(token)}`;
      const ws = new WebSocket(wsUrl);
      wsRef.current = ws;

      ws.onopen = () => {
        console.log("WebSocket connected.");
      };

      ws.onmessage = (event) => {
        console.log("WebSocket message:", event.data);
        try {
          const message = JSON.parse(event.data);
          const price = message.price;
          const timestamp = new Date().toLocaleString();

          // Update the UI
          setLatestPrice(price);
          setPriceUpdates((prev) => [...prev, { price, timestamp }]);
        } catch (err) {
          console.error("Error parsing WebSocket message:", err);
        }
      };

      ws.onerror = (error) => {
        console.error("WebSocket error:", error);
      };

      ws.onclose = () => {
        console.log("WebSocket closed.");
      };
    } catch (err) {
      console.error("Error during login or WebSocket setup:", err);
    }
  };

  // Clean up WebSocket when the component unmounts
  useEffect(() => {
    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, []);

  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h1>Stock Price Viewer</h1>
      <form onSubmit={handleSubmit} style={{ marginBottom: "20px" }}>
        <div style={{ marginBottom: "10px" }}>
          <label>
            Email: <br />
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              style={{ width: "300px", padding: "5px" }}
            />
          </label>
        </div>
        <div style={{ marginBottom: "10px" }}>
          <label>
            Password: <br />
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              style={{ width: "300px", padding: "5px" }}
            />
          </label>
        </div>
        <div style={{ marginBottom: "10px" }}>
          <label>
            Stock Symbol: <br />
            <input
              type="text"
              value={stockSymbol}
              onChange={(e) => setStockSymbol(e.target.value)}
              required
              style={{ width: "300px", padding: "5px" }}
            />
          </label>
        </div>
        <button type="submit" style={{ padding: "10px 20px" }}>
          Get Stock Price
        </button>
      </form>

      <h2>
        Latest Price: <span>{latestPrice}</span>
      </h2>
      <div>
        <h3>Price Updates:</h3>
        <ul>
          {priceUpdates.map((update, index) => (
            <li key={index}>
              {update.timestamp}: {update.price}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
