const api = "http://localhost:8000/api/contactAndReport/contacts";

export const GET_UNANSWERED_CONTACT_MESSAGES =
	api + "/getUnansweredContactMessages";
export const GET_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID =
	api + "/getUnansweredContactMessagesByUserId";
export const ANSWER_CONTACT_MESSAGE = api + "/answerContactMessage";
export const GET_CONTACT_MESSAGES_BY_USER_ID =
	api + "/getContactMessagesByUserId";
export const SEND_CONTACT_MESSAGE = api + "/sendContactMessage";
