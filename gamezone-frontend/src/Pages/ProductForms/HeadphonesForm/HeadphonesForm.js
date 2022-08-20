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
	CardHeader,
	CardTitle,
	CardBody,
	FormFeedback,
	Container,
	Row,
	Col,
} from "reactstrap";
import { Helmet } from "react-helmet";
import axios from "axios";
import { toast } from "react-toastify";
import ProductForm from "../../../Components/ProductForm/ProductForm";
import * as headphonesAPI from "../../../APIs/ProductMicroservice/headphones_api";
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const HeadphonesForm = (props) => {
	const customId = "headphonesForm";

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
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	const add = (data) => {
		data.Product.Image.Content = base64Image;
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
				if (product.Product.Quantity === 0 && data.product.Quantity > 0) {
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
		<>
			{product === null && props.addButton && (
				<Helmet>
					<title>{props.title} | GameZone</title>
				</Helmet>
			)}
			{product !== null && props.updateButton && (
				<Helmet>
					<title>Updating {product.Product.Name} | GameZone</title>
				</Helmet>
			)}
			<Container>
				<Row>
					<Col className="form-card-col">
						<Card>
							<CardHeader>
								<CardTitle className="title" tag="h2">
									{props.title}
								</CardTitle>
							</CardHeader>
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
												<FormGroup>
													<Label>Connection</Label>
													<Input
														className="input-field"
														type="text"
														name="Connection"
														invalid={
															methods.formState.errors.Connection?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Connection !== null
																? product.Connection
																: ""
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
														value={
															product === null || microphone !== null
																? microphone
																: product.Microphone
														}
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
														value={
															product === null || microphone !== null
																? microphone
																: product.Microphone
														}
														innerRef={methods.register}
														onChange={() => setMicrophone(false)}
													/>
												</span>
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
														invalid={
															methods.formState.errors.DriverSize?.message
														}
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
															methods.formState.errors.FrequencyResponse
																?.message
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
														{
															methods.formState.errors.FrequencyResponse
																?.message
														}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										{props.addButton && (
											<Row>
												<Col>
													<Button
														className="my-button"
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
														className="my-button"
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
		</>
	);
};

export default HeadphonesForm;
