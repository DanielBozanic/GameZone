import * as yup from "yup";

export const graphicsCardFormSchema = yup.object({
	ChipManufacturer: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Chip Manufacturer is required"),
	ModelName: yup
		.string()
		.max(100, "Maximum number of characters allowed is 100")
		.required("Model Name is required"),
	BusWidth: yup.string().max(30, "Maximum number of characters allowed is 30"),
	MemorySize: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	MemoryType: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	PCIInterface: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
	GPUSpeed: yup.string().max(20, "Maximum number of characters allowed is 20"),
	CUDAStreamProcessors: yup
		.number()
		.typeError("CUDA Stream Processors must be a number")
		.test(
			"positive",
			"CUDA Stream Processors must be a positive number",
			(value) => value >= 0
		)
		.integer("CUDA Stream Processors must be a non decimal value")
		.nullable()
		.transform((_, val) => (val !== "" ? Number(val) : null)),
	Cooling: yup.string().max(20, "Maximum number of characters allowed is 20"),
	HDMI: yup
		.number()
		.typeError("HDMI must be a number")
		.test("positive", "HDMI must be a positive number", (value) => value >= 0)
		.integer("HDMI must be a non decimal value")
		.nullable()
		.transform((_, val) => (val !== "" ? Number(val) : null)),
	DisplayPort: yup
		.number()
		.typeError("Display Port must be a number")
		.test(
			"positive",
			"Display Port must be a positive number",
			(value) => value >= 0
		)
		.integer("Display Port must be a non decimal value")
		.nullable()
		.transform((_, val) => (val !== "" ? Number(val) : null)),
	TDP: yup.string().max(30, "Maximum number of characters allowed is 30"),
	PowerConnector: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Dimensions: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
});
