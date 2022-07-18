import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { mouseFormSchema } from "./MouseFormSchema";
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
import * as mouseAPI from "../../../APIs/ProductMicroservice/mouse_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const MouseForm = (props) => {
	const customId = "mouseForm";

	const [wireless, setWireless] = useState(false);
	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, mouseFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${mouseAPI.GET_BY_ID}/${id}`)
				.then((res) => {
					setWireless(null);
					setProduct(res.data);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	const add = (data) => {
		data.Product.Image.Content = base64Image;
		data.Wireless = helperFunctions.str2Bool(data.Wireless);
		axios
			.post(mouseAPI.CREATE, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
				setFileName("");
				setBase64Image("");
				methods.reset();
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	const update = (data) => {
		data.Product.Id = product.Product.Id;
		data.Product.Type = product.Product.Type;
		data.Product.Image.Content = base64Image;
		data.Wireless = helperFunctions.str2Bool(data.Wireless);
		axios
			.put(mouseAPI.UPDATE, data)
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
									<ProductForm
										product={product}
										fileName={fileName}
										setFileName={setFileName}
										setBase64Image={setBase64Image}
									/>
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
													checked={
														product === null || wireless !== null
															? wireless
															: product.Wireless
													}
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
													checked={
														product === null || wireless !== null
															? !wireless
															: !product.Wireless
													}
													value={wireless}
													innerRef={methods.register}
													onChange={() => setWireless(false)}
												/>
											</span>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Connection</Label>
												<Input
													className="input-field"
													type="text"
													name="Connection"
													invalid={methods.formState.errors.Connection?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null ? product.Connection : ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Connection?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Sensor</Label>
												<Input
													className="input-field"
													type="text"
													name="Sensor"
													invalid={methods.formState.errors.Sensor?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.Sensor !== null
															? product.Sensor
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Sensor?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>DPI</Label>
												<Input
													className="input-field"
													type="text"
													name="DPI"
													invalid={methods.formState.errors.DPI?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.DPI !== null
															? product.DPI
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.DPI?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Polling Rate</Label>
												<Input
													className="input-field"
													type="text"
													name="PollingRate"
													invalid={
														methods.formState.errors.PollingRate?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.PollingRate !== null
															? product.PollingRate
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.PollingRate?.message}
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
													defaultValue={
														product !== null && product.Color !== null
															? product.Color
															: ""
													}
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
												<Label>Tracking Speed</Label>
												<Input
													className="input-field"
													type="text"
													name="TrackingSpeed"
													invalid={
														methods.formState.errors.TrackingSpeed?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.TrackingSpeed !== null
															? product.TrackingSpeed
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.TrackingSpeed?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Acceleration</Label>
												<Input
													className="input-field"
													type="text"
													name="Acceleration"
													invalid={
														methods.formState.errors.Acceleration?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.Acceleration !== null
															? product.Acceleration
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Acceleration?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Buttons</Label>
												<Input
													className="input-field"
													type="number"
													min="0"
													name="Buttons"
													invalid={methods.formState.errors.Buttons?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.Buttons !== null
															? product.Buttons
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Buttons?.message}
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
													defaultValue={
														product !== null && product.Weight !== null
															? product.Weight
															: ""
													}
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
												<Label>Lifespan</Label>
												<Input
													className="input-field"
													type="text"
													name="Lifespan"
													invalid={methods.formState.errors.Lifespan?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.Lifespan !== null
															? product.Lifespan
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Lifespan?.message}
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
													onClick={methods.handleSubmit(update)}
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

export default MouseForm;
