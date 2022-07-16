import * as yup from "yup";

export const psuFormSchema = yup.object({
	Power: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Power is required"),
	Type: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Type is required"),
	FormFactor: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
});
