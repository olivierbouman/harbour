import React, { useState } from "react";
import { Routes, Route, Link, Outlet } from "react-router-dom";
import Footer from "./AppFooter";
import AppContent from "./AppContent";
import Navbar from "./Navbar";

import Home from './components/Home'
import Boats from './components/Boats'
import Boat from "./components/Boat";
import Admin from './components/Admin'
import Categories from "./components/Categories";
import Category from "./components/Category";


const App = () => {
  const [counter, setCounter] = useState(0);

  const myProps = {
    title: "De Levant",
    boten: "De Juliette",
  };

  return (
    <div>
      <Navbar {...myProps} />
      <AppContent counter={counter} setCounter={setCounter} />

        <div className="container">
          <div className="row">
            <h1 className="mt-3">Go find your boat!</h1>
            <hr className="mb-3" />

            <div className="row">
              <div className="col-md-2">
                <nav>
                  <ul className="list-group">
                    <li className="list-group-item">
                      <Link to="/">Home</Link>
                    </li>
                    <li className="list-group-item">
                      <Link to="/boats">Boats</Link>
                    </li>
                    <li className="list-group-item">
                      <Link to="/by-category">Categories</Link>
                    </li>                    
                    <li className="list-group-item">
                      <Link to="/admin">Admin</Link>
                    </li>
                  </ul>
                </nav>
              </div>
              <div className="col-md-10">
                <Routes>
                  <Route path="/" element={<Home />} />
                  <Route path="boats" element={<Boats />}>
                    <Route path=":id" element={<Boat />} />
                  </Route>
                  <Route path="by-category" element={<Categories />} />
                  <Route 
                    path="by-category/yachts"
                    element={<Category title={`Yachts`} />}
                  />
                  <Route 
                    path="by-category/motorboats"
                    element={<Category title={`Motor Boats`} />}
                  />
                  <Route 
                    path="by-category/other-categories"
                    element={<Category title={`Other Categories`} />}
                  />
                  <Route exact path="admin" element={<Admin />} />
                </Routes>
              </div>
            </div>
          </div>
        </div>

      <Footer name="Juliette" counter={counter} />
    </div>
  );
};

export default App;