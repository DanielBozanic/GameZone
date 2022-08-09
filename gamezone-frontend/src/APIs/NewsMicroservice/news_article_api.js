const api = "http://localhost:8000/api/news/newsArticles";

export const GET_PUBLISHED_ARTICLES = api + "/getPublishedArticles";
export const GET_NUMBER_OF_RECORDS_PUBLISHED_ARTICLES =
	api + "/getNumberOfRecordsPublishedArticles";
export const GET_BY_ID = api;

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const ADD_NEWS_ARTICLE = api + "/addNewsArticle";
export const EDIT_NEWS_ARTICLE = api + "/editNewsArticle";
export const DELETE_NEWS_ARTICLE = api + "/deleteNewsArticle";
export const PUBLISH_NEWS_ARTICLE = api + "/publishNewsArticle";
