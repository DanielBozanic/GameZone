import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { signUpSchema } from "./SignUpSchema";
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
import "../../Assets/css/sign-up.css";
import * as userAPI from "../../APIs/UserMicroservice/user_api";

toast.configure();
const SignUp = () => {
	const customId = "signup";
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({
		resolver: yupResolver(signUpSchema),
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
			.then((_res) => {
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

	return (
		<Container>
			<Row>
				<Col md="10">
					<Card className="sign-up-card">
						<CardTitle tag="h2">Sign Up</CardTitle>
						<CardBody>
							<Form className="sign-up-card-form">
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
										<FormGroup>
											<Label>Email</Label>
											<Input
												className="input-field"
												type="email"
												placeholder="exmaple@email.com"
												name="email"
												invalid={errors.email?.message}
												innerRef={register}
											/>
											<FormFeedback className="input-field-error-msg">
												{errors.email?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<FormGroup>
											<Label>Name</Label>
											<Input
												className="input-field"
												type="text"
												name="name"
												invalid={errors.name?.message}
												innerRef={register}
											/>
											<FormFeedback className="input-field-error-msg">
												{errors.name?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<FormGroup>
											<Label>Surname</Label>
											<Input
												className="input-field"
												type="text"
												name="surname"
												invalid={errors.surname?.message}
												innerRef={register}
											/>
											<FormFeedback className="input-field-error-msg">
												{errors.surname?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<Button
											className="signup-btn"
											type="button"
											onClick={handleSubmit(signUp)}
										>
											Sign up
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

export default SignUp;
