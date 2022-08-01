const api = "http://localhost:7003/api/contactAndReport/contacts";

const adminAndEmployeeProtected = api + "/adminAndEmployeeProtected";
const userProtected = api + "/userProtected";

export const GET_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID =
	adminAndEmployeeProtected + "/getUnansweredContactMessagesByUserId";
export const ANSWER_CONTACT_MESSAGE =
	adminAndEmployeeProtected + "/answerContactMessage";
export const GET_ANSWERED_CONTACT_MESSAGES_BY_USER_ID =
	userProtected + "/getAnsweredContactMessagesByUserId";
export const SEND_CONTACT_MESSAGE = userProtected + "/sendContactMessage";
