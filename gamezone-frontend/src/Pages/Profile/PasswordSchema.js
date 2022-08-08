import * as yup from "yup";

export const passwordSchema = yup.object().shape({
	password: yup
		.string()
		.max(120, "Maximum number of characters allowed is 120")
		.required("Password is required"),
	confirmPassword: yup
		.string()
		.oneOf([yup.ref("password"), null], "Passwords must match")
		.required("You must enter the same password again"),
});
