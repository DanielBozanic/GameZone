import "../../Assets/css/filter.css";
import {
	Button,
	Modal,
	ModalHeader,
	ModalBody,
	ModalFooter,
	Form,
	FormGroup,
	Label,
	Col,
} from "reactstrap";
import { useState, useEffect } from "react";
import Select from "react-select";
import { useForm, Controller } from "react-hook-form";
import axios from "axios";
import * as keyboardAPI from "../../APIs/ProductMicroservice/keyboard_api";

const KeyboardFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [keyboardConnectors, setKeyboardConnectors] = useState([]);
	const [keyTypes, setKeyTypes] = useState([]);

	useEffect(() => {
		getManufacturers();
		getKeyboardConnectors();
		getKeyTypes();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			KeyboardConnectors: [],
			KeyTypes: [],
		},
	});

	const onFilterClick = (filterData) => {
		toggle();
		props.onFilterClick && props.onFilterClick(filterData);
	};

	const filter = (filterData) => {
		onFilterClick(filterData);
	};

	const getManufacturers = () => {
		axios
			.get(keyboardAPI.GET_MANUFACTURERS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setManufacturers(temp);
				setValue("Manufacturers", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getKeyboardConnectors = () => {
		axios
			.get(keyboardAPI.GET_KEYBOARD_CONNECTORS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setKeyboardConnectors(temp);
				setValue("KeyboardConnectors", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getKeyTypes = () => {
		axios
			.get(keyboardAPI.GET_KEY_TYPES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setKeyTypes(temp);
				setValue("KeyTypes", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			<Col md="2">
				<Button className="filter-button" onClick={toggle}>
					Filter
				</Button>
			</Col>

			<Modal isOpen={modal} toggle={toggle}>
				<ModalHeader toggle={toggle}>Filter</ModalHeader>
				<ModalBody>
					<Form>
						<FormGroup>
							<Label>Filter by manufacturer</Label>
							<Controller
								name={"Manufacturers"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={manufacturers}
											isMulti={true}
											onChange={(val) => {
												onChange(
													manufacturers
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={manufacturers.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by keyboard connector</Label>
							<Controller
								name={"KeyboardConnectors"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={keyboardConnectors}
											isMulti={true}
											onChange={(val) => {
												onChange(
													keyboardConnectors
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={keyboardConnectors.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by key type</Label>
							<Controller
								name={"KeyTypes"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={keyTypes}
											isMulti={true}
											onChange={(val) => {
												onChange(
													keyTypes
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={keyTypes.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
					</Form>
				</ModalBody>
				<ModalFooter>
					<Button className="confirm-filter-btn" onClick={handleSubmit(filter)}>
						Filter
					</Button>
				</ModalFooter>
			</Modal>
		</>
	);
};

export default KeyboardFilter;
