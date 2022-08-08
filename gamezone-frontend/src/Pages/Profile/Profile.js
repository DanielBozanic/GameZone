import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
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
} from "reactstrap";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { profileSchema } from "./ProfileSchema";
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

	const navigate = useNavigate();

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({
		resolver: yupResolver(profileSchema),
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
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const update = (data) => {
		user.email = data.email;
		user.password = data.password;
		user.name = data.name;
		user.surname = data.surname;
		axios
			.put(`${userAPI.UPDATE}`, user)
			.then((_res) => {
				setUpdateMode(false);
				getUserById();
			})
			.catch((err) => {
				toast.error(err.response.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	const purchaseHistory = () => {
		navigate("/purchaseHistory/" + authService.getId());
	};

	return (
		<Container>
			<Row>
				<Col className="profile-card-col" md="8">
					{user !== null && (
						<Card className="card">
							<CardHeader>
								<CardTitle className="title" tag="h5">
									{user.user_name}
								</CardTitle>
							</CardHeader>
							<CardBody>
								{!updateMode && (
									<CardText>
										<Row>
											<Col>
												<CardText>
													<Label style={{ fontWeight: "bold" }}>Email:</Label>{" "}
													{user.email}
												</CardText>
											</Col>
										</Row>
										<Row>
											<Col>
												<CardText>
													<Label style={{ fontWeight: "bold" }}>Name:</Label>{" "}
													{user.name}
												</CardText>
											</Col>
										</Row>
										<Row>
											<Col>
												<CardText>
													<Label style={{ fontWeight: "bold" }}>Surname:</Label>{" "}
													{user.surname}
												</CardText>
											</Col>
										</Row>
									</CardText>
								)}
								{updateMode && (
									<Form className="form">
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
														defaultValue={user.password}
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
														defaultValue={user.email}
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
														defaultValue={user.name}
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
														defaultValue={user.surname}
													/>
													<FormFeedback className="input-field-error-msg">
														{errors.surname?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
									</Form>
								)}
							</CardBody>
							<CardFooter>
								{!updateMode && (
									<Button
										style={{ marginRight: "5px" }}
										className="my-button"
										type="button"
										onClick={() => setUpdateMode(true)}
									>
										Update
									</Button>
								)}
								{updateMode && (
									<Button
										style={{ marginRight: "5px" }}
										className="my-button"
										type="button"
										onClick={handleSubmit(update)}
									>
										Save
									</Button>
								)}
								{updateMode && (
									<Button
										style={{ marginRight: "5px" }}
										className="my-button"
										type="button"
										onClick={() => setUpdateMode(false)}
									>
										Cancel
									</Button>
								)}
								{authService.isUser() && (
									<Button
										className="my-button"
										type="button"
										onClick={purchaseHistory}
									>
										Purchase history
									</Button>
								)}
							</CardFooter>
						</Card>
					)}
				</Col>
			</Row>
		</Container>
	);
};

export default Profile;
