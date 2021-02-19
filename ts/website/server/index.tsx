import fs from "fs";
import path from "path";
import React from "react";
import { renderToString } from "react-dom/server";
import express from "express";

import App from "../src/App";
const app = express();

const indexHTML = fs.readFileSync(path.resolve(__dirname, "index.html"), {
  encoding: "utf8",
});

app.use(
  /\.(js|css|map|ico|jpe?g|png)$/,
  express.static(path.resolve(__dirname, "."))
);

app.get("*", (req, res) => {
  let appHTML = renderToString(<App />);

  const html = indexHTML.replace(
    '<div id="root"></div>',
    `<div id="root">${appHTML}</div>`
  );

  res.contentType("text/html");
  res.status(200);

  return res.send(html);
});

app.listen("80", () => {
  console.log("Express server started at http://localhost:80");
});
