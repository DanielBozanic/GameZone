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
} from "reactstrap";
import axios from "axios";
import "../Assets/css/product-detail.css";
import AppNavbar from "../Layout/AppNavbar";

const ProductDetail = (props) => {
	const { id } = useParams();
	const [product, setProduct] = useState(null);

	useEffect(() => {
		getProductById();
	}, []);

	const getProductById = () => {
		axios
			.get(`${props.GET_PRODUCT_BY_ID}/${id}`)
			.then((res) => {
				setProduct(res.data);
			})
			.catch((err) => {
				console.error(err);
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
								</CardBody>
							</Card>
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
