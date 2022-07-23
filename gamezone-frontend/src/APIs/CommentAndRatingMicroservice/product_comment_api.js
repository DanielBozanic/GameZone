const api = "http://localhost:7001/api/comments/productComments";

const userProtected = api + "/userProtected";
const authProtected = api + "/authProtected";

export const GET_ALL = api;
export const GET_BY_ID = api;
export const GET_BY_PRODUCT_NAME = api + "/getByProductName";
export const GET_BY_USERNAME = api + "/getByUsername";
export const GET_BY_PRODUCT_NAME_AND_USERNAME =
	api + "/getByProductNameAndUsername";

export const ADD_COMMENT = userProtected + "/addComment";
export const EDIT_COMMENT = userProtected + "/editComment";

export const DELETE_COMMENT = authProtected + "/deleteComment";
