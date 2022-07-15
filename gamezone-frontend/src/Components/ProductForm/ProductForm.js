import {
	FormGroup,
	Label,
	Input,
	Button,
	FormFeedback,
	Row,
	Col,
} from "reactstrap";
import { useState } from "react";
import { useFormContext } from "react-hook-form";
import "../../Assets/css/forms.css";

const ProductForm = (props) => {
	const {
		register,
		formState: { errors },
	} = useFormContext();
	const [fileName, setFileName] = useState("");

	const handleClickFile = () => {
		document.getElementById("imageInput").click();
	};

	const handleFileChange = (event) => {
		const fileUploaded = event.target.files[0];
		setFileName(event.target.files[0].name);
		convertToBase64(fileUploaded);
	};

	const convertToBase64 = (file) => {
		const reader = new FileReader();
		reader.onloadend = () => {
			props.setBase64Image && props.setBase64Image(reader.result);
		};
		reader.readAsDataURL(file);
	};

	return (
		<div className="form-border">
			<Row>
				<Col>
					<FormGroup>
						<Label>Name</Label>
						<Input
							className="input-field"
							type="text"
							name="Product.Name"
							invalid={errors.Product?.Name && errors.Product?.Name.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.Product?.Name && errors.Product?.Name.message}
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
							name="Product.Description"
							invalid={
								errors.Product?.Description &&
								errors.Product?.Description.message
							}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.Product?.Description &&
								errors.Product?.Description.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Manufacturer</Label>
						<Input
							className="input-field"
							type="text"
							name="Product.Manufacturer"
							invalid={
								errors.Product?.Manufacturer &&
								errors.Product?.Manufacturer.message
							}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.Product?.Manufacturer &&
								errors.Product?.Manufacturer.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Price</Label>
						<Input
							className="input-field"
							type="number"
							name="Product.Price"
							min="0"
							invalid={errors.Product?.Price && errors.Product?.Price.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.Product?.Price && errors.Product?.Price.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Amount</Label>
						<Input
							className="input-field"
							type="number"
							name="Product.Amount"
							min="0"
							invalid={errors.Product?.Amount && errors.Product?.Amount.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.Product?.Amount && errors.Product?.Amount.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Button
							className="form-btn"
							type="button"
							onClick={handleClickFile}
						>
							Choose image
						</Button>
						<Label style={{ marginLeft: "10px", marginTop: "10px" }}>
							{fileName}
						</Label>
						<input
							id="imageInput"
							style={{ display: "none" }}
							type="file"
							accept=".jpg,.jpeg,.png"
							name="Product.Image"
							invalid={errors.Product?.Image && errors.Product?.Image.message}
							ref={register}
							onChange={handleFileChange}
						/>
						<p className="input-field-error-msg">
							{errors.Product?.Image && errors.Product?.Image.message}
						</p>
					</FormGroup>
				</Col>
			</Row>
		</div>
	);
};

export default ProductForm;
