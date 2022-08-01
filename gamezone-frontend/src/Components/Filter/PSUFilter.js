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
import * as psuAPI from "../../APIs/ProductMicroservice/psu_api";

const PSUFilter = (props) => {
	const [modal, setModal] = useState(false);
	const toggle = () => setModal(!modal);
	const [manufacturers, setManufacturers] = useState([]);
	const [powers, setPowers] = useState([]);
	const [types, setTypes] = useState([]);
	const [formFactors, setFormFactors] = useState([]);

	useEffect(() => {
		getManufacturers();
		getPowers();
		getTypes();
		getFormFactors();
	}, []);

	const { handleSubmit, setValue, control } = useForm({
		mode: "onBlur",
		reValidateMode: "onChange",
		shouldUnregister: true,
		defaultValues: {
			Manufacturers: [],
			Powers: [],
			Types: [],
			FormFactors: [],
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
			.get(psuAPI.GET_MANUFACTURERS)
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

	const getPowers = () => {
		axios
			.get(psuAPI.GET_POWERS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setPowers(temp);
				setValue("Powers", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getTypes = () => {
		axios
			.get(psuAPI.GET_TYPES)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setTypes(temp);
				setValue("Types", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const getFormFactors = () => {
		axios
			.get(psuAPI.GET_FORM_FACTORS)
			.then((res) => {
				const temp = res.data.map((value) => ({
					label: value,
					value: value,
				}));
				setFormFactors(temp);
				setValue("FormFactors", res.data);
			})
			.catch((err) => {
				console.log(err);
			});
	};

	return (
		<>
			<Col md="2">
				<Button className="my-button filter-button" onClick={toggle}>
					Filter
				</Button>
			</Col>

			<Modal isOpen={modal} toggle={toggle}>
				<ModalHeader className="title" toggle={toggle}>
					Filter
				</ModalHeader>
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
							<Label>Filter by power</Label>
							<Controller
								name={"Powers"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={powers}
											isMulti={true}
											onChange={(val) => {
												onChange(
													powers
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={powers.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by type</Label>
							<Controller
								name={"Types"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={types}
											isMulti={true}
											onChange={(val) => {
												onChange(
													types
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={types.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
						<FormGroup>
							<Label>Filter by form factor</Label>
							<Controller
								name={"FormFactors"}
								control={control}
								render={({ value, onChange }) => {
									return (
										<Select
											options={formFactors}
											isMulti={true}
											onChange={(val) => {
												onChange(
													formFactors
														.filter((o) => val.includes(o))
														.map((p) => p.value)
												);
											}}
											value={formFactors.find((c) => c.value === value)}
										/>
									);
								}}
							/>
						</FormGroup>
					</Form>
				</ModalBody>
				<ModalFooter>
					<Button className="my-button" onClick={handleSubmit(filter)}>
						Filter
					</Button>
				</ModalFooter>
			</Modal>
		</>
	);
};

export default PSUFilter;
