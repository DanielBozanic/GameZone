import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Fragment } from "react";
import ProtectedRoute from "./Auth/ProtectedRoute";
import "bootstrap/dist/css/bootstrap.min.css";
import "./Assets/css/index.css";
import "./Assets/fonts/batmfa__.ttf";
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
import ProductDetail from "./Pages/ProductDetail";
import SignUp from "./Pages/SignUp/SignUp";
import SignIn from "./Pages/SignIn/SignIn";
import ShoppingCart from "./Pages/ShoppingCart";

import ConsoleForm from "./Pages/ProductForms/ConsoleForm/ConsoleForm";
import GraphicsCardForm from "./Pages/ProductForms/GraphicsCardForm/GraphicsCardForm";
import HDDForm from "./Pages/ProductForms/HDDForm/HDDForm";
import HeadphonesForm from "./Pages/ProductForms/HeadphonesForm/HeadphonesForm";
import KeyboardForm from "./Pages/ProductForms/KeyboardForm/KeyboardForm";
import MonitorForm from "./Pages/ProductForms/MonitorForm/MonitorForm";
import MotherboardForm from "./Pages/ProductForms/MotherboardForm/MotherboardForm";
import MouseForm from "./Pages/ProductForms/MouseForm/MouseForm";
import ProcessorForm from "./Pages/ProductForms/ProcessorForm/ProcessorForm";
import PSUForm from "./Pages/ProductForms/PSUForm/PSUForm";
import RAMForm from "./Pages/ProductForms/RAMForm/RAMForm";
import SSDForm from "./Pages/ProductForms/SSDForm/SSDForm";
import VideoGameForm from "./Pages/ProductForms/VideoGameForm/VideoGameForm";

import * as role from "./Utils/Role";

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
								path="/videoGames/:id"
								element={
									<ProductDetail
										key="/videoGames/:id"
										GET_PRODUCT_BY_ID={videoGameAPI.GET_BY_ID}
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
								path="/consoles/:id"
								element={
									<ProductDetail
										key="/consoles/:id"
										GET_PRODUCT_BY_ID={consoleAPI.GET_BY_ID}
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
								path="/graphicsCards/:id"
								element={
									<ProductDetail
										key="/graphicsCards/:id"
										GET_PRODUCT_BY_ID={graphicsCardsAPI.GET_BY_ID}
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
								path="/rams/:id"
								element={
									<ProductDetail
										key="/rams/:id"
										GET_PRODUCT_BY_ID={ramAPI.GET_BY_ID}
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
								path="/motherboards/:id"
								element={
									<ProductDetail
										key="/motherboards/:id"
										GET_PRODUCT_BY_ID={motherboardAPI.GET_BY_ID}
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
								path="/processors/:id"
								element={
									<ProductDetail
										key="/processors/:id"
										GET_PRODUCT_BY_ID={processorAPI.GET_BY_ID}
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
								path="/hdds/:id"
								element={
									<ProductDetail
										key="/hdds/:id"
										GET_PRODUCT_BY_ID={hddAPI.GET_BY_ID}
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
								key="/ssds/:id"
								path="/ssds/:id"
								element={
									<ProductDetail
										key="/ssds/:id"
										GET_PRODUCT_BY_ID={ssdAPI.GET_BY_ID}
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
								path="/monitors/:id"
								element={
									<ProductDetail
										key="/monitors/:id"
										GET_PRODUCT_BY_ID={monitorAPI.GET_BY_ID}
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
								path="/psus/:id"
								element={
									<ProductDetail
										key="/psus/:id"
										GET_PRODUCT_BY_ID={psuAPI.GET_BY_ID}
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
								path="/headphones/:id"
								element={
									<ProductDetail
										key="/headphones/:id"
										GET_PRODUCT_BY_ID={headphonesAPI.GET_BY_ID}
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
								path="/keyboards/:id"
								element={
									<ProductDetail
										key="/keyboards/:id"
										GET_PRODUCT_BY_ID={keyboardAPI.GET_BY_ID}
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
							<Route
								path="/mice/:id"
								element={
									<ProductDetail
										key="/mice/:id"
										GET_PRODUCT_BY_ID={mouseAPI.GET_BY_ID}
									/>
								}
							/>
							<Route path="/signUp" element={<SignUp />} />
							<Route path="/signIn" element={<SignIn />} />
							<Route
								path="/shoppingCart"
								element={<ProtectedRoute roles={[role.ROLE_USER]} />}
							>
								<Route path="/shoppingCart" element={<ShoppingCart />} />
							</Route>
							<Route
								path="/addNewVideoGame"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewVideoGame"
									element={
										<VideoGameForm
											key="/addNewVideoGame"
											title={"Add new video game"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateVideoGame/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateVideoGame/:id"
									element={
										<VideoGameForm
											key="/updateVideoGame/:id"
											title={"Update video game"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewConsole"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewConsole"
									element={
										<ConsoleForm
											key="/addNewConsole"
											title={"Add new console"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateConsole/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateConsole/:id"
									element={
										<ConsoleForm
											key="/updateConsole/:id"
											title={"Update console"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewGraphicsCard"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewGraphicsCard"
									element={
										<GraphicsCardForm
											key="/addNewGraphicsCard"
											title={"Add new graphics card"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateGraphicsCard/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateGraphicsCard/:id"
									element={
										<GraphicsCardForm
											key="/updateGraphicsCard/:id"
											title={"Update graphics card"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewProcessor"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewProcessor"
									element={
										<ProcessorForm
											key="/addNewProcessor"
											title={"Add new processor"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateProcessor/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateProcessor/:id"
									element={
										<ProcessorForm
											key="/updateProcessor/:id"
											title={"Update processor"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewMotherboard"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewMotherboard"
									element={
										<MotherboardForm
											key="/addNewMotherboard"
											title={"Add new motherboard"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateMotherboard/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateMotherboard/:id"
									element={
										<MotherboardForm
											key="/updateMotherboard/:id"
											title={"Update motherboard"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewRam"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewRam"
									element={
										<RAMForm
											key="/addNewRam"
											title={"Add new RAM"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateRam/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateRam/:id"
									element={
										<RAMForm
											key="/updateRam/:id"
											title={"Update RAM"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewSsd"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewSsd"
									element={
										<SSDForm
											key="/addNewSsd"
											title={"Add new solid state drive"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateSsd/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateSsd/:id"
									element={
										<SSDForm
											key="/updateSsd/:id"
											title={"Update solid state drive"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewHdd"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewHdd"
									element={
										<HDDForm
											key="/addNewHdd"
											title={"Add new hard disk drive"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateHdd/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateHdd/:id"
									element={
										<HDDForm
											key="/updateHdd/:id"
											title={"Update hard disk drive"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewMonitor"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewMonitor"
									element={
										<MonitorForm
											key="/addNewMonitor"
											title={"Add new monitor"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateMonitor/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateMonitor/:id"
									element={
										<MonitorForm
											key="/updateMonitor/:id"
											title={"Update monitor"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewPsu"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewPsu"
									element={
										<PSUForm
											key="/addNewPsu"
											title={"Add new power supply unit"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updatePsu/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updatePsu/:id"
									element={
										<PSUForm
											key="/updatePsu/:id"
											title={"Update power supply unit"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewHeadphones"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewHeadphones"
									element={
										<HeadphonesForm
											key="/addNewHeadphones"
											title={"Add new headphones"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateHeadphones/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateHeadphones/:id"
									element={
										<HeadphonesForm
											key="/updateHeadphones/:id"
											title={"Update headphones"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewKeyboard"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewKeyboard"
									element={
										<KeyboardForm
											key="/addNewKeyboard"
											title={"Add new keyboard"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateKeyboard/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateKeyboard/:id"
									element={
										<KeyboardForm
											key="/updateKeyboard/:id"
											title={"Update keyboard"}
											updateButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/addNewMouse"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/addNewMouse"
									element={
										<MouseForm
											key="/addNewMouse"
											title={"Add new mouse"}
											addButton={true}
										/>
									}
								/>
							</Route>
							<Route
								path="/updateMouse/:id"
								element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
							>
								<Route
									path="/updateMouse/:id"
									element={
										<MouseForm
											key="/updateMouse/:id"
											title={"Update mouse"}
											updateButton={true}
										/>
									}
								/>
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
