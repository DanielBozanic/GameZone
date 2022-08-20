import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { monitorFormSchema } from "./MonitorFormSchema";
import {
	Button,
	Form,
	FormGroup,
	Label,
	Input,
	Card,
	CardTitle,
	CardHeader,
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
import * as monitorAPI from "../../../APIs/ProductMicroservice/monitor_api";
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const MonitorForm = (props) => {
	const customId = "monitorForm";

	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, monitorFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${monitorAPI.GET_BY_ID}/${id}`)
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
			.post(monitorAPI.CREATE, data)
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
			.put(monitorAPI.UPDATE, data)
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
													<Label>Resolution</Label>
													<Input
														className="input-field"
														type="text"
														name="Resolution"
														invalid={
															methods.formState.errors.Resolution?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.Resolution : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Resolution?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Refresh Rate</Label>
													<Input
														className="input-field"
														type="text"
														name="RefreshRate"
														invalid={
															methods.formState.errors.RefreshRate?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.RefreshRate : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.RefreshRate?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Size</Label>
													<Input
														className="input-field"
														type="text"
														name="Size"
														invalid={methods.formState.errors.Size?.message}
														innerRef={methods.register}
														defaultValue={product !== null ? product.Size : ""}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Size?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Aspect Ratio</Label>
													<Input
														className="input-field"
														type="text"
														name="AspectRatio"
														invalid={
															methods.formState.errors.AspectRatio?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.AspectRatio !== null
																? product.AspectRatio
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.AspectRatio?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Contrast Ratio</Label>
													<Input
														className="input-field"
														type="text"
														name="ContrastRatio"
														invalid={
															methods.formState.errors.ContrastRatio?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.ContrastRatio !== null
																? product.ContrastRatio
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ContrastRatio?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Response Time</Label>
													<Input
														className="input-field"
														type="text"
														name="ResponseTime"
														invalid={
															methods.formState.errors.ResponseTime?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.ResponseTime !== null
																? product.ResponseTime
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ResponseTime?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Panel Type</Label>
													<Input
														className="input-field"
														type="text"
														name="PanelType"
														invalid={
															methods.formState.errors.PanelType?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.PanelType !== null
																? product.PanelType
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.PanelType?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Viewing Angle</Label>
													<Input
														className="input-field"
														type="text"
														name="ViewingAngle"
														invalid={
															methods.formState.errors.ViewingAngle?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.ViewingAngle !== null
																? product.ViewingAngle
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ViewingAngle?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Brightness</Label>
													<Input
														className="input-field"
														type="text"
														name="Brightness"
														invalid={
															methods.formState.errors.Brightness?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.Brightness !== null
																? product.Brightness
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Brightness?.message}
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

export default MonitorForm;
