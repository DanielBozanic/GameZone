import * as yup from "yup";

export const newsEditorSchema = yup.object().shape({
	UnpublishedTitle: yup
		.string()
		.max(250, "Maximum number of characters allowed is 250")
		.required("Title is required"),
	UnpublishedDescription: yup
		.string()
		.max(200, "Maximum number of characters allowed is 200"),
});
