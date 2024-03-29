import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import ProtectedRoute from "./Routes/ProtectedRoute";
import "bootstrap/dist/css/bootstrap.min.css";
import "react-toastify/dist/ReactToastify.css";
import "./Assets/css/index.css";
import "./Assets/fonts/batmfa__.ttf";
import "swiper/css/bundle";
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import "./Assets/css/swiper.css";
import Header from "./Layout/Header";

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
import SignUp from "./Pages/SignUp";
import GetVerificationCode from "./Pages/AccountVerification/GetVerificationCode";
import VerifyAccount from "./Pages/AccountVerification/VerifyAccount";
import SignIn from "./Pages/SignIn/SignIn";
import CreateAdminAndEmployee from "./Pages/CreateAdminAndEmployee";
import ShoppingCart from "./Pages/ShoppingCart";
import NewsEditor from "./Pages/NewsEditor/NewsEditor";
import NewsList from "./Pages/NewsList";
import NewsDetail from "./Pages/NewsDetail";
import RegisteredUsersList from "./Pages/RegisteredUsersList";
import PurchaseHistory from "./Pages/PurchaseHistory";
import ManageUser from "./Pages/ManageUser/ManageUser";
import ContactForm from "./Pages/ContactForm/ContactForm";
import ContactMessages from "./Pages/ContactMessages/ContactMessages";
import BusinessReports from "./Pages/BusinessReports";
import Profile from "./Pages/Profile/Profile";

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
							<Route
								path="/getVerificationCode"
								element={<GetVerificationCode />}
							/>
							<Route path="/verify/:email" element={<VerifyAccount />} />
							<Route path="/signIn" element={<SignIn />} />
							<Route
								path="/createAdminAndEmployee"
								element={<ProtectedRoute roles={[role.ROLE_ADMIN]} />}
							>
								<Route
									path="/createAdminAndEmployee"
									element={<CreateAdminAndEmployee />}
								/>
							</Route>
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

							<Route
								path="/addNewsArticle"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewsArticle"
									element={<NewsEditor key="/addNewsArticle" />}
								/>
							</Route>
							<Route
								path="/editNewsArticle/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/editNewsArticle/:id"
									element={<NewsEditor key="/editNewsArticle/:id" />}
								/>
							</Route>
							<Route path="/viewNews" element={<NewsList />} />
							<Route path="/viewNews/:id" element={<NewsDetail />} />
							<Route
								path="/registeredUsers"
								element={<ProtectedRoute roles={[role.ROLE_ADMIN]} />}
							>
								<Route
									path="/registeredUsers"
									element={<RegisteredUsersList key="/registeredUsers" />}
								/>
							</Route>
							<Route
								path="/purchaseHistory/:id"
								element={
									<ProtectedRoute roles={[role.ROLE_ADMIN, role.ROLE_USER]} />
								}
							>
								<Route
									path="/purchaseHistory/:id"
									element={<PurchaseHistory key="/purchaseHistory/:id" />}
								/>
							</Route>
							<Route
								path="/manageUser/:id"
								element={<ProtectedRoute roles={[role.ROLE_ADMIN]} />}
							>
								<Route
									path="/manageUser/:id"
									element={<ManageUser key="/manageUser/:id" />}
								/>
							</Route>
							<Route
								path="/contact"
								element={<ProtectedRoute roles={[role.ROLE_USER]} />}
							>
								<Route
									path="/contact"
									element={<ContactForm key="/contact" />}
								/>
							</Route>
							<Route
								path="/contactMessages/:id"
								element={
									<ProtectedRoute roles={[role.ROLE_USER, role.ROLE_ADMIN]} />
								}
							>
								<Route
									path="/contactMessages/:id"
									element={<ContactMessages key="/contactMessages/:id" />}
								/>
							</Route>
							<Route
								path="/contactMessages"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/contactMessages"
									element={<ContactMessages key="/contactMessages" />}
								/>
							</Route>
							<Route
								path="/businessReports"
								element={<ProtectedRoute roles={[role.ROLE_ADMIN]} />}
							>
								<Route
									path="/businessReports"
									element={<BusinessReports key="/businessReports" />}
								/>
							</Route>
							<Route
								path="/profile"
								element={
									<ProtectedRoute
										roles={[
											role.ROLE_USER,
											role.ROLE_EMPLOYEE,
											role.ROLE_ADMIN,
										]}
									/>
								}
							>
								<Route path="/profile" element={<Profile key="/profile" />} />
							</Route>
						</Routes>
					</div>
				</Fragment>
			</Router>
		</div>
	);
}

export default App;
