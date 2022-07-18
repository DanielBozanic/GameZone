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
	Pagination,
	PaginationItem,
	PaginationLink,
} from "reactstrap";
import cn from "classnames";
import { useNavigate } from "react-router-dom";
import * as productType from "../Utils/ProductType";

const ProductsView = (props) => {
	const navigate = useNavigate();

	const onHandleClick = (e, nextOrPrev) => {
		props.handleClick && props.handleClick(e, nextOrPrev);
	};

	const viewProductDetails = (product) => {
		if (product.Product !== undefined) {
			navigate(window.location.pathname + "/" + product.Product.Id);
		} else {
			navigationBasedOnType(product);
		}
	};

	const navigationBasedOnType = (product) => {
		switch (product.Type) {
			case productType.CONSOLE:
				navigate("/consoles/" + product.Id);
				break;
			case productType.GRAPHICS_CARD:
				navigate("/graphicsCards/" + product.Id);
				break;
			case productType.HARD_DISK_DRIVE:
				navigate("/hdds/" + product.Id);
				break;
			case productType.HEADPHONES:
				navigate("/headphones/" + product.Id);
				break;
			case productType.KEYBOARD:
				navigate("/keyboards/" + product.Id);
				break;
			case productType.MONITOR:
				navigate("/monitors/" + product.Id);
				break;
			case productType.MOTHERBOARD:
				navigate("/motherboards/" + product.Id);
				break;
			case productType.MOUSE:
				navigate("/mice/" + product.Id);
				break;
			case productType.POWER_SUPPLY_UNIT:
				navigate("/psus/" + product.Id);
				break;
			case productType.PROCESSOR:
				navigate("/processors/" + product.Id);
				break;
			case productType.RAM:
				navigate("/rams/" + product.Id);
				break;
			case productType.SOLID_STATE_DRIVE:
				navigate("/ssds/" + product.Id);
				break;
			case productType.VIDEO_GAME:
				navigate("/videoGames/" + product.Id);
				break;
		}
	};

	return (
		<>
			{props.products.length > 0 && (
				<Container>
					<Row className="card-row">
						{props.products.map((product, index) =>
							index < 4 ? (
								<Col style={{ paddingTop: "5px" }} md={3}>
									<Card
										className="card-with-image"
										onClick={() => viewProductDetails(product)}
									>
										<CardImg
											className="card-image"
											alt="No image"
											src={
												product.Product === undefined
													? product.Image.Content
													: product.Product.Image.Content
											}
										/>
										<CardBody>
											<CardTitle tag="h5">
												{product.Product === undefined
													? product.Name
													: product.Product.Name}
											</CardTitle>
											<CardText>
												{product.Product === undefined
													? product.Price
													: product.Product.Price}
												RSD
											</CardText>
										</CardBody>
									</Card>
								</Col>
							) : (
								""
							)
						)}
					</Row>
					<Row className="card-row">
						{props.products.map((product, index) =>
							index > 3 ? (
								<Col style={{ paddingTop: "5px" }} md={3}>
									<Card
										className="card-with-image"
										onClick={() => viewProductDetails(product)}
									>
										<CardImg
											className="card-image"
											alt="No image"
											src={
												product.Product === undefined
													? product.Image.Content
													: product.Product.Image.Content
											}
										/>
										<CardBody>
											<CardTitle tag="h5">
												{product.Product === undefined
													? product.Name
													: product.Product.Name}
											</CardTitle>
											<CardText>
												{product.Product === undefined
													? product.Price
													: product.Product.Price}
												RSD
											</CardText>
										</CardBody>
									</Card>
								</Col>
							) : (
								""
							)
						)}
					</Row>
					<Row
						className={cn(
							"pagination",
							props.products.length < 4
								? "pagination-padding"
								: "pagination-padding-normal"
						)}
					>
						<Col md="12">
							<Pagination size="lg">
								<PaginationItem disabled={props.currentPage <= 1}>
									<PaginationLink
										onClick={(e) => onHandleClick(e, props.currentPage - 1)}
										previous
									/>
								</PaginationItem>

								{[...Array(props.pageCount)].map((page, i) => (
									<PaginationItem active={i === props.currentPage - 1} key={i}>
										<PaginationLink onClick={(e) => onHandleClick(e, i + 1)}>
											{i + 1}
										</PaginationLink>
									</PaginationItem>
								))}
								<PaginationItem
									disabled={props.currentPage - 1 >= props.pageCount - 1}
								>
									<PaginationLink
										onClick={(e) => onHandleClick(e, props.currentPage + 1)}
										next
									/>
								</PaginationItem>
							</Pagination>
						</Col>
					</Row>
				</Container>
			)}
		</>
	);
};

export default ProductsView;
