import * as yup from "yup";

export const processorFormSchema = yup.object({
	Type: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Type is required"),
	Socket: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Socket is required"),
	NumberOfCores: yup
		.number()
		.typeError("Number Of Cores must be a number")
		.positive("Number Of Cores must be greater than zero")
		.integer("Number Of Cores must be a non decimal value")
		.nullable()
		.transform((_, val) => (val !== "" ? Number(val) : null)),
	Threads: yup
		.number()
		.typeError("Threads must be a number")
		.positive("Threads must be greater than zero")
		.integer("Threads must be a non decimal value")
		.nullable()
		.transform((_, val) => (val !== "" ? Number(val) : null)),
	TDP: yup.string().max(30, "Maximum number of characters allowed is 30"),
	IntegratedGraphics: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	BaseClockRate: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
	TurboClockRate: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
});
