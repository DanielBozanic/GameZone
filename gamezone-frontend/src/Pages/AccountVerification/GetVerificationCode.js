import { useForm } from "react-hook-form";
import {
	Button,
	Form,
	FormGroup,
	Label,
	Input,
	Card,
	CardTitle,
	CardBody,
	Container,
	Row,
	Col,
	Spinner,
} from "reactstrap";
import { Helmet } from "react-helmet";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { toast } from "react-toastify";
import * as userAPI from "../../APIs/UserMicroservice/user_api";

const GetVerificationCode = () => {
	const customId = "GetVerificationCode";
	const [loading, setLoading] = useState(false);

	const { register, handleSubmit } = useForm();
	const navigate = useNavigate();

	const getVerificationCode = (data) => {
		setLoading(true);
		axios
			.get(`${userAPI.GET_VERIFICATION_CODE}?email=${data.email}`)
			.then((res) => {
				toast.success(res.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
				navigate("/verify/" + data.email);
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

	return (
		<>
			<Helmet>
				<title>Get Verification Code | GameZone</title>
			</Helmet>
			<Container>
				<Row>
					<Col md="10">
						<Card className="form-card">
							<CardTitle className="title" tag="h2">
								Get verification code {loading && <Spinner />}
							</CardTitle>
							<CardBody>
								<Form className="form">
									<Row>
										<Col>
											<FormGroup>
												<Label>Email</Label>
												<Input
													className="input-field"
													type="text"
													name="email"
													innerRef={register}
												/>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<Button
												className="my-button"
												type="button"
												onClick={handleSubmit(getVerificationCode)}
												disabled={loading}
											>
												Submit
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

export default GetVerificationCode;
