export const getToken = () => {
	return sessionStorage.getItem("token");
};

export const getEmail = () => {
	return sessionStorage.getItem("email");
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
	sessionStorage.removeItem("role");
	sessionStorage.removeItem("id");
};

export const storeToken = (data) => {
	sessionStorage.setItem("token", data.token);
	sessionStorage.setItem("email", data.user.email);
	sessionStorage.setItem("role", data.user.role);
	sessionStorage.setItem("id", data.user.id);
};
