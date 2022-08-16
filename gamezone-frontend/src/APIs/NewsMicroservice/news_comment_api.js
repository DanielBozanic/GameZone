const api = "http://localhost:8000/api/news/newsComments";

export const GET_BY_NEWS_ARTICLE = api + "/getByNewsArticle";

export const ADD_NEWS_COMMENT = api + "/addNewsComment";
export const EDIT_NEWS_COMMENT = api + "/editNewsComment";

export const DELETE_NEWS_COMMENT = api + "/deleteNewsComment";
export const DELETE_NEWS_COMMENTS_BY_NEWS_ARTICLE_ID =
	api + "/deleteNewsCommentsByNewsArticleId";
export const GET_BY_USER_ID = api + "/getByUserId";
