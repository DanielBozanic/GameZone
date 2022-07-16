import * as yup from "yup";

export const headphonesFormSchema = yup.object({
	ConnectionType: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Connection Type is required"),
	VirtualSurroundEncoding: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Sensitivity: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	DriverSize: yup
		.string()
		.max(20, "Maximum number of characters allowed is 20"),
	Color: yup.string().max(20, "Maximum number of characters allowed is 20"),
	Weight: yup.string().max(20, "Maximum number of characters allowed is 20"),
	FrequencyResponse: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
});
