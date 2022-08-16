import "../Assets/css/registered-users-list.css";
import {
	Button,
	CardText,
	CardTitle,
	CardHeader,
	CardBody,
	Card,
	CardFooter,
	Container,
	Row,
	Col,
	Label,
	Spinner,
} from "reactstrap";
import Pagination from "react-js-pagination";
import { Helmet } from "react-helmet";
import { useState, useEffect } from "react";
import cn from "classnames";
import axios from "axios";
import * as userAPI from "../APIs/UserMicroservice/user_api";
import * as contactAPI from "../APIs/ContactAndReportMicroservice/contact_api";
import Search from "../Components/Search";

const RegisteredUsersList = () => {
	const [registeredUsers, setRegisteredUsers] = useState([]);
	const [unansweredContactMsgs, setUnansweredContactMsgs] = useState(new Map());
	const [currentPage, setCurrentPage] = useState(1);
	const [numberOfRecords, setNumberOfRecords] = useState(0);
	const pageSize = 10;
	const [loading, setLoading] = useState(true);
	const [searchTerm, setSearchTerm] = useState("");

	useEffect(() => {
		if (searchTerm === "") {
			getRegisteredUsers();
			getNumberOfRecords();
		} else {
			searchRegisteredUsers();
			getNumberOfRecordsSearch();
		}
	}, [currentPage, searchTerm]);

	const getRegisteredUsers = () => {
		axios
			.get(
				`${userAPI.GET_ALL_REGISTERED_USERS}?page=${currentPage}&pageSize=${pageSize}`
			)
			.then((res) => {
				setLoading(false);
				setRegisteredUsers(res.data.users);
				getNumberOfUnansweredContactMessagesForByUserId(res.data.users);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getNumberOfRecords = () => {
		axios
			.get(`${userAPI.GET_NUMBER_OF_RECORDS_REGISTERED_USERS}`)
			.then((res) => {
				setNumberOfRecords(res.data.count);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const searchRegisteredUsers = () => {
		axios
			.get(
				`${userAPI.SEARCH_REGISTERED_USERS}?page=${currentPage}&pageSize=${pageSize}&searchTerm=${searchTerm}`
			)
			.then((res) => {
				setLoading(false);
				setRegisteredUsers(res.data.users);
				getNumberOfUnansweredContactMessagesForByUserId(res.data.users);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getNumberOfRecordsSearch = () => {
		axios
			.get(
				`${userAPI.GET_NUMBER_OF_RECORDS_REGISTERED_USERS_SEARCH}?searchTerm=${searchTerm}`
			)
			.then((res) => {
				setNumberOfRecords(res.data.count);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getNumberOfUnansweredContactMessagesForByUserId = (users) => {
		for (let user of users) {
			axios
				.get(
					`${contactAPI.GET_NUMBER_OF_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID}/${user.id}`
				)
				.then((res) => {
					setUnansweredContactMsgs(
						new Map(unansweredContactMsgs.set(user.id, res.data))
					);
				})
				.catch((err) => {
					console.log(err);
					setLoading(false);
				});
		}
	};

	const viewPurchaseHistory = (userId) => {
		return "/purchaseHistory/" + userId;
	};

	const manageUser = (userId) => {
		return "/manageUser/" + userId;
	};

	const viewContactMessages = (userId) => {
		return "/contactMessages/" + userId;
	};

	const handleClick = (index) => {
		setRegisteredUsers([]);
		if (index !== currentPage) {
			setLoading(true);
		}
		setCurrentPage(index);
	};

	const onSearchClick = (st) => {
		setSearchTerm(st);
		setRegisteredUsers([]);
		setLoading(true);
		setCurrentPage(1);
	};

	return (
		<Container fluid>
			<Helmet>
				<title>Registered Users | GameZone</title>
			</Helmet>
			{loading && (
				<div className="div-spinner">
					<Spinner className="spinner" />
				</div>
			)}
			{registeredUsers.length > 0 && (
				<>
					<Row className="registered-users-row">
						<Col md="10" style={{ marginBottom: "5px" }}>
							<Row>
								<Search
									onSearchClick={onSearchClick}
									searchPlaceholder={"Search by username"}
								/>
							</Row>
						</Col>
						{!loading &&
							registeredUsers.map((registeredUser) => {
								return (
									<>
										<Col
											style={{
												paddingBottom: "10px",
											}}
											md="6"
										>
											<Card className="card">
												<CardHeader>
													<Row>
														<Col>
															<CardTitle className="title" tag="h5">
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
													<Row>
														<Col>
															<CardText>
																<Label style={{ fontWeight: "bold" }}>
																	Unanswered messages:
																</Label>{" "}
																{unansweredContactMsgs.get(registeredUser.id) <=
																	0 && "No messages to answer"}
																{unansweredContactMsgs.get(registeredUser.id) >
																	0 &&
																	unansweredContactMsgs.get(registeredUser.id)}
															</CardText>
														</Col>
													</Row>
												</CardBody>
												<CardFooter>
													<Row>
														<Col>
															<a href={viewPurchaseHistory(registeredUser.id)}>
																<Button
																	className="my-button"
																	style={{
																		marginRight: "5px",
																		marginBottom: "5px",
																	}}
																	type="button"
																>
																	Purchase history
																</Button>
															</a>
															<a href={manageUser(registeredUser.id)}>
																<Button
																	style={{
																		marginRight: "5px",
																		marginBottom: "5px",
																	}}
																	className="my-button"
																	type="button"
																>
																	Manage
																</Button>
															</a>
															<a href={viewContactMessages(registeredUser.id)}>
																{" "}
																<Button
																	style={{
																		marginRight: "5px",
																		marginBottom: "5px",
																	}}
																	className="my-button"
																	type="button"
																>
																	Contact messages
																</Button>
															</a>
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
							"pagination",
							registeredUsers.length < 4
								? "registered-users-pagination-padding"
								: "registered-users-pagination-padding-normal"
						)}
					>
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
				</>
			)}
		</Container>
	);
};

export default RegisteredUsersList;
