import {
	Row,
	Col,
	Container,
	Spinner,
	Card,
	CardTitle,
	CardBody,
} from "reactstrap";
import Pagination from "react-js-pagination";
import { Helmet } from "react-helmet";
import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";
import * as authService from "../Auth/AuthService";
import * as productPurchaseAPI from "../APIs/ProductMicroservice/product_purchase_api";
import * as userAPI from "../APIs/UserMicroservice/user_api";
import Purchase from "../Components/PurchaseHistory/Purchase";

const PurchaseHistory = () => {
	const [purchases, setPurchases] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [numberOfRecords, setNumberOfRecords] = useState(0);
	const pageSize = 10;
	const [loading, setLoading] = useState(true);
	const [user, setUser] = useState(null);

	const { id } = useParams();

	const navigate = useNavigate();

	useEffect(() => {
		if (
			authService.isAdmin() ||
			(authService.isUser() && Number(id) === Number(authService.getId()))
		) {
			getPurchaseHistory();
			getUserById();
		} else {
			navigate(-1);
		}
	}, [currentPage]);

	const getPurchaseHistory = () => {
		axios
			.get(
				`${productPurchaseAPI.GET_PURCHASE_HISTORY}?userId=${id}&page=${currentPage}&pageSize=${pageSize}`
			)
			.then((res) => {
				setPurchases(res.data);
				getNumberOfRecords();
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getNumberOfRecords = () => {
		axios
			.get(
				`${productPurchaseAPI.GET_NUMBER_OF_RECORDS_PURCHASE_HISTORY}?userId=${id}`
			)
			.then((res) => {
				setNumberOfRecords(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getUserById = () => {
		if (authService.isAdmin()) {
			axios
				.get(`${userAPI.GET_USER_BY_ID}?userId=${id}`)
				.then((res) => {
					setUser(res.data.user);
					console.log(res.data.user);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	const handleClick = (index) => {
		if (index !== currentPage) {
			setLoading(true);
		}
		setCurrentPage(index);
	};

	return (
		<>
			{user === null && authService.isUser() && (
				<Helmet>
					<title>Purchase history | GameZone</title>
				</Helmet>
			)}
			{user !== null && authService.isAdmin() && (
				<Helmet>
					<title>{user.user_name}'s purchase history | GameZone</title>
				</Helmet>
			)}
			{loading && (
				<div className="div-spinner">
					<Spinner className="spinner" />
				</div>
			)}
			<Container>
				<Row>
					<Col>
						{!loading && purchases.length <= 0 && authService.isUser() && (
							<Card className="card">
								<CardBody>
									<Container>
										<Row>
											<Col>
												<CardTitle className="title " tag="h2">
													You have no purchases.
												</CardTitle>
											</Col>
										</Row>
									</Container>
								</CardBody>
							</Card>
						)}
						{!loading &&
							purchases.length <= 0 &&
							user !== null &&
							authService.isAdmin() && (
								<Card className="card">
									<CardBody>
										<Container>
											<Row>
												<Col>
													<CardTitle className="title " tag="h2">
														{user.user_name} has no purchases.
													</CardTitle>
												</Col>
											</Row>
										</Container>
									</CardBody>
								</Card>
							)}
					</Col>
				</Row>
			</Container>

			{!loading && purchases.length > 0 && (
				<Container>
					{purchases.map((purchase) => {
						return (
							<Purchase
								purchase={purchase}
								getPurchaseHistory={getPurchaseHistory}
							/>
						);
					})}
					<Row className="pagination">
						<Col>
							<Pagination
								className="pagination"
								activePage={currentPage}
								itemsCountPerPage={pageSize}
								totalItemsCount={numberOfRecords}
								onChange={handleClick}
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

export default PurchaseHistory;
