import { useState, useEffect } from "react";
import {
	CardText,
	CardTitle,
	Row,
	Col,
	Input,
	Button,
	FormFeedback,
} from "reactstrap";
import axios from "axios";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { newsCommentSchema } from "./NewsCommentSchema";
import * as newsCommentAPI from "../../APIs/NewsMicroservice/news_comment_api";
import * as authService from "../../Auth/AuthService";
import ReportModal from "../ReportModal/ReportModal";

const NewsComment = (props) => {
	const [newsComments, setNewsComments] = useState([]);
	const [currentUserComments, setCurrentUserComments] = useState([]);
	const [selectedComment, setSelectedComment] = useState(null);

	const newComment = useForm({
		resolver: yupResolver(newsCommentSchema),
		mode: "onChange",
	});
	const existingComment = useForm({
		resolver: yupResolver(newsCommentSchema),
		mode: "onChange",
	});

	useEffect(() => {
		getNewsComments();
	}, []);

	const getNewsComments = () => {
		axios
			.get(`${newsCommentAPI.GET_BY_NEWS_ARTICLE}/${props.newsArticle.Id}`)
			.then((res) => {
				const userComments = res.data.filter(
					(comment) => comment.UserId === Number(authService.getId())
				);
				if (userComments.length > 0) {
					setCurrentUserComments(userComments);
				}
				res.data = res.data.filter(
					(review) => review.UserId !== Number(authService.getId())
				);
				setNewsComments(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const addComment = (data) => {
		data.NewsArticle = props.newsArticle;
		data.UserId = Number(authService.getId());
		axios
			.post(`${newsCommentAPI.ADD_NEWS_COMMENT}`, data)
			.then((res) => {
				console.log(res);
				newComment.setValue("Comment", "");
				getNewsComments();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const editComment = (data) => {
		selectedComment.Comment = data.Comment;
		axios
			.put(`${newsCommentAPI.EDIT_NEWS_COMMENT}`, selectedComment)
			.then((res) => {
				console.log(res);
				setSelectedComment(null);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const deleteComment = (commentId) => {
		axios
			.delete(`${newsCommentAPI.DELETE_NEWS_COMMENT}/${commentId}`)
			.then((res) => {
				console.log(res);
				getNewsComments();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			{newsComments.length > 0 && (
				<Row>
					<Col>
						<CardTitle className="title" tag="h3">
							Comment section
						</CardTitle>
						{authService.isUser() && (
							<Row>
								<Col>
									<Input
										type="textarea"
										className="input-field"
										style={{ resize: "none", height: "120px" }}
										name="Comment"
										innerRef={newComment.register}
										invalid={newComment.errors.Comment?.message}
									/>
									<FormFeedback className="input-field-error-msg">
										{newComment.errors.Comment?.message}
									</FormFeedback>
									<Button
										style={{ marginTop: "10px" }}
										className="my-button"
										type="button"
										onClick={newComment.handleSubmit(addComment)}
									>
										Post
									</Button>
								</Col>
							</Row>
						)}
						{currentUserComments.length > 0 &&
							currentUserComments.map((currentUserComment) => {
								return (
									<Row style={{ marginTop: "15px" }}>
										<Col>
											<Row>
												<Col>
													<CardText style={{ fontWeight: "bold" }}>
														{currentUserComment.Username}
													</CardText>
												</Col>
											</Row>
											<Row>
												<Col>
													{selectedComment !== null &&
														selectedComment.Id === currentUserComment.Id && (
															<>
																<Input
																	type="textarea"
																	className="input-field"
																	name="Comment"
																	defaultValue={
																		currentUserComment !== null
																			? currentUserComment.Comment
																			: ""
																	}
																	innerRef={existingComment.register}
																	invalid={
																		existingComment.errors.Comment?.message
																	}
																/>
																<FormFeedback className="input-field-error-msg">
																	{existingComment.errors.Comment?.message}
																</FormFeedback>
															</>
														)}
													{(selectedComment === null ||
														selectedComment.Id !== currentUserComment.Id) && (
														<div style={{ fontWeight: "normal" }}>
															{currentUserComment.Comment}
														</div>
													)}
												</Col>
											</Row>
											<Row>
												<Col>
													{selectedComment !== null &&
														selectedComment.Id === currentUserComment.Id && (
															<Button
																style={{ marginRight: "5px", marginTop: "5px" }}
																className="my-button"
																type="button"
																onClick={existingComment.handleSubmit(
																	editComment
																)}
															>
																Post
															</Button>
														)}
													{(selectedComment === null ||
														selectedComment.Id !== currentUserComment.Id) && (
														<Button
															style={{ marginRight: "5px", marginTop: "5px" }}
															className="my-button"
															type="button"
															onClick={() =>
																setSelectedComment(currentUserComment)
															}
														>
															Edit
														</Button>
													)}

													<Button
														style={{ marginRight: "5px", marginTop: "5px" }}
														className="my-button"
														type="button"
														onClick={() => deleteComment(currentUserComment.Id)}
													>
														Delete
													</Button>
												</Col>
											</Row>
										</Col>
									</Row>
								);
							})}
						{newsComments.length > 0 &&
							newsComments.map((newsComment) => {
								return (
									<Row style={{ marginTop: "15px" }}>
										<Col>
											<Row>
												<Col>
													<CardText style={{ fontWeight: "bold" }}>
														{newsComment.Username}
													</CardText>
												</Col>
											</Row>
											<Row>
												<Col>
													<div style={{ fontWeight: "normal" }}>
														{newsComment.Comment}
													</div>
												</Col>
											</Row>
											{(authService.isEmployee() || authService.isUser()) && (
												<ReportModal
													userId={newsComment.UserId}
													username={newsComment.Username}
												/>
											)}
											{authService.isAdmin() && (
												<Row>
													<Col>
														<Button
															style={{ marginTop: "5px" }}
															className="my-button"
															type="button"
															onClick={() => deleteComment(newsComment.Id)}
														>
															Delete
														</Button>
													</Col>
												</Row>
											)}
										</Col>
									</Row>
								);
							})}
					</Col>
				</Row>
			)}
		</>
	);
};

export default NewsComment;
