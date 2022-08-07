import * as role from "../Utils/Role";

export const getToken = () => {
	return localStorage.getItem("token");
};

export const getEmail = () => {
	return localStorage.getItem("email");
};

export const getUsername = () => {
	return localStorage.getItem("username");
};

export const getRole = () => {
	return localStorage.getItem("role");
};

export const getId = () => {
	return localStorage.getItem("id");
};

export const removeToken = () => {
	localStorage.removeItem("token");
	localStorage.removeItem("email");
	localStorage.removeItem("username");
	localStorage.removeItem("role");
	localStorage.removeItem("id");
};

export const storeToken = (data) => {
	localStorage.setItem("token", data.token);
	localStorage.setItem("email", data.user.email);
	localStorage.setItem("username", data.user.user_name);
	localStorage.setItem("role", data.user.role);
	localStorage.setItem("id", data.user.id);
};

export const isUser = () => {
	return getToken() != null && getRole() === role.ROLE_USER;
};

export const isEmployee = () => {
	return getToken() != null && getRole() === role.ROLE_EMPLOYEE;
};

export const isAdmin = () => {
	return getToken() != null && getRole() === role.ROLE_ADMIN;
};
