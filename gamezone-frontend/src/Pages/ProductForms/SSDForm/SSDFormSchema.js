import * as yup from "yup";

export const ssdFormSchema = yup.object({
	Capacity: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Capacity is required"),
	Form: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Form is required"),
	Interface: yup.string().max(30, "Maximum number of characters allowed is 30"),
	MaxSequentialRead: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	MaxSequentialWrite: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Dimensions: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
});
