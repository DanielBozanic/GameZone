import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
	CardTitle,
	Card,
	Row,
	Col,
	Container,
	Table,
	Button,
	Input,
} from "reactstrap";
import axios from "axios";
import "../Assets/css/shopping-cart.css";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import * as videoGameAPI from "../APIs/ProductMicroservice/video_game_api";
import * as productPurchaseAPI from "../APIs/ProductMicroservice/product_purchase_api";
import * as productType from "../Utils/ProductType";
import AppNavbar from "../Layout/AppNavbar";
import BuyerInfo from "../Components/Checkout/BuyerInfo/BuyerInfo";
import PaymentType from "../Components/Checkout/PaymentType/PaymentType";
import ReviewCheckoutInfo from "../Components/Checkout/ReviewCheckoutInfo";

toast.configure();
const ShoppingCart = () => {
	const customId = "shopping-cart";

	const [shoppingCart, setShoppingCart] = useState([]);
	const [totalPrice, setTotalPrice] = useState(0);
	const [confirmedCheckout, setConfirmedCheckout] = useState(false);

	const [buyerInfo, setBuyerInfo] = useState(null);
	const [paymentType, setPaymentType] = useState(null);

	const [allDigital, setAllDigital] = useState(false);

	const navigate = useNavigate();

	useEffect(() => {
		getShoppingCart();
	}, []);

	const getShoppingCart = () => {
		const cart = JSON.parse(localStorage.getItem("cart") || "[]");
		setShoppingCart(cart);
		calculateTotalPrice(cart);
		cartContainsOnlyDigitalItems(cart);
	};

	const calculateTotalPrice = (cart) => {
		let temp = 0;
		for (let product of cart) {
			temp += Number(product.Product.Product.Price) * Number(product.Quantity);
		}
		setTotalPrice(temp);
	};

	const cartContainsOnlyDigitalItems = (cart) => {
		let everyItemDigital = true;
		setAllDigital(everyItemDigital);
		let videoGames = cart.filter(
			(pic) => pic.Product.Product.Type === productType.VIDEO_GAME
		);
		let otherProducts = cart.filter(
			(pic) => pic.Product.Product.Type !== productType.VIDEO_GAME
		);
		if (otherProducts.length > 0) {
			everyItemDigital = false;
			setAllDigital(everyItemDigital);
		} else if (videoGames.length > 0) {
			for (let videoGame of videoGames) {
				axios
					.get(`${videoGameAPI.GET_BY_ID}/${videoGame.Product.Product.Id}`)
					.then((res) => {
						if (!res.data.Digital) {
							setAllDigital(false);
						}
					})
					.catch((err) => {
						console.log(err);
					});
			}
		} else {
			setAllDigital(false);
		}
	};

	const updateProductPurchase = (productInCart, newQuantity) => {
		productInCart.Quantity = newQuantity;
		const newCart = shoppingCart.filter(
			(product) =>
				product.Product.Product.Id !== productInCart.Product.Product.Id
		);
		newCart.push(productInCart);
		localStorage.setItem("cart", JSON.stringify(newCart));
		setShoppingCart(newCart);
		calculateTotalPrice(newCart);
	};

	const removeProductFromCart = (productInCart) => {
		const newCart = shoppingCart.filter(
			(product) =>
				product.Product.Product.Id !== productInCart.Product.Product.Id
		);
		if (newCart.length === 0) {
			setConfirmedCheckout(false);
			setBuyerInfo(null);
			setPaymentType(null);
		}
		localStorage.setItem("cart", JSON.stringify(newCart));
		setShoppingCart(newCart);
		calculateTotalPrice(newCart);
		cartContainsOnlyDigitalItems(newCart);
	};

	const checkout = () => {
		setConfirmedCheckout(true);
	};

	const purchaseComplete = () => {
		const productPurchaseDetail = [];
		for (let productInCart of shoppingCart) {
			productPurchaseDetail.push({
				ProductId: productInCart.Product.Product.Id,
				ProductName: productInCart.Product.Product.Name,
				ProductPrice: productInCart.Product.Product.Price,
				ProductQuantity: productInCart.Quantity,
			});
		}
		const finalPurchase = {
			ProductPurchaseDetail: productPurchaseDetail,
			DeliveryAddress: buyerInfo.deliveryAddress,
			City: buyerInfo.city,
			MobilePhoneNumber: buyerInfo.mobilePhoneNumber,
			TypeOfPayment: paymentType,
			TotalPrice: totalPrice,
		};
		axios
			.post(productPurchaseAPI.CONFIRM_PURCHASE, finalPurchase)
			.then((res) => {
				sendPurchaseConfirmationMail(finalPurchase);
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				navigate("/");
				localStorage.setItem("cart", JSON.stringify([]));
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const sendPurchaseConfirmationMail = (finalPurchase) => {
		axios.post(
			`${productPurchaseAPI.SEND_PURCHASE_CONFIRMATION_MAIL}`,
			finalPurchase
		);
	};

	return (
		<>
			<AppNavbar />
			<Container>
				<Row>
					<Col style={{ paddingTop: "5px" }}>
						<Card className="shopping-cart-card">
							<CardTitle className="shopping-cart-card-title" tag="h2">
								My Cart
							</CardTitle>
							<Table className="shopping-cart-table">
								<tr>
									<th>Name</th>
									<th>Price</th>
									<th>Quantity</th>
									<th>Amount</th>
								</tr>
								<tbody>
									{shoppingCart.map((productInCart, idx) => {
										return (
											<tr key={idx}>
												<td>{productInCart.Product.Product.Name}</td>
												<td>{productInCart.Product.Product.Price} RSD</td>
												<td>{productInCart.Quantity}</td>
												<td>
													{Number(productInCart.Product.Product.Price) *
														Number(productInCart.Quantity)}{" "}
													RSD
												</td>
												<td>
													<Input
														className="amount-select"
														type="select"
														onChange={(e) =>
															updateProductPurchase(
																productInCart,
																Number(e.target.value)
															)
														}
													>
														<option hidden>Select quantity</option>
														<option>1</option>
														<option>2</option>
														<option>3</option>
														<option>4</option>
														<option>5</option>
													</Input>
												</td>
												<td>
													<Button
														type="button"
														className="remove-product-from-cart-btn"
														onClick={() => removeProductFromCart(productInCart)}
													>
														Remove
													</Button>
												</td>
											</tr>
										);
									})}
								</tbody>
								<tfoot>
									<tr>
										<td className="total-price-td">Total: {totalPrice} RSD</td>
									</tr>
								</tfoot>
							</Table>
							<Row>
								<Col md={10}>
									<Button
										type="button"
										className="checkout-btn"
										disabled={shoppingCart.length === 0 ? true : false}
										onClick={checkout}
									>
										Checkout
									</Button>
								</Col>
							</Row>
						</Card>
					</Col>
				</Row>
				{confirmedCheckout && <BuyerInfo setBuyerInfo={setBuyerInfo} />}
				{buyerInfo !== null && (
					<PaymentType
						allDigital={allDigital}
						setPaymentType={setPaymentType}
					/>
				)}
				{buyerInfo != null && paymentType !== null && (
					<ReviewCheckoutInfo
						buyerInfo={buyerInfo}
						paymentType={paymentType}
						purchaseComplete={purchaseComplete}
					/>
				)}
			</Container>
		</>
	);
};

export default ShoppingCart;
