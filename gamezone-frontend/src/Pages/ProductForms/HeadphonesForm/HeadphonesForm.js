import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
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
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const HeadphonesForm = (props) => {
	const customId = "headphonesForm";

	const [wireless, setWireless] = useState(false);
	const [microphone, setMicrophone] = useState(false);
	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, headphonesFormSchema)
		),
		mode: "onChange",
	});

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${headphonesAPI.GET_BY_ID}/${id}`)
				.then((res) => {
					setProduct(res.data);
					setMicrophone(null);
					setWireless(null);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	const add = (data) => {
		data.Product.Image.Content = base64Image;
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
		data.Microphone = helperFunctions.str2Bool(data.Microphone);
		console.log(data);
		axios
			.put(headphonesAPI.UPDATE, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
				if (product.Product.Amount === 0 && data.Product.Amount > 0) {
					notifyProductAvailability(product.Product.Id);
				}
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	const notifyProductAvailability = (productId) => {
		axios.get(
			`${productPurchaseAPI.NOTIFY_PRODUCT_AVAILABILITY}?productId=${productId}`
		);
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
													defaultValue={
														product !== null ? product.ConnectionType : ""
													}
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
											<div>
												<Label>Microphone</Label>
											</div>
											<span>
												<Label>Yes</Label>
												<Input
													className="ml-2"
													type="radio"
													name="Microphone"
													checked={
														product === null || microphone !== null
															? microphone
															: product.Microphone
													}
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
													checked={
														product === null || microphone !== null
															? !microphone
															: !product.Microphone
													}
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
													defaultValue={
														product !== null &&
														product.VirtualSurroundEncoding !== null
															? product.VirtualSurroundEncoding
															: ""
													}
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
													defaultValue={
														product !== null && product.Sensitivity !== null
															? product.Sensitivity
															: ""
													}
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
													defaultValue={
														product !== null && product.DriverSize !== null
															? product.DriverSize
															: ""
													}
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
												<Label>Frequency Response</Label>
												<Input
													className="input-field"
													type="text"
													name="FrequencyResponse"
													invalid={
														methods.formState.errors.FrequencyResponse?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null &&
														product.FrequencyResponse !== null
															? product.FrequencyResponse
															: ""
													}
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

export default HeadphonesForm;
