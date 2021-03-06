import "../Assets/css/registered-users-list.css";
import {
	Button,
	CardText,
	CardTitle,
	CardHeader,
	CardBody,
	Card,
	CardFooter,
	Row,
	Col,
	Pagination,
	PaginationItem,
	PaginationLink,
	Label,
} from "reactstrap";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import cn from "classnames";
import axios from "axios";
import * as userAPI from "../APIs/UserMicroservice/user_api";

const RegisteredUsersList = () => {
	const [registeredUsers, setRegisteredUsers] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [pageCount, setPageCount] = useState([]);
	const pageSize = 10;

	const navigate = useNavigate();

	useEffect(() => {
		getRegisteredUsers();
		getPageCount();
	}, [currentPage]);

	const getRegisteredUsers = () => {
		axios
			.get(
				`${userAPI.GET_ALL_REGISTERED_USERS}?page=${currentPage}&pageSize=${pageSize}`
			)
			.then((res) => {
				setRegisteredUsers(res.data.users);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPageCount = () => {
		axios
			.get(`${userAPI.GET_NUMBER_OF_RECORDS_REGISTERED_USERS}`)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data.count) / pageSize));
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const viewPurchaseHistory = (userId) => {
		navigate("/purchaseHistory/" + userId);
	};

	const handleClick = (e, index) => {
		e.preventDefault();
		setCurrentPage(index);
	};

	return (
		<>
			{registeredUsers.length > 0 && (
				<>
					<Row className="registered-users-row">
						{registeredUsers.map((registeredUser) => {
							return (
								<>
									<Col
										md="4"
										style={{
											paddingBottom: "10px",
										}}
									>
										<Card className="registered-user-card">
											<CardHeader>
												<Row>
													<Col>
														<CardTitle
															className="registered-user-card-title"
															tag="h5"
														>
															{registeredUser.user_name}
														</CardTitle>
													</Col>
												</Row>
											</CardHeader>
											<CardBody>
												<Row>
													<Col>
														<CardText>
															<Label style={{ fontWeight: "bold" }}>
																Email:
															</Label>{" "}
															{registeredUser.email}
														</CardText>
													</Col>
												</Row>
												<Row>
													<Col>
														<CardText>
															<Label style={{ fontWeight: "bold" }}>
																Name:
															</Label>{" "}
															{registeredUser.name}
														</CardText>
													</Col>
												</Row>
												<Row>
													<Col>
														<CardText>
															<Label style={{ fontWeight: "bold" }}>
																Surname:
															</Label>{" "}
															{registeredUser.surname}
														</CardText>
													</Col>
												</Row>
											</CardBody>
											<CardFooter>
												<Row>
													<Col>
														<Button
															className="registered-user-buttons"
															type="button"
															onClick={() =>
																viewPurchaseHistory(registeredUser.id)
															}
														>
															Purchase history
														</Button>
														<Button
															className="registered-user-buttons"
															type="button"
														>
															Report log
														</Button>
													</Col>
												</Row>
											</CardFooter>
										</Card>
									</Col>
								</>
							);
						})}
					</Row>
					<Row
						className={cn(
							"registered-users-pagination",
							registeredUsers.length < 4
								? "registered-users-pagination-padding"
								: "registered-users-pagination-padding-normal"
						)}
					>
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
				</>
			)}
		</>
	);
};

export default RegisteredUsersList;
