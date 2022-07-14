import * as yup from "yup";

export const buyerInfoSchema = yup.object().shape({
	deliveryAddress: yup
		.string()
		.max(50, "Maximum number of characters allowed is 50")
		.required("Delivery address is required"),
	city: yup
		.string()
		.max(50, "Maximum number of characters allowed is 50")
		.required("City is required"),
	mobilePhoneNumber: yup
		.string()
		.required("Mobile phone number is required")
		.test(
			"len",
			"Mobile phone number must have exactly 10 characters",
			(val) => val?.length === 10
		)
		.matches(
			/^\d+$/,
			"Mobile phone number must only contain numerical characters"
		),
});
