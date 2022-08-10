import { useForm, FormProvider } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { createFormSchema } from "../Components/UserForms/CreateForm/CreateFormSchema";
import CreateForm from "../Components/UserForms/CreateForm/CreateForm";
import {
	Button,
	Form,
	Card,
	CardTitle,
	CardBody,
	Container,
	Row,
	Col,
	FormGroup,
	Label,
	Input,
	Spinner,
} from "reactstrap";
import { useState } from "react";
import axios from "axios";
import { toast } from "react-toastify";
import * as userAPI from "../APIs/UserMicroservice/user_api";
import * as role from "../Utils/Role";

toast.configure();
const CreateAdminAndEmployee = () => {
	const customId = "CreateAdminAndEmployeeForm";
	const [loading, setLoading] = useState(false);

	const methods = useForm({
		resolver: yupResolver(createFormSchema),
		mode: "onChange",
	});

	const add = (data) => {
		setLoading(true);
		axios
			.post(userAPI.ADD_EMPLOYEE_AND_ADMIN, data, {
				headers: {
					"Content-Type": "application/json",
				},
			})
			.then((res) => {
				toast.success(res.data.message, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				setLoading(false);
				methods.reset();
			})
			.catch((err) => {
				toast.error(err.response.data.message, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
				setLoading(false);
			});
	};

	return (
		<Container>
			<Row>
				<Col md="10">
					<Card className="form-card">
						<CardTitle className="title" tag="h2">
							Add employee/admin
						</CardTitle>
						{loading && (
							<div className="div-spinner">
								<Spinner className="spinner" />
							</div>
						)}
						{!loading && (
							<CardBody>
								<FormProvider {...methods}>
									<Form className="form">
										<CreateForm />
										<Row>
											<Col>
												<FormGroup>
													<Label>Role</Label>
													<Input
														className="input-field"
														name="role"
														type="select"
														innerRef={methods.register}
													>
														<option value={role.ROLE_EMPLOYEE}>Employee</option>
														<option value={role.ROLE_ADMIN}>Admin</option>
													</Input>
												</FormGroup>
											</Col>
										</Row>
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
									</Form>
								</FormProvider>
							</CardBody>
						)}
					</Card>
				</Col>
			</Row>
		</Container>
	);
};

export default CreateAdminAndEmployee;
