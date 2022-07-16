import { useForm, FormProvider } from "react-hook-form";
import { useState } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { headphonesFormSchema } from "./HeadphonesFormSchema";
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
import "../../../Assets/css/forms.css";
import ProductForm from "../../../Components/ProductForm/ProductForm";
import * as headphonesAPI from "../../../APIs/ProductMicroservice/headphones_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const HeadphonesForm = (props) => {
	const customId = "headphonesForm";

	const [wireless, setWireless] = useState(false);
	const [microphone, setMicrophone] = useState(false);
	const [base64Image, setBase64Image] = useState("");

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, headphonesFormSchema)
		),
		mode: "onChange",
	});

	const add = (data) => {
		data.Product.Image = base64Image;
		data.Wireless = helperFunctions.str2Bool(data.Wireless);
		data.Microphone = helperFunctions.str2Bool(data.Microphone);
		axios
			.post(headphonesAPI.CREATE, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
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
		<Container>
			<Row>
				<Col>
					<Card className="form-card">
						<CardTitle className="form-title" tag="h2">
							{props.title}
						</CardTitle>
						<CardBody>
							<FormProvider {...methods}>
								<Form className="form">
									<ProductForm setBase64Image={setBase64Image} />

									<Row>
										<Col>
											<FormGroup>
												<Label>Connection Type</Label>
												<Input
													className="input-field"
													type="text"
													name="ConnectionType"
													invalid={
														methods.formState.errors.ConnectionType?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.ConnectionType?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<div>
												<Label>Wireless</Label>
											</div>
											<span>
												<Label>Yes</Label>
												<Input
													className="ml-2"
													type="radio"
													name="Wireless"
													checked={wireless}
													value={wireless}
													innerRef={methods.register}
													onChange={() => setWireless(true)}
												/>
											</span>
											<span className="pl-5">
												<Label>No</Label>
												<Input
													className="ml-2"
													type="radio"
													name="Wireless"
													checked={!wireless}
													value={wireless}
													innerRef={methods.register}
													onChange={() => setWireless(false)}
												/>
											</span>
										</Col>
									</Row>
									<Row>
										<Col>
											<div>
												<Label>Microphone</Label>
											</div>
											<span>
												<Label>Yes</Label>
												<Input
													className="ml-2"
													type="radio"
													name="Microphone"
													checked={microphone}
													value={microphone}
													innerRef={methods.register}
													onChange={() => setMicrophone(true)}
												/>
											</span>
											<span className="pl-5">
												<Label>No</Label>
												<Input
													className="ml-2"
													type="radio"
													name="Microphone"
													checked={!microphone}
													value={microphone}
													innerRef={methods.register}
													onChange={() => setMicrophone(false)}
												/>
											</span>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Virtual Surround Encoding</Label>
												<Input
													className="input-field"
													type="text"
													name="VirtualSurroundEncoding"
													invalid={
														methods.formState.errors.VirtualSurroundEncoding
															?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{
														methods.formState.errors.VirtualSurroundEncoding
															?.message
													}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Sensitivity</Label>
												<Input
													className="input-field"
													type="text"
													name="Sensitivity"
													invalid={
														methods.formState.errors.Sensitivity?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Sensitivity?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Driver Size</Label>
												<Input
													className="input-field"
													type="text"
													name="DriverSize"
													invalid={methods.formState.errors.DriverSize?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.DriverSize?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Color</Label>
												<Input
													className="input-field"
													type="text"
													name="Color"
													invalid={methods.formState.errors.Color?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Color?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Weight</Label>
												<Input
													className="input-field"
													type="text"
													name="Weight"
													invalid={methods.formState.errors.Weight?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Weight?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Frequency Response</Label>
												<Input
													className="input-field"
													type="text"
													name="FrequencyResponse"
													invalid={
														methods.formState.errors.FrequencyResponse?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.FrequencyResponse?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									{props.addButton && (
										<Row>
											<Col>
												<Button
													className="confirm-form-btn"
													type="button"
													onClick={methods.handleSubmit(add)}
												>
													Add
												</Button>
											</Col>
										</Row>
									)}
									{props.updateButton && (
										<Row>
											<Col>
												<Button
													className="confirm-form-btn"
													type="button"
													onClick={methods.handleSubmit(add)}
												>
													Update
												</Button>
											</Col>
										</Row>
									)}
								</Form>
							</FormProvider>
						</CardBody>
					</Card>
				</Col>
			</Row>
		</Container>
	);
};

export default HeadphonesForm;
