import * as yup from "yup";

export const signUpSchema = yup.object().shape({
	user_name: yup
		.string()
		.max(120, "Maximum number of characters allowed is 120")
		.required("Username is required"),
	password: yup
		.string()
		.max(120, "Maximum number of characters allowed is 120")
		.required("Password is required"),
	email: yup
		.string()
		.email("Email format is invalid")
		.max(120, "Maximum number of characters allowed is 120")
		.required("Email is required"),
	name: yup
		.string()
		.max(120, "Maximum number of characters allowed is 120")
		.required("Name is required"),
	surname: yup
		.string()
		.max(120, "Maximum number of characters allowed is 120")
		.required("Surname is required"),
});
