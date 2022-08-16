import "react-quill/dist/quill.snow.css";
import "../../Assets/css/news-editor.css";
import ReactQuill from "react-quill";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useParams, useNavigate } from "react-router-dom";
import { newsEditorSchema } from "./NewsEditorSchema";
import NewsEditorToolbar, {
	modules,
	formats,
} from "../../Components/NewsEditor/NewsEditorToolbar";
import {
	Card,
	CardBody,
	Row,
	Col,
	Container,
	Button,
	Form,
	Label,
	FormGroup,
	Input,
	FormFeedback,
} from "reactstrap";
import { Helmet } from "react-helmet";
import DOMPurify from "dompurify";
import { toast } from "react-toastify";
import * as newsArticleAPI from "../../APIs/NewsMicroservice/news_article_api";
import axios from "axios";

toast.configure();
const NewsEditor = () => {
	const customId = "NewsEditor";
	const [newsArticle, setNewsArticle] = useState(null);
	const {
		register,
		handleSubmit,
		formState: { errors },
		watch,
		setValue,
		getValues,
	} = useForm({ resolver: yupResolver(newsEditorSchema), mode: "onChange" });

	const navigate = useNavigate();
	const { id } = useParams();

	useEffect(() => {
		register("UnpublishedContent");
		setValue("UnpublishedContent", "");
		getNewsArticle();
	}, []);

	const unpublishedContent = watch("UnpublishedContent");

	const onEditorStateChange = (value) => {
		setValue("UnpublishedContent", value);
	};

	const getNewsArticle = () => {
		axios
			.get(`${newsArticleAPI.GET_BY_ID}/${id}`)
			.then((res) => {
				setNewsArticle(res.data);
				setValue("UnpublishedContent", res.data.UnpublishedContent);
			})
			.catch((err) => {
				console.log(err);
				setValue("UnpublishedContent", "");
			});
	};

	const isQuillEmpty = () => {
		const value = getValues("UnpublishedContent");
		return (
			value.replace(/<(.|\n)*?>/g, "").trim().length === 0 &&
			!value.includes("<img")
		);
	};

	const add = (data) => {
		data.UnpublishedContent = DOMPurify.sanitize(data.UnpublishedContent);
		if (!isQuillEmpty()) {
			axios
				.post(newsArticleAPI.ADD_NEWS_ARTICLE, data)
				.then((res) => {
					toast.success("News article added successfully", {
						position: toast.POSITION.TOP_CENTER,
						autoClose: 5000,
						toastId: customId,
					});
					navigate("/editNewsArticle/" + res.data.Id);
				})
				.catch((err) => {
					console.log(err);
				});
		} else {
			toast.error("News article must not be empty", {
				position: toast.POSITION.TOP_CENTER,
				autoClose: 10000,
				toastId: customId,
			});
		}
	};

	const edit = (data) => {
		if (!isQuillEmpty()) {
			newsArticle.UnpublishedTitle = data.UnpublishedTitle;
			newsArticle.UnpublishedDescription = data.UnpublishedDescription;
			newsArticle.UnpublishedContent = DOMPurify.sanitize(
				data.UnpublishedContent
			);
			axios
				.put(newsArticleAPI.EDIT_NEWS_ARTICLE, newsArticle)
				.then((res) => {
					toast.success(res.data, {
						position: toast.POSITION.TOP_CENTER,
						autoClose: 5000,
						toastId: customId,
					});
				})
				.catch((err) => {
					console.log(err);
				});
		} else {
			toast.error("News article must not be empty", {
				position: toast.POSITION.TOP_CENTER,
				autoClose: 10000,
				toastId: customId,
			});
		}
	};

	const publishNewsArticle = (data) => {
		if (!isQuillEmpty()) {
			let publishData = data;
			if (newsArticle !== null) {
				newsArticle.UnpublishedTitle = data.UnpublishedTitle;
				newsArticle.UnpublishedDescription = data.UnpublishedDescription;
				newsArticle.UnpublishedContent = DOMPurify.sanitize(
					data.UnpublishedContent
				);
				publishData = newsArticle;
			}
			console.log(publishData);
			axios
				.put(newsArticleAPI.PUBLISH_NEWS_ARTICLE, publishData)
				.then((res) => {
					toast.success(res.data, {
						position: toast.POSITION.TOP_CENTER,
						autoClose: 5000,
						toastId: customId,
					});
					navigate("/viewNews");
				})
				.catch((err) => {
					console.log(err);
				});
		} else {
			toast.error("News article must not be empty", {
				position: toast.POSITION.TOP_CENTER,
				autoClose: 10000,
				toastId: customId,
			});
		}
	};

	return (
		<>
			{newsArticle === null && (
				<Helmet>
					<title>Add news article | GameZone</title>
				</Helmet>
			)}
			{newsArticle !== null && (
				<Helmet>
					<title>Editing {newsArticle.UnpublishedTitle} | GameZone</title>
				</Helmet>
			)}
			<Container>
				<Row>
					<Col>
						<div className="news-editor">
							<NewsEditorToolbar />
							<ReactQuill
								theme="snow"
								modules={modules}
								formats={formats}
								value={unpublishedContent}
								onChange={onEditorStateChange}
							/>
						</div>
					</Col>
				</Row>
				<Row>
					<Col>
						<Card className="card">
							<CardBody>
								<Form className="form">
									<Row>
										<Col>
											<FormGroup>
												<Label>Title</Label>
												<Input
													className="input-field"
													type="text"
													name="UnpublishedTitle"
													innerRef={register}
													invalid={errors.UnpublishedTitle?.message}
													defaultValue={
														newsArticle !== null
															? newsArticle.UnpublishedTitle
															: null
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{errors.UnpublishedTitle?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										<Col>
											<FormGroup>
												<Label>Description</Label>
												<Input
													className="input-field"
													type="textarea"
													style={{ height: "100px", resize: "none" }}
													name="UnpublishedDescription"
													innerRef={register}
													invalid={errors.UnpublishedDescription?.message}
													defaultValue={
														newsArticle !== null
															? newsArticle.UnpublishedDescription
															: null
													}
												/>
												<FormFeedback className="input-field-error-msg">
													{errors.UnpublishedDescription?.message}
												</FormFeedback>
											</FormGroup>
										</Col>
									</Row>
									<Row>
										{newsArticle === null && (
											<Col md="1">
												<Button
													type="button"
													className="my-button"
													onClick={handleSubmit(add)}
												>
													Add
												</Button>
											</Col>
										)}
										{newsArticle !== null && (
											<Col md="1">
												<Button
													type="button"
													className="my-button"
													onClick={handleSubmit(edit)}
												>
													Save
												</Button>
											</Col>
										)}
										<Col md="1">
											<Button
												type="button"
												className="my-button"
												onClick={handleSubmit(publishNewsArticle)}
											>
												Publish
											</Button>
										</Col>
									</Row>
								</Form>
							</CardBody>
						</Card>
					</Col>
				</Row>
			</Container>
		</>
	);
};

export default NewsEditor;
