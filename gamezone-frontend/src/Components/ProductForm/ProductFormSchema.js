import * as yup from "yup";

export const productFormSchema = yup.object({
	Product: yup.object({
		Name: yup
			.string()
			.max(100, "Maximum number of characters allowed is 100")
			.required("Name is required"),
		Description: yup
			.string()
			.required("Description is required")
			.max(1500, "Maximum number of characters allowed is 1500"),
		Manufacturer: yup
			.string()
			.max(40, "Maximum number of characters allowed is 40")
			.required("Manufacturer is required"),
		Price: yup
			.number()
			.test(
				"positive",
				"Price must be a positive number",
				(value) => value >= 0
			)
			.required("Price is required")
			.typeError("Price must be a number"),
		Amount: yup
			.number()
			.test(
				"positive",
				"Amount must be a positive number",
				(value) => value >= 0
			)
			.required("Amount is required")
			.typeError("Amount must be a number")
			.integer("Amount must be a non decimal value"),
		Image: yup
			.mixed()
			.test("required", "You need to provide an image", (value) => {
				return value !== undefined && value.length > 0 && value[0].size > 0;
			})
			.test("type", "File type must be image", (value) => {
				return (
					value !== undefined &&
					value.length > 0 &&
					(value[0].type === "image/jpeg" ||
						value[0].type === "image/jpg" ||
						value[0].type === "image/png")
				);
			}),
	}),
});
