import { useState, useEffect } from "react";
import {
	CardText,
	CardTitle,
	CardBody,
	Card,
	Row,
	Col,
	Label,
	Input,
	Button,
	FormFeedback,
} from "reactstrap";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { commentRatingSchema } from "./CommentRatingSchema";
import axios from "axios";
import * as authService from "../../Auth/AuthService";
import * as productCommentAPI from "../../APIs/CommentAndRatingMicroservice/product_comment_api";
import * as productPurchaseAPI from "../../APIs/ProductMicroservice/product_purchase_api";
import "../../Assets/css/comment-rating.css";

const CommentRating = (props) => {
	const [reviews, setReviews] = useState([]);
	const [rating, setRating] = useState("1");
	const [currentUserReview, setCurrentUserReview] = useState(null);
	const [productIsPaidFor, setProductIsPaidFor] = useState(false);
	const [readOnlyMode, setReadOnlyMode] = useState(true);

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({ resolver: yupResolver(commentRatingSchema), mode: "onChange" });

	useEffect(() => {
		getAll();
		checkIfProductIsPaidFor();
	}, []);

	const getAll = () => {
		axios
			.get(
				`${productCommentAPI.GET_BY_PRODUCT_NAME}/${props.product.Product.Name}`
			)
			.then((res) => {
				if (authService.isUser()) {
					const userReview = res.data.filter(
						(review) => review.Username === authService.getUsername()
					);
					res.data = res.data.filter(
						(review) => review.Username !== authService.getUsername()
					);
					setReviews(res.data);
					if (userReview[0] !== undefined) {
						setCurrentUserReview(userReview[0]);
					}
				} else {
					setReviews(res.data);
				}
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const checkIfProductIsPaidFor = () => {
		axios
			.get(
				`${productPurchaseAPI.CHECK_IF_PRODUCT_IS_PAID_FOR}?productId=${props.product.Product.Id}`
			)
			.then((res) => {
				setProductIsPaidFor(res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const post = (data) => {
		if (currentUserReview === null) {
			data.Rating = Number(rating);
			data.ProductName = props.product.Product.Name;
			data.Username = authService.getUsername();
			addComment(data);
		} else {
			currentUserReview.Rating = Number(rating);
			currentUserReview.Comment = data.Comment;
			editComment(currentUserReview);
		}
	};

	const addComment = (data) => {
		axios
			.post(`${productCommentAPI.ADD_COMMENT}`, data)
			.then((res) => {
				console.log(res);
				getAll();
				setReadOnlyMode(true);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const editComment = (data) => {
		axios
			.put(`${productCommentAPI.EDIT_COMMENT}`, data)
			.then((res) => {
				console.log(res);
				setReadOnlyMode(true);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const deleteComment = (commentId) => {
		axios
			.delete(`${productCommentAPI.DELETE_COMMENT}/${commentId}`)
			.then((res) => {
				console.log(res);
				getAll();
				setCurrentUserReview(null);
				setRating("1");
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			{(reviews.length > 0 ||
				currentUserReview !== null ||
				productIsPaidFor) && (
				<Row>
					<Col>
						<Card className="review-card">
							<CardBody>
								<CardTitle className="review-title" tag="h3">
									Customer impressions
								</CardTitle>
								{(currentUserReview !== null ||
									(productIsPaidFor && currentUserReview === null)) && (
									<Row>
										<Col>
											<Row>
												<Col style={{ marginBottom: "5px" }}>
													{currentUserReview !== null && (
														<CardText>
															{currentUserReview.Username} -{" "}
															{currentUserReview.Rating} ★
														</CardText>
													)}
												</Col>

												{(!readOnlyMode || currentUserReview === null) && (
													<div class="rating">
														<Label>
															<Input
																type="radio"
																name="Rating"
																value="1"
																checked={rating === "1"}
																onChange={() => setRating("1")}
															/>
															<span class="icon">★</span>
														</Label>
														<Label>
															<Input
																type="radio"
																name="Rating"
																value="2"
																checked={rating === "2"}
																onChange={() => setRating("2")}
															/>
															<span class="icon">★</span>
															<span class="icon">★</span>
														</Label>
														<Label>
															<Input
																type="radio"
																name="Rating"
																value="3"
																checked={rating === "3"}
																onChange={() => setRating("3")}
															/>
															<span class="icon">★</span>
															<span class="icon">★</span>
															<span class="icon">★</span>
														</Label>
														<Label>
															<Input
																type="radio"
																name="Rating"
																value="4"
																checked={rating === "4"}
																onChange={() => setRating("4")}
															/>
															<span class="icon">★</span>
															<span class="icon">★</span>
															<span class="icon">★</span>
															<span class="icon">★</span>
														</Label>
														<Label>
															<Input
																type="radio"
																name="Rating"
																value="5"
																checked={rating === "5"}
																onChange={() => setRating("5")}
															/>
															<span class="icon">★</span>
															<span class="icon">★</span>
															<span class="icon">★</span>
															<span class="icon">★</span>
															<span class="icon">★</span>
														</Label>
													</div>
												)}
											</Row>
											<Row>
												<Col>
													{(!readOnlyMode || currentUserReview === null) && (
														<>
															<Input
																type="textarea"
																className="review-input"
																name="Comment"
																defaultValue={
																	currentUserReview !== null
																		? currentUserReview.Comment
																		: ""
																}
																placeholder="Write a product review"
																innerRef={register}
																invalid={errors.Comment?.message}
															/>
															<FormFeedback className="input-field-error-msg">
																{errors.Comment?.message}
															</FormFeedback>
														</>
													)}
													{readOnlyMode && currentUserReview !== null && (
														<div style={{ fontWeight: "normal" }}>
															{currentUserReview.Comment}
														</div>
													)}
												</Col>
											</Row>
											<Row>
												<Col>
													{(!readOnlyMode || currentUserReview === null) && (
														<Button
															className="comment-btn"
															type="button"
															onClick={handleSubmit(post)}
														>
															Post
														</Button>
													)}
													{readOnlyMode && currentUserReview !== null && (
														<Button
															className="comment-btn"
															type="button"
															onClick={() => setReadOnlyMode(false)}
														>
															Edit
														</Button>
													)}
													{currentUserReview !== null && (
														<Button
															className="comment-btn"
															type="button"
															onClick={() =>
																deleteComment(currentUserReview.Id)
															}
														>
															Delete
														</Button>
													)}
												</Col>
											</Row>
										</Col>
									</Row>
								)}
								{reviews.length > 0 &&
									reviews.map((review) => {
										return (
											<Row>
												<Col>
													<Row>
														<Col
															style={{
																marginBottom: "5px",
																marginTop: "15px",
															}}
														>
															<CardText>
																{review.Username} - {review.Rating} ★
															</CardText>
														</Col>
													</Row>
													<Row>
														<Col>
															<div style={{ fontWeight: "normal" }}>
																{review.Comment}
															</div>
														</Col>
													</Row>
													{(authService.isEmployee() ||
														authService.isUser()) && (
														<Row>
															<Col>
																<Button className="comment-btn" type="button">
																	Report
																</Button>
															</Col>
														</Row>
													)}
													{authService.isAdmin() && (
														<Row>
															<Col>
																<Button
																	className="comment-btn"
																	type="button"
																	onClick={() => deleteComment(review.Id)}
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
							</CardBody>
						</Card>
					</Col>
				</Row>
			)}
		</>
	);
};

export default CommentRating;
