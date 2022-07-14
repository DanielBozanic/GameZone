import {
	Button,
	Card,
	CardTitle,
	CardBody,
	Label,
	CardImg,
	Row,
	Col,
	Modal,
	ModalHeader,
	ModalBody,
	ModalFooter,
} from "reactstrap";
import { useState } from "react";
import { saveAs } from "file-saver";
import "../../Assets/css/checkout.css";
import paymentSlipImage from "../../Assets/images/payment_slip.png";

const ReviewCheckoutInfo = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);

	const downloadSlip = () => {
		saveAs(paymentSlipImage, "payment_slip.png");
	};

	const onPurchaseComplete = () => {
		props.purchaseComplete && props.purchaseComplete();
	};

	return (
		<>
			<Row>
				<Col>
					<Card className="checkout-card">
						<CardTitle className="checkout-card-title" tag="h2">
							Review checkout information
						</CardTitle>
						<CardBody>
							<Row>
								<Col>
									<Label style={{ fontWeight: "bold" }}>Delivery address</Label>
									: {props.buyerInfo.deliveryAddress}
								</Col>
							</Row>
							<Row>
								<Col md="2">
									<Label style={{ fontWeight: "bold" }}>City</Label>:{" "}
									{props.buyerInfo.city}
								</Col>
							</Row>
							<Row>
								<Col>
									<Label style={{ fontWeight: "bold" }}>
										Mobile phone number
									</Label>
									: {props.buyerInfo.mobilePhoneNumber}
								</Col>
							</Row>
							<Row>
								<Col>
									<Label style={{ fontWeight: "bold" }}>Payment method</Label>:{" "}
									{props.paymentType === 1 && <>Cash on delivery</>}
									{props.paymentType === 2 && <>Payment slip</>}
								</Col>
							</Row>
							<Row>
								<Col md="9">
									<Button
										type="button"
										className="checkout-btn"
										onClick={onPurchaseComplete}
									>
										Confirm purchase
									</Button>
								</Col>
								{props.paymentType === 2 && (
									<>
										<Col md="3">
											<Button
												type="button"
												className="checkout-btn"
												onClick={toggle}
											>
												View payment slip
											</Button>
										</Col>
									</>
								)}
							</Row>
						</CardBody>
					</Card>
				</Col>
			</Row>
			<Modal size="lg" isOpen={modal} toggle={toggle}>
				<ModalHeader className="payment-slip-modal-title" toggle={toggle}>
					Payment slip
				</ModalHeader>
				<ModalBody>
					<CardImg src={paymentSlipImage} />
				</ModalBody>
				<ModalFooter>
					<Button className="download-slip-btn" onClick={downloadSlip}>
						Download
					</Button>
				</ModalFooter>
			</Modal>
		</>
	);
};

export default ReviewCheckoutInfo;
