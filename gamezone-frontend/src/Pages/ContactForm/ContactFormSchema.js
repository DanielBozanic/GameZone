import * as yup from "yup";

export const contactFormSchema = yup.object().shape({
	Message: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000")
		.required("Message is required"),
});
