import * as yup from "yup";

export const contactFormSchema = yup.object().shape({
	Subject: yup
		.string()
		.max(70, "Maximum number of characters allowed is 70")
		.required("Subject is required"),
	Message: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000")
		.required("Message is required"),
});
