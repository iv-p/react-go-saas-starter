import React from "react";
import { hydrate } from "react-dom";
import "./index.css";
import App from "./App";

hydrate(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
