import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { hddFormSchema } from "./HDDFormSchema";
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
import * as hddAPI from "../../../APIs/ProductMicroservice/hard_disk_drive_api";
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const HDDForm = (props) => {
	const customId = "HDDForm";

	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, hddFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${hddAPI.GET_BY_ID}/${id}`)
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
			.post(hddAPI.CREATE, data)
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
			.put(hddAPI.UPDATE, data)
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
													<Label>Disk Speed</Label>
													<Input
														className="input-field"
														type="text"
														name="DiskSpeed"
														invalid={
															methods.formState.errors.DiskSpeed?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.DiskSpeed !== null
																? product.DiskSpeed
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.DiskSpeed?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Interface</Label>
													<Input
														className="input-field"
														type="text"
														name="Interface"
														invalid={
															methods.formState.errors.Interface?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Interface !== null
																? product.Interface
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Interface?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Transfer rate</Label>
													<Input
														className="input-field"
														type="text"
														name="TransferRate"
														invalid={
															methods.formState.errors.TransferRate?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.TransferRate !== null
																? product.TransferRate
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.TransferRate?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Form</Label>
													<Input
														className="input-field"
														type="text"
														name="Form"
														invalid={methods.formState.errors.Form?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Form !== null
																? product.Form
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Form?.message}
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

export default HDDForm;
