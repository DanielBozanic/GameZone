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
	Row,
	Col,
} from "reactstrap";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { buyerInfoSchema } from "./BuyerInfoSchema";

const BuyerInfo = (props) => {
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(buyerInfoSchema), mode: "onChange" });

	const buyerInfoDone = (data) => {
		props.setBuyerInfo && props.setBuyerInfo(data);
	};

	return (
		<>
			<Row>
				<Col>
					<Card className="card shopping-cart-card">
						<CardTitle className="title" tag="h2">
							Buyer information
						</CardTitle>
						<CardBody>
							<Form className="form">
								<Row>
									<Col>
										<FormGroup>
											<Label>Delivery address</Label>
											<Input
												name="deliveryAddress"
												type="text"
												className="input-field"
												innerRef={register}
												invalid={errors.deliveryAddress?.message}
											/>
											<FormFeedback className="input-field-error-msg">
												{errors.deliveryAddress?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<FormGroup>
											<Label>City</Label>
											<Input
												name="city"
												type="text"
												className="input-field"
												innerRef={register}
												invalid={errors.city?.message}
											/>
											<FormFeedback className="input-field-error-msg">
												{errors.city?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<FormGroup>
											<Label>Mobile phone number</Label>
											<Input
												name="mobilePhoneNumber"
												type="text"
												className="input-field"
												innerRef={register}
												invalid={errors.mobilePhoneNumber?.message}
											/>
											<FormFeedback className="input-field-error-msg">
												{errors.mobilePhoneNumber?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<Button
											type="button"
											className="my-button"
											onClick={handleSubmit(buyerInfoDone)}
										>
											Confirm
										</Button>
									</Col>
								</Row>
							</Form>
						</CardBody>
					</Card>
				</Col>
			</Row>
		</>
	);
};

export default BuyerInfo;
