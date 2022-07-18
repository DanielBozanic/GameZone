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

toast.configure();
const ProductDetail = (props) => {
	const customId = "product-detail";

	const { id } = useParams();
	const [product, setProduct] = useState(null);
	const [amount, setAmount] = useState(1);
	const [available, setAvailable] = useState("");
	const [disabled, setDisabled] = useState(false);

	const navigate = useNavigate();

	useEffect(() => {
		getProductById();
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

	const pageSetup = (p) => {
		if (p.Product.Type === productType.VIDEO_GAME) {
			if (p.Digital) {
				setAvailable("This product is available for purchase");
			} else if (p.Product.Amount > 0) {
				setAvailable("This product is available for purchase");
			} else {
				setAvailable("This product is unavailable for purchase");
				setDisabled(true);
			}
		} else {
			if (p.Product.Amount > 0) {
				setAvailable("This product is available for purchase");
			} else {
				setAvailable("This product is unavailable for purchase");
				setDisabled(true);
			}
		}

		if (
			authService.getToken() == null ||
			(authService.getToken() != null &&
				authService.getRole() !== role.ROLE_USER)
		) {
			setDisabled(true);
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
		switch (product.Product.Type) {
			case productType.CONSOLE:
				navigate("/updateConsole/" + product.Product.Id);
				break;
			case productType.GRAPHICS_CARD:
				navigate("/updateGraphicsCard/" + product.Product.Id);
				break;
			case productType.HARD_DISK_DRIVE:
				navigate("/updateHdd/" + product.Product.Id);
				break;
			case productType.HEADPHONES:
				navigate("/updateHeadphones/" + product.Product.Id);
				break;
			case productType.KEYBOARD:
				navigate("/updateKeyboard/" + product.Product.Id);
				break;
			case productType.MONITOR:
				navigate("/updateMonitor/" + product.Product.Id);
				break;
			case productType.MOTHERBOARD:
				navigate("/updateMotherboard/" + product.Product.Id);
				break;
			case productType.MOUSE:
				navigate("/updateMouse/" + product.Product.Id);
				break;
			case productType.POWER_SUPPLY_UNIT:
				navigate("/updatePsu/" + product.Product.Id);
				break;
			case productType.PROCESSOR:
				navigate("/updateProcessor/" + product.Product.Id);
				break;
			case productType.RAM:
				navigate("/updateRam/" + product.Product.Id);
				break;
			case productType.SOLID_STATE_DRIVE:
				navigate("/updateSsd/" + product.Product.Id);
				break;
			case productType.VIDEO_GAME:
				navigate("/updateVideoGame/" + product.Product.Id);
				break;
		}
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
				navigate("/");
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
							{!disabled && (
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
