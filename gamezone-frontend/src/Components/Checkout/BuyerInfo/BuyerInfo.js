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
import "../../../Assets/css/checkout.css";

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
					<Card className="checkout-card">
						<CardTitle className="checkout-card-title" tag="h2">
							Buyer information
						</CardTitle>
						<CardBody>
							<Form>
								<Row>
									<Col>
										<FormGroup>
											<Label>Delivery address</Label>
											<Input
												name="deliveryAddress"
												type="text"
												className="checkout-input-field"
												innerRef={register}
												invalid={errors.deliveryAddress?.message}
											/>
											<FormFeedback className="checkout-error-msg">
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
												className="checkout-input-field"
												innerRef={register}
												invalid={errors.city?.message}
											/>
											<FormFeedback className="checkout-error-msg">
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
												className="checkout-input-field"
												innerRef={register}
												invalid={errors.mobilePhoneNumber?.message}
											/>
											<FormFeedback className="checkout-error-msg">
												{errors.mobilePhoneNumber?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<Button
											type="button"
											className="next-step-checkout-btn"
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
