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
	Pagination,
	PaginationItem,
	PaginationLink,
} from "reactstrap";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import * as authService from "../Auth/AuthService";
import * as newsArticleAPI from "../APIs/NewsMicroservice/news_article_api";

const NewsList = () => {
	const [newsArticles, setNewsArticles] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [pageCount, setPageCount] = useState([]);
	const pageSize = 10;

	useEffect(() => {
		if (authService.isEmployee()) {
			getNewsArticles();
			getPageCount();
		} else {
			getPublishedNewsArticles();
			getPublishedArticlesPageCount();
		}
	}, [currentPage]);

	const navigate = useNavigate();

	const getNewsArticles = () => {
		axios
			.get(`${newsArticleAPI.GET_ALL}?page=${currentPage}&pageSize=${pageSize}`)
			.then((res) => {
				setNewsArticles(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPageCount = () => {
		axios
			.get(`${newsArticleAPI.GET_NUMBER_OF_RECORDS}`)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
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
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const getPublishedArticlesPageCount = () => {
		axios
			.get(`${newsArticleAPI.GET_NUMBER_OF_RECORDS_PUBLISHED_ARTICLES}`)
			.then((res) => {
				setPageCount(Math.ceil(Number(res.data) / pageSize));
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const moreDetails = (id) => {
		navigate("/viewNews/" + id);
	};

	const editNewsArticle = (id) => {
		navigate("/editNewsArticle/" + id);
	};

	const deleteNewsArticle = (id) => {
		axios
			.delete(`${newsArticleAPI.DELETE_NEWS_ARTICLE}/${id}`)
			.then((res) => {
				console.log(res);
				getNewsArticles();
				getPageCount();
			})
			.catch((err) => {
				console.error(err);
			});
	};

	const handleClick = (e, index) => {
		e.preventDefault();
		setCurrentPage(index);
	};

	return (
		<>
			{newsArticles.length > 0 && (
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
											<Button
												style={{ marginRight: "5px" }}
												className="my-button"
												type="button"
												onClick={() => moreDetails(newsArticle.Id)}
											>
												More details
											</Button>
											{authService.isEmployee() && (
												<>
													<Button
														style={{ marginRight: "5px" }}
														className="my-button"
														type="button"
														onClick={() => editNewsArticle(newsArticle.Id)}
													>
														Edit
													</Button>
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

export default NewsList;
