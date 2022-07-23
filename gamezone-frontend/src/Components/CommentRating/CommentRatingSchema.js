import * as yup from "yup";

export const commentRatingSchema = yup.object({
	Comment: yup
		.string()
		.max(490, "Maximum number of characters allowed is 490")
		.required("Comment is required"),
});
