import * as yup from "yup";

export const reportSchema = yup.object().shape({
	Description: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
});
