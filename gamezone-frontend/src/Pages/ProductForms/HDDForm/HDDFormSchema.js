import * as yup from "yup";

export const hddFormSchema = yup.object({
	Capacity: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Capacity is required"),
	DiskSpeed: yup.string().max(30, "Maximum number of characters allowed is 30"),
	Interface: yup.string().max(30, "Maximum number of characters allowed is 30"),
	TransferRate: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Form: yup.string().max(30, "Maximum number of characters allowed is 30"),
});
