import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import ProtectedRoute from "./Auth/ProtectedRoute";
import "bootstrap/dist/css/bootstrap.min.css";
import "./Assets/css/index.css";
import Header from "./Layout/Header";
import Footer from "./Layout/Footer";

import * as videoGameAPI from "./APIs/ProductMicroservice/video_game_api";
import * as consoleAPI from "./APIs/ProductMicroservice/console_api";
import * as hddAPI from "./APIs/ProductMicroservice/hard_disk_drive_api";
import * as headphonesAPI from "./APIs/ProductMicroservice/headphones_api";
import * as keyboardAPI from "./APIs/ProductMicroservice/keyboard_api";
import * as monitorAPI from "./APIs/ProductMicroservice/monitor_api";
import * as motherboardAPI from "./APIs/ProductMicroservice/motherboard_api";
import * as mouseAPI from "./APIs/ProductMicroservice/mouse_api";
import * as processorAPI from "./APIs/ProductMicroservice/processor_api";
import * as psuAPI from "./APIs/ProductMicroservice/psu_api";
import * as ramAPI from "./APIs/ProductMicroservice/ram_api";
import * as ssdAPI from "./APIs/ProductMicroservice/solid_state_drive_api";
import * as graphicsCardsAPI from "./APIs/ProductMicroservice/graphics_card_api";

import VideoGameFilter from "./Components/Filter/VideoGamesFilter";
import ConsoleFilter from "./Components/Filter/ConsoleFilter";
import GraphicsCardFilter from "./Components/Filter/GraphicsCardFilter";
import RAMFilter from "./Components/Filter/RAMFilter";
import MotherboardFilter from "./Components/Filter/MotherboardFilter";
import ProcessorFilter from "./Components/Filter/ProcessorFilter";
import HDDFilter from "./Components/Filter/HDDFilter";
import SSDFilter from "./Components/Filter/SSDFilter";
import MonitorFilter from "./Components/Filter/MonitorFilter";
import PSUFilter from "./Components/Filter/PSUFilter";
import HeadphonesFilter from "./Components/Filter/HeadphonesFilter";
import KeyboardFilter from "./Components/Filter/KeyboardFilter";
import MouseFilter from "./Components/Filter/MouseFilter";

import Main from "./Pages/Main";
import ProductList from "./Pages/ProductList";
import SignUp from "./Pages/SignUp/SignUp";
import SignIn from "./Pages/SignIn/SignIn";
import ShoppingCart from "./Pages/ShoppingCart/ShoppingCart";

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
										key="/videoGames"
										GET_PRODUCTS={videoGameAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={videoGameAPI.GET_NUMBER_OF_RECORDS}
										FILTER={videoGameAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											videoGameAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={videoGameAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											videoGameAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={VideoGameFilter}
									/>
								}
							/>
							<Route
								path="/consoles"
								element={
									<ProductList
										key="/consoles"
										GET_PRODUCTS={consoleAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={consoleAPI.GET_NUMBER_OF_RECORDS}
										FILTER={consoleAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											consoleAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={consoleAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											consoleAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={ConsoleFilter}
									/>
								}
							/>
							<Route
								path="/graphicsCards"
								element={
									<ProductList
										key="/graphicsCards"
										GET_PRODUCTS={graphicsCardsAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={
											graphicsCardsAPI.GET_NUMBER_OF_RECORDS
										}
										FILTER={graphicsCardsAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											graphicsCardsAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={graphicsCardsAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											graphicsCardsAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={GraphicsCardFilter}
									/>
								}
							/>
							<Route
								path="/rams"
								element={
									<ProductList
										key="/rams"
										GET_PRODUCTS={ramAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={ramAPI.GET_NUMBER_OF_RECORDS}
										FILTER={ramAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											ramAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={ramAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											ramAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={RAMFilter}
									/>
								}
							/>
							<Route
								path="/motherboards"
								element={
									<ProductList
										key="/motherboards"
										GET_PRODUCTS={motherboardAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={motherboardAPI.GET_NUMBER_OF_RECORDS}
										FILTER={motherboardAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											motherboardAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={motherboardAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											motherboardAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={MotherboardFilter}
									/>
								}
							/>
							<Route
								path="/processors"
								element={
									<ProductList
										key="/processors"
										GET_PRODUCTS={processorAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={processorAPI.GET_NUMBER_OF_RECORDS}
										FILTER={processorAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											processorAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={processorAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											processorAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={ProcessorFilter}
									/>
								}
							/>
							<Route
								path="/hdds"
								element={
									<ProductList
										key="/hdds"
										GET_PRODUCTS={hddAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={hddAPI.GET_NUMBER_OF_RECORDS}
										FILTER={hddAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											hddAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={hddAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											hddAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={HDDFilter}
									/>
								}
							/>
							<Route
								path="/ssds"
								element={
									<ProductList
										key="/ssds"
										GET_PRODUCTS={ssdAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={ssdAPI.GET_NUMBER_OF_RECORDS}
										FILTER={ssdAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											ssdAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={ssdAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											ssdAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={SSDFilter}
									/>
								}
							/>
							<Route
								path="/monitors"
								element={
									<ProductList
										key="/monitors"
										GET_PRODUCTS={monitorAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={monitorAPI.GET_NUMBER_OF_RECORDS}
										FILTER={monitorAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											monitorAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={monitorAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											monitorAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={MonitorFilter}
									/>
								}
							/>
							<Route
								path="/psus"
								element={
									<ProductList
										key="/psus"
										GET_PRODUCTS={psuAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={psuAPI.GET_NUMBER_OF_RECORDS}
										FILTER={psuAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											psuAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={psuAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											psuAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={PSUFilter}
									/>
								}
							/>
							<Route
								path="/headphones"
								element={
									<ProductList
										key="/headphones"
										GET_PRODUCTS={headphonesAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={headphonesAPI.GET_NUMBER_OF_RECORDS}
										FILTER={headphonesAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											headphonesAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={headphonesAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											headphonesAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={HeadphonesFilter}
									/>
								}
							/>
							<Route
								path="/keyboards"
								element={
									<ProductList
										key="/keyboards"
										GET_PRODUCTS={keyboardAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={keyboardAPI.GET_NUMBER_OF_RECORDS}
										FILTER={keyboardAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											keyboardAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={keyboardAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											keyboardAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={KeyboardFilter}
									/>
								}
							/>
							<Route
								path="/mice"
								element={
									<ProductList
										key="/mice"
										GET_PRODUCTS={mouseAPI.GET_ALL}
										GET_NUMBER_OF_RECORDS={mouseAPI.GET_NUMBER_OF_RECORDS}
										FILTER={mouseAPI.FILTER}
										GET_NUMBER_OF_RECORDS_FILTER={
											mouseAPI.GET_NUMBER_OF_RECORDS_FILTER
										}
										SEARCH_BY_NAME={mouseAPI.SEARCH_BY_NAME}
										GET_NUMBER_OF_RECORDS_SEARCH={
											mouseAPI.GET_NUMBER_OF_RECORDS_SEARCH
										}
										filter={MouseFilter}
									/>
								}
							/>
							<Route path="/signUp" element={<SignUp />} />
							<Route path="/signIn" element={<SignIn />} />
							<Route
								path="/shoppingCart"
								element={<ProtectedRoute roles={["ROLE_USER"]} />}
							>
								<Route path="/shoppingCart" element={<ShoppingCart />} />
							</Route>
						</Routes>
					</Fragment>
				</div>
				<Footer />
			</Router>
		</div>
	);
}

export default App;
