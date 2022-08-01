const api = "http://localhost:7003/api/contactAndReport/bans";

const adminProtected = api + "/adminProtected";

export const GET_USER_BAN_HISTORY = adminProtected + "/getUserBanHistory";
export const ADD_BAN = adminProtected + "/addBan";
export const SEND_EMAIL_TO_BANNED_USER =
	adminProtected + "/sendEmailToBannedUser";
