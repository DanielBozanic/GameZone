import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "./Assets/css/index.css";
import Header from "./Layout/Header";
import Footer from "./Layout/Footer";
import VideoGames from "./Pages/VideoGames";

function App() {
  return (
    <div className="full-page">
      <Header />
      <div className="app-content">
        <Router>
          <Fragment>
            <Routes>
              <Route path="/" element={<VideoGames />} />
            </Routes>
          </Fragment>
        </Router>
      </div>
      <Footer />
    </div>
  );
}

export default App;
