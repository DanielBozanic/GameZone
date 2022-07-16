import * as yup from "yup";

export const videoGameFormSchema = yup.object({
	Platform: yup
		.string()
		.max(40, "Maximum number of characters allowed is 40")
		.required("Platform is required"),
	Rating: yup
		.number()
		.typeError("Rating must be a number")
		.positive("Rating must be greater than zero")
		.integer("Rating must be a non decimal value")
		.required("Rating is required"),
	Genre: yup
		.string()
		.max(50, "Maximum number of characters allowed is 50")
		.required("Genre is required"),
	ReleaseDate: yup.string(),
});
