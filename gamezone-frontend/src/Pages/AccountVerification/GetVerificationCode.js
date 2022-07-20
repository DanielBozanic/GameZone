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
} from "reactstrap";
import axios from "axios";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "../../Assets/css/forms.css";
import * as userAPI from "../../APIs/UserMicroservice/user_api";

const GetVerificationCode = () => {
	const customId = "GetVerificationCode";

	const { register, handleSubmit } = useForm();

	const getVerificationCode = (data) => {
		axios
			.get(`${userAPI.GET_VERIFICATION_CODE}?email=${data.email}`)
			.then((res) => {
				toast.success(res.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
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
							Get verification code
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
											className="confirm-form-btn"
											type="button"
											onClick={handleSubmit(getVerificationCode)}
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
	);
};

export default GetVerificationCode;
