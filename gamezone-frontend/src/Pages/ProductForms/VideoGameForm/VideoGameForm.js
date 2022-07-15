import { useForm, FormProvider } from "react-hook-form";
import { useState } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { productFormSchema } from "../../../Components/ProductForm/ProductFormSchema";
import { videoGameFormSchema } from "./VideoGameFormSchema";
import {
	Button,
	Form,
	FormGroup,
	Label,
	Input,
	Card,
	CardTitle,
	CardBody,
	FormFeedback,
	Container,
	Row,
	Col,
} from "reactstrap";
import axios from "axios";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "../../../Assets/css/forms.css";
import ProductForm from "../../../Components/ProductForm/ProductForm";
import * as videoGameAPI from "../../../APIs/ProductMicroservice/video_game_api";
import * as helperFunctions from "../../../Utils/HelperFunctions";

toast.configure();
const VideoGameForm = () => {
	const customId = "videoGameForm";

	const [digital, setDigital] = useState(false);
	const [base64Image, setBase64Image] = useState("");

	const methods = useForm({
		resolver: yupResolver(
			helperFunctions.merge(productFormSchema, videoGameFormSchema)
		),
		mode: "onChange",
	});

	const addNewVideoGame = (data) => {
		data.Product.Image = base64Image;
		data.Digital = helperFunctions.str2Bool(data.Digital);
		console.log(data);
		axios
			.post(videoGameAPI.CREATE, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: 5000,
					toastId: customId,
				});
				methods.reset();
			})
			.catch((err) => {
				toast.error(err.response.data, {
					position: toast.POSITION.TOP_CENTER,
					autoClose: false,
					toastId: customId,
				});
			});
	};

	return (
		<Container>
			<Row>
				<Col>
					<Card className="form-card">
						<CardTitle className="form-title" tag="h2">
							Add new video game
						</CardTitle>
						<CardBody>
							<FormProvider {...methods}>
								<Form className="form">
									<ProductForm setBase64Image={setBase64Image} />
									<div className="form-border">
										<Row>
											<Col>
												<FormGroup>
													<Label>Platform</Label>
													<Input
														className="input-field"
														type="text"
														name="Platform"
														invalid={methods.formState.errors.Platform?.message}
														innerRef={methods.register}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Platform?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<div>
													<Label>Digital</Label>
												</div>
												<span>
													<Label>Yes</Label>
													<Input
														className="ml-2"
														type="radio"
														name="Digital"
														checked={digital}
														value={digital}
														innerRef={methods.register}
														onChange={() => setDigital(true)}
													/>
												</span>
												<span className="pl-5">
													<Label>No</Label>
													<Input
														className="ml-2"
														type="radio"
														name="Digital"
														checked={!digital}
														value={digital}
														innerRef={methods.register}
														onChange={() => setDigital(false)}
													/>
												</span>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Rating</Label>
													<Input
														className="input-field"
														type="number"
														name="Rating"
														min="1"
														invalid={methods.formState.errors.Rating?.message}
														innerRef={methods.register}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Rating?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Genre</Label>
													<Input
														className="input-field"
														type="text"
														name="Genre"
														invalid={methods.formState.errors.Genre?.message}
														innerRef={methods.register}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.Genre?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
										<Row>
											<Col>
												<FormGroup>
													<Label>Release date</Label>
													<Input
														className="input-field"
														type="date"
														name="ReleaseDate"
														invalid={
															methods.formState.errors.ReleaseDate?.message
														}
														innerRef={methods.register}
													/>
													<FormFeedback className="input-field-error-msg">
														{methods.formState.errors.ReleaseDate?.message}
													</FormFeedback>
												</FormGroup>
											</Col>
										</Row>
									</div>
									<Row>
										<Col>
											<Button
												className="confirm-form-btn"
												type="button"
												onClick={methods.handleSubmit(addNewVideoGame)}
											>
												Add
											</Button>
										</Col>
									</Row>
								</Form>
							</FormProvider>
						</CardBody>
					</Card>
				</Col>
			</Row>
		</Container>
	);
};

export default VideoGameForm;
