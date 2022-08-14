import { useForm, FormProvider } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { createFormSchema } from "../Components/UserForms/CreateForm/CreateFormSchema";
import CreateForm from "../Components/UserForms/CreateForm/CreateForm";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import {
	Button,
	Form,
	Card,
	CardTitle,
	CardBody,
	Container,
	Row,
	Col,
	Spinner,
} from "reactstrap";
import { Helmet } from "react-helmet";
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
	const [loading, setLoading] = useState(false);

	const navigate = useNavigate();

	const signUp = (data) => {
		setLoading(true);
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
				setLoading(false);
			});
	};

	const getVerificationCode = (email) => {
		axios.get(`${userAPI.GET_VERIFICATION_CODE}?email=${email}`);
	};

	return (
		<>
			<Helmet>
				<title>Sign Up | GameZone</title>
			</Helmet>
			<Container>
				<Row>
					<Col md="10">
						<Card className="form-card">
							<CardTitle className="title" tag="h2">
								Sign Up
							</CardTitle>
							{loading && (
								<div className="div-spinner">
									<Spinner className="spinner" />
								</div>
							)}
							{!loading && (
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
							)}
						</Card>
					</Col>
				</Row>
			</Container>
		</>
	);
};

export default SignUp;
