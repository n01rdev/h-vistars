import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./components/pages/Home.jsx";
import Authentication from "./components/pages/Authentication.jsx";

function App() {
    return (
        <Router>
            <Routes>
                <Route exact path="/" element={<Home />} />
                <Route path="/login" element={<Authentication option="login" />} />
                <Route path="/register" element={<Authentication option="!login" />} />
            </Routes>
        </Router>
    );
}

ReactDOM.render(<App />, document.getElementById("root"));

