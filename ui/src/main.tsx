import React from "react";
import {createRoot} from 'react-dom/client';
import "./index.css";
// import App from "app/app";
import Model from "app/model"

const akEl = document.getElementById("ak");
if (akEl == null) {
  throw new Error("Not found root node(#ak)!")
}

const root = createRoot(akEl);

root.render(
  <React.StrictMode>
    {/*<App/>*/}
    <Model/>
  </React.StrictMode>
)

