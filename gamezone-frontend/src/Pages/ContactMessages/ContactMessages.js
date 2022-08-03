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
import * as authService from "../../Auth/AuthService";

toast.configure();
const ContactMessages = () => {
	const customId = "ContactMessages";

	const { id } = useParams();
	const [contactMessages, setContactMessages] = useState([]);
	const [selectedMessage, setSelectedMessage] = useState(null);
	const [loading, setLoading] = useState(true);

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
			getUnansweredMessagesByUserId();
		} else if (authService.isEmployee()) {
			getUnansweredMessages();
		} else if (authService.isUser()) {
			getContactMessagesByUserId();
		}
	}, []);

	const getUnansweredMessagesByUserId = () => {
		axios
			.get(`${contactAPI.GET_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID}/${id}`)
			.then((res) => {
				setContactMessages(res.data);
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getUnansweredMessages = () => {
		axios
			.get(`${contactAPI.GET_UNANSWERED_CONTACT_MESSAGES}`)
			.then((res) => {
				setContactMessages(res.data);
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

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
					getUnansweredMessagesByUserId();
				} else if (authService.isEmployee()) {
					getUnansweredMessages();
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

	const handleSelectedCard = (sm) => {
		setSelectedMessage(sm);
	};

	return (
		<>
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
													[
													{new Date(
														contactMessage.DateTime
													).toLocaleDateString()}{" "}
													{new Date(
														contactMessage.DateTime
													).toLocaleTimeString()}
													]
												</span>
											</Col>
											{authService.isEmployee() && (
												<Col>
													<span style={{ fontWeight: "bold" }}>
														{contactMessage.Username}
													</span>
												</Col>
											)}
										</Row>
									</CardHeader>
									<CardBody>
										<CardText>{contactMessage.Message}</CardText>
									</CardBody>
									<CardFooter>
										{authService.isUser() && contactMessage.Answer === "" && (
											<span style={{ marginLeft: "5px", fontWeight: "bold" }}>
												Answer pending
											</span>
										)}
										{authService.isUser() && contactMessage.Answer !== "" && (
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
								{contactMessages.length <= 0 &&
									(authService.isEmployee() || authService.isAdmin()) && (
										<Card>
											<CardBody>
												<Container>
													<Row>
														<Col>
															<CardTitle className="title " tag="h2">
																No messages to answer
															</CardTitle>
														</Col>
													</Row>
												</Container>
											</CardBody>
										</Card>
									)}
								{selectedMessage !== null &&
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
								{contactMessages.length <= 0 && authService.isUser() && (
									<Card>
										<CardBody>
											<Container>
												<Row>
													<Col>
														<CardTitle className="title " tag="h2">
															You have not sent any messages
														</CardTitle>
													</Col>
												</Row>
											</Container>
										</CardBody>
									</Card>
								)}
								{selectedMessage !== null &&
									selectedMessage.Answer !== null &&
									selectedMessage.Answer !== "" &&
									authService.isUser() && (
										<Card className="card">
											<CardHeader>
												<CardTitle className="title" tag="h5">
													Answer
												</CardTitle>
											</CardHeader>
											<CardBody>
												<Row>
													<Col>
														<CardText>{selectedMessage.Answer}</CardText>
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
