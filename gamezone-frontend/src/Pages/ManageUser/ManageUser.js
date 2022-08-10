import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { banSchema } from "./BanSchema";
import {
	CardHeader,
	CardTitle,
	CardBody,
	CardText,
	Card,
	Row,
	Col,
	Button,
	Form,
	Label,
	FormGroup,
	Input,
	FormFeedback,
	Spinner,
} from "reactstrap";
import axios from "axios";
import { toast } from "react-toastify";
import * as productCommentAPI from "../../APIs/CommentAndRatingMicroservice/product_comment_api";
import * as newsCommentAPI from "../../APIs/NewsMicroservice/news_comment_api";
import * as reportAPI from "../../APIs/ContactAndReportMicroservice/report_api";
import * as banAPI from "../../APIs/ContactAndReportMicroservice/ban_api";

toast.configure();
const ManageUser = () => {
	const customId = "ManagerUser";
	const reasons = [
		"Offensive language",
		"Inappropriate language",
		"Harrasment",
		"Spam",
		"Misinformation",
	];
	const { id } = useParams();
	const [productComments, setProductComments] = useState([]);
	const [newsComments, setNewsComments] = useState([]);
	const [reports, setReports] = useState([]);
	const [banHistory, setBanHistory] = useState([]);
	const [loadingComments, setLoadingComments] = useState(true);
	const [loadingBanHistory, setLoadingBanHistory] = useState(true);
	const [loadingReports, setLoadingReports] = useState(true);

	const {
		register,
		handleSubmit,
		formState: { errors },
		reset,
	} = useForm({
		resolver: yupResolver(banSchema),
		mode: "onChange",
	});

	useEffect(() => {
		getProductCommentsByUserId();
		getNewsCommentsByUserId();
		getReportsByUserId();
		getBanHistory();
	}, []);

	const getProductCommentsByUserId = () => {
		axios
			.get(`${productCommentAPI.GET_BY_USER_ID}/${id}`)
			.then((res) => {
				setProductComments(res.data);
				setLoadingComments(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getNewsCommentsByUserId = () => {
		axios
			.get(`${newsCommentAPI.GET_BY_USER_ID}/${id}`)
			.then((res) => {
				setNewsComments(res.data);
				setLoadingComments(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getReportsByUserId = () => {
		axios
			.get(`${reportAPI.GET_REPORTS_BY_USER_ID}/${id}`)
			.then((res) => {
				setReports(res.data);
				setLoadingReports(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getBanHistory = () => {
		axios
			.get(`${banAPI.GET_USER_BAN_HISTORY}/${id}`)
			.then((res) => {
				setBanHistory(res.data);
				setLoadingBanHistory(false);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const banUser = (data) => {
		setLoadingBanHistory(true);
		data.UserId = Number(id);
		axios
			.post(`${banAPI.ADD_BAN}`, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				getBanHistory();
				reset();
				sendEmailToBannedUser(data);
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: false,
				});
				setLoadingBanHistory(false);
			});
	};

	const sendEmailToBannedUser = (ban) => {
		axios.post(`${banAPI.SEND_EMAIL_TO_BANNED_USER}`, ban);
	};

	return (
		<>
			<Row style={{ margin: "auto", padding: "5px" }}>
				<Col>
					<Card style={{ height: "534px" }} className="card">
						<CardHeader>
							<CardTitle className="title" tag="h5">
								Reports Received By Other Users
							</CardTitle>
						</CardHeader>
						{loadingReports && (
							<div className="div-spinner">
								<Spinner className="spinner" />
							</div>
						)}
						{!loadingReports && (
							<CardBody>
								{reports.length > 0 &&
									reports.map((report) => {
										return (
											<CardText>
												<span style={{ fontWeight: "bold" }}>
													[{new Date(report.DateTime).toLocaleDateString()}{" "}
													{new Date(report.DateTime).toLocaleTimeString()}] [
													{report.Reason}]
												</span>
												- {report.Description}
											</CardText>
										);
									})}
								{reports.length <= 0 && (
									<CardTitle className="title" tag="h5">
										User has not been reported
									</CardTitle>
								)}
							</CardBody>
						)}
					</Card>
				</Col>
				<Col>
					<Card className="card">
						<CardBody>
							<Form className="form">
								<FormGroup>
									<Label>Reason For Ban</Label>
									<Input
										className="input-field"
										type="select"
										name="Reason"
										innerRef={register}
									>
										{reasons.map((reason) => {
											return <option value={reason}>{reason}</option>;
										})}
									</Input>
								</FormGroup>
								<FormGroup>
									<Label>Ban Description</Label>
									<Input
										style={{ height: "200px", resize: "none" }}
										className="input-field"
										type="textarea"
										name="Description"
										innerRef={register}
										invalid={errors.Description?.message}
									/>
									<FormFeedback className="input-field-error-msg">
										{errors.Description?.message}
									</FormFeedback>
								</FormGroup>
								<FormGroup>
									<Label>Ban Expiration</Label>
									<Input
										className="input-field"
										type="datetime-local"
										name="ExpirationDate"
										innerRef={register}
										invalid={errors.ExpirationDate?.message}
									/>
									<FormFeedback className="input-field-error-msg">
										{errors.ExpirationDate?.message}
									</FormFeedback>
								</FormGroup>
								<Button
									className="my-button"
									type="button"
									onClick={handleSubmit(banUser)}
								>
									Ban
								</Button>
							</Form>
						</CardBody>
					</Card>
				</Col>
			</Row>
			<Row style={{ margin: "auto", padding: "5px" }}>
				<Col md="6">
					<Card
						style={{ marginTop: "10px", marginBottom: "10px" }}
						className="card"
					>
						<CardHeader>
							<CardTitle className="title" tag="h5">
								Comment History
							</CardTitle>
						</CardHeader>
						{loadingComments && (
							<div className="div-spinner">
								<Spinner className="spinner" />
							</div>
						)}
						{!loadingComments && (
							<CardBody>
								{(productComments.length > 0 || newsComments.length > 0) &&
									productComments
										.concat(newsComments)
										.sort((a, b) => (a.DateTime > b.DateTime ? -1 : 1))
										.map((comment) => {
											return (
												<CardText>
													<span style={{ fontWeight: "bold" }}>
														[{new Date(comment.DateTime).toLocaleDateString()}{" "}
														{new Date(comment.DateTime).toLocaleTimeString()}]
													</span>{" "}
													{comment.Comment}
												</CardText>
											);
										})}
								{productComments.length <= 0 && newsComments.length <= 0 && (
									<CardTitle className="title" tag="h5">
										User has not posted any comments
									</CardTitle>
								)}
							</CardBody>
						)}
					</Card>
				</Col>

				<Col md="6">
					<Card
						style={{ marginTop: "10px", marginBottom: "10px" }}
						className="card"
					>
						<CardHeader>
							<CardTitle className="title" tag="h5">
								Ban History
							</CardTitle>
						</CardHeader>
						{loadingBanHistory && (
							<div className="div-spinner">
								<Spinner className="spinner" />
							</div>
						)}
						{!loadingBanHistory && (
							<CardBody>
								{banHistory.length > 0 &&
									banHistory.map((ban) => {
										return (
											<CardText>
												Banned until{" "}
												<span style={{ fontWeight: "bold" }}>
													[{new Date(ban.ExpirationDate).toLocaleDateString()}{" "}
													{new Date(ban.ExpirationDate).toLocaleTimeString()}]{" "}
												</span>
												for{" "}
												<span style={{ fontWeight: "bold" }}>{ban.Reason}</span>
												.
											</CardText>
										);
									})}
								{banHistory.length <= 0 && (
									<CardTitle className="title" tag="h5">
										User has not been banned
									</CardTitle>
								)}
							</CardBody>
						)}
					</Card>
				</Col>
			</Row>
		</>
	);
};

export default ManageUser;
