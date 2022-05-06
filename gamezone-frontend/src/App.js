import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import Header from "./Layout/Header";
import Footer from "./Layout/Footer";
import VideoGames from "./Pages/VideoGames";

function App() {
  return (
    <div className="App">
      <Header />
      <Router>
        <Fragment>
          <Routes>
            <Route path="/" element={<VideoGames />} />
          </Routes>
        </Fragment>
      </Router>
      <Footer />
    </div>
  );
}

export default App;
