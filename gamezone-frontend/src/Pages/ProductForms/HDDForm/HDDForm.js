import { useForm, FormProvider } from "react-hook-form";
import { useState } from "react";
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
import axios from "axios";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "../../../Assets/css/forms.css";
import ProductForm from "../../../Components/ProductForm/ProductForm";
import * as hddAPI from "../../../APIs/ProductMicroservice/hard_disk_drive_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const HDDForm = (props) => {
	const customId = "hDDForm";

	const [base64Image, setBase64Image] = useState("");

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, hddFormSchema)
		),
		mode: "onChange",
	});

	const add = (data) => {
		data.Product.Image = base64Image;
		axios
			.post(hddAPI.CREATE, data)
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
												<Label>Disk Speed</Label>
												<Input
													className="input-field"
													type="text"
													name="DiskSpeed"
													invalid={methods.formState.errors.DiskSpeed?.message}
													innerRef={methods.register}
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
													invalid={methods.formState.errors.Interface?.message}
													innerRef={methods.register}
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

export default HDDForm;
