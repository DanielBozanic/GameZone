const api = "http://localhost:7002/api/news/newsArticles";

const employeeProtected = api + "/employeeProtected";

export const GET_PUBLISHED_ARTICLES = api + "/getPublishedArticles";
export const GET_NUMBER_OF_RECORDS_PUBLISHED_ARTICLES =
	api + "/getNumberOfRecordsPublishedArticles";
export const GET_BY_ID = api;

export const GET_UNPUBLISHED_ARTICLES =
	employeeProtected + "/getUnpublishedArticles";
export const GET_NUMBER_OF_RECORDS_UNPUBLISHED_ARTICLES =
	employeeProtected + "/getNumberOfRecordsUnpublishedArticles";
export const ADD_NEWS_ARTICLE = employeeProtected + "/addNewsArticle";
export const EDIT_NEWS_ARTICLE = employeeProtected + "/editNewsArticle";
export const DELETE_NEWS_ARTICLE = employeeProtected + "/deleteNewsArticle";
export const PUBLISH_NEWS_ARTICLE = employeeProtected + "/publishNewsArticle";
