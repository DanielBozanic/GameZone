import * as yup from "yup";

export const consoleFormSchema = yup.object({
	Platform: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Platform is required"),
});
