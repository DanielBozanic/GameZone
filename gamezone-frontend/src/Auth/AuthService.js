export const getAccessToken = () => {
  return localStorage.getItem("accessToken");
};

export const getEmail = () => {
  return localStorage.getItem("email");
};

export const getRole = () => {
  return localStorage.getItem("role");
};

export const getJMBG = () => {
  return localStorage.getItem("id");
};

export const removeToken = () => {
  localStorage.removeItem("accessToken");
  localStorage.removeItem("email");
  localStorage.removeItem("role");
  localStorage.removeItem("id");
};

export const storeToken = (token) => {
  localStorage.setItem("accessToken", token.accessToken);
  localStorage.setItem("email", token.email);
  localStorage.setItem("role", token.role);
  localStorage.setItem("id", token.id);
};
