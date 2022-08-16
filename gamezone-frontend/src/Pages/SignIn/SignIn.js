import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { signInSchema } from "./SignInSchema";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import {
	Button,
	Form,
	FormGroup,
	Label,
	Input,
	Card,
	CardTitle,
	CardBody,
	CardFooter,
	FormFeedback,
	Container,
	Row,
	Col,
	Spinner,
} from "reactstrap";
import { Helmet } from "react-helmet";
import axios from "axios";
import { toast } from "react-toastify";
import * as authAPI from "../../APIs/UserMicroservice/auth_api";
import * as authService from "../../Auth/AuthService";

toast.configure();
const SignIn = () => {
	const customId = "signin";
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({
		resolver: yupResolver(signInSchema),
		mode: "onChange",
	});
	const [loading, setLoading] = useState(false);

	const navigate = useNavigate();

	const signIn = (data) => {
		setLoading(true);
		axios
			.post(authAPI.LOGIN, data, {
				headers: {
					"Content-Type": "application/json",
				},
			})
			.then((res) => {
				authService.storeToken(res.data);
				navigate("/");
				window.location.reload();
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

	const getVerificationCode = () => {
		return "/getVerificationCode";
	};

	return (
		<>
			<Helmet>
				<title>Sign In | GameZone</title>
			</Helmet>
			<Container>
				<Row>
					<Col md="10">
						<Card className="form-card">
							<CardTitle className="title" tag="h2">
								Sign In {loading && <Spinner />}
							</CardTitle>
							<CardBody>
								<Form className="form">
									<Row>
										<Col>
											<FormGroup>
												<Label>Username</Label>
												<Input
													className="input-field"
													type="text"
													name="user_name"
													invalid={errors.user_name?.message}
													innerRef={register}
												/>
												<FormFeedback className="input-field-error-msg">
													{errors.user_name?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Password</Label>
												<Input
													className="input-field"
													type="password"
													name="password"
													invalid={errors.password?.message}
													innerRef={register}
												/>
												<FormFeedback className="input-field-error-msg">
													{errors.password?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<Button
												className="my-button"
												type="button"
												onClick={handleSubmit(signIn)}
												disabled={loading}
											>
												Sign in
											</Button>
										</Col>
									</Row>
								</Form>
							</CardBody>
							<CardFooter>
								<Row>
									<Col>
										<a
											style={{ textDecoration: "none", color: "inherit" }}
											href={getVerificationCode()}
										>
											Get verification code
										</a>
									</Col>
								</Row>
							</CardFooter>
						</Card>
					</Col>
				</Row>
			</Container>
		</>
	);
};

export default SignIn;
