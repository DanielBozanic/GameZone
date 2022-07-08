import { Navigate, Outlet } from "react-router-dom";
import * as authService from "./AuthService";

const ProtectedRoute = ({ roles }) => {
	let currentUserRole = authService.getRole();
	let token = authService.getToken();

	return roles.includes(currentUserRole) && token ? (
		<Outlet />
	) : (
		<Navigate to="/" />
	);
};

export default ProtectedRoute;
