import {
	Button,
	Modal,
	ModalHeader,
	ModalBody,
	ModalFooter,
	Form,
	FormGroup,
	Label,
	Input,
	FormFeedback,
	Row,
	Col,
} from "reactstrap";
import { toast } from "react-toastify";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { reportSchema } from "./ReportSchema";
import axios from "axios";
import * as reportAPI from "../../APIs/ContactAndReportMicroservice/report_api";

toast.configure();
const ReportModal = (props) => {
	const customId = "ReportModal";

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm({
		resolver: yupResolver(reportSchema),
		mode: "onChange",
	});

	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const reasons = [
		"Offensive language",
		"Inappropriate language",
		"Harrasment",
		"Spam",
		"Misinformation",
	];

	const report = (data) => {
		data.UserId = props.userId;
		axios
			.post(`${reportAPI.ADD_REPORT}`, data)
			.then((res) => {
				toast.success(res.data, {
					position: toast.POSITION.TOP_CENTER,
					toastId: customId,
					autoClose: 5000,
				});
				toggle();
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<Row>
			<Col>
				<Button
					style={{
						marginRight: "5px",
						marginTop: "5px",
					}}
					className="my-button"
					type="button"
					onClick={toggle}
				>
					Report
				</Button>
			</Col>
			<Modal isOpen={modal} toggle={toggle}>
				<ModalHeader className="title" toggle={toggle}>
					Report user {props.username}
				</ModalHeader>
				<ModalBody>
					<Form>
						<FormGroup>
							<Label>Select reason</Label>
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
							<Label>Description</Label>
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
					</Form>
				</ModalBody>
				<ModalFooter>
					<Button className="my-button" onClick={handleSubmit(report)}>
						Report
					</Button>
				</ModalFooter>
			</Modal>
		</Row>
	);
};

export default ReportModal;
