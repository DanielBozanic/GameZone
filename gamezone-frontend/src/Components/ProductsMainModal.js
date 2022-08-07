import {
	Button,
	Modal,
	ModalHeader,
	ModalBody,
	Row,
	Col,
	Card,
	CardHeader,
	CardTitle,
	CardBody,
	CardImg,
	CardFooter,
} from "reactstrap";
import { useState, useEffect } from "react";
import axios from "axios";
import "../Assets/css/products-main-modal.css";
import * as productAPI from "../APIs/ProductMicroservice/product_api";

const ProductsMainModal = () => {
	const [mainPageProducts, setMainPageProducts] = useState([]);
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);

	useEffect(() => {
		getMainPageProducts();
	}, []);

	const getMainPageProducts = () => {
		axios
			.get(`${productAPI.GET_MAIN_PAGE_PRODUCTS}`)
			.then((res) => {
				setMainPageProducts(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const removeProductFromMainPage = (product) => {
		product.MainPage = false;
		axios
			.put(`${productAPI.REMOVE_PRODUCT_FROM_MAIN_PAGE}/${product.Id}`)
			.then((_res) => {
				getMainPageProducts();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			<Button
				style={{ marginTop: "10px", marginRight: "10px", marginBottom: "10px" }}
				className="my-button"
				type="button"
				onClick={toggle}
			>
				Settings
			</Button>
			<Modal
				size="lg"
				style={{ maxWidth: "1500px", width: "100%" }}
				isOpen={modal}
				toggle={toggle}
			>
				<ModalHeader className="title" toggle={toggle}>
					Current products on main page
				</ModalHeader>
				<ModalBody>
					<Row>
						{mainPageProducts.length > 0 &&
							mainPageProducts.map((product) => {
								return (
									<Col md="4">
										<Card className="single-product-card">
											<CardHeader>
												<CardTitle>{product.Name}</CardTitle>
											</CardHeader>
											<CardBody>
												<CardImg
													className="single-product-card-image"
													src={product.Image.Content}
													alt="No image"
												/>
											</CardBody>
											<CardFooter>
												<Row>
													<Col>
														<Button
															className="my-button"
															type="button"
															onClick={() => removeProductFromMainPage(product)}
														>
															Remove
														</Button>
													</Col>
												</Row>
											</CardFooter>
										</Card>
									</Col>
								);
							})}
					</Row>
				</ModalBody>
			</Modal>
		</>
	);
};

export default ProductsMainModal;
