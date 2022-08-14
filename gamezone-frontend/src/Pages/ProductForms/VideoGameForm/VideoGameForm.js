import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { videoGameFormSchema } from "./VideoGameFormSchema";
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
import * as videoGameAPI from "../../../APIs/ProductMicroservice/video_game_api";
import * as productPurchaseAPI from "../../../APIs/ProductMicroservice/product_purchase_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const VideoGameForm = (props) => {
	const customId = "videoGameForm";

	const [digital, setDigital] = useState(false);
	const [isDigital, setIsDigital] = useState(false);
	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, videoGameFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${videoGameAPI.GET_BY_ID}/${id}`)
				.then((res) => {
					setDigital(null);
					setIsDigital(res.data.Digital);
					setProduct(res.data);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	const add = (data) => {
		data.Product.Image.Content = base64Image;
		data.Digital = helperFunctions.str2Bool(data.Digital);
		data.ReleaseDate = new Date(data.ReleaseDate);
		axios
			.post(videoGameAPI.CREATE, data)
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
		data.Digital = helperFunctions.str2Bool(data.Digital);
		data.ReleaseDate = new Date(data.ReleaseDate);
		axios
			.put(videoGameAPI.UPDATE, data)
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
											isDigital={isDigital}
											fileName={fileName}
											setFileName={setFileName}
											setBase64Image={setBase64Image}
										/>
										<Row>
											<Col>
												<FormGroup>
													<Label>Platform</Label>
													<Input
														className="input-field"
														type="text"
														name="Platform"
														invalid={methods.formState.errors.Platform?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.Platform : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Platform?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<div>
													<Label>Digital</Label>
												</div>
												<span>
													<Label>Yes</Label>
													<Input
														className="ml-2"
														type="radio"
														name="Digital"
														checked={
															product === null || digital !== null
																? digital
																: product.Digital
														}
														value={
															product === null || digital !== null
																? digital
																: product.Digital
														}
														innerRef={methods.register}
														onChange={() => {
															setIsDigital(true);
															setDigital(true);
														}}
													/>
												</span>
												<span className="pl-5">
													<Label>No</Label>
													<Input
														className="ml-2"
														type="radio"
														name="Digital"
														checked={
															product === null || digital !== null
																? !digital
																: !product.Digital
														}
														value={
															product === null || digital !== null
																? digital
																: product.Digital
														}
														innerRef={methods.register}
														onChange={() => {
															setIsDigital(false);
															setDigital(false);
														}}
													/>
												</span>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Rating</Label>
													<Input
														className="input-field"
														type="number"
														name="Rating"
														min="1"
														invalid={methods.formState.errors.Rating?.message}
														innerRef={methods.register}
														defaultValue={
															product !== null ? product.Rating : ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Rating?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Genre</Label>
													<Input
														className="input-field"
														type="text"
														name="Genre"
														invalid={methods.formState.errors.Genre?.message}
														innerRef={methods.register}
														defaultValue={product !== null ? product.Genre : ""}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Genre?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Release Date</Label>
													<Input
														className="input-field"
														type="date"
														name="ReleaseDate"
														invalid={
															methods.formState.errors.ReleaseDate?.message
														}
														innerRef={methods.register}
														defaultValue={
															product !== null && product.ReleaseDate !== null
																? product.ReleaseDate.toLocaleString().substring(
																		0,
																		10
																  )
																: ""
														}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ReleaseDate?.message}
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

export default VideoGameForm;
