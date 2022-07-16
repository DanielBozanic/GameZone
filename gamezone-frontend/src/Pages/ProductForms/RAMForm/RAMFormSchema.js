import * as yup from "yup";

export const ramFormSchema = yup.object({
	MemoryType: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Memory Type is required"),
	Capacity: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Capacity is required"),
	Speed: yup.string().max(30, "Maximum number of characters allowed is 30"),
	Voltage: yup.string().max(30, "Maximum number of characters allowed is 30"),
	Latency: yup.string().max(30, "Maximum number of characters allowed is 30"),
});
