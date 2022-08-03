const api = "http://localhost:7003/api/contactAndReport/contacts";

const adminProtected = api + "/adminProtected";
const employeeProtected = api + "/employeeProtected";
const adminAndEmployeeProtected = api + "/adminAndEmployeeProtected";
const userProtected = api + "/userProtected";

export const GET_UNANSWERED_CONTACT_MESSAGES =
	employeeProtected + "/getUnansweredContactMessages";
export const GET_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID =
	adminProtected + "/getUnansweredContactMessagesByUserId";
export const ANSWER_CONTACT_MESSAGE =
	adminAndEmployeeProtected + "/answerContactMessage";
export const GET_CONTACT_MESSAGES_BY_USER_ID =
	userProtected + "/getContactMessagesByUserId";
export const SEND_CONTACT_MESSAGE = userProtected + "/sendContactMessage";
