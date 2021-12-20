import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.bundle.min.js'
import { Routes, Route, Link, Outlet } from "react-router-dom";


import App from "./App"

ReactDOM.render(
  <BrowserRouter>
    <App>
      <Outlet />
    </App>
  </BrowserRouter>,
  document.getElementById("root")
)

