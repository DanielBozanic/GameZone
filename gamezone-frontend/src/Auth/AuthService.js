import * as role from "../Utils/Role";

export const getToken = () => {
	return sessionStorage.getItem("token");
};

export const getEmail = () => {
	return sessionStorage.getItem("email");
};

export const getUsername = () => {
	return sessionStorage.getItem("username");
};

export const getRole = () => {
	return sessionStorage.getItem("role");
};

export const getId = () => {
	return sessionStorage.getItem("id");
};

export const removeToken = () => {
	sessionStorage.removeItem("token");
	sessionStorage.removeItem("email");
	sessionStorage.removeItem("username");
	sessionStorage.removeItem("role");
	sessionStorage.removeItem("id");
};

export const storeToken = (data) => {
	sessionStorage.setItem("token", data.token);
	sessionStorage.setItem("email", data.user.email);
	sessionStorage.setItem("username", data.user.user_name);
	sessionStorage.setItem("role", data.user.role);
	sessionStorage.setItem("id", data.user.id);
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
