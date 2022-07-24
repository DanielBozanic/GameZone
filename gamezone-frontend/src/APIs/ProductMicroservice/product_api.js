const api = "http://localhost:7000/api/products";

const employeeProtectedProducts = "/employeeProtected";

export const GET_PRODUCT_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";

export const DELETE_PRODUCT =
	api + employeeProtectedProducts + "/deleteProduct";
