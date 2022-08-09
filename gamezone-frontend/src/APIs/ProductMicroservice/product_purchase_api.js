const api = "http://localhost:8000/api/products/productPurchases";

export const CONFIRM_PURCHASE = api + "/confirmPurchase";
export const SEND_PURCHASE_CONFIRMATION_MAIL =
	api + "/sendPurchaseConfirmationMail";
export const CHECK_IF_PRODUCT_IS_PAID_FOR = api + "/checkIfProductIsPaidFor";
export const GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_USER_ID =
	api + "/getProductAlertByProductIdAndUserId";
export const ADD_PRODUCT_ALERT = api + "/addProductAlert";
export const GET_PURCHASE_HISTORY = api + "/getPurchaseHistory";
export const GET_NUMBER_OF_RECORDS_PURCHASE_HISTORY =
	api + "/getNumberOfRecordsPurchaseHistory";
export const NOTIFY_PRODUCT_AVAILABILITY = api + "/notifyProductAvailability";
export const CONFIRM_PAYMENT = api + "/confirmPayment";
export const SEND_PURCHASED_DIGITAL_VIDEO_GAMES =
	api + "/sendPurchasedDigitalVideoGames";
