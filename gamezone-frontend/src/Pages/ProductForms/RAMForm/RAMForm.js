import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { ramFormSchema } from "./RAMFormSchema";
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
import * as ramAPI from "../../../APIs/ProductMicroservice/ram_api";
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const RAMForm = (props) => {
	const customId = "RAMForm";

	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, ramFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${ramAPI.GET_BY_ID}/${id}`)
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
			.post(ramAPI.CREATE, data)
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
			.put(ramAPI.UPDATE, data)
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
															product !== null ? product.MemoryType : ""
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
													<Label>Capacity</Label>
													<Input
														className="input-field"
														type="text"
														name="Capacity"
														invalid={methods.formState.errors.Capacity?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.Capacity : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Capacity?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Speed</Label>
													<Input
														className="input-field"
														type="text"
														name="Speed"
														invalid={methods.formState.errors.Speed?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Speeed !== null
																? product.Speed
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Speed?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Voltage</Label>
													<Input
														className="input-field"
														type="text"
														name="Voltage"
														invalid={methods.formState.errors.Voltage?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Voltage !== null
																? product.Voltage
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Voltage?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Latency</Label>
													<Input
														className="input-field"
														type="text"
														name="Latency"
														invalid={methods.formState.errors.Latency?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Latency !== null
																? product.Latency
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Latency?.message}
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

export default RAMForm;
