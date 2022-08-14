import {
	Row,
	Col,
	Container,
	Pagination,
	PaginationItem,
	PaginationLink,
	Spinner,
} from "reactstrap";
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
	const [pageCount, setPageCount] = useState([]);
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
				getPageCount();
				setLoading(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getPageCount = () => {
		axios
			.get(
				`${productPurchaseAPI.GET_NUMBER_OF_RECORDS_PURCHASE_HISTORY}?userId=${id}`
			)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
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

	const handleClick = (e, index) => {
		setLoading(true);
		e.preventDefault();
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
							<Pagination size="lg">
								<PaginationItem disabled={currentPage <= 1}>
									<PaginationLink
										onClick={(e) => handleClick(e, currentPage - 1)}
										previous
									/>
								</PaginationItem>

								{[...Array(pageCount)].map((page, i) => (
									<PaginationItem active={i === currentPage - 1} key={i}>
										<PaginationLink onClick={(e) => handleClick(e, i + 1)}>
											{i + 1}
										</PaginationLink>
									</PaginationItem>
								))}
								<PaginationItem disabled={currentPage - 1 >= pageCount - 1}>
									<PaginationLink
										onClick={(e) => handleClick(e, currentPage + 1)}
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

export default PurchaseHistory;
