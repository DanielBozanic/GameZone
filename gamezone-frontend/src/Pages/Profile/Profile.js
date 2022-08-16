import { useState, useEffect } from "react";
import {
	CardText,
	CardTitle,
	CardBody,
	CardFooter,
	CardHeader,
	Card,
	Row,
	Col,
	Container,
	Button,
	Input,
	Label,
	Form,
	FormGroup,
	FormFeedback,
	Spinner,
} from "reactstrap";
import { Helmet } from "react-helmet";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { updateSchema } from "./UpdateSchema";
import { passwordSchema } from "./PasswordSchema";
import axios from "axios";
import { toast } from "react-toastify";
import * as userAPI from "../../APIs/UserMicroservice/user_api";
import * as authService from "../../Auth/AuthService";
import "../../Assets/css/profile.css";

toast.configure();
const Profile = () => {
	const customId = "Profile";

	const [user, setUser] = useState(null);
	const [updateMode, setUpdateMode] = useState(false);
	const [changePasswordMode, setChangePasswordMode] = useState(false);
	const [bigSpinnerLoading, setBigSpinnerLoading] = useState(true);
	const [smallSpinnerLoading, setSmallSpinnerLoading] = useState(false);

	const updateForm = useForm({
		resolver: yupResolver(updateSchema),
		mode: "onChange",
	});

	const changePasswordForm = useForm({
		resolver: yupResolver(passwordSchema),
		mode: "onChange",
	});

	useEffect(() => {
		getUserById();
	}, []);

	const getUserById = () => {
		axios
			.get(`${userAPI.GET_USER_BY_ID}?userId=${Number(authService.getId())}`)
			.then((res) => {
				setUser(res.data.user);
				setBigSpinnerLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const update = (data) => {
		setSmallSpinnerLoading(true);
		user.email = data.email;
		user.name = data.name;
		user.surname = data.surname;
		axios
			.put(`${userAPI.UPDATE}`, user)
			.then((_res) => {
				setSmallSpinnerLoading(false);
				setUpdateMode(false);
				getUserById();
			})
			.catch((err) => {
				toast.error(err.response.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
				setSmallSpinnerLoading(false);
			});
	};

	const changePassword = (data) => {
		setSmallSpinnerLoading(true);
		user.oldPassword = data.oldPassword;
		user.password = data.newPassword;
		axios
			.put(`${userAPI.CHANGE_PASSWORD}`, user)
			.then((res) => {
				toast.error(res.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
				setSmallSpinnerLoading(false);
				setChangePasswordMode(false);
				getUserById();
			})
			.catch((err) => {
				toast.error(err.response.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
				setSmallSpinnerLoading(false);
			});
	};

	const purchaseHistory = () => {
		return "/purchaseHistory/" + authService.getId();
	};

	return (
		<>
			<Helmet>
				<title>Profile | GameZone</title>
			</Helmet>
			<Container>
				<Row>
					<Col className="profile-card-col" md="8">
						<Card className="card">
							{bigSpinnerLoading && (
								<div className="div-spinner">
									<Spinner className="spinner" />
								</div>
							)}
							{user !== null && (
								<>
									<CardHeader>
										<CardTitle className="title" tag="h2">
											{user.user_name} {smallSpinnerLoading && <Spinner />}
										</CardTitle>
									</CardHeader>
									<CardBody>
										{!updateMode && !changePasswordMode && (
											<CardText>
												<Row>
													<Col>
														<CardText>
															<Label style={{ fontWeight: "bold" }}>
																Email:
															</Label>{" "}
															{user.email}
														</CardText>
													</Col>
												</Row>
												<Row>
													<Col>
														<CardText>
															<Label style={{ fontWeight: "bold" }}>
																Name:
															</Label>{" "}
															{user.name}
														</CardText>
													</Col>
												</Row>
												<Row>
													<Col>
														<CardText>
															<Label style={{ fontWeight: "bold" }}>
																Surname:
															</Label>{" "}
															{user.surname}
														</CardText>
													</Col>
												</Row>
											</CardText>
										)}
										<Form className="form">
											{updateMode && (
												<>
													<Row>
														<Col>
															<FormGroup>
																<Label>Email</Label>
																<Input
																	className="input-field"
																	type="email"
																	placeholder="exmaple@email.com"
																	name="email"
																	invalid={updateForm.errors.email?.message}
																	innerRef={updateForm.register}
																	defaultValue={user.email}
																/>
																<FormFeedback className="input-field-error-msg">
																	{updateForm.errors.email?.message}
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
																	invalid={updateForm.errors.name?.message}
																	innerRef={updateForm.register}
																	defaultValue={user.name}
																/>
																<FormFeedback className="input-field-error-msg">
																	{updateForm.errors.name?.message}
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
																	invalid={updateForm.errors.surname?.message}
																	innerRef={updateForm.register}
																	defaultValue={user.surname}
																/>
																<FormFeedback className="input-field-error-msg">
																	{updateForm.errors.surname?.message}
																</FormFeedback>
															</FormGroup>
														</Col>
													</Row>
													<Row>
														<Col>
															<Button
																style={{ marginRight: "5px" }}
																className="my-button"
																type="button"
																onClick={updateForm.handleSubmit(update)}
																disabled={smallSpinnerLoading}
															>
																Update
															</Button>

															<Button
																className="my-button"
																type="button"
																onClick={() => setUpdateMode(false)}
																disabled={smallSpinnerLoading}
															>
																Cancel
															</Button>
														</Col>
													</Row>
												</>
											)}
											{changePasswordMode && (
												<>
													<Row>
														<Col>
															<FormGroup>
																<Label>Old Password</Label>
																<Input
																	className="input-field"
																	type="password"
																	name="oldPassword"
																	invalid={
																		changePasswordForm.errors.oldPassword
																			?.message
																	}
																	innerRef={changePasswordForm.register}
																/>
																<FormFeedback className="input-field-error-msg">
																	{
																		changePasswordForm.errors.oldPassword
																			?.message
																	}
																</FormFeedback>
															</FormGroup>
														</Col>
													</Row>
													<Row>
														<Col>
															<FormGroup>
																<Label>New Password</Label>
																<Input
																	className="input-field"
																	type="password"
																	name="newPassword"
																	invalid={
																		changePasswordForm.errors.newPassword
																			?.message
																	}
																	innerRef={changePasswordForm.register}
																/>
																<FormFeedback className="input-field-error-msg">
																	{
																		changePasswordForm.errors.newPassword
																			?.message
																	}
																</FormFeedback>
															</FormGroup>
														</Col>
													</Row>
													<Row>
														<Col>
															<FormGroup>
																<Label>Confirm Password</Label>
																<Input
																	className="input-field"
																	type="password"
																	name="confirmPassword"
																	invalid={
																		changePasswordForm.errors.confirmPassword
																			?.message
																	}
																	innerRef={changePasswordForm.register}
																/>
																<FormFeedback className="input-field-error-msg">
																	{
																		changePasswordForm.errors.confirmPassword
																			?.message
																	}
																</FormFeedback>
															</FormGroup>
														</Col>
													</Row>
													<Row>
														<Col>
															<Button
																style={{ marginRight: "5px" }}
																className="my-button"
																type="button"
																onClick={changePasswordForm.handleSubmit(
																	changePassword
																)}
																disabled={smallSpinnerLoading}
															>
																Confirm
															</Button>
															<Button
																className="my-button"
																type="button"
																onClick={() => setChangePasswordMode(false)}
																disabled={smallSpinnerLoading}
															>
																Cancel
															</Button>
														</Col>
													</Row>
												</>
											)}
										</Form>
										{!updateMode && !changePasswordMode && (
											<Button
												style={{ marginRight: "5px" }}
												className="my-button"
												type="button"
												onClick={() => setUpdateMode(true)}
											>
												Update
											</Button>
										)}
										{!updateMode && !changePasswordMode && (
											<Button
												style={{ marginRight: "5px" }}
												className="my-button"
												type="button"
												onClick={() => setChangePasswordMode(true)}
											>
												Change password
											</Button>
										)}
									</CardBody>
									{authService.isUser() && (
										<CardFooter>
											<Row>
												<Col>
													<a href={purchaseHistory}>
														<Button
															style={{ marginTop: "5px" }}
															className="my-button"
															type="button"
														>
															Purchase history
														</Button>
													</a>
												</Col>
											</Row>
										</CardFooter>
									)}
								</>
							)}
						</Card>
					</Col>
				</Row>
			</Container>
		</>
	);
};

export default Profile;
