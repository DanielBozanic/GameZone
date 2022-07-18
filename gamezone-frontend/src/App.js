import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import ProtectedRoute from "./Routes/ProtectedRoute";
import "bootstrap/dist/css/bootstrap.min.css";
import "./Assets/css/index.css";
import "./Assets/fonts/batmfa__.ttf";
import Header from "./Layout/Header";
import Footer from "./Layout/Footer";

import {
	ProductListVideoGames,
	ProductListConsoles,
	ProductListGraphicsCards,
	ProductListHDDS,
	ProductListHeadphones,
	ProductListKeyboards,
	ProductListMice,
	ProductListMonitors,
	ProductListMotherboards,
	ProductListPSUS,
	ProductListProcessors,
	ProductListRAMS,
	ProductListSSDS,
} from "./Routes/ProductViewRoutes/ProductListRoutes";
import {
	ProductDetailConsole,
	ProductDetailGraphicsCard,
	ProductDetailHDD,
	ProductDetailHeadphones,
	ProductDetailKeyboard,
	ProductDetailMonitor,
	ProductDetailMotherboard,
	ProductDetailMouse,
	ProductDetailPSU,
	ProductDetailProcessor,
	ProductDetailRAM,
	ProductDetailSSD,
	ProductDetailVideoGame,
} from "./Routes/ProductViewRoutes/ProductDetailRoutes";
import {
	CreateProductConsole,
	CreateProductGraphicsCard,
	CreateProductHDD,
	CreateProductHeadphones,
	CreateProductKeyboard,
	CreateProductMonitor,
	CreateProductMotherboard,
	CreateProductMouse,
	CreateProductPSU,
	CreateProductProcessor,
	CreateProductRAM,
	CreateProductSSD,
	CreateProductVideoGame,
} from "./Routes/ProductFormRoutes/CreateProductRoutes";
import {
	UpdateProductConsole,
	UpdateProductGraphicsCard,
	UpdateProductHDD,
	UpdateProductHeadphones,
	UpdateProductKeyboard,
	UpdateProductMonitor,
	UpdateProductMotherboard,
	UpdateProductMouse,
	UpdateProductPSU,
	UpdateProductProcessor,
	UpdateProductRAM,
	UpdateProductSSD,
	UpdateProductVideoGame,
} from "./Routes/ProductFormRoutes/UpdateProductRoutes";

import Main from "./Pages/Main";
import SignUp from "./Pages/SignUp/SignUp";
import SignIn from "./Pages/SignIn/SignIn";
import ShoppingCart from "./Pages/ShoppingCart";

import * as role from "./Utils/Role";

function App() {
	return (
		<div className="full-page">
			<Router>
				<Fragment>
					<Header />
					<div className="app-content">
						<Routes>
							<Route index path="/" element={<Main />} />
							<Route path="/signUp" element={<SignUp />} />
							<Route path="/signIn" element={<SignIn />} />
							<Route
								path="/shoppingCart"
								element={<ProtectedRoute roles={[role.ROLE_USER]} />}
							>
								<Route path="/shoppingCart" element={<ShoppingCart />} />
							</Route>
							{ProductListVideoGames()}
							{ProductListConsoles()}
							{ProductListGraphicsCards()}
							{ProductListHDDS()}
							{ProductListHeadphones()}
							{ProductListKeyboards()}
							{ProductListMice()}
							{ProductListMonitors()}
							{ProductListMotherboards()}
							{ProductListProcessors()}
							{ProductListPSUS()}
							{ProductListRAMS()}
							{ProductListSSDS()}

							{ProductDetailConsole()}
							{ProductDetailGraphicsCard()}
							{ProductDetailHDD()}
							{ProductDetailHeadphones()}
							{ProductDetailKeyboard()}
							{ProductDetailMonitor()}
							{ProductDetailMotherboard()}
							{ProductDetailMouse()}
							{ProductDetailPSU()}
							{ProductDetailProcessor()}
							{ProductDetailRAM()}
							{ProductDetailSSD()}
							{ProductDetailVideoGame()}

							{CreateProductConsole()}
							{CreateProductGraphicsCard()}
							{CreateProductHDD()}
							{CreateProductHeadphones()}
							{CreateProductKeyboard()}
							{CreateProductMonitor()}
							{CreateProductMotherboard()}
							{CreateProductMouse()}
							{CreateProductPSU()}
							{CreateProductProcessor()}
							{CreateProductRAM()}
							{CreateProductSSD()}
							{CreateProductVideoGame()}

							{UpdateProductConsole()}
							{UpdateProductGraphicsCard()}
							{UpdateProductHDD()}
							{UpdateProductHeadphones()}
							{UpdateProductKeyboard()}
							{UpdateProductMonitor()}
							{UpdateProductMotherboard()}
							{UpdateProductMouse()}
							{UpdateProductPSU()}
							{UpdateProductProcessor()}
							{UpdateProductRAM()}
							{UpdateProductSSD()}
							{UpdateProductVideoGame()}
						</Routes>
					</div>
					<Footer />
				</Fragment>
			</Router>
		</div>
	);
}

export default App;
