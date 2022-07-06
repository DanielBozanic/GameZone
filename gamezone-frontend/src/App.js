import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "./Assets/css/index.css";
import Header from "./Layout/Header";
import Footer from "./Layout/Footer";
import * as videoGameAPI from "./APIs/ProductMicroservice/video_game_api";
import Main from "./Pages/Main";
import ProductList from "./Pages/ProductList";

function App() {
	return (
		<div className="full-page">
			<Router>
				<Header />
				<div className="app-content">
					<Fragment>
						<Routes>
							<Route path="/" element={<Main />} />
							<Route
								path="/videoGames"
								element={
									<ProductList
										GET_PRODUCTS={videoGameAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={videoGameAPI.GET_NUMBER_OF_RECORDS}
										FILTER={videoGameAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											videoGameAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
									/>
								}
							/>
						</Routes>
					</Fragment>
				</div>
				<Footer />
			</Router>
		</div>
	);
}

export default App;
