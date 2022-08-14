import * as yup from "yup";

export const mouseFormSchema = yup.object({
	Connection: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Connection is required"),
	Sensor: yup.string().max(30, "Maximum number of characters allowed is 30"),
	BusWidth: yup.string().max(30, "Maximum number of characters allowed is 30"),
	DPI: yup.string().max(40, "Maximum number of characters allowed is 40"),
	PollingRate: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Color: yup.string().max(20, "Maximum number of characters allowed is 20"),
	TrackingSpeed: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Acceleration: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Weight: yup.string().max(20, "Maximum number of characters allowed is 20"),
	Lifespan: yup.string().max(30, "Maximum number of characters allowed is 30"),
});
