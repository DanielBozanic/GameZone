import {
	Button,
	Form,
	FormGroup,
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
import { paymentTypeSchema } from "./PaymentTypeSchema";
import "../../../Assets/css/checkout.css";

const PaymentType = (props) => {
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(paymentTypeSchema), mode: "onChange" });

	const paymentStepDone = (data) => {
		props.setPaymentType && props.setPaymentType(Number(data.paymentType));
		props.setConfirmedPurchase && props.setConfirmedPurchase(true);
	};

	return (
		<>
			<Row>
				<Col>
					<Card className="checkout-card">
						<CardTitle className="checkout-card-title" tag="h2">
							Payment method
						</CardTitle>
						<CardBody>
							<Form>
								<Row>
									<Col>
										<FormGroup>
											<Input
												className="payment-method-type-select"
												name="paymentType"
												type="select"
												innerRef={register}
												invalid={errors.paymentType?.message}
											>
												<option value="" hidden>
													Select type of payment
												</option>
												{!props.allDigital && (
													<>
														<option value={1}>Cash on delivery</option>
														<option value={2}>Payment slip</option>
													</>
												)}
												{props.allDigital && (
													<>
														<option value={2}>Payment slip</option>
													</>
												)}
											</Input>
											<FormFeedback className="checkout-error-msg">
												{errors.paymentType?.message}
											</FormFeedback>
										</FormGroup>
									</Col>
								</Row>
								<Row>
									<Col>
										<Button
											type="button"
											className="next-step-checkout-btn"
											onClick={handleSubmit(paymentStepDone)}
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

export default PaymentType;
