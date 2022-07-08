import * as yup from "yup";

export const signInSchema = yup.object().shape({
	user_name: yup.string().required("Username is required"),
	password: yup.string().required("Password is required"),
});
