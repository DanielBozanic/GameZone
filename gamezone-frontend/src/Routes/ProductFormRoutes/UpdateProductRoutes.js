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

export const UpdateProductVideoGame = () => (
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
);

export const UpdateProductConsole = () => (
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
);

export const UpdateProductGraphicsCard = () => (
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
);

export const UpdateProductProcessor = () => (
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
);

export const UpdateProductMotherboard = () => (
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
);

export const UpdateProductRAM = () => (
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
);

export const UpdateProductSSD = () => (
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
);

export const UpdateProductHDD = () => (
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
);

export const UpdateProductMonitor = () => (
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
);

export const UpdateProductPSU = () => (
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
);

export const UpdateProductHeadphones = () => (
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
);

export const UpdateProductKeyboard = () => (
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
);

export const UpdateProductMouse = () => (
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
);
