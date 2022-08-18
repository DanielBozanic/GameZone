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
import { Helmet } from "react-helmet";
import { toast } from "react-toastify";
import "../Assets/css/product-detail.css";
import AppNavbar from "../Layout/AppNavbar";
import CommentRating from "../Components/CommentRating/CommentRating";
import * as productAPI from "../APIs/ProductMicroservice/product_api";
import * as productPurchaseAPI from "../APIs/ProductMicroservice/product_purchase_api";
import * as productCommentAPI from "../APIs/CommentAndRatingMicroservice/product_comment_api";
import * as productType from "../Utils/ProductType";
import * as authService from "../Auth/AuthService";
import * as role from "../Utils/Role";
import * as helperFunctions from "../Utils/HelperFunctions";
import moment from "moment";

toast.configure();
const ProductDetail = (props) => {
	const customId = "product-detail";

	const { id } = useParams();
	const [product, setProduct] = useState(null);
	const [quantity, setQuantity] = useState(1);
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
				`${productPurchaseAPI.GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_USER_ID}?productId=${id}`
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
			if (p.Digital || p.Product.Quantity > 0) {
				setAvailable("This product is available for purchase");
				setDisableNotify(true);
			} else {
				setAvailable("This product is unavailable for purchase");
				setDisableAddToCart(true);
			}
		} else {
			if (p.Product.Quantity > 0) {
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
		let msg = "Product added to cart.";
		const productInCart = {
			Product: product,
			Quantity: quantity,
		};
		let productsInCart = JSON.parse(localStorage.getItem("cart") || "[]");
		const productInCartExists = productsInCart.filter(
			(pic) => pic.Product.Product.Id === productInCart.Product.Product.Id
		);
		if (productInCartExists.length > 0) {
			productsInCart = productsInCart.filter(
				(pic) => pic.Product.Product.Id !== productInCart.Product.Product.Id
			);
			msg = "Cart updated.";
		}
		productsInCart.push(productInCart);
		localStorage.setItem("cart", JSON.stringify(productsInCart));
		toast.success(msg, {
			position: toast.POSITION.TOP_CENTER,
			toastId: customId,
			autoClose: 5000,
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
				deleteProductComments(product.Product.Id);
				const route = helperFunctions.getProductListRoute(product);
				navigate(route);
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: false,
				});
			});
	};

	const deleteProductComments = (productId) => {
		axios.delete(
			`${productCommentAPI.DELETE_COMMENTS_BY_PRODUCT_ID}/${productId}`
		);
	};

	const addProductAlert = () => {
		axios
			.post(
				`${productPurchaseAPI.ADD_PRODUCT_ALERT}?productId=${product.Product.Id}`
			)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				getProductAlertByProductIdAndEmail();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const addProductOnMainPage = () => {
		axios
			.put(`${productAPI.ADD_PRODUCT_TO_MAIN_PAGE}/${product.Product.Id}`)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				getProductById();
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: false,
				});
			});
	};

	const removeProductFromMainPage = () => {
		axios
			.put(`${productAPI.REMOVE_PRODUCT_FROM_MAIN_PAGE}/${product.Product.Id}`)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				getProductById();
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: false,
				});
			});
	};

	return (
		<>
			<AppNavbar />
			{product !== null && (
				<>
					<Helmet>
						<title>{product.Product.Name} | GameZone</title>
					</Helmet>
					<Container>
						<Row>
							<Col style={{ marginTop: "10px", marginBottom: "10px" }} md={5}>
								<Card className="product-detail-card">
									<CardImg
										className="product-detail-card-image"
										src={product.Product.Image.Content}
										alt="No image"
									/>
									<CardBody>
										<CardTitle tag="h5">{product.Product.Name}</CardTitle>
										<CardText>{product.Product.Price} RSD</CardText>
										<CardText>{available}</CardText>
									</CardBody>
								</Card>
							</Col>

							<>
								<Col>
									<Card style={{ marginTop: "10px" }} className="card">
										<CardTitle
											className="title product-detail-description-title"
											tag="h4"
										>
											Description
										</CardTitle>
										<CardBody style={{ whiteSpace: "pre-line" }}>
											{product.Product.Description}
										</CardBody>
										<CardFooter className="product-detail-description-card-footer">
											More information on the manufacturer's website
										</CardFooter>
									</Card>
									{!disableAddToCart && (
										<>
											<Input
												className="input-field quantity-select"
												type="select"
												onChange={(e) => setQuantity(Number(e.target.value))}
											>
												<option hidden>Select quantity</option>
												<option>1</option>
												<option>2</option>
												<option>3</option>
												<option>4</option>
												<option>5</option>
											</Input>
											<Button
												style={{ marginTop: "10px" }}
												className="my-button"
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
												style={{ marginTop: "10px" }}
												className="my-button"
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
												{!product.Product.MainPage && (
													<Button
														style={{ marginTop: "10px", marginRight: "10px" }}
														className="my-button"
														type="button"
														onClick={addProductOnMainPage}
													>
														Put on main page
													</Button>
												)}
												{product.Product.MainPage && (
													<Button
														style={{ marginTop: "10px", marginRight: "10px" }}
														className="my-button"
														type="button"
														onClick={removeProductFromMainPage}
													>
														Remove from main page
													</Button>
												)}
												<Button
													style={{ marginTop: "10px", marginRight: "10px" }}
													className="my-button"
													type="button"
													onClick={updateProduct}
												>
													Update
												</Button>
												<Button
													style={{ marginTop: "10px", marginRight: "10px" }}
													className="my-button"
													type="button"
													onClick={deleteProduct}
												>
													Delete
												</Button>
											</>
										)}
								</Col>
							</>
						</Row>
						<Row>
							<Col>
								<Card
									style={{ marginTop: "20px", marginBottom: "10px" }}
									className="card"
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
													moment(
														product[value],
														"YYYY-MM-DDThh:mm:ssZ",
														true
													).isValid()
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
																{new Date(product[value]).toLocaleDateString()}
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
															<td style={{ wordBreak: "break-word" }}>
																{product[value].toString()}
															</td>
														</tr>
													);
												}
											}
										})}
									</Table>
								</Card>
							</Col>
						</Row>
						<CommentRating product={product} />
					</Container>
				</>
			)}
		</>
	);
};

export default ProductDetail;
