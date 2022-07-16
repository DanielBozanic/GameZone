import * as yup from "yup";

export const keyboardFormSchema = yup.object({
	KeyboardConnector: yup
		.string()
		.max(30, "Maximum number of characters allowed is 30")
		.required("Keyboard Connector is required"),
	KeyType: yup.string().max(30, "Maximum number of characters allowed is 30"),
	LetterLayout: yup
		.string()
		.max(20, "Maximum number of characters allowed is 20"),
	KeyboardColor: yup
		.string()
		.max(20, "Maximum number of characters allowed is 20"),
});
