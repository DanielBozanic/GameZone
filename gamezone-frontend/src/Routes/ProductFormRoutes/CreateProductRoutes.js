import { Route } from "react-router-dom";
import ProtectedRoute from "../ProtectedRoute";
import ConsoleForm from "../../Pages/ProductForms/ConsoleForm/ConsoleForm";
import GraphicsCardForm from "../../Pages/ProductForms/GraphicsCardForm/GraphicsCardForm";
import HDDForm from "../../Pages/ProductForms/HDDForm/HDDForm";
import HeadphonesForm from "../../Pages/ProductForms/HeadphonesForm/HeadphonesForm";
import KeyboardForm from "../../Pages/ProductForms/KeyboardForm/KeyboardForm";
import MonitorForm from "../../Pages/ProductForms/MonitorForm/MonitorForm";
import MotherboardForm from "../../Pages/ProductForms/MotherboardForm/MotherboardForm";
import MouseForm from "../../Pages/ProductForms/MouseForm/MouseForm";
import ProcessorForm from "../../Pages/ProductForms/ProcessorForm/ProcessorForm";
import PSUForm from "../../Pages/ProductForms/PSUForm/PSUForm";
import RAMForm from "../../Pages/ProductForms/RAMForm/RAMForm";
import SSDForm from "../../Pages/ProductForms/SSDForm/SSDForm";
import VideoGameForm from "../../Pages/ProductForms/VideoGameForm/VideoGameForm";

import * as role from "../../Utils/Role";

export const CreateProductVideoGame = () => (
	<Route
		path="/addNewVideoGame"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewVideoGame"
			element={
				<VideoGameForm
					key="/addNewVideoGame"
					title={"Add video game"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductConsole = () => (
	<Route
		path="/addNewConsole"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewConsole"
			element={
				<ConsoleForm
					key="/addNewConsole"
					title={"Add console"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductGraphicsCard = () => (
	<Route
		path="/addNewGraphicsCard"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewGraphicsCard"
			element={
				<GraphicsCardForm
					key="/addNewGraphicsCard"
					title={"Add graphics card"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductProcessor = () => (
	<Route
		path="/addNewProcessor"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewProcessor"
			element={
				<ProcessorForm
					key="/addNewProcessor"
					title={"Add processor"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductMotherboard = () => (
	<Route
		path="/addNewMotherboard"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewMotherboard"
			element={
				<MotherboardForm
					key="/addNewMotherboard"
					title={"Add motherboard"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductRAM = () => (
	<Route
		path="/addNewRam"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewRam"
			element={<RAMForm key="/addNewRam" title={"Add RAM"} addButton={true} />}
		/>
	</Route>
);

export const CreateProductSSD = () => (
	<Route
		path="/addNewSsd"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewSsd"
			element={
				<SSDForm
					key="/addNewSsd"
					title={"Add solid state drive"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductHDD = () => (
	<Route
		path="/addNewHdd"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewHdd"
			element={
				<HDDForm
					key="/addNewHdd"
					title={"Add hard disk drive"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductMonitor = () => (
	<Route
		path="/addNewMonitor"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewMonitor"
			element={
				<MonitorForm
					key="/addNewMonitor"
					title={"Add monitor"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductPSU = () => (
	<Route
		path="/addNewPsu"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewPsu"
			element={
				<PSUForm
					key="/addNewPsu"
					title={"Add power supply unit"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductHeadphones = () => (
	<Route
		path="/addNewHeadphones"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewHeadphones"
			element={
				<HeadphonesForm
					key="/addNewHeadphones"
					title={"Add headphones"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductKeyboard = () => (
	<Route
		path="/addNewKeyboard"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewKeyboard"
			element={
				<KeyboardForm
					key="/addNewKeyboard"
					title={"Add keyboard"}
					addButton={true}
				/>
			}
		/>
	</Route>
);

export const CreateProductMouse = () => (
	<Route
		path="/addNewMouse"
		element={<ProtectedRoute roles={[role.ROLE_EMPLOYEE]} />}
	>
		<Route
			path="/addNewMouse"
			element={
				<MouseForm key="/addNewMouse" title={"Add mouse"} addButton={true} />
			}
		/>
	</Route>
);
