import * as yup from "yup";

export const passwordSchema = yup.object().shape({
	oldPassword: yup.string().required("Old Password is required"),
	newPassword: yup
		.string()
		.max(120, "Maximum number of characters allowed is 120")
		.required("New Password is required"),
	confirmPassword: yup
		.string()
		.oneOf([yup.ref("newPassword"), null], "Passwords must match")
		.required("You must enter the same password again"),
});
