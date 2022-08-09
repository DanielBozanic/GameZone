const api = "http://localhost:8000/api/products";

export const GET_PRODUCT_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";

export const DELETE_PRODUCT = api + "/deleteProduct";
export const ADD_PRODUCT_TO_MAIN_PAGE = api + "/addProductToMainPage";
export const REMOVE_PRODUCT_FROM_MAIN_PAGE = api + "/removeProductFromMainPage";
export const GET_MAIN_PAGE_PRODUCTS = api + "/getMainPageProducts";
export const GET_POPULAR_PRODUCTS = api + "/getPopularProducts";
export const GET_RECOMMENDED_PRODUCTS = api + "/getRecommendedProducts";
