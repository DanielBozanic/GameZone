import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import {
	CardText,
	CardImg,
	CardTitle,
	CardBody,
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

toast.configure();
const ProductDetail = (props) => {
	const customId = "product-detail";

	const { id } = useParams();
	const [product, setProduct] = useState(null);
	const [amount, setAmount] = useState(1);
	const [available, setAvailable] = useState("");
	const [disabled, setDisabled] = useState(false);

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
			(authService.getToken() != null && authService.getRole() !== "ROLE_USER")
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

	return (
		<>
			<AppNavbar />
			{product !== null && (
				<Container>
					<Row className="product-detail-card">
						<Col style={{ paddingTop: "5px" }} md={5}>
							<Card className="product-detail-card-with-image">
								<CardImg
									className="product-detail-card-image"
									src={product.Product.Image}
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
						</Col>
						<Col style={{ paddingTop: "5px", paddingBottom: "10px" }} md={7}>
							<Card className="product-detail-table-card">
								<Table className="product-detail-table">
									{product !== null &&
										Object.keys(product).map(function (value, idx) {
											if (value !== "Product") {
												if (typeof product[value] == "boolean") {
													return (
														<tr key={idx}>
															<th>{value.replace(/([A-Z])/g, " $1").trim()}</th>
															<td>{product[value] ? "Yes" : "No"}</td>
														</tr>
													);
												} else {
													return (
														<tr key={idx}>
															<th>{value.replace(/([A-Z])/g, " $1").trim()}</th>
															<td>{product[value].toString()}</td>
														</tr>
													);
												}
											}
										})}
								</Table>
							</Card>
						</Col>
					</Row>
				</Container>
			)}
		</>
	);
};

export default ProductDetail;
