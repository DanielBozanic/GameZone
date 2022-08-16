import "../Assets/css/products-view.css";
import {
	CardText,
	CardImg,
	CardTitle,
	CardBody,
	Card,
	Row,
	Col,
	Container,
	CardFooter,
} from "reactstrap";
import Pagination from "react-js-pagination";
import { Link } from "react-router-dom";
import cn from "classnames";
import * as helperFunctions from "../Utils/HelperFunctions";

const ProductsView = (props) => {
	const onHandleClick = (nextOrPrev) => {
		props.handleClick && props.handleClick(nextOrPrev);
	};

	const viewProductDetails = (product) => {
		if (product.Product !== undefined) {
			return window.location.pathname + "/" + product.Product.Id;
		} else {
			const route = helperFunctions.getProductDetailRoute(product);
			return route;
		}
	};

	return (
		<>
			{props.products.length > 0 && (
				<Container>
					<Row style={{ marginTop: "20px" }}>
						{props.products.map((product) => (
							<Col className="products-view-card-col" md="3">
								<Card className="products-view-card">
									<Link
										to={viewProductDetails(product)}
										style={{ textDecoration: "none", color: "inherit" }}
									>
										<CardImg
											className="products-view-card-image"
											alt="No image"
											src={
												product.Product === undefined
													? product.Image.Content
													: product.Product.Image.Content
											}
										/>
										<CardBody className="products-view-card-body">
											<CardTitle tag="h5">
												{product.Product === undefined
													? product.Name
													: product.Product.Name}
											</CardTitle>
										</CardBody>
										<CardFooter className="products-view-card-footer">
											<CardText className="products-view-card-footer-text">
												{product.Product === undefined
													? product.Price
													: product.Product.Price}{" "}
												RSD
											</CardText>
										</CardFooter>
									</Link>
								</Card>
							</Col>
						))}
					</Row>
					<Row
						className={cn(
							"pagination",
							props.products.length < 4
								? "products-view-pagination-padding"
								: "products-view-pagination-padding-normal"
						)}
					>
						<Col>
							<Pagination
								className="pagination"
								activePage={props.currentPage}
								itemsCountPerPage={props.pageSize}
								totalItemsCount={props.numberOfRecords}
								onChange={onHandleClick}
								itemClass="page-item"
								linkClass="page-link"
								pageRangeDisplayed={5}
							/>
						</Col>
					</Row>
				</Container>
			)}
		</>
	);
};

export default ProductsView;
