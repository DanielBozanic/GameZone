import { useParams } from "react-router-dom";
import { useForm, FormProvider } from "react-hook-form";
import { useState, useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { processorFormSchema } from "./ProcessorFormSchema";
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
import * as processorAPI from "../../../APIs/ProductMicroservice/processor_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const ProcessorForm = (props) => {
	const customId = "processorForm";

	const [base64Image, setBase64Image] = useState("");
	const [fileName, setFileName] = useState("");
	const [product, setProduct] = useState(null);

	const { id } = useParams();

	useEffect(() => {
		getProductById();
	}, []);

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, processorFormSchema)
		),
		mode: "onChange",
	});

	const getProductById = () => {
		if (id !== undefined) {
			axios
				.get(`${processorAPI.GET_BY_ID}/${id}`)
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
			.post(processorAPI.CREATE, data)
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
			.put(processorAPI.UPDATE, data)
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
											<FormGroup>
												<Label>Type</Label>
												<Input
													className="input-field"
													type="text"
													name="Type"
													invalid={methods.formState.errors.Type?.message}
													innerRef={methods.register}
													defaultValue={product !== null ? product.Type : ""}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Type?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Socket</Label>
												<Input
													className="input-field"
													type="text"
													name="Socket"
													invalid={methods.formState.errors.Socket?.message}
													innerRef={methods.register}
													defaultValue={product !== null ? product.Socket : ""}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Socket?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Number Of Cores</Label>
												<Input
													className="input-field"
													type="number"
													name="NumberOfCores"
													min="1"
													invalid={
														methods.formState.errors.NumberOfCores?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.NumberOfCores !== null
															? product.NumberOfCores
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.NumberOfCores?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Threads</Label>
												<Input
													className="input-field"
													type="number"
													name="Threads"
													min="1"
													invalid={methods.formState.errors.Threads?.message}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.Threads !== null
															? product.Threads
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Threads?.message}
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
												<Label>Integrated Graphics</Label>
												<Input
													className="input-field"
													type="text"
													name="IntegratedGraphics"
													invalid={
														methods.formState.errors.IntegratedGraphics?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null &&
														product.IntegratedGraphics !== null
															? product.IntegratedGraphics
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.IntegratedGraphics?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Base Clock Rate</Label>
												<Input
													className="input-field"
													type="text"
													name="BaseClockRate"
													invalid={
														methods.formState.errors.BaseClockRate?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.BaseClockRate !== null
															? product.BaseClockRate
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.BaseClockRate?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Turbo Clock Rate</Label>
												<Input
													className="input-field"
													type="text"
													name="TurboClockRate"
													invalid={
														methods.formState.errors.TurboClockRate?.message
													}
													innerRef={methods.register}
													defaultValue={
														product !== null && product.TurboClockRate !== null
															? product.TurboClockRate
															: ""
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.TurboClockRate?.message}
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

export default ProcessorForm;
