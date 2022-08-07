const api = "http://localhost:7000/api/products";

const employeeProtectedProducts = "/employeeProtected";

export const GET_PRODUCT_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";

export const DELETE_PRODUCT =
	api + employeeProtectedProducts + "/deleteProduct";
export const ADD_PRODUCT_TO_MAIN_PAGE =
	api + employeeProtectedProducts + "/addProductToMainPage";
export const REMOVE_PRODUCT_FROM_MAIN_PAGE =
	api + employeeProtectedProducts + "/removeProductFromMainPage";
export const GET_MAIN_PAGE_PRODUCTS = api + "/getMainPageProducts";
