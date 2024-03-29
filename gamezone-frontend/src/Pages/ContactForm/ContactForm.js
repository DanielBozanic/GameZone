import {
	Form,
	FormGroup,
	Label,
	Input,
	FormFeedback,
	Container,
	Row,
	Col,
	Card,
	CardHeader,
	CardTitle,
	CardBody,
	Button,
} from "reactstrap";
import { Helmet } from "react-helmet";
import axios from "axios";
import { toast } from "react-toastify";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { contactFormSchema } from "./ContactFormSchema";
import * as contactAPI from "../../APIs/ContactAndReportMicroservice/contact_api";

toast.configure();
const ContactForm = () => {
	const customId = "ContactForm";

	const {
		register,
		handleSubmit,
		formState: { errors },
		reset,
	} = useForm({
		resolver: yupResolver(contactFormSchema),
		mode: "onChange",
	});

	const sendMessage = (data) => {
		axios
			.post(`${contactAPI.SEND_CONTACT_MESSAGE}`, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				reset();
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	return (
		<>
			<Helmet>
				<title>Contact | GameZone</title>
			</Helmet>
			<Container>
				<Row>
					<Col className="form-card-col" md="10">
						<Card>
							<CardHeader>
								<CardTitle className="title" tag="h2">
									Contact
								</CardTitle>
							</CardHeader>
							<CardBody>
								<Form className="form">
									<Row>
										<Col>
											<FormGroup>
												<Label>Subject</Label>
												<Input
													className="input-field"
													type="text"
													name="Subject"
													invalid={errors.Subject?.message}
													innerRef={register}
												/>
												<FormFeedback className="input-field-error-msg">
													{errors.Subject?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Message</Label>
												<Input
													style={{ height: "450px", resize: "none" }}
													className="input-field"
													type="textarea"
													name="Message"
													invalid={errors.Message?.message}
													innerRef={register}
												/>
												<FormFeedback className="input-field-error-msg">
													{errors.Message?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<Button
												className="my-button"
												type="button"
												onClick={handleSubmit(sendMessage)}
											>
												Send
											</Button>
										</Col>
									</Row>
								</Form>
							</CardBody>
						</Card>
					</Col>
				</Row>
			</Container>
		</>
	);
};

export default ContactForm;
