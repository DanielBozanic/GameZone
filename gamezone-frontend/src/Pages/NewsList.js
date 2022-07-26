import "../Assets/css/news-list.css";
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
import DOMPurify from "dompurify";
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

	const editNewsArticle = (id) => {
		navigate("/editNewsArticle/" + id);
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
							<Row className="news-list-row">
								<Col style={{ paddingTop: "5px" }}>
									<Card className="news-article-card">
										<CardHeader>
											<Row>
												<Col md="10">
													{authService.isEmployee() && (
														<CardTitle tag="h5">
															{newsArticle.UnpublishedTitle}
														</CardTitle>
													)}
													{!authService.isEmployee() && (
														<CardTitle tag="h5">
															{newsArticle.PublishedTitle}
														</CardTitle>
													)}
												</Col>
												<Col md="2">
													<CardTitle tag="h5">
														{newsArticle.DateTime.toString().split("T")[0]}
													</CardTitle>
												</Col>
											</Row>
										</CardHeader>
										{((newsArticle.UnpublishedDescription !== null &&
											newsArticle.UnpublishedDescription !== "") ||
											(newsArticle.PublishedDescription !== null &&
												newsArticle.PublishedDescription !== "")) && (
											<CardBody>
												{authService.isEmployee() && (
													<CardText>
														{newsArticle.UnpublishedDescription}
													</CardText>
												)}
												{!authService.isEmployee() && (
													<CardText>
														{newsArticle.PublishedDescription}
													</CardText>
												)}
											</CardBody>
										)}
										{/* <div
											dangerouslySetInnerHTML={{
												__html: DOMPurify.sanitize(
													newsArticle.UnpublishedContent
												),
											}}
										></div> */}
										<CardFooter>
											<Button className="news-list-buttons" type="button">
												More details
											</Button>
											{authService.isEmployee() && (
												<>
													<Button
														className="news-list-buttons"
														type="button"
														onClick={() => editNewsArticle(newsArticle.Id)}
													>
														Edit
													</Button>
													<Button className="news-list-buttons" type="button">
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
