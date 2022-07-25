const api = "http://localhost:7002/api/news/newsComments";

const userProtected = api + "/userProtected";
const userAndAdminProtected = api + "/userAndAdminProtected";

export const GET_BY_NEWS_ARTICLE = api + "/getByNewsArticle";

export const ADD_NEWS_COMMENT = userProtected + "/addNewsComment";
export const EDIT_NEWS_COMMENT = userProtected + "/editNewsComment";

export const DELETE_NEWS_COMMENT = userAndAdminProtected + "/deleteNewsComment";
