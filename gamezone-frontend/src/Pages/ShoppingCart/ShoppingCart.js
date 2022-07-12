import { useState, useEffect } from "react";
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
import "../../Assets/css/shopping-cart.css";
import * as productAPI from "../../APIs/ProductMicroservice/product_api";
import AppNavbar from "../../Layout/AppNavbar";

const ShoppingCart = () => {
	const [shoppingCart, setShoppingCart] = useState([]);
	const [totalPrice, setTotalPrice] = useState(0);

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

	const checkout = () => {};

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
								</tr>
								{shoppingCart.map((product, idx) => {
									return (
										<tr key={idx}>
											{/* <td>
												<img className="product-img" src={product.ProductImage} />
											</td> */}
											<td>{product.ProductName}</td>
											<td>{product.ProductPrice} RSD</td>
											<td>{product.Amount}</td>
											<td>
												<Input
													className="amount-select"
													type="select"
													onChange={(e) =>
														updateProductPurchase(
															product,
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
													onClick={() => removeProductFromCart(product)}
												>
													Remove
												</Button>
											</td>
										</tr>
									);
								})}
								<tr>
									<td className="total-price-td">Total: {totalPrice} RSD</td>
								</tr>
							</Table>
							<Button
								type="button"
								className="checkout-btn"
								disabled={shoppingCart.length === 0 ? true : false}
								onClick={checkout}
							>
								Checkout
							</Button>
						</Card>
					</Col>
				</Row>
			</Container>
		</>
	);
};

export default ShoppingCart;
