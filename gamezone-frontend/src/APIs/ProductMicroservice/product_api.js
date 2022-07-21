const api = "http://localhost:7000/api/products";

const adminAndUserProtectedProducts = "/adminAndUserProtectedProducts";
const userProtectedProducts = "/userProtectedProducts";
const employeeProtectedProducts = "/employeeProtectedProducts";

export const ADD_PRODUCT_TO_CART =
	api + userProtectedProducts + "/addProductToCart";
export const GET_CURRENT_CART = api + userProtectedProducts + "/getCurrentCart";
export const CART_CONTAINS_ONLY_DIGITAL_ITEMS =
	api + userProtectedProducts + "/cartContainsOnlyDigitalItems";
export const UPDATE_PURCHASE = api + userProtectedProducts + "/updatePurchase";
export const REMOVE_PRODUCT_FROM_CART =
	api + userProtectedProducts + "/removeProductFromCart";
export const CONFIRM_PURCHASE =
	api + userProtectedProducts + "/confirmPurchase";
export const GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_EMAIL =
	api + userProtectedProducts + "/getProductAlertByProductIdAndEmail";
export const ADD_PRODUCT_ALERT =
	api + userProtectedProducts + "/addProductAlert";
export const GET_PURCHASE_HISTORY =
	api + +adminAndUserProtectedProducts + "/getPurchaseHistory";

export const GET_PRODUCT_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";

export const DELETE_PRODUCT =
	api + employeeProtectedProducts + "/deleteProduct";
export const NOTIFY_PRODUCT_AVAILABILITY =
	api + employeeProtectedProducts + "/notifyProductAvailability";
