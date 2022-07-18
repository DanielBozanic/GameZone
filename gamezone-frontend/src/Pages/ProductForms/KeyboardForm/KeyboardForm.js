import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { keyboardFormSchema } from "./KeyboardFormSchema";
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
import * as keyboardAPI from "../../../APIs/ProductMicroservice/keyboard_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const KeyboardForm = (props) => {
	const customId = "keyboardForm";

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
			helperFunctions.merge(productFormSchema, keyboardFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${keyboardAPI.GET_BY_ID}/${id}`)
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
			.post(keyboardAPI.CREATE, data)
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
			.put(keyboardAPI.UPDATE, data)
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
												<Label>Keyboard Connector</Label>
												<Input
													className="input-field"
													type="text"
													name="KeyboardConnector"
													invalid={
														methods.formState.errors.KeyboardConnector?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null ? product.KeyboardConnector : ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.KeyboardConnector?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Key Type</Label>
												<Input
													className="input-field"
													type="text"
													name="KeyType"
													invalid={methods.formState.errors.KeyType?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.KeyType !== null
															? product.KeyType
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.KeyType?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Letter Layout</Label>
												<Input
													className="input-field"
													type="text"
													name="LetterLayout"
													invalid={
														methods.formState.errors.LetterLayout?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.LetterLayout !== null
															? product.LetterLayout
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.LetterLayout?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Keyboard Color</Label>
												<Input
													className="input-field"
													type="text"
													name="KeyboardColor"
													invalid={
														methods.formState.errors.KeyboardColor?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.KeyboardColor !== null
															? product.KeyboardColor
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.KeyboardColor?.message}
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

export default KeyboardForm;
