import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { signInSchema } from "./SignInSchema";
import { useNavigate } from "react-router-dom";
import {
	Button,
	Form,
	FormGroup,
	Label,
	Input,
	Card,
	CardTitle,
	CardBody,
	FormFeedback,
	Container,
	Row,
	Col,
} from "reactstrap";
import axios from "axios";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "../../Assets/css/forms.css";
import * as authAPI from "../../APIs/UserMicroservice/auth_api";
import * as authService from "../../Auth/AuthService";

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

	const navigate = useNavigate();

	const signIn = (data) => {
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
			});
	};

	return (
		<Container>
			<Row>
				<Col md="10">
					<Card className="form-card">
						<CardTitle className="form-title" tag="h2">
							Sign In
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
											className="confirm-form-btn"
											type="button"
											onClick={handleSubmit(signIn)}
										>
											Sign in
										</Button>
									</Col>
								</Row>
							</Form>
						</CardBody>
					</Card>
				</Col>
			</Row>
		</Container>
	);
};

export default SignIn;
