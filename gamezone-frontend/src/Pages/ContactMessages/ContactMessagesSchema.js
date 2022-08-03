import * as yup from "yup";

export const contactMessagesSchema = yup.object().shape({
	Answer: yup
		.string()
		.required("Answer is required")
		.max(1000, "Maximum number of characters allowed is 1000"),
});
