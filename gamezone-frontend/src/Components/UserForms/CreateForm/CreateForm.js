import { FormGroup, Label, Input, FormFeedback, Row, Col } from "reactstrap";
import { useFormContext } from "react-hook-form";
import "../../../Assets/css/forms.css";

const CreateForm = () => {
	const {
		register,
		formState: { errors },
	} = useFormContext();

	return (
		<>
			<Row>
				<Col>
					<FormGroup>
						<Label>Username</Label>
						<Input
							className="input-field"
							type="text"
							name="user_name"
							invalid={errors.user_name?.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.user_name?.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Password</Label>
						<Input
							className="input-field"
							type="password"
							name="password"
							invalid={errors.password?.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.password?.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Email</Label>
						<Input
							className="input-field"
							type="email"
							placeholder="exmaple@email.com"
							name="email"
							invalid={errors.email?.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.email?.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Name</Label>
						<Input
							className="input-field"
							type="text"
							name="name"
							invalid={errors.name?.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.name?.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
			<Row>
				<Col>
					<FormGroup>
						<Label>Surname</Label>
						<Input
							className="input-field"
							type="text"
							name="surname"
							invalid={errors.surname?.message}
							innerRef={register}
						/>
						<FormFeedback className="input-field-error-msg">
							{errors.surname?.message}
						</FormFeedback>
					</FormGroup>
				</Col>
			</Row>
		</>
	);
};

export default CreateForm;
