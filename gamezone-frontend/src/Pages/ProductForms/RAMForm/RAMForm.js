import { useForm, FormProvider } from "react-hook-form";
import { useState } from "react";
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
import axios from "axios";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "../../../Assets/css/forms.css";
import ProductForm from "../../../Components/ProductForm/ProductForm";
import * as ramAPI from "../../../APIs/ProductMicroservice/ram_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const RAMForm = (props) => {
	const customId = "RAMForm";

	const [base64Image, setBase64Image] = useState("");

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, ramFormSchema)
		),
		mode: "onChange",
	});

	const add = (data) => {
		data.Product.Image = base64Image;
		axios
			.post(ramAPI.CREATE, data)
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
									<ProductForm setBase64Image={setBase64Image} />
									<Row>
										<Col>
											<FormGroup>
												<Label>Memory Type</Label>
												<Input
													className="input-field"
													type="text"
													name="MemoryType"
													invalid={methods.formState.errors.MemoryType?.message}
													innerRef={methods.register}
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
													onClick={methods.handleSubmit(add)}
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

export default RAMForm;
