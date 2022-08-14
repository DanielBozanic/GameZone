import * as yup from "yup";

export const motherboardFormSchema = yup.object({
	ProcessorType: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Processor Type is required"),
	Socket: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Socket is required"),
	SupportedProcessors: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	Chipset: yup.string().max(30, "Maximum number of characters allowed is 30"),
	Memory: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	ExpansionSlots: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	StorageInterface: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	Audio: yup.string().max(1000, "Maximum number of characters allowed is 1000"),
	USB: yup.string().max(1000, "Maximum number of characters allowed is 1000"),
	BackPanelConnectors: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	InternalConnectors: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	BIOS: yup.string().max(1000, "Maximum number of characters allowed is 1000"),
	FormFactor: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
});
