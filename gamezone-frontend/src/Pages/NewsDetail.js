import {
	CardText,
	CardHeader,
	CardTitle,
	CardBody,
	Card,
	Row,
	Col,
	Container,
} from "reactstrap";
import { Helmet } from "react-helmet";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import DOMPurify from "dompurify";
import * as authService from "../Auth/AuthService";
import * as newsArticleAPI from "../APIs/NewsMicroservice/news_article_api";
import NewsComment from "../Components/NewsComment/NewsComment";

const NewsDetail = () => {
	const [newsArticle, setNewsArticle] = useState(null);
	const { id } = useParams();

	useEffect(() => {
		getNewArticle();
	}, []);

	const getNewArticle = () => {
		axios
			.get(`${newsArticleAPI.GET_BY_ID}/${id}`)
			.then((res) => {
				setNewsArticle(res.data);
			})
			.catch((err) => {
				console.error(err);
			});
	};

	return (
		<>
			{newsArticle !== null && (
				<>
					<Helmet>
						<title>
							{authService.isEmployee()
								? newsArticle.UnpublishedTitle
								: newsArticle.PublishedTitle}{" "}
							| GameZone
						</title>
					</Helmet>
					<Container>
						<Row style={{ marginTop: "10px", marginBottom: "10px" }}>
							<Col>
								<Card className="card">
									<CardHeader>
										<Row>
											<Col md="10">
												<CardTitle className="title" tag="h3">
													{authService.isEmployee()
														? newsArticle.UnpublishedTitle
														: newsArticle.PublishedTitle}
												</CardTitle>
											</Col>
											<Col md="2">
												<CardTitle tag="h5">
													{new Date(newsArticle.DateTime).toLocaleDateString()}
												</CardTitle>
											</Col>
										</Row>
									</CardHeader>
									<CardBody>
										<CardText>
											<Container>
												<Row>
													<Col className="news-content">
														<div
															dangerouslySetInnerHTML={{
																__html: DOMPurify.sanitize(
																	authService.isEmployee()
																		? newsArticle.UnpublishedContent
																		: newsArticle.PublishedContent
																),
															}}
														></div>
													</Col>
												</Row>
											</Container>
										</CardText>
										<NewsComment newsArticle={newsArticle} />
									</CardBody>
								</Card>
							</Col>
						</Row>
					</Container>
				</>
			)}
		</>
	);
};

export default NewsDetail;
