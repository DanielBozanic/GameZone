import { useForm, FormProvider } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { createFormSchema } from "../Components/UserForms/CreateForm/CreateFormSchema";
import CreateForm from "../Components/UserForms/CreateForm/CreateForm";
import { useNavigate } from "react-router-dom";
import {
	Button,
	Form,
	Card,
	CardTitle,
	CardBody,
	Container,
	Row,
	Col,
} from "reactstrap";
import axios from "axios";
import { toast } from "react-toastify";
import * as userAPI from "../APIs/UserMicroservice/user_api";

toast.configure();
const SignUp = () => {
	const customId = "signup";
	const methods = useForm({
		resolver: yupResolver(createFormSchema),
		mode: "onChange",
	});

	const navigate = useNavigate();

	const signUp = (data) => {
		axios
			.post(userAPI.REGISTER, data, {
				headers: {
					"Content-Type": "application/json",
				},
			})
			.then((res) => {
				toast.success(res.data.message, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: false,
				});
				getVerificationCode(data.email);
				navigate("/signIn");
			})
			.catch((err) => {
				toast.error(err.response.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	const getVerificationCode = (email) => {
		axios.get(`${userAPI.GET_VERIFICATION_CODE}?email=${email}`);
	};

	return (
		<Container>
			<Row>
				<Col md="10">
					<Card className="form-card">
						<CardTitle className="title" tag="h2">
							Sign Up
						</CardTitle>
						<CardBody>
							<FormProvider {...methods}>
								<Form className="form">
									<CreateForm />
									<Row>
										<Col>
											<Button
												className="my-button"
												type="button"
												onClick={methods.handleSubmit(signUp)}
											>
												Sign up
											</Button>
										</Col>
									</Row>
								</Form>
							</FormProvider>
						</CardBody>
					</Card>
				</Col>
			</Row>
		</Container>
	);
};

export default SignUp;
