import * as yup from "yup";

export const monitorFormSchema = yup.object({
	Resolution: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Resolution is required"),
	RefreshRate: yup
		.string()
		.max(20, "Maximum number of characters allowed is 20")
		.required("Refresh Rate is required"),
	Size: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Size is required"),
	AspectRatio: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
	ContrastRatio: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40"),
	ResponseTime: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	PanelType: yup.string().max(30, "Maximum number of characters allowed is 30"),
	ViewingAngle: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30"),
	Brightness: yup
		.string()
		.max(20, "Maximum number of characters allowed is 20"),
});
