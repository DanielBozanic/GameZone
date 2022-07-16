import { useForm, FormProvider } from "react-hook-form";
import { useState } from "react";
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
import * as monitorAPI from "../../../APIs/ProductMicroservice/monitor_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const MonitorForm = (props) => {
	const customId = "monitorForm";
	const [base64Image, setBase64Image] = useState("");

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, monitorFormSchema)
		),
		mode: "onChange",
	});

	const add = (data) => {
		data.Product.Image = base64Image;
		axios
			.post(monitorAPI.CREATE, data)
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
												<Label>Resolution</Label>
												<Input
													className="input-field"
													type="text"
													name="Resolution"
													invalid={methods.formState.errors.Resolution?.message}
													innerRef={methods.register}
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
													invalid={methods.formState.errors.PanelType?.message}
													innerRef={methods.register}
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
													invalid={methods.formState.errors.Brightness?.message}
													innerRef={methods.register}
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

export default MonitorForm;
