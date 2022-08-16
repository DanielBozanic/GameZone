const api = "http://localhost:8000/api/comments/productComments";

export const GET_ALL = api;
export const GET_BY_ID = api;
export const GET_BY_PRODUCT_ID = api + "/getByProductId";

export const ADD_COMMENT = api + "/addComment";
export const EDIT_COMMENT = api + "/editComment";

export const GET_BY_USER_ID = api + "/getByUserId";
export const DELETE_COMMENT = api + "/deleteComment";
export const DELETE_COMMENTS_BY_PRODUCT_ID = api + "/deleteCommentsByProductId";
