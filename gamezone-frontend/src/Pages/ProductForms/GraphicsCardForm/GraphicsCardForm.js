import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { graphicsCardFormSchema } from "./GraphicsCardFormSchema";
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
import { Helmet } from "react-helmet";
import axios from "axios";
import { toast } from "react-toastify";
import ProductForm from "../../../Components/ProductForm/ProductForm";
import * as graphicsCardAPI from "../../../APIs/ProductMicroservice/graphics_card_api";
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const GraphicsCardForm = (props) => {
	const customId = "graphicsCardForm";

	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, graphicsCardFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${graphicsCardAPI.GET_BY_ID}/${id}`)
				.then((res) => {
					setProduct(res.data);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	const add = (data) => {
		data.Product.Image.Content = base64Image;
		axios
			.post(graphicsCardAPI.CREATE, data)
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
		axios
			.put(graphicsCardAPI.UPDATE, data)
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
					<Col>
						<Card className="form-card">
							<CardTitle className="title" tag="h2">
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
													<Label>Chip Manufacturer</Label>
													<Input
														className="input-field"
														type="text"
														name="ChipManufacturer"
														invalid={
															methods.formState.errors.ChipManufacturer?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.ChipManufacturer : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ChipManufacturer?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Model Name</Label>
													<Input
														className="input-field"
														type="text"
														name="ModelName"
														invalid={
															methods.formState.errors.ModelName?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.ModelName : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ModelName?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Bus Width</Label>
													<Input
														className="input-field"
														type="text"
														name="BusWidth"
														invalid={methods.formState.errors.BusWidth?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.BusWidth !== null
																? product.BusWidth
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.BusWidth?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Memory Size</Label>
													<Input
														className="input-field"
														type="text"
														name="MemorySize"
														invalid={
															methods.formState.errors.MemorySize?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.MemorySize !== null
																? product.MemorySize
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.MemorySize?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Memory Type</Label>
													<Input
														className="input-field"
														type="text"
														name="MemoryType"
														invalid={
															methods.formState.errors.MemoryType?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.MemoryType !== null
																? product.MemoryType
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.MemoryType?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>PCI Interface</Label>
													<Input
														className="input-field"
														type="text"
														name="PCIInterface"
														invalid={
															methods.formState.errors.PCIInterface?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.PCIInterface !== null
																? product.PCIInterface
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.PCIInterface?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>GPU Speed</Label>
													<Input
														className="input-field"
														type="text"
														name="GPUSpeed"
														invalid={methods.formState.errors.GPUSpeed?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.GPUSpeed !== null
																? product.GPUSpeed
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.GPUSpeed?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Stream Processors</Label>
													<Input
														className="input-field"
														type="number"
														name="StreamProcessors"
														min="0"
														invalid={
															methods.formState.errors.StreamProcessors?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null &&
															product.StreamProcessors !== null
																? product.StreamProcessors
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.StreamProcessors?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Cooling</Label>
													<Input
														className="input-field"
														type="text"
														name="Cooling"
														invalid={methods.formState.errors.Cooling?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Cooling !== null
																? product.Cooling
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Cooling?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>HDMI</Label>
													<Input
														className="input-field"
														type="number"
														name="HDMI"
														min="0"
														invalid={methods.formState.errors.HDMI?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.HDMI !== null
																? product.HDMI
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.HDMI?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Display Port</Label>
													<Input
														className="input-field"
														type="number"
														name="DisplayPort"
														min="0"
														invalid={
															methods.formState.errors.DisplayPort?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.DisplayPort !== null
																? product.DisplayPort
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.DisplayPort?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>TDP</Label>
													<Input
														className="input-field"
														type="text"
														name="TDP"
														invalid={methods.formState.errors.TDP?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.TDP !== null
																? product.TDP
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.TDP?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Power Connector</Label>
													<Input
														className="input-field"
														type="text"
														name="PowerConnector"
														invalid={
															methods.formState.errors.PowerConnector?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null &&
															product.PowerConnector !== null
																? product.PowerConnector
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.PowerConnector?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Dimensions</Label>
													<Input
														className="input-field"
														type="text"
														name="Dimensions"
														invalid={
															methods.formState.errors.Dimensions?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Dimensions !== null
																? product.Dimensions
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Dimensions?.message}
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

export default GraphicsCardForm;
