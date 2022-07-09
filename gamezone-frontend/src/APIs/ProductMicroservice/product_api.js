const api = "http://localhost:7000/api/products";

const adminAndUserProtectedProducts = "/adminAndUserProtectedProducts";
const userProtectedProducts = "/userProtectedProducts";

export const ADD_PRODUCT_TO_CART =
	api + userProtectedProducts + "/addProductToCart";
export const GET_CURRENT_CART = api + userProtectedProducts + "/getCurrentCart";
export const GET_PURCHASE_HISTORY =
	api + +adminAndUserProtectedProducts + "/getPurchaseHistory";
export const UPDATE_PURCHASE = api + userProtectedProducts + "/updatePurchase";
export const REMOVE_PRODUCT_FROM_CART =
	api + userProtectedProducts + "/removeProductFromCart";
export const CONFIRM_PURCHASE =
	api + userProtectedProducts + "/confirmPurchase";

export const GET_PRODUCT_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
