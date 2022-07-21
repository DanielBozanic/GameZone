import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import {
	CardText,
	CardImg,
	CardTitle,
	CardBody,
	CardFooter,
	Card,
	Row,
	Col,
	Container,
	Table,
	Button,
	Input,
} from "reactstrap";
import axios from "axios";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "../Assets/css/product-detail.css";
import AppNavbar from "../Layout/AppNavbar";
import * as productAPI from "../APIs/ProductMicroservice/product_api";
import * as productType from "../Utils/ProductType";
import * as authService from "../Auth/AuthService";
import * as role from "../Utils/Role";
import * as helperFunctions from "../Utils/HelperFunctions";

toast.configure();
const ProductDetail = (props) => {
	const customId = "product-detail";

	const { id } = useParams();
	const [product, setProduct] = useState(null);
	const [amount, setAmount] = useState(1);
	const [available, setAvailable] = useState("");
	const [disableAddToCart, setDisableAddToCart] = useState(false);
	const [disableNotify, setDisableNotify] = useState(false);

	const navigate = useNavigate();

	useEffect(() => {
		getProductById();
		getProductAlertByProductIdAndEmail();
	}, []);

	const getProductById = () => {
		axios
			.get(`${props.GET_PRODUCT_BY_ID}/${id}`)
			.then((res) => {
				setProduct(res.data);
				pageSetup(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getProductAlertByProductIdAndEmail = () => {
		axios
			.get(
				`${productAPI.GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_EMAIL}?productId=${id}`
			)
			.then((res) => {
				if (res.data.ProductId === Number(id)) {
					setDisableNotify(true);
				}
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const pageSetup = (p) => {
		if (p.Product.Type === productType.VIDEO_GAME) {
			if (p.Digital || p.Product.Amount > 0) {
				setAvailable("This product is available for purchase");
				setDisableNotify(true);
			} else {
				setAvailable("This product is unavailable for purchase");
				setDisableAddToCart(true);
			}
		} else {
			if (p.Product.Amount > 0) {
				setAvailable("This product is available for purchase");
				setDisableNotify(true);
			} else {
				setAvailable("This product is unavailable for purchase");
				setDisableAddToCart(true);
			}
		}

		if (
			authService.getToken() == null ||
			(authService.getToken() != null &&
				authService.getRole() !== role.ROLE_USER)
		) {
			setDisableAddToCart(true);
			setDisableNotify(true);
		}
	};

	const addToCart = () => {
		const productPurchase = {
			Product: { Id: product.Product.Id },
			Amount: amount,
		};
		axios
			.post(productAPI.ADD_PRODUCT_TO_CART, productPurchase)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const updateProduct = () => {
		const route = helperFunctions.getUpdateProductRoute(product);
		navigate(route);
	};

	const deleteProduct = () => {
		axios
			.delete(`${productAPI.DELETE_PRODUCT}/${product.Product.Id}`)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				const route = helperFunctions.getProductListRoute(product);
				navigate(route);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const addProductAlert = () => {
		axios
			.post(`${productAPI.ADD_PRODUCT_ALERT}?productId=${product.Product.Id}`)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			<AppNavbar />
			{product !== null && (
				<Container>
					<Row className="product-detail-card">
						<Col style={{ marginTop: "10px", marginBottom: "10px" }} md={5}>
							<Card className="product-detail-card-with-image">
								<CardImg
									className="product-detail-card-image"
									src={product.Product.Image.Content}
									alt="No image"
								/>
								<CardBody>
									<CardTitle tag="h5">{product.Product.Name}</CardTitle>
									<CardText>
										{product.Product.Price}
										RSD
									</CardText>
									<CardText>{available}</CardText>
								</CardBody>
							</Card>
							{!disableAddToCart && (
								<>
									<Input
										className="amount-product-select"
										type="select"
										onChange={(e) => setAmount(Number(e.target.value))}
									>
										<option hidden>Select quantity</option>
										<option>1</option>
										<option>2</option>
										<option>3</option>
										<option>4</option>
										<option>5</option>
									</Input>
									<Button
										className="add-to-cart-btn"
										type="button"
										onClick={addToCart}
									>
										Add to cart
									</Button>
								</>
							)}
							{!disableNotify && (
								<>
									<Button
										className="notify-btn"
										type="button"
										onClick={addProductAlert}
									>
										Alert Me When In Stock
									</Button>
								</>
							)}
							{authService.getToken() != null &&
								authService.getRole() === role.ROLE_EMPLOYEE && (
									<>
										<Button
											className="update-btn"
											type="button"
											onClick={updateProduct}
										>
											Update
										</Button>
										<Button
											className="delete-btn"
											type="button"
											onClick={deleteProduct}
										>
											Delete
										</Button>
									</>
								)}
						</Col>
						{product !== null && (
							<>
								<Col>
									<Card
										style={{ marginTop: "10px", marginBottom: "10px" }}
										className="product-detail-description-card"
									>
										<CardTitle
											className="product-detail-description-card-title"
											tag="h4"
										>
											Description
										</CardTitle>
										<CardBody>{product.Product.Description}</CardBody>
										<CardFooter>
											More information on the manufacturer's website
										</CardFooter>
									</Card>
									<Card
										style={{ marginTop: "20px", marginBottom: "10px" }}
										className="product-detail-table-card"
									>
										<Table className="product-detail-table">
											<tr>
												<th>Manufacturer</th>
												<td>{product.Product.Manufacturer}</td>
											</tr>
											{Object.keys(product).map(function (value, idx) {
												if (
													value !== "Product" &&
													product[value] !== null &&
													product[value] !== ""
												) {
													if (typeof product[value] == "boolean") {
														return (
															<tr key={idx}>
																<th>
																	{value
																		.replace(/([A-Z]+)/g, " $1")
																		.replace(/([A-Z][a-z])/g, " $1")
																		.trim()}
																</th>
																<td>{product[value] ? "Yes" : "No"}</td>
															</tr>
														);
													} else if (
														new Date(product[value]) !== "Invalid Date"
													) {
														return (
															<tr key={idx}>
																<th>
																	{value
																		.replace(/([A-Z]+)/g, " $1")
																		.replace(/([A-Z][a-z])/g, " $1")
																		.trim()}
																</th>
																<td>
																	{product[value].toString().split("T")[0]}
																</td>
															</tr>
														);
													} else {
														return (
															<tr key={idx}>
																<th>
																	{value
																		.replace(/([A-Z]+)/g, " $1")
																		.replace(/([A-Z][a-z])/g, " $1")
																		.trim()}
																</th>
																<td>{product[value].toString()}</td>
															</tr>
														);
													}
												}
											})}
										</Table>
									</Card>
								</Col>
							</>
						)}
					</Row>
				</Container>
			)}
		</>
	);
};

export default ProductDetail;
