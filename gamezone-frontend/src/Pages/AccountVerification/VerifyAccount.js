import { useForm } from "react-hook-form";
import { useParams, useNavigate } from "react-router-dom";
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
import axios from "axios";
import { useState } from "react";
import { toast } from "react-toastify";
import * as userAPI from "../../APIs/UserMicroservice/user_api";

const VerifyAccount = () => {
	const customId = "VerifyAccount";

	const { email } = useParams();
	const { register, handleSubmit } = useForm();
	const [loading, setLoading] = useState(false);

	const navigate = useNavigate();

	const verifyAccount = (data) => {
		setLoading(true);
		data.email = email;
		axios
			.put(`${userAPI.VERIFY_ACCOUNT}`, data)
			.then((res) => {
				toast.success(res.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
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

	return (
		<>
			<Helmet>
				<title>Verify Account | GameZone</title>
			</Helmet>
			<Container>
				<Row>
					<Col md="10">
						<Card className="form-card">
							<CardTitle className="title" tag="h2">
								Verify account {loading && <Spinner />}
							</CardTitle>
							<CardBody>
								<Form className="form">
									<Row>
										<Col>
											<FormGroup>
												<Label>Code</Label>
												<Input
													className="input-field"
													type="text"
													name="code"
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
												onClick={handleSubmit(verifyAccount)}
												disabled={loading}
											>
												Verify
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

export default VerifyAccount;
