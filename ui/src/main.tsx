import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "app/app";

// mount to metis node
ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("metis")
);
