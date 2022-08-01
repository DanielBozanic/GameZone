import * as yup from "yup";

export const banSchema = yup.object().shape({
	Description: yup
		.string()
		.max(1000, "Maximum number of characters allowed is 1000"),
	ExpirationDate: yup
		.date()
		.nullable()
		.default(undefined)
		.transform((curr, orig) => (orig === "" ? null : curr))
		.min(new Date(), "Expiration date must not be in the past")
		.typeError("Invalid Date")
		.required("Expiration Date is required"),
});
