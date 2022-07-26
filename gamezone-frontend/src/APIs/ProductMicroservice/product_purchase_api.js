const api = "http://localhost:7000/api/products/productPurchases";

const adminAndUserProtectedProductPurchases = "/adminAndUserProtected";
const userProtectedProductPurchases = "/userProtected";
const employeeProtectedProductPurchases = "/employeeProtected";

export const CONFIRM_PURCHASE =
	api + userProtectedProductPurchases + "/confirmPurchase";
export const CHECK_IF_PRODUCT_IS_PAID_FOR =
	api + userProtectedProductPurchases + "/checkIfProductIsPaidFor";
export const GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_USER_ID =
	api + userProtectedProductPurchases + "/getProductAlertByProductIdAndUserId";
export const ADD_PRODUCT_ALERT =
	api + userProtectedProductPurchases + "/addProductAlert";
export const GET_PURCHASE_HISTORY =
	api + adminAndUserProtectedProductPurchases + "/getPurchaseHistory";
export const NOTIFY_PRODUCT_AVAILABILITY =
	api + employeeProtectedProductPurchases + "/notifyProductAvailability";
