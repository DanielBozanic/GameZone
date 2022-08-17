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
	Container,
	Spinner,
} from "reactstrap";
import Pagination from "react-js-pagination";
import { Helmet } from "react-helmet";
import { useState, useEffect } from "react";
import axios from "axios";
import * as authService from "../Auth/AuthService";
import * as newsArticleAPI from "../APIs/NewsMicroservice/news_article_api";
import * as newsCommentAPI from "../APIs/NewsMicroservice/news_comment_api";

const NewsList = () => {
	const [newsArticles, setNewsArticles] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [numberOfRecords, setNumberOfRecords] = useState(0);
	const [loading, setLoading] = useState(true);
	const pageSize = 8;

	useEffect(() => {
		if (authService.isEmployee()) {
			getNewsArticles();
			getNumberOfRecords();
		} else {
			getPublishedNewsArticles();
			getNumberOfRecordsPublishedArticles();
		}
	}, [currentPage]);

	const getNewsArticles = () => {
		axios
			.get(`${newsArticleAPI.GET_ALL}?page=${currentPage}&pageSize=${pageSize}`)
			.then((res) => {
				setNewsArticles(res.data);
				setLoading(false);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getNumberOfRecords = () => {
		axios
			.get(`${newsArticleAPI.GET_NUMBER_OF_RECORDS}`)
			.then((res) => {
				setNumberOfRecords(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPublishedNewsArticles = () => {
		axios
			.get(
				`${newsArticleAPI.GET_PUBLISHED_ARTICLES}?page=${currentPage}&pageSize=${pageSize}`
			)
			.then((res) => {
				setNewsArticles(res.data);
				setLoading(false);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getNumberOfRecordsPublishedArticles = () => {
		axios
			.get(`${newsArticleAPI.GET_NUMBER_OF_RECORDS_PUBLISHED_ARTICLES}`)
			.then((res) => {
				setNumberOfRecords(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const moreDetails = (id) => {
		return "/viewNews/" + id;
	};

	const editNewsArticle = (id) => {
		return "/editNewsArticle/" + id;
	};

	const deleteNewsArticle = (id) => {
		axios
			.delete(`${newsArticleAPI.DELETE_NEWS_ARTICLE}/${id}`)
			.then((_res) => {
				deleteNewsCommentsByNewsArticleId(id);
				getNewsArticles();
				getNumberOfRecords();
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const deleteNewsCommentsByNewsArticleId = (newsArticleId) => {
		axios.delete(
			`${newsCommentAPI.DELETE_NEWS_COMMENTS_BY_NEWS_ARTICLE_ID}/${newsArticleId}`
		);
	};

	const handleClick = (index) => {
		if (index !== currentPage) {
			setLoading(true);
		}
		setCurrentPage(index);
	};

	return (
		<>
			<Helmet>
				<title>News | GameZone</title>
			</Helmet>
			{loading && (
				<div className="div-spinner">
					<Spinner className="spinner" />
				</div>
			)}
			{!loading && newsArticles.length > 0 && (
				<Container>
					{newsArticles.map((newsArticle) => {
						return (
							<Row>
								<Col style={{ paddingTop: "5px" }}>
									<Card className="card">
										<CardHeader>
											<Row>
												<Col md="10">
													<CardTitle className="title" tag="h5">
														{authService.isEmployee()
															? newsArticle.UnpublishedTitle
															: newsArticle.PublishedTitle}
													</CardTitle>
												</Col>
												<Col md="2">
													<CardTitle tag="h5">
														{new Date(
															newsArticle.DateTime
														).toLocaleDateString()}
													</CardTitle>
												</Col>
											</Row>
										</CardHeader>
										{((newsArticle.UnpublishedDescription !== null &&
											newsArticle.UnpublishedDescription !== "") ||
											(newsArticle.PublishedDescription !== null &&
												newsArticle.PublishedDescription !== "")) && (
											<CardBody>
												<CardText>
													{authService.isEmployee()
														? newsArticle.UnpublishedDescription
														: newsArticle.PublishedDescription}
												</CardText>
											</CardBody>
										)}
										<CardFooter>
											<a href={moreDetails(newsArticle.Id)}>
												<Button
													style={{ marginRight: "5px" }}
													className="my-button"
													type="button"
												>
													More details
												</Button>
											</a>
											{authService.isEmployee() && (
												<>
													<a href={editNewsArticle(newsArticle.Id)}>
														<Button
															style={{ marginRight: "5px" }}
															className="my-button"
															type="button"
														>
															Edit
														</Button>
													</a>
													<Button
														style={{ marginRight: "5px" }}
														className="my-button"
														type="button"
														onClick={() => deleteNewsArticle(newsArticle.Id)}
													>
														Delete
													</Button>
												</>
											)}
										</CardFooter>
									</Card>
								</Col>
							</Row>
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

export default NewsList;
