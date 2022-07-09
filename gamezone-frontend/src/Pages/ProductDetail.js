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
} from "reactstrap";
import axios from "axios";
import "../Assets/css/product-detail.css";
import AppNavbar from "../Layout/AppNavbar";
import * as productAPI from "../APIs/ProductMicroservice/product_api";

import VideoGameDetail from "../Components/ProductDetails/VideoGameDetail";

const ProductDetail = () => {
	const { id } = useParams();
	const [product, setProduct] = useState(null);

	useEffect(() => {
		getProductById();
	}, []);

	const getProductById = () => {
		axios
			.get(`${productAPI.GET_PRODUCT_BY_ID}/${id}`)
			.then((res) => {
				setProduct(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getSpecificProductDetails = () => {
		switch (product.Type) {
			case 1:
				break;
			case 2:
				break;
			case 3:
				break;
			case 4:
				break;
			case 5:
				break;
			case 6:
				break;
			case 7:
				break;
			case 8:
				break;
			case 9:
				break;
			case 10:
				break;
			case 11:
				break;
			case 12:
				break;
			case 13:
				return <VideoGameDetail product={product} />;
		}
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
									src={product.Image}
									alt="No image"
								/>
								<CardBody>
									<CardTitle tag="h5">{product.Name}</CardTitle>

									<CardText>
										{product.Price}
										RSD
									</CardText>
								</CardBody>
							</Card>
						</Col>
						{getSpecificProductDetails()}
					</Row>
				</Container>
			)}
		</>
	);
};

export default ProductDetail;
