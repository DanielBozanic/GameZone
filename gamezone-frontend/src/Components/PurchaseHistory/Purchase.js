import "../../Assets/css/purchase-history.css";
import {
	CardText,
	CardHeader,
	CardBody,
	Card,
	CardTitle,
	CardFooter,
	Row,
	Col,
	Label,
	Button,
} from "reactstrap";
import { useState } from "react";
import axios from "axios";
import { toast } from "react-toastify";
import * as authService from "../../Auth/AuthService";
import * as productPurchaseAPI from "../../APIs/ProductMicroservice/product_purchase_api";
import PurchaseDetail from "./PurchaseDetail";

toast.configure();
const Purchase = (props) => {
	const customId = "Purchase";

	const [isOpen, setIsOpen] = useState(false);
	const toggleItem = () => setIsOpen(!isOpen);

	const confirmPayment = () => {
		axios
			.put(`${productPurchaseAPI.CONFIRM_PAYMENT}`, props.purchase)
			.then((_res) => {
				sendPurchasedDigitalVideoGames();
				props.getPurchaseHistory && props.getPurchaseHistory();
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: false,
				});
			});
	};

	const sendPurchasedDigitalVideoGames = () => {
		axios.post(
			`${productPurchaseAPI.SEND_PURCHASED_DIGITAL_VIDEO_GAMES}`,
			props.purchase
		);
	};

	return (
		<Row>
			<Col>
				<Card className="card">
					<CardHeader>
						<Row>
							<Col md="10">
								<CardTitle tag="h5">Purchase ID #{props.purchase.Id}</CardTitle>
							</Col>
							<Col md="2">
								<CardTitle tag="h5">
									{new Date(props.purchase.PurchaseDate).toLocaleDateString()}
								</CardTitle>
							</Col>
						</Row>
					</CardHeader>
					<CardBody>
						<CardText>
							<Row>
								<Col>
									<CardText>
										<Label style={{ fontWeight: "bold" }}>
											Delivery Address:
										</Label>{" "}
										{props.purchase.DeliveryAddress}
									</CardText>
								</Col>
							</Row>
							<Row>
								<Col>
									<CardText>
										<Label style={{ fontWeight: "bold" }}>City:</Label>{" "}
										{props.purchase.City}
									</CardText>
								</Col>
							</Row>
							<Row>
								<Col>
									<CardText>
										<Label style={{ fontWeight: "bold" }}>
											Mobile Phone Number:
										</Label>{" "}
										{props.purchase.MobilePhoneNumber}
									</CardText>
								</Col>
							</Row>
							<Row>
								<Col>
									<CardText>
										<Label style={{ fontWeight: "bold" }}>
											Total Purchase Price:
										</Label>{" "}
										{props.purchase.TotalPrice} RSD
									</CardText>
								</Col>
							</Row>
						</CardText>
					</CardBody>
					<CardFooter>
						<Row>
							<Button
								style={{
									marginRight: "5px",
									marginLeft: "5px",
								}}
								className="my-button"
								type="button"
								onClick={toggleItem}
							>
								View purchased items
							</Button>
							{!props.purchase.IsPaidFor && authService.isAdmin() && (
								<Button
									className="my-button"
									type="button"
									onClick={confirmPayment}
								>
									Confirm payment
								</Button>
							)}
							{props.purchase.IsPaidFor && (
								<div
									style={{
										marginLeft: "10px",
										marginTop: "7px",
										fontWeight: "bold",
									}}
								>
									Purchase complete
								</div>
							)}
							{!props.purchase.IsPaidFor && (
								<div
									style={{
										marginLeft: "10px",
										marginTop: "7px",
										marginRight: "10px",
										fontWeight: "bold",
									}}
								>
									Payment incomplete
								</div>
							)}
						</Row>
					</CardFooter>
					<PurchaseDetail
						isOpen={isOpen}
						toggleItem={toggleItem}
						purchase={props.purchase}
					/>
				</Card>
			</Col>
		</Row>
	);
};

export default Purchase;
