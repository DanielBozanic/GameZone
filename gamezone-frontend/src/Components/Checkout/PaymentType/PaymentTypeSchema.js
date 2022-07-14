import * as yup from "yup";

export const paymentTypeSchema = yup.object().shape({
	paymentType: yup.string().required("Payment type is required"),
});
