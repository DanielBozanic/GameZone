import { useEffect } from "react";
import {
	FormGroup,
	Label,
	Input,
	Button,
	FormFeedback,
	Row,
	Col,
} from "reactstrap";
import { useFormContext } from "react-hook-form";

const ProductForm = (props) => {
	const {
		register,
		setValue,
		trigger,
		formState: { errors },
	} = useFormContext();

	useEffect(() => {
		if (props.product !== null) {
			setFileAttr(
				props.product.Product.Image.Name,
				props.product.Product.Image.Type,
				props.product.Product.Image.Size
			);
			props.setFileName && props.setFileName(props.product.Product.Image.Name);
			props.setBase64Image &&
				props.setBase64Image(props.product.Product.Image.Content);
		}
	}, [props.product]);

	const handleClickFile = () => {
		document.getElementById("imageInput").click();
	};

	const handleFileChange = (event) => {
		const fileUploaded = event.target.files[0];
		props.setFileName && props.setFileName(event.target.files[0].name);
		setFileAttr(
			event.target.files[0].name,
			event.target.files[0].type,
			event.target.files[0].size
		);
		convertToBase64(fileUploaded);
		trigger("Product.Image.Name");
		trigger("Product.Image.Type");
		trigger("Product.Image.Size");
	};

	const setFileAttr = (name, type, size) => {
		setValue("Product.Image.Name", name);
		setValue("Product.Image.Type", type);
		setValue("Product.Image.Size", size);
	};

	const convertToBase64 = (file) => {
		const reader = new FileReader();
		reader.onloadend = () => {
			props.setBase64Image && props.setBase64Image(reader.result);
		};
		reader.readAsDataURL(file);
	};

	return (
		<>
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
							defaultValue={
								props.product !== null ? props.product.Product.Name : ""
							}
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
							defaultValue={
								props.product !== null ? props.product.Product.Description : ""
							}
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
							defaultValue={
								props.product !== null ? props.product.Product.Manufacturer : ""
							}
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
							defaultValue={
								props.product !== null ? props.product.Product.Price : ""
							}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.Product?.Price && errors.Product?.Price.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			{(!props.isDigital || props.isDigital === undefined) && (
				<Row>
					<Col>
						<FormGroup>
							<Label>Quantity</Label>
							<Input
								className="input-field"
								type="number"
								name="Product.Quantity"
								min="0"
								invalid={
									errors.Product?.Quantity && errors.Product?.Quantity.message
								}
								innerRef={register}
								defaultValue={
									props.product !== null ? props.product.Product.Quantity : ""
								}
							/>
							<FormFeedback className="input-field-error-msg">
								{errors.Product?.Quantity && errors.Product?.Quantity.message}
							</FormFeedback>
						</FormGroup>
					</Col>
				</Row>
			)}
			{props.isDigital && (
				<Input
					style={{ display: "none" }}
					type="number"
					name="Product.Quantity"
					innerRef={register}
					value={0}
				/>
			)}
			<Row>
				<Col>
					<FormGroup>
						<Button
							className="my-button"
							type="button"
							onClick={handleClickFile}
						>
							Choose image
						</Button>
						<Label style={{ marginLeft: "10px", marginTop: "10px" }}>
							{props.fileName}
						</Label>
						<input
							id="imageInput"
							style={{ display: "none" }}
							type="file"
							accept=".jpg,.jpeg,.png"
							name="Product.Image"
							onChange={handleFileChange}
						/>
						<input
							style={{ display: "none" }}
							name="Product.Image.Name"
							invalid={
								errors.Product?.Image?.Name &&
								errors.Product?.Image.Name.message
							}
							ref={register}
						/>

						<input
							style={{ display: "none" }}
							name="Product.Image.Type"
							invalid={
								errors.Product?.Image?.Type &&
								errors.Product?.Image.Type.message
							}
							ref={register}
						/>

						<input
							style={{ display: "none" }}
							type="number"
							name="Product.Image.Size"
							invalid={
								errors.Product?.Image?.Size &&
								errors.Product?.Image.Size.message
							}
							ref={register}
						/>

						<div className="input-field-error-msg">
							{errors.Product?.Image?.Name &&
								errors.Product?.Image.Name.message}
						</div>
						<div className="input-field-error-msg">
							{errors.Product?.Image?.Type &&
								errors.Product?.Image.Type.message}
						</div>
						<div className="input-field-error-msg">
							{errors.Product?.Image?.Size &&
								errors.Product?.Image.Size.message}
						</div>
					</FormGroup>
				</Col>
			</Row>
		</>
	);
};

export default ProductForm;
