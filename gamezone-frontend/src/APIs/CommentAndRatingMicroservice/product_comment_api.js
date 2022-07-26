const api = "http://localhost:7001/api/comments/productComments";

const userProtected = api + "/userProtected";
const userAndAdminProtected = api + "/userAndAdminProtected";

export const GET_ALL = api;
export const GET_BY_ID = api;
export const GET_BY_PRODUCT_ID = api + "/getByProductId";
export const GET_BY_USER_ID = api + "/getByUserId";
export const GET_BY_PRODUCT_AND_USER = api + "/getByProductAndUser";

export const ADD_COMMENT = userProtected + "/addComment";
export const EDIT_COMMENT = userProtected + "/editComment";

export const DELETE_COMMENT = userAndAdminProtected + "/deleteComment";
