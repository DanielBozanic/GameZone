const api = "http://localhost:7000/api/products/productPurchases";

const adminAndUserProtectedProductPurchases = "/adminAndUserProtected";
const userProtectedProductPurchases = "/userProtected";
const employeeProtectedProductPurchases = "/employeeProtected";
const adminProtectedProductPurchases = "/adminProtected";

export const CONFIRM_PURCHASE =
	api + userProtectedProductPurchases + "/confirmPurchase";
export const SEND_PURCHASE_CONFIRMATION_MAIL =
	api + userProtectedProductPurchases + "/sendPurchaseConfirmationMail";
export const CHECK_IF_PRODUCT_IS_PAID_FOR =
	api + userProtectedProductPurchases + "/checkIfProductIsPaidFor";
export const GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_USER_ID =
	api + userProtectedProductPurchases + "/getProductAlertByProductIdAndUserId";
export const ADD_PRODUCT_ALERT =
	api + userProtectedProductPurchases + "/addProductAlert";
export const GET_PURCHASE_HISTORY =
	api + adminAndUserProtectedProductPurchases + "/getPurchaseHistory";
export const GET_NUMBER_OF_RECORDS_PURCHASE_HISTORY =
	api +
	adminAndUserProtectedProductPurchases +
	"/getNumberOfRecordsPurchaseHistory";
export const NOTIFY_PRODUCT_AVAILABILITY =
	api + employeeProtectedProductPurchases + "/notifyProductAvailability";
export const CONFIRM_PAYMENT =
	api + adminProtectedProductPurchases + "/confirmPayment";
export const SEND_PURCHASED_DIGITAL_VIDEO_GAMES =
	api + adminProtectedProductPurchases + "/sendPurchasedDigitalVideoGames";
