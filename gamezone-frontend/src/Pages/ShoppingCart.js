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
import * as productAPI from "../APIs/ProductMicroservice/product_api";
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
		axios
			.get(productAPI.GET_CURRENT_CART)
			.then((res) => {
				let temp = 0;
				for (let product of res.data) {
					temp += Number(product.TotalPrice);
				}
				setTotalPrice(temp);
				setShoppingCart(res.data);
				cartContainsOnlyDigitalItems();
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const cartContainsOnlyDigitalItems = () => {
		axios
			.get(productAPI.CART_CONTAINS_ONLY_DIGITAL_ITEMS)
			.then((res) => {
				setAllDigital(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const updateProductPurchase = (product, newAmount) => {
		product.Amount = newAmount;
		axios
			.put(productAPI.UPDATE_PURCHASE, product)
			.then((res) => {
				console.log(res.data);
				getShoppingCart();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const removeProductFromCart = (product) => {
		axios
			.delete(`${productAPI.REMOVE_PRODUCT_FROM_CART}/${product.Id}`)
			.then((res) => {
				console.log(res.data);
				getShoppingCart();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const checkout = () => {
		setConfirmedCheckout(true);
	};

	const purchaseComplete = () => {
		const finalPurchase = {
			DeliveryAddress: buyerInfo.deliveryAddress,
			City: buyerInfo.city,
			MobilePhoneNumber: buyerInfo.mobilePhoneNumber,
			TypeOfPayment: paymentType,
		};
		axios
			.put(productAPI.CONFIRM_PURCHASE, finalPurchase)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				navigate("/");
			})
			.catch((err) => {
				console.log(err);
			});
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
												{/* <td>
												<img className="product-img" src={productInCart.Product.Image} />
											</td> */}
												<td>{productInCart.Product.Name}</td>
												<td>{productInCart.Product.Price} RSD</td>
												<td>{productInCart.Amount}</td>
												<td>{productInCart.TotalPrice} RSD</td>
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
