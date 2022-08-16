const api = "http://localhost:8000/api/contactAndReport/contacts";

export const GET_CONTACT_MESSAGES = api + "/getContactMessages";
export const GET_CONTACT_MESSAGES_BY_USER_ID =
	api + "/getContactMessagesByUserId";
export const GET_NUMBER_OF_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID =
	api + "/getNumberOfUnansweredContactMessagesByUserId";
export const ANSWER_CONTACT_MESSAGE = api + "/answerContactMessage";
export const SEND_CONTACT_MESSAGE = api + "/sendContactMessage";
