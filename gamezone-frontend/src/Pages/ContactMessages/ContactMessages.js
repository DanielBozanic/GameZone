import {
	Container,
	Row,
	Col,
	Card,
	CardHeader,
	CardTitle,
	CardBody,
	CardText,
	CardFooter,
	Input,
	Button,
	FormFeedback,
	Spinner,
} from "reactstrap";
import { Helmet } from "react-helmet";
import { ScrollMenu } from "react-horizontal-scrolling-menu";
import "../../Assets/css/contact-messages.css";
import axios from "axios";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { contactMessagesSchema } from "./ContactMessagesSchema";
import { toast } from "react-toastify";
import * as contactAPI from "../../APIs/ContactAndReportMicroservice/contact_api";
import * as userAPI from "../../APIs/UserMicroservice/user_api";
import * as authService from "../../Auth/AuthService";

toast.configure();
const ContactMessages = () => {
	const customId = "ContactMessages";

	const { id } = useParams();
	const [contactMessages, setContactMessages] = useState([]);
	const [selectedMessage, setSelectedMessage] = useState(null);
	const [loading, setLoading] = useState(true);
	const [user, setUser] = useState(null);

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({
		resolver: yupResolver(contactMessagesSchema),
		mode: "onChange",
	});

	useEffect(() => {
		if (authService.isAdmin()) {
			getContactMessagesByUserId();
			getUserById();
		} else if (authService.isUser()) {
			getContactMessagesByUserId();
		} else if (authService.isEmployee()) {
			getContactMessages();
		}
	}, []);

	const getContactMessagesByUserId = () => {
		axios
			.get(`${contactAPI.GET_CONTACT_MESSAGES_BY_USER_ID}/${id}`)
			.then((res) => {
				setContactMessages(res.data);
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getContactMessages = () => {
		axios
			.get(`${contactAPI.GET_CONTACT_MESSAGES}`)
			.then((res) => {
				setContactMessages(res.data);
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const answerMessage = (data) => {
		selectedMessage.Answer = data.Answer;
		axios
			.put(`${contactAPI.ANSWER_CONTACT_MESSAGE}`, selectedMessage)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
				if (authService.isAdmin()) {
					getContactMessagesByUserId();
				} else if (authService.isEmployee()) {
					getContactMessages();
				}
				setSelectedMessage(null);
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	const getUserById = () => {
		axios
			.get(`${userAPI.GET_USER_BY_ID}?userId=${id}`)
			.then((res) => {
				setUser(res.data.user);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const handleSelectedCard = (sm) => {
		setSelectedMessage(sm);
	};

	return (
		<>
			{authService.isAdmin() && user !== null && (
				<Helmet>
					<title>{user.user_name}'s contact messages | GameZone</title>
				</Helmet>
			)}
			{!authService.isAdmin() && (
				<Helmet>
					<title>Contact Messages | GameZone</title>
				</Helmet>
			)}
			{loading && (
				<div className="div-spinner">
					<Spinner className="spinner" />
				</div>
			)}
			{!loading && (
				<>
					{contactMessages.length > 0 && (
						<ScrollMenu>
							{contactMessages.map((contactMessage) => (
								<Card
									className={
										"card message-card " +
										(selectedMessage !== null &&
										selectedMessage.Id === contactMessage.Id
											? "selected-message-card"
											: "")
									}
									itemId={contactMessage}
									onClick={() => handleSelectedCard(contactMessage)}
								>
									<CardHeader>
										<Row>
											<Col>
												<span style={{ fontWeight: "bold" }}>
													{contactMessage.Subject} {"\t"}[
													{new Date(
														contactMessage.DateTime
													).toLocaleDateString()}{" "}
													{new Date(
														contactMessage.DateTime
													).toLocaleTimeString()}
													]
												</span>
											</Col>
										</Row>
										<Row>
											{authService.isEmployee() && (
												<Col>
													<span style={{ fontWeight: "bold" }}>
														Sent by {contactMessage.Username}
													</span>
												</Col>
											)}
										</Row>
									</CardHeader>
									<CardBody>
										<CardText style={{ whiteSpace: "pre-line" }}>
											{contactMessage.Message}
										</CardText>
									</CardBody>
									<CardFooter>
										{!authService.isUser() && contactMessage.Answer === "" && (
											<span style={{ marginLeft: "5px", fontWeight: "bold" }}>
												Unanswered
											</span>
										)}
										{authService.isUser() && contactMessage.Answer === "" && (
											<span style={{ marginLeft: "5px", fontWeight: "bold" }}>
												Answer pending
											</span>
										)}
										{contactMessage.Answer !== "" && (
											<span style={{ marginLeft: "5px", fontWeight: "bold" }}>
												Message answered
											</span>
										)}
									</CardFooter>
								</Card>
							))}
						</ScrollMenu>
					)}
					<Container>
						<Row className="answer-card-row">
							<Col>
								{selectedMessage !== null &&
									(selectedMessage.Answer === "" ||
										selectedMessage.Answer === null) &&
									(authService.isEmployee() || authService.isAdmin()) && (
										<Card className="card">
											<CardBody>
												<Row>
													<Col>
														<Input
															className="input-field answer-text-area"
															type="textarea"
															name="Answer"
															innerRef={register}
															invalid={errors.Answer?.message}
														/>
														<FormFeedback className="input-field-error-msg">
															{errors.Answer?.message}
														</FormFeedback>
													</Col>
												</Row>
												<Row>
													<Col>
														<Button
															style={{ marginTop: "10px" }}
															className="my-button"
															type="button"
															onClick={handleSubmit(answerMessage)}
														>
															Send
														</Button>
													</Col>
												</Row>
											</CardBody>
										</Card>
									)}
								{contactMessages.length <= 0 && (
									<Card>
										<CardBody>
											<Container>
												<Row>
													<Col>
														<CardTitle className="title " tag="h2">
															{authService.isUser() &&
																"You have not sent any messages"}
															{(authService.isEmployee() ||
																authService.isAdmin()) &&
																"There are currently no contact messages"}
														</CardTitle>
													</Col>
												</Row>
											</Container>
										</CardBody>
									</Card>
								)}
								{selectedMessage !== null &&
									selectedMessage.Answer !== null &&
									selectedMessage.Answer !== "" && (
										<Card className="card">
											<CardHeader>
												<CardTitle className="title" tag="h5">
													Answer
												</CardTitle>
											</CardHeader>
											<CardBody>
												<Row>
													<Col>
														<CardText style={{ whiteSpace: "pre-line" }}>
															{selectedMessage.Answer}
														</CardText>
													</Col>
												</Row>
											</CardBody>
										</Card>
									)}
							</Col>
						</Row>
					</Container>
				</>
			)}
		</>
	);
};

export default ContactMessages;
