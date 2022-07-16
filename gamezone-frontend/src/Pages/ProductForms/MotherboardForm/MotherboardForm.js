import { useForm, FormProvider } from "react-hook-form";
import { useState } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { motherboardFormSchema } from "./MotherboardFormSchema";
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
import * as motherboardAPI from "../../../APIs/ProductMicroservice/motherboard_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const MotherboardForm = (props) => {
	const customId = "motherboardForm";
	const [base64Image, setBase64Image] = useState("");

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, motherboardFormSchema)
		),
		mode: "onChange",
	});

	const add = (data) => {
		data.Product.Image = base64Image;
		axios
			.post(motherboardAPI.CREATE, data)
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
												<Label>Processor Type</Label>
												<Input
													className="input-field"
													type="text"
													name="ProcessorType"
													invalid={
														methods.formState.errors.ProcessorType?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.ProcessorType?.message}
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
												<Label>Supported Processors</Label>
												<Input
													className="input-field"
													type="textarea"
													name="SupportedProcessors"
													invalid={
														methods.formState.errors.SupportedProcessors
															?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{
														methods.formState.errors.SupportedProcessors
															?.message
													}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Chipset</Label>
												<Input
													className="input-field"
													type="text"
													name="Chipset"
													invalid={methods.formState.errors.Chipset?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Chipset?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Memory</Label>
												<Input
													className="input-field"
													type="textarea"
													name="Memory"
													invalid={methods.formState.errors.Memory?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Memory?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Expansion Slots</Label>
												<Input
													className="input-field"
													type="textarea"
													name="ExpansionSlots"
													invalid={
														methods.formState.errors.ExpansionSlots?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.ExpansionSlots?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Storage Interface</Label>
												<Input
													className="input-field"
													type="textarea"
													name="StorageInterface"
													invalid={
														methods.formState.errors.StorageInterface?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.StorageInterface?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Audio</Label>
												<Input
													className="input-field"
													type="textarea"
													name="Audio"
													invalid={methods.formState.errors.Audio?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.Audio?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>USB</Label>
												<Input
													className="input-field"
													type="textarea"
													name="USB"
													invalid={methods.formState.errors.USB?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.USB?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Back Panel Connectors</Label>
												<Input
													className="input-field"
													type="textarea"
													name="BackPanelConnectors"
													invalid={
														methods.formState.errors.BackPanelConnectors
															?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{
														methods.formState.errors.BackPanelConnectors
															?.message
													}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Internal Connectors</Label>
												<Input
													className="input-field"
													type="textarea"
													name="InternalConnectors"
													invalid={
														methods.formState.errors.InternalConnectors?.message
													}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.InternalConnectors?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>BIOS</Label>
												<Input
													className="input-field"
													type="textarea"
													name="BIOS"
													invalid={methods.formState.errors.BIOS?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.BIOS?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Form Factor</Label>
												<Input
													className="input-field"
													type="text"
													name="FormFactor"
													invalid={methods.formState.errors.FormFactor?.message}
													innerRef={methods.register}
												/>
												<FormFeedback className="input-field-error-msg">
													{methods.formState.errors.FormFactor?.message}
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

export default MotherboardForm;
